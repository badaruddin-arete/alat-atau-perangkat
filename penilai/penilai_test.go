package penilai

import (
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/badaruddin-arete/alat-atau-perangkat/kasus"
	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

func bawaan(t *testing.T) []rubrik.Kasus {
	t.Helper()
	d, err := kasus.Semua()
	if err != nil {
		t.Fatal(err)
	}
	return d
}

// isi membuat lembar CSV untuk kasus yang diberikan dengan skor dari fungsi f.
func isi(ids []string, f func(id, sumbu string) int) string {
	var b strings.Builder
	b.WriteString(strings.Join(Kolom, ",") + "\n")
	for _, id := range ids {
		for _, s := range rubrik.DaftarSumbu {
			b.WriteString(id + "," + s.Kode + ",," + string(rune('0'+f(id, s.Kode))) + ",alasan penilai\n")
		}
	}
	return b.String()
}

func salinSkor(daftar []rubrik.Kasus) func(id, sumbu string) int {
	return func(id, sumbu string) int {
		k, _ := kasus.Cari(daftar, id)
		return k.Penilaian[sumbu].Skor
	}
}

func TestLembarKosongTanpaSkorPenyusun(t *testing.T) {
	daftar := bawaan(t)
	var b bytes.Buffer
	if err := TulisLembar(&b, daftar); err != nil {
		t.Fatal(err)
	}
	teks := b.String()
	baris, err := csv.NewReader(strings.NewReader(teks)).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(baris) != 1+len(daftar)*len(rubrik.DaftarSumbu) {
		t.Fatalf("jumlah baris = %d", len(baris))
	}
	for _, r := range baris[1:] {
		if r[3] != "" || r[4] != "" {
			t.Fatalf("lembar kosong memuat skor atau alasan: %v", r)
		}
	}
	// Lembar kosong harus terbaca tanpa galat dan tanpa isian.
	l, err := BacaLembar(strings.NewReader(teks))
	if err != nil || len(l) != 0 {
		t.Fatalf("lembar kosong: %v, %v", l, err)
	}
}

func TestBacaLembarMenolak(t *testing.T) {
	kepala := strings.Join(Kolom, ",") + "\n"
	uji := map[string]string{
		"kepala salah":     "a,b,c,d,e\n",
		"sumbu asing":      kepala + "kopi-tubruk,cepat,,2,x\n",
		"skor bukan angka": kepala + "kopi-tubruk,tubuh,,dua,x\n",
		"skor di luar":     kepala + "kopi-tubruk,tubuh,,5,x\n",
		"alasan kosong":    kepala + "kopi-tubruk,tubuh,,2,\n",
		"ganda":            kepala + "kopi-tubruk,tubuh,,2,x\nkopi-tubruk,tubuh,,3,y\n",
	}
	for nama, isi := range uji {
		if _, err := BacaLembar(strings.NewReader(isi)); err == nil {
			t.Errorf("%s: harus ditolak", nama)
		}
	}
}

func TestTerapkanMenolakKasusTidakLengkap(t *testing.T) {
	kepala := strings.Join(Kolom, ",") + "\n"
	l, err := BacaLembar(strings.NewReader(kepala + "kopi-tubruk,tubuh,,2,x\n"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Terapkan(l, bawaan(t)); err == nil {
		t.Fatal("kasus yang hanya diisi satu sumbu harus ditolak")
	}
	l2 := Lembar{"teh-celup": {}}
	if _, err := Terapkan(l2, bawaan(t)); err == nil {
		t.Fatal("kasus tidak dikenal harus ditolak")
	}
}

func TestPenilaiYangSetujuDanYangMembalik(t *testing.T) {
	daftar := bawaan(t)
	ids := []string{"kopi-tubruk", "kopi-kapsul", "surat-tangan"}
	fs := fstest.MapFS{
		"p/lembar-kosong.csv": {Data: []byte(strings.Join(Kolom, ",") + "\n")},
		"p/setuju.csv":        {Data: []byte(isi(ids, salinSkor(daftar)))},
		// Penilai yang membalik segalanya: skor = 4 - skor penyusun.
		"p/membalik.csv": {Data: []byte(isi(ids, func(id, s string) int { return 4 - salinSkor(daftar)(id, s) }))},
		"p/catatan.txt":  {Data: []byte("bukan lembar")},
	}
	semua, err := MuatSemua(fs, "p", daftar)
	if err != nil {
		t.Fatal(err)
	}
	if len(semua) != 2 || semua[0].Nama != "membalik" || semua[1].Nama != "setuju" {
		t.Fatalf("penilai terbaca: %+v", semua)
	}
	r := Bandingkan(semua, daftar, kasus.Pasangan(daftar))
	balik, setuju := r[0], r[1]
	if setuju.SelisihRata != 0 || setuju.JauhBerbeda != 0 || setuju.PasanganSama != 1 || setuju.PasanganDinilai != 1 {
		t.Errorf("penilai setuju: %+v", setuju)
	}
	if balik.Kasus != 3 || balik.Sumbu != 24 || balik.PasanganSama != 0 || balik.SelisihRata == 0 {
		t.Errorf("penilai membalik: %+v", balik)
	}
}

func TestMuatSemuaMelaporkanBerkasRusak(t *testing.T) {
	fs := fstest.MapFS{"p/rusak.csv": {Data: []byte("x\n")}}
	_, err := MuatSemua(fs, "p", bawaan(t))
	if err == nil || !strings.Contains(err.Error(), "rusak.csv") {
		t.Fatalf("galat harus menyebut berkasnya, dapat %v", err)
	}
}
