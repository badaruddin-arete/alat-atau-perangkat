package kasus

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

func TestDuaBelasKasusBawaanSah(t *testing.T) {
	daftar, err := Semua()
	if err != nil {
		t.Fatal(err)
	}
	if len(daftar) != 12 {
		t.Fatalf("jumlah kasus = %d, ingin 12", len(daftar))
	}
	if got := len(Pasangan(daftar)); got != 6 {
		t.Errorf("jumlah pasangan = %d, ingin 6", got)
	}
}

func TestSetiapKasusPunyaKisah(t *testing.T) {
	daftar, _ := Semua()
	for _, k := range daftar {
		if len(strings.TrimSpace(k.Kisah)) < 40 {
			t.Errorf("kasus %q tidak punya kisah yang cukup", k.ID)
		}
	}
}

// Dalam setiap pasangan, satu sisi harus lebih condong ke perangkat.
// Kalau tidak, pasangan itu tidak mengajarkan apa pun.
func TestSetiapPasanganBerbeda(t *testing.T) {
	daftar, _ := Semua()
	for _, p := range Pasangan(daftar) {
		a, b := rubrik.Indeks(p[0], nil), rubrik.Indeks(p[1], nil)
		if a == b {
			t.Errorf("pasangan %s dan %s berindeks sama", p[0].ID, p[1].ID)
		}
	}
}

func TestCari(t *testing.T) {
	daftar, _ := Semua()
	if k, ada := Cari(daftar, "kopi-tubruk"); !ada || k.Pasangan != "kopi-kapsul" {
		t.Fatalf("Cari(kopi-tubruk) = %+v, %v", k, ada)
	}
	if _, ada := Cari(daftar, "teh-celup"); ada {
		t.Error("kasus yang tidak ada tidak boleh ditemukan")
	}
}

func kasusJSON(id, pasangan string) string {
	var b strings.Builder
	b.WriteString(`{"id":"` + id + `","nama":"N","pasangan":"` + pasangan + `","kisah":"k","penilaian":{`)
	for i, s := range rubrik.DaftarSumbu {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(`"` + s.Kode + `":{"skor":2,"alasan":"a"}`)
	}
	b.WriteString("}}")
	return b.String()
}

func TestMuatDariMenolak(t *testing.T) {
	uji := []struct {
		nama  string
		fs    fstest.MapFS
		pesan string
	}{
		{"pasangan tidak ada", fstest.MapFS{
			"d/a.json": {Data: []byte(kasusJSON("a", "b"))},
		}, "tidak ada"},
		{"tidak saling menunjuk", fstest.MapFS{
			"d/a.json": {Data: []byte(kasusJSON("a", "b"))},
			"d/b.json": {Data: []byte(kasusJSON("b", "c"))},
			"d/c.json": {Data: []byte(kasusJSON("c", "b"))},
		}, "tidak saling menunjuk"},
		{"id beda dari nama berkas", fstest.MapFS{
			"d/x.json": {Data: []byte(kasusJSON("a", ""))},
		}, "nama berkas"},
		{"kasus tidak sah", fstest.MapFS{
			"d/a.json": {Data: []byte(`{"id":"a"}`)},
		}, "tidak sah"},
	}
	for _, u := range uji {
		t.Run(u.nama, func(t *testing.T) {
			_, err := MuatDari(u.fs, "d")
			if err == nil || !strings.Contains(err.Error(), u.pesan) {
				t.Fatalf("ingin galat memuat %q, dapat %v", u.pesan, err)
			}
		})
	}
}

func TestMuatDariMengabaikanBerkasLain(t *testing.T) {
	fs := fstest.MapFS{
		"d/a.json":   {Data: []byte(kasusJSON("a", "b"))},
		"d/b.json":   {Data: []byte(kasusJSON("b", "a"))},
		"d/CATATAN":  {Data: []byte("bukan kasus")},
		"d/sub/c.md": {Data: []byte("bukan kasus")},
	}
	daftar, err := MuatDari(fs, "d")
	if err != nil || len(daftar) != 2 {
		t.Fatalf("dapat %d kasus, galat %v", len(daftar), err)
	}
}

// Pasangan penguji harus benar-benar menantang: sisi perangkat menaikkan
// setidaknya satu sumbu keterlibatan. Kalau tidak, labelnya menyesatkan.
func TestPasanganPengujiBenarMenantang(t *testing.T) {
	daftar, _ := Semua()
	n := 0
	for _, p := range Pasangan(daftar) {
		if !p[0].Penguji && !p[1].Penguji {
			continue
		}
		if p[0].Penguji != p[1].Penguji {
			t.Errorf("pasangan %s/%s: tanda penguji harus di kedua sisi", p[0].ID, p[1].ID)
		}
		n++
		b := rubrik.Bandingkan(p[0], p[1], nil)
		naik := false
		for _, s := range b.Lainnya {
			if s.Sumbu.Kelompok == rubrik.Keterlibatan && s.Beda > 0 {
				naik = true
			}
		}
		if !naik {
			t.Errorf("pasangan penguji %s/%s tidak menaikkan sumbu keterlibatan apa pun", p[0].ID, p[1].ID)
		}
	}
	if n != 2 {
		t.Errorf("jumlah pasangan penguji = %d, ingin 2", n)
	}
}
