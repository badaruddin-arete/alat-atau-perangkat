// Paket laporan menyusun laporan perbandingan dalam Markdown dari kasus bawaan.
// Laporan hanya memuat angka dan alasan dari data; tafsirannya ada di README.
package laporan

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

// urutanGolongan menjaga kolom golongan tampil dari fokal ke perangkat.
var urutanGolongan = []string{"benda fokal", "condong benda fokal", "di tengah", "condong perangkat", "perangkat"}

func persen(a, b int) string {
	return strings.Replace(fmt.Sprintf("%.1f%%", 100*float64(a)/float64(b)), ".", ",", 1)
}

func angka(x float64) string {
	return strings.Replace(fmt.Sprintf("%+.2f", x), ".", ",", 1)
}

func rata(x float64) string {
	return strings.Replace(fmt.Sprintf("%.2f", x), ".", ",", 1)
}

func sel(s string) string { return strings.ReplaceAll(s, "|", "\\|") }

// Tulis menulis laporan lengkap. Keluarannya pasti: data yang sama selalu
// menghasilkan teks yang sama, sehingga docs/laporan.md dapat diuji.
func Tulis(w io.Writer, daftar []rubrik.Kasus, pasangan [][2]rubrik.Kasus) {
	fmt.Fprintln(w, "# Laporan perbandingan")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Dibuat otomatis dengan `go run . laporan > docs/laporan.md`. Jangan disunting dengan tangan; uji akan gagal bila laporan ini tidak sama dengan keluaran program.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Semua kasus fiktif. Skor 0 sampai 4, bobot bawaan 1 untuk setiap sumbu. Indeks −1 berarti benda fokal penuh, +1 perangkat penuh.")

	hasil := make([]rubrik.Hasil, len(daftar))
	for i, k := range daftar {
		hasil[i] = rubrik.Nilai(k, nil)
	}
	sort.SliceStable(hasil, func(i, j int) bool { return hasil[i].Indeks < hasil[j].Indeks })

	fmt.Fprintln(w, "\n## Spektrum")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Kasus | Keterlibatan | Ketersediaan | Indeks | Golongan | Indeks median |")
	fmt.Fprintln(w, "|---|---:|---:|---:|---|---:|")
	for _, h := range hasil {
		nama := sel(h.Kasus.Nama)
		if h.Kasus.Penguji {
			nama += " *"
		}
		fmt.Fprintf(w, "| %s | %s | %s | %s | %s | %s |\n", nama, rata(h.Keterlibatan), rata(h.Ketersediaan), angka(h.Indeks), h.Golongan, angka(rubrik.IndeksMedian(h.Kasus)))
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Tanda * menandai kasus penguji: pasangan yang sengaja ditulis untuk menantang rubrik, karena sisi perangkatnya bisa menambah kebersamaan.")

	fmt.Fprintln(w, "\n## Pasangan: apa yang didapat, apa yang tergeser")
	for _, p := range pasangan {
		b := rubrik.Bandingkan(p[0], p[1], nil)
		judul := b.Fokal.Kasus.Nama + " → " + b.Lawan.Kasus.Nama
		if b.Fokal.Kasus.Penguji || b.Lawan.Kasus.Penguji {
			judul += " (penguji)"
		}
		fmt.Fprintf(w, "\n### %s\n\n", judul)
		fmt.Fprintf(w, "Indeks %s → %s. Titik balik: %d perubahan skor satu poin.\n\n", angka(b.Fokal.Indeks), angka(b.Lawan.Indeks), rubrik.TitikBalik(b.Fokal.Kasus, b.Lawan.Kasus))
		fmt.Fprintln(w, "| Arah | Sumbu | Skor | Alasan di sisi perangkat |")
		fmt.Fprintln(w, "|---|---|---|---|")
		baris := func(arah string, xs []rubrik.Selisih) {
			for _, s := range xs {
				fmt.Fprintf(w, "| %s | %s | %d → %d | %s |\n", arah, s.Sumbu.Nama, s.Fokal.Skor, s.Lawan.Skor, sel(s.Lawan.Alasan))
			}
		}
		baris("didapat", b.Didapat)
		baris("tergeser", b.Tergeser)
		var lain []rubrik.Selisih
		for _, s := range b.Lainnya {
			if s.Beda != 0 {
				lain = append(lain, s)
			}
		}
		baris("berlawanan dugaan", lain)
	}

	rentang, kp := rubrik.Kepekaan(daftar, pasangan)
	fmt.Fprintln(w, "\n## Kepekaan terhadap bobot")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Setiap sumbu diberi bobot 0,5, 1, atau 2, sehingga ada %d kombinasi bobot.\n\n", kp[0].JumlahBobot)
	fmt.Fprintln(w, "| Pasangan | Urutan bertahan | Selisih indeks terkecil |")
	fmt.Fprintln(w, "|---|---:|---:|")
	for _, p := range kp {
		fmt.Fprintf(w, "| %s → %s | %s | %s |\n", sel(p.Fokal.Nama), sel(p.Lawan.Nama), persen(p.Bertahan, p.JumlahBobot), angka(p.SelisihMin))
	}
	sort.SliceStable(rentang, func(i, j int) bool { return rentang[i].Bawaan < rentang[j].Bawaan })
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Kasus | Indeks bawaan | Rentang indeks | Golongan tetap | Sebaran golongan |")
	fmt.Fprintln(w, "|---|---:|---|---:|---|")
	for _, r := range rentang {
		var sebar []string
		for _, g := range urutanGolongan {
			if n := r.Golongan[g]; n > 0 {
				sebar = append(sebar, fmt.Sprintf("%s %s", g, persen(n, r.JumlahBobot)))
			}
		}
		fmt.Fprintf(w, "| %s | %s | %s sampai %s | %s | %s |\n", sel(r.Kasus.Nama), angka(r.Bawaan), angka(r.Min), angka(r.Maks), persen(r.GolonganTetap, r.JumlahBobot), strings.Join(sebar, "; "))
	}

	fmt.Fprintln(w, "\n## Kepekaan terhadap sumbu dan cara merangkum")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Sumbu penentu: sumbu yang bila dibuang sendirian membuat urutan pasangan hilang. Median: indeks dihitung dengan median, bukan rata-rata, karena skor 0 sampai 4 sebenarnya peringkat.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Pasangan | Sumbu penentu | Indeks median | Urutan dengan median |")
	fmt.Fprintln(w, "|---|---|---|---|")
	for _, p := range kp {
		penentu := "tidak ada"
		if s := rubrik.SumbuPenentu(p.Fokal, p.Lawan); len(s) > 0 {
			penentu = strings.Join(s, ", ")
		}
		mf, ml := rubrik.IndeksMedian(p.Fokal), rubrik.IndeksMedian(p.Lawan)
		urut := "bertahan"
		if ml <= mf {
			urut = "hilang"
		}
		fmt.Fprintf(w, "| %s → %s | %s | %s → %s | %s |\n", sel(p.Fokal.Nama), sel(p.Lawan.Nama), penentu, angka(mf), angka(ml), urut)
	}
}
