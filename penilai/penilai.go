// Paket penilai membaca skor dari penilai lain dan membandingkannya dengan
// skor penyusun. Tujuannya menguji celah terbesar karya ini: semua kasus
// dinilai oleh orang yang juga merancang rubriknya.
package penilai

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"math"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

// Kolom lembar isian. Kolom pertanyaan hanya bantuan bagi pengisi dan diabaikan saat dibaca.
var Kolom = []string{"kasus", "sumbu", "pertanyaan", "skor", "alasan"}

// NamaLembarKosong adalah berkas contoh yang tidak dihitung sebagai penilai.
const NamaLembarKosong = "lembar-kosong.csv"

// TulisLembar menulis lembar isian kosong untuk semua kasus. Skor penyusun
// sengaja tidak disertakan supaya pengisi tidak terpancing angka yang sudah ada.
func TulisLembar(w io.Writer, daftar []rubrik.Kasus) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(Kolom); err != nil {
		return err
	}
	for _, k := range daftar {
		for _, s := range rubrik.DaftarSumbu {
			if err := cw.Write([]string{k.ID, s.Kode, s.Pertanyaan, "", ""}); err != nil {
				return err
			}
		}
	}
	cw.Flush()
	return cw.Error()
}

// Lembar adalah isian satu penilai: id kasus -> kode sumbu -> penilaian.
type Lembar map[string]map[string]rubrik.Penilaian

// BacaLembar membaca satu lembar. Baris dengan skor kosong dilewati, sehingga
// penilai boleh mengisi sebagian kasus saja; tetapi setiap kasus yang diisi
// harus lengkap delapan sumbu, diperiksa oleh Terapkan.
func BacaLembar(r io.Reader) (Lembar, error) {
	cr := csv.NewReader(r)
	baris, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(baris) == 0 || strings.Join(baris[0], ",") != strings.Join(Kolom, ",") {
		return nil, fmt.Errorf("kepala kolom harus %s", strings.Join(Kolom, ","))
	}
	l := Lembar{}
	for i, b := range baris[1:] {
		no := i + 2
		id, kode, teks, alasan := strings.TrimSpace(b[0]), strings.TrimSpace(b[1]), strings.TrimSpace(b[3]), strings.TrimSpace(b[4])
		if teks == "" {
			continue
		}
		if _, ada := rubrik.CariSumbu(kode); !ada {
			return nil, fmt.Errorf("baris %d: sumbu tidak dikenal %q", no, kode)
		}
		skor, err := strconv.Atoi(teks)
		if err != nil || skor < rubrik.SkorMin || skor > rubrik.SkorMaks {
			return nil, fmt.Errorf("baris %d: skor %q harus bilangan %d..%d", no, teks, rubrik.SkorMin, rubrik.SkorMaks)
		}
		if alasan == "" {
			return nil, fmt.Errorf("baris %d: alasan kosong", no)
		}
		if l[id] == nil {
			l[id] = map[string]rubrik.Penilaian{}
		}
		if _, ganda := l[id][kode]; ganda {
			return nil, fmt.Errorf("baris %d: %s/%s diisi dua kali", no, id, kode)
		}
		l[id][kode] = rubrik.Penilaian{Skor: skor, Alasan: alasan}
	}
	return l, nil
}

// Terapkan mengganti skor kasus penyusun dengan skor penilai untuk kasus yang
// diisi penilai. Hasilnya hanya memuat kasus yang diisi, dan setiap kasus divalidasi.
func Terapkan(l Lembar, daftar []rubrik.Kasus) (map[string]rubrik.Kasus, error) {
	hasil := map[string]rubrik.Kasus{}
	dikenal := map[string]rubrik.Kasus{}
	for _, k := range daftar {
		dikenal[k.ID] = k
	}
	for id, isian := range l {
		asli, ada := dikenal[id]
		if !ada {
			return nil, fmt.Errorf("kasus tidak dikenal %q", id)
		}
		k := asli
		k.Penilaian = isian
		if err := k.Validasi(); err != nil {
			return nil, err
		}
		hasil[id] = k
	}
	return hasil, nil
}

// Penilai adalah satu penilai beserta kasus yang ia nilai.
type Penilai struct {
	Nama  string
	Kasus map[string]rubrik.Kasus
}

