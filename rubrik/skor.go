package rubrik

import (
	"fmt"
	"sort"
)

// Bobot memberi bobot pada sumbu. Sumbu yang tidak disebut berbobot 1.
// Bobot bawaan semuanya 1; bobot dibuka supaya kepekaan hasil bisa diuji.
type Bobot map[string]float64

func (b Bobot) dari(kode string) float64 {
	if w, ada := b[kode]; ada {
		return w
	}
	return 1
}

// Validasi menolak bobot negatif, bobot untuk sumbu yang tidak dikenal,
// dan kelompok yang seluruh bobotnya nol.
func (b Bobot) Validasi() error {
	for kode, w := range b {
		if _, ada := CariSumbu(kode); !ada {
			return fmt.Errorf("bobot untuk sumbu tidak dikenal: %q", kode)
		}
		if w < 0 {
			return fmt.Errorf("bobot sumbu %q negatif", kode)
		}
	}
	for _, kel := range []Kelompok{Keterlibatan, Ketersediaan} {
		total := 0.0
		for _, s := range SumbuKelompok(kel) {
			total += b.dari(s.Kode)
		}
		if total == 0 {
			return fmt.Errorf("semua bobot kelompok %s nol", kel)
		}
	}
	return nil
}

// RataKelompok adalah rata-rata berbobot skor satu kelompok, dalam skala 0..4.
func RataKelompok(k Kasus, kel Kelompok, b Bobot) float64 {
	jumlah, total := 0.0, 0.0
	for _, s := range SumbuKelompok(kel) {
		w := b.dari(s.Kode)
		jumlah += w * float64(k.Penilaian[s.Kode].Skor)
		total += w
	}
	return jumlah / total
}

// Indeks menempatkan kasus pada spektrum -1 (benda fokal) sampai +1 (perangkat):
// rata-rata ketersediaan dikurangi rata-rata keterlibatan, dibagi skor maksimum.
func Indeks(k Kasus, b Bobot) float64 {
	return (RataKelompok(k, Ketersediaan, b) - RataKelompok(k, Keterlibatan, b)) / SkorMaks
}

// Golongan memberi nama kasar pada indeks. Ambangnya pilihan penyusun,
// bukan batas alami, dan ditulis terang di README.
func Golongan(indeks float64) string {
	switch {
	case indeks <= -0.5:
		return "benda fokal"
	case indeks < -0.15:
		return "condong benda fokal"
	case indeks <= 0.15:
		return "di tengah"
	case indeks < 0.5:
		return "condong perangkat"
	default:
		return "perangkat"
	}
}

// Hasil merangkum penilaian satu kasus.
type Hasil struct {
	Kasus        Kasus
	Keterlibatan float64
	Ketersediaan float64
	Indeks       float64
	Golongan     string
}

// Nilai menghitung hasil satu kasus dengan bobot tertentu.
func Nilai(k Kasus, b Bobot) Hasil {
	i := Indeks(k, b)
	return Hasil{
		Kasus:        k,
		Keterlibatan: RataKelompok(k, Keterlibatan, b),
		Ketersediaan: RataKelompok(k, Ketersediaan, b),
		Indeks:       i,
		Golongan:     Golongan(i),
	}
}

// Selisih adalah perbedaan skor satu sumbu antara dua kasus.
type Selisih struct {
	Sumbu Sumbu
	Fokal Penilaian
	Lawan Penilaian
	Beda  int // skor lawan dikurangi skor kasus fokal
}

// Perbandingan membaca pasangan kasus dengan pertanyaan Borgmann:
// apa yang didapat dari pergeseran ke arah perangkat, dan apa yang tergeser.
type Perbandingan struct {
	Fokal, Lawan Hasil
	Didapat      []Selisih // sumbu ketersediaan yang naik
	Tergeser     []Selisih // sumbu keterlibatan yang turun
	Lainnya      []Selisih // sumbu yang sama atau bergerak berlawanan arah
}

// Bandingkan menempatkan kasus dengan indeks lebih rendah sebagai sisi fokal.
func Bandingkan(a, c Kasus, b Bobot) Perbandingan {
	ha, hc := Nilai(a, b), Nilai(c, b)
	if hc.Indeks < ha.Indeks {
		ha, hc = hc, ha
	}
	p := Perbandingan{Fokal: ha, Lawan: hc}
	for _, s := range DaftarSumbu {
		f, l := ha.Kasus.Penilaian[s.Kode], hc.Kasus.Penilaian[s.Kode]
		sel := Selisih{Sumbu: s, Fokal: f, Lawan: l, Beda: l.Skor - f.Skor}
		switch {
		case s.Kelompok == Ketersediaan && sel.Beda > 0:
			p.Didapat = append(p.Didapat, sel)
		case s.Kelompok == Keterlibatan && sel.Beda < 0:
			p.Tergeser = append(p.Tergeser, sel)
		default:
			p.Lainnya = append(p.Lainnya, sel)
		}
	}
	urut := func(xs []Selisih) {
		sort.SliceStable(xs, func(i, j int) bool { return abs(xs[i].Beda) > abs(xs[j].Beda) })
	}
	urut(p.Didapat)
	urut(p.Tergeser)
	return p
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
