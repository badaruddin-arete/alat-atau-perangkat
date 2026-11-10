package laporan

import (
	"bytes"
	"strings"
	"testing"

	"github.com/badaruddin-arete/alat-atau-perangkat/kasus"
)

func tulisBawaan(t *testing.T) string {
	t.Helper()
	daftar, err := kasus.Semua()
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	Tulis(&b, daftar, kasus.Pasangan(daftar))
	return b.String()
}

func TestLaporanMemuatSemuaBagian(t *testing.T) {
	isi := tulisBawaan(t)
	for _, kata := range []string{"## Spektrum", "## Pasangan", "## Kepekaan terhadap bobot", "6561 kombinasi", "Titik balik", "## Kepekaan terhadap sumbu", "(penguji)", "berlawanan dugaan"} {
		if !strings.Contains(isi, kata) {
			t.Errorf("laporan tidak memuat %q", kata)
		}
	}
	if n := strings.Count(isi, "\n### "); n != 6 {
		t.Errorf("jumlah bagian pasangan = %d, ingin 6", n)
	}
}

func TestLaporanPasti(t *testing.T) {
	if tulisBawaan(t) != tulisBawaan(t) {
		t.Fatal("laporan harus sama setiap kali dibuat")
	}
}

func TestAngkaMemakaiKomaDesimal(t *testing.T) {
	if got := angka(-0.4375); got != "-0,44" {
		t.Errorf("angka = %q", got)
	}
	if got := persen(1, 3); got != "33,3%" {
		t.Errorf("persen = %q", got)
	}
	if got := sel("a|b"); got != `a\|b` {
		t.Errorf("sel = %q", got)
	}
}