// MuatSemua membaca semua lembar .csv di dir, kecuali lembar kosong.
// Nama penilai diambil dari nama berkas; pakai nama samaran, bukan nama asli.
func MuatSemua(fsys fs.FS, dir string, daftar []rubrik.Kasus) ([]Penilai, error) {
	entri, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	var hasil []Penilai
	for _, e := range entri {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".csv") || e.Name() == NamaLembarKosong {
			continue
		}
		f, err := fsys.Open(path.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		l, err := BacaLembar(f)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		ks, err := Terapkan(l, daftar)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		if len(ks) > 0 {
			hasil = append(hasil, Penilai{Nama: strings.TrimSuffix(e.Name(), ".csv"), Kasus: ks})
		}
	}
	sort.Slice(hasil, func(i, j int) bool { return hasil[i].Nama < hasil[j].Nama })
	return hasil, nil
}

// Ringkasan membandingkan satu penilai dengan penyusun.
type Ringkasan struct {
	Nama            string
	Kasus           int     // jumlah kasus yang dinilai
	SelisihRata     float64 // rata-rata |skor penilai - skor penyusun| per sumbu
	JauhBerbeda     int     // sumbu yang selisihnya 2 poin atau lebih
	Sumbu           int     // jumlah sumbu yang dibandingkan
	PasanganSama    int     // pasangan yang urutannya sama dengan penyusun
	PasanganDinilai int     // pasangan yang kedua sisinya dinilai
}

// Bandingkan meringkas setiap penilai terhadap penyusun.
func Bandingkan(semua []Penilai, daftar []rubrik.Kasus, pasangan [][2]rubrik.Kasus) []Ringkasan {
	asli := map[string]rubrik.Kasus{}
	for _, k := range daftar {
		asli[k.ID] = k
	}
	var hasil []Ringkasan
	for _, p := range semua {
		r := Ringkasan{Nama: p.Nama, Kasus: len(p.Kasus)}
		total := 0.0
		for id, k := range p.Kasus {
			for _, s := range rubrik.DaftarSumbu {
				d := math.Abs(float64(k.Penilaian[s.Kode].Skor - asli[id].Penilaian[s.Kode].Skor))
				total += d
				r.Sumbu++
				if d >= 2 {
					r.JauhBerbeda++
				}
			}
		}
		if r.Sumbu > 0 {
			r.SelisihRata = total / float64(r.Sumbu)
		}
		for _, ps := range pasangan {
			a, okA := p.Kasus[ps[0].ID]
			b, okB := p.Kasus[ps[1].ID]
			if !okA || !okB {
				continue
			}
			r.PasanganDinilai++
			arahPenyusun := rubrik.Indeks(ps[1], nil) > rubrik.Indeks(ps[0], nil)
			arahPenilai := rubrik.Indeks(b, nil) > rubrik.Indeks(a, nil)
			if arahPenyusun == arahPenilai {
				r.PasanganSama++
			}
		}
		hasil = append(hasil, r)
	}
	return hasil
}

// TulisKisah menulis kisah semua kasus untuk dibaca pengisi lembar,
// tanpa skor dan tanpa golongan supaya tidak memengaruhi penilaian.
func TulisKisah(w io.Writer, daftar []rubrik.Kasus) {
	fmt.Fprintln(w, "# Kisah kasus untuk penilai")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Dibuat otomatis dengan `go run . kisah > penilaian/kisah.md`. Semua tokoh dan tempat fiktif.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Baca setiap kisah, lalu isi `lembar-kosong.csv`: skor 0 sampai 4 dan satu kalimat alasan untuk setiap sumbu. Boleh mengisi sebagian kasus saja, asal setiap kasus yang diisi lengkap delapan sumbu.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "## Sumbu")
	fmt.Fprintln(w)
	for _, s := range rubrik.DaftarSumbu {
		fmt.Fprintf(w, "- `%s`: %s\n", s.Kode, s.Pertanyaan)
	}
	for _, k := range daftar {
		fmt.Fprintf(w, "\n## %s (`%s`)\n\n%s\n", k.Nama, k.ID, k.Kisah)
	}
}
