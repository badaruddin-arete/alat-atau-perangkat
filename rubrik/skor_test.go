package rubrik

import (
	"math"
	"testing"
)

// kasusDua membuat kasus dengan skor keterlibatan dan ketersediaan seragam.
func kasusDua(id string, libat, sedia int) Kasus {
	p := map[string]Penilaian{}
	for _, s := range DaftarSumbu {
		skor := libat
		if s.Kelompok == Ketersediaan {
			skor = sedia
		}
		p[s.Kode] = Penilaian{Skor: skor, Alasan: "uji"}
	}
	return Kasus{ID: id, Nama: id, Penilaian: p}
}

func dekat(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

func TestIndeksUjung(t *testing.T) {
	uji := []struct {
		libat, sedia int
		ingin        float64
	}{
		{4, 0, -1}, {0, 4, 1}, {2, 2, 0}, {3, 1, -0.5}, {1, 2, 0.25},
	}
	for _, u := range uji {
		if got := Indeks(kasusDua("x", u.libat, u.sedia), nil); !dekat(got, u.ingin) {
			t.Errorf("Indeks(%d,%d) = %v, ingin %v", u.libat, u.sedia, got, u.ingin)
		}
	}
}

func TestIndeksSelaluDalamRentang(t *testing.T) {
	for l := SkorMin; l <= SkorMaks; l++ {
		for s := SkorMin; s <= SkorMaks; s++ {
			i := Indeks(kasusDua("x", l, s), Bobot{"tubuh": 3, "mudah": 0.5})
			if i < -1 || i > 1 {
				t.Fatalf("indeks %v di luar [-1, 1]", i)
			}
		}
	}
}

func TestBobotMengubahRataBerbobot(t *testing.T) {
	k := kasusDua("x", 0, 0)
	k.Penilaian["tubuh"] = Penilaian{4, "uji"}
	if got := RataKelompok(k, Keterlibatan, nil); !dekat(got, 1) {
		t.Errorf("tanpa bobot = %v, ingin 1", got)
	}
	// tubuh berbobot 4, tiga sumbu lain berbobot 1: 16/7.
	if got := RataKelompok(k, Keterlibatan, Bobot{"tubuh": 4}); !dekat(got, 16.0/7) {
		t.Errorf("dengan bobot = %v, ingin %v", got, 16.0/7)
	}
}

func TestValidasiBobot(t *testing.T) {
	sah := []Bobot{nil, {}, {"tubuh": 2, "mudah": 0}}
	for _, b := range sah {
		if err := b.Validasi(); err != nil {
			t.Errorf("%v: %v", b, err)
		}
	}
	nol := Bobot{}
	for _, s := range SumbuKelompok(Ketersediaan) {
		nol[s.Kode] = 0
	}
	salah := []Bobot{{"cepat": 1}, {"tubuh": -1}, nol}
	for _, b := range salah {
		if err := b.Validasi(); err == nil {
			t.Errorf("%v harus ditolak", b)
		}
	}
}

func TestGolongan(t *testing.T) {
	uji := map[float64]string{
		-1: "benda fokal", -0.5: "benda fokal", -0.3: "condong benda fokal",
		-0.15: "di tengah", 0: "di tengah", 0.15: "di tengah",
		0.3: "condong perangkat", 0.5: "perangkat", 1: "perangkat",
	}
	for i, ingin := range uji {
		if got := Golongan(i); got != ingin {
			t.Errorf("Golongan(%v) = %q, ingin %q", i, got, ingin)
		}
	}
}

func TestBandingkanMenempatkanSisiFokal(t *testing.T) {
	fokal, perangkat := kasusDua("fokal", 4, 0), kasusDua("perangkat", 0, 4)
	for _, p := range []Perbandingan{Bandingkan(fokal, perangkat, nil), Bandingkan(perangkat, fokal, nil)} {
		if p.Fokal.Kasus.ID != "fokal" || p.Lawan.Kasus.ID != "perangkat" {
			t.Fatalf("urutan keliru: %s lalu %s", p.Fokal.Kasus.ID, p.Lawan.Kasus.ID)
		}
		if len(p.Didapat) != 4 || len(p.Tergeser) != 4 || len(p.Lainnya) != 0 {
			t.Errorf("didapat %d, tergeser %d, lainnya %d", len(p.Didapat), len(p.Tergeser), len(p.Lainnya))
		}
	}
}

func TestBandingkanMemisahkanArahBerlawanan(t *testing.T) {
	fokal, lawan := kasusDua("a", 3, 1), kasusDua("b", 1, 3)
	// Kebersamaan justru naik di sisi perangkat: bukan "tergeser".
	lawan.Penilaian["kebersamaan"] = Penilaian{4, "uji"}
	p := Bandingkan(fokal, lawan, nil)
	if len(p.Tergeser) != 3 || len(p.Lainnya) != 1 || p.Lainnya[0].Sumbu.Kode != "kebersamaan" {
		t.Fatalf("tergeser %d, lainnya %+v", len(p.Tergeser), p.Lainnya)
	}
}

func TestBandingkanMengurutkanSelisihTerbesar(t *testing.T) {
	fokal, lawan := kasusDua("a", 4, 0), kasusDua("b", 4, 0)
	lawan.Penilaian["aman"] = Penilaian{1, "uji"}
	lawan.Penilaian["mudah"] = Penilaian{4, "uji"}
	p := Bandingkan(fokal, lawan, nil)
	if len(p.Didapat) != 2 || p.Didapat[0].Sumbu.Kode != "mudah" {
		t.Fatalf("urutan didapat keliru: %+v", p.Didapat)
	}
}
