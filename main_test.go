package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTanpaArgumenMenampilkanBantuan(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan(nil, &keluar, &galat); kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	if !strings.Contains(keluar.String(), "Pemakaian") {
		t.Errorf("bantuan tidak tampil: %q", keluar.String())
	}
}

func TestPerintahSumbu(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan([]string{"sumbu"}, &keluar, &galat); kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	for _, kata := range []string{"[keterlibatan]", "[ketersediaan]", "tubuh", "mudah"} {
		if !strings.Contains(keluar.String(), kata) {
			t.Errorf("keluaran tidak memuat %q", kata)
		}
	}
}

func TestPerintahTidakDikenal(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan([]string{"terbang"}, &keluar, &galat); kode != 2 {
		t.Fatalf("kode keluar = %d, ingin 2", kode)
	}
	if !strings.Contains(galat.String(), "tidak dikenal") {
		t.Errorf("pesan galat keliru: %q", galat.String())
	}
}

func jalankanUji(t *testing.T, arg ...string) (string, string, int) {
	t.Helper()
	var keluar, galat bytes.Buffer
	kode := jalankan(arg, &keluar, &galat)
	return keluar.String(), galat.String(), kode
}

func TestDaftarUrutDariFokal(t *testing.T) {
	keluar, _, kode := jalankanUji(t, "daftar")
	if kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	baris := strings.Split(strings.TrimSpace(keluar), "\n")
	if len(baris) != 13 {
		t.Fatalf("ingin 1 judul + 12 kasus, dapat %d baris", len(baris))
	}
	if !strings.HasPrefix(baris[1], "tungku-kayu") {
		t.Errorf("kasus paling fokal seharusnya tungku-kayu, dapat %q", baris[1])
	}
}

func TestNilaiMenampilkanAlasan(t *testing.T) {
	keluar, _, kode := jalankanUji(t, "nilai", "kopi-tubruk")
	if kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	for _, kata := range []string{"Kopi tubruk di teras", "gilingan tangan", "Indeks", "condong benda fokal"} {
		if !strings.Contains(keluar, kata) {
			t.Errorf("keluaran tidak memuat %q", kata)
		}
	}
}

func TestBandingDariSisiManaPun(t *testing.T) {
	a, _, _ := jalankanUji(t, "banding", "kopi-tubruk")
	b, _, _ := jalankanUji(t, "banding", "kopi-kapsul")
	if a != b {
		t.Error("banding harus sama dari kedua sisi pasangan")
	}
	for _, kata := range []string{"Yang didapat", "Yang tergeser", "Instan 1 -> 4"} {
		if !strings.Contains(a, kata) {
			t.Errorf("keluaran tidak memuat %q", kata)
		}
	}
}

func TestNilaiDanBandingMenolakMasukanSalah(t *testing.T) {
	uji := [][]string{{"nilai"}, {"banding"}, {"nilai", "teh-celup"}, {"banding", "teh-celup"}}
	for _, arg := range uji {
		_, galat, kode := jalankanUji(t, arg...)
		if kode != 2 || galat == "" {
			t.Errorf("%v: kode %d, galat %q", arg, kode, galat)
		}
	}
}

func TestGarisSpektrum(t *testing.T) {
	uji := map[float64]string{
		-1: "fokal [o--------------------] perangkat",
		0:  "fokal [----------o----------] perangkat",
		1:  "fokal [--------------------o] perangkat",
	}
	for i, ingin := range uji {
		if got := garis(i); got != ingin {
			t.Errorf("garis(%v) = %q", i, got)
		}
	}
}

// Laporan di docs/ harus sama persis dengan keluaran program,
// supaya angka di README dan laporan tidak pernah basi.
func TestLaporanTersimpanMutakhir(t *testing.T) {
	keluar, _, kode := jalankanUji(t, "laporan")
	if kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	tersimpan, err := os.ReadFile("docs/laporan.md")
	if err != nil {
		t.Fatal(err)
	}
	if string(tersimpan) != keluar {
		t.Fatal("docs/laporan.md basi; jalankan: go run . laporan > docs/laporan.md")
	}
}

func TestPerintahKepekaan(t *testing.T) {
	keluar, _, kode := jalankanUji(t, "kepekaan")
	if kode != 0 || !strings.Contains(keluar, "6561 kombinasi") || !strings.Contains(keluar, "poin") {
		t.Fatalf("kode %d, keluaran %q", kode, keluar)
	}
}

func TestLembarDanKisahTersimpanMutakhir(t *testing.T) {
	for perintah, berkas := range map[string]string{"lembar": "penilaian/lembar-kosong.csv", "kisah": "penilaian/kisah.md"} {
		keluar, _, kode := jalankanUji(t, perintah)
		if kode != 0 {
			t.Fatalf("%s: kode keluar = %d", perintah, kode)
		}
		tersimpan, err := os.ReadFile(berkas)
		if err != nil {
			t.Fatal(err)
		}
		if string(tersimpan) != keluar {
			t.Errorf("%s basi; jalankan: go run . %s > %s", berkas, perintah, berkas)
		}
	}
}

func TestAntarpenilai(t *testing.T) {
	keluar, _, kode := jalankanUji(t, "antarpenilai", t.TempDir())
	if kode != 0 || !strings.Contains(keluar, "Belum ada lembar") {
		t.Fatalf("folder kosong: kode %d, %q", kode, keluar)
	}
	_, galat, kode := jalankanUji(t, "antarpenilai", "folder-yang-tidak-ada")
	if kode != 1 || galat == "" {
		t.Fatalf("folder tidak ada: kode %d, %q", kode, galat)
	}
	// Lembar kosong yang tersimpan tidak dihitung sebagai penilai.
	keluar, _, kode = jalankanUji(t, "antarpenilai")
	if kode != 0 || !strings.Contains(keluar, "Belum ada lembar") {
		t.Fatalf("folder penilaian: kode %d, %q", kode, keluar)
	}
}
