package kasus

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

func TestDelapanKasusBawaanSah(t *testing.T) {
	daftar, err := Semua()
	if err != nil {
		t.Fatal(err)
	}
	if len(daftar) != 8 {
		t.Fatalf("jumlah kasus = %d, ingin 8", len(daftar))
	}
	if got := len(Pasangan(daftar)); got != 4 {
		t.Errorf("jumlah pasangan = %d, ingin 4", got)
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
