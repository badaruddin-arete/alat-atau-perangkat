package rubrik

import (
	"strings"
	"testing"
)

func TestDelapanSumbuEmpatPerKelompok(t *testing.T) {
	if got := len(DaftarSumbu); got != 8 {
		t.Fatalf("jumlah sumbu = %d, ingin 8", got)
	}
	for _, k := range []Kelompok{Keterlibatan, Ketersediaan} {
		if got := len(SumbuKelompok(k)); got != 4 {
			t.Errorf("kelompok %s punya %d sumbu, ingin 4", k, got)
		}
	}
}

func TestKodeSumbuUnikDanLengkap(t *testing.T) {
	dilihat := map[string]bool{}
	for _, s := range DaftarSumbu {
		if dilihat[s.Kode] {
			t.Errorf("kode %q ganda", s.Kode)
		}
		dilihat[s.Kode] = true
		if s.Nama == "" || s.Pertanyaan == "" {
			t.Errorf("sumbu %q tanpa nama atau pertanyaan", s.Kode)
		}
	}
}

func TestCariSumbu(t *testing.T) {
	s, ada := CariSumbu("mudah")
	if !ada || s.Kelompok != Ketersediaan {
		t.Fatalf("CariSumbu(mudah) = %+v, %v", s, ada)
	}
	if _, ada := CariSumbu("cepat"); ada {
		t.Error("kode yang tidak ada tidak boleh ditemukan")
	}
}

// kasusContoh membuat kasus sah dengan skor yang sama di semua sumbu.
func kasusContoh(skor int) Kasus {
	p := map[string]Penilaian{}
	for _, s := range DaftarSumbu {
		p[s.Kode] = Penilaian{Skor: skor, Alasan: "alasan uji"}
	}
	return Kasus{ID: "uji", Nama: "Kasus uji", Pasangan: "lawan", Penilaian: p}
}

func TestValidasiKasusSah(t *testing.T) {
	for skor := SkorMin; skor <= SkorMaks; skor++ {
		if err := kasusContoh(skor).Validasi(); err != nil {
			t.Errorf("skor %d: %v", skor, err)
		}
	}
}

func TestValidasiMenolak(t *testing.T) {
	uji := []struct {
		nama  string
		ubah  func(*Kasus)
		pesan string
	}{
		{"id kosong", func(k *Kasus) { k.ID = " " }, "id kosong"},
		{"nama kosong", func(k *Kasus) { k.Nama = "" }, "nama kosong"},
		{"pasangan diri sendiri", func(k *Kasus) { k.Pasangan = k.ID }, "dirinya sendiri"},
		{"sumbu hilang", func(k *Kasus) { delete(k.Penilaian, "aman") }, `"aman" belum dinilai`},
		{"skor terlalu tinggi", func(k *Kasus) { k.Penilaian["tubuh"] = Penilaian{5, "x"} }, "di luar"},
		{"skor negatif", func(k *Kasus) { k.Penilaian["tubuh"] = Penilaian{-1, "x"} }, "di luar"},
		{"alasan kosong", func(k *Kasus) { k.Penilaian["mudah"] = Penilaian{2, "  "} }, "alasan kosong"},
		{"sumbu asing", func(k *Kasus) { k.Penilaian["cepat"] = Penilaian{2, "x"} }, "tidak dikenal"},
	}
	for _, u := range uji {
		t.Run(u.nama, func(t *testing.T) {
			k := kasusContoh(2)
			u.ubah(&k)
			err := k.Validasi()
			if err == nil {
				t.Fatal("harus ditolak")
			}
			if !strings.Contains(err.Error(), u.pesan) {
				t.Errorf("galat %q tidak memuat %q", err, u.pesan)
			}
		})
	}
}

func TestValidasiMengumpulkanSemuaMasalah(t *testing.T) {
	k := kasusContoh(2)
	k.Nama = ""
	delete(k.Penilaian, "instan")
	err := k.Validasi()
	if err == nil || !strings.Contains(err.Error(), "nama kosong") || !strings.Contains(err.Error(), "instan") {
		t.Fatalf("ingin kedua masalah dilaporkan, dapat %v", err)
	}
}

const jsonSah = `{
  "id": "uji", "nama": "Kasus uji", "pasangan": "lawan", "kisah": "fiktif",
  "penilaian": {
    "tubuh": {"skor": 3, "alasan": "a"}, "keterampilan": {"skor": 3, "alasan": "a"},
    "keterlihatan": {"skor": 3, "alasan": "a"}, "kebersamaan": {"skor": 3, "alasan": "a"},
    "instan": {"skor": 1, "alasan": "a"}, "di_mana_saja": {"skor": 1, "alasan": "a"},
    "aman": {"skor": 1, "alasan": "a"}, "mudah": {"skor": 1, "alasan": "a"}
  }
}`

func TestMuatJSONSah(t *testing.T) {
	k, err := Muat(strings.NewReader(jsonSah))
	if err != nil {
		t.Fatal(err)
	}
	if k.ID != "uji" || k.Penilaian["tubuh"].Skor != 3 {
		t.Errorf("hasil muat keliru: %+v", k)
	}
}

func TestMuatMenolakBidangAsing(t *testing.T) {
	rusak := strings.Replace(jsonSah, `"kisah"`, `"bobot_rahasia": 9, "kisah"`, 1)
	if _, err := Muat(strings.NewReader(rusak)); err == nil {
		t.Fatal("bidang asing harus ditolak")
	}
}

func TestMuatMenolakJSONRusak(t *testing.T) {
	if _, err := Muat(strings.NewReader(`{"id": `)); err == nil {
		t.Fatal("JSON rusak harus ditolak")
	}
}

func TestMuatMenolakKasusTidakSah(t *testing.T) {
	rusak := strings.Replace(jsonSah, `"skor": 3`, `"skor": 7`, 1)
	if _, err := Muat(strings.NewReader(rusak)); err == nil {
		t.Fatal("skor di luar rentang harus ditolak")
	}
}
