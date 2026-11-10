// Program alat-atau-perangkat menilai praktik sehari-hari fiktif di antara
// benda fokal dan perangkat menurut Albert Borgmann.
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/badaruddin-arete/alat-atau-perangkat/kasus"
	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

func main() {
	os.Exit(jalankan(os.Args[1:], os.Stdout, os.Stderr))
}

const bantuan = `Pemakaian: alat-atau-perangkat <perintah> [id]

  daftar        semua kasus beserta posisinya di spektrum
  nilai <id>    rincian skor dan alasan satu kasus
  banding <id>  bandingkan kasus dengan pasangannya: apa yang didapat, apa yang tergeser
  sumbu         delapan sumbu rubrik
  bantuan       pesan ini
`

// jalankan dipisah dari main supaya perintah bisa diuji tanpa keluar dari proses.
func jalankan(arg []string, keluar, galat io.Writer) int {
	if len(arg) == 0 || arg[0] == "bantuan" {
		fmt.Fprint(keluar, bantuan)
		return 0
	}
	if arg[0] == "sumbu" {
		cetakSumbu(keluar)
		return 0
	}
	daftar, err := kasus.Semua()
	if err != nil {
		fmt.Fprintln(galat, "gagal memuat kasus:", err)
		return 1
	}
	switch arg[0] {
	case "daftar":
		cetakDaftar(keluar, daftar)
		return 0
	case "nilai", "banding":
		if len(arg) < 2 {
			fmt.Fprintf(galat, "perintah %s butuh id kasus; lihat: daftar\n", arg[0])
			return 2
		}
		k, ada := kasus.Cari(daftar, arg[1])
		if !ada {
			fmt.Fprintf(galat, "kasus tidak dikenal: %q; lihat: daftar\n", arg[1])
			return 2
		}
		if arg[0] == "nilai" {
			cetakNilai(keluar, rubrik.Nilai(k, nil))
			return 0
		}
		p, ada := kasus.Cari(daftar, k.Pasangan)
		if !ada {
			fmt.Fprintf(galat, "kasus %q tidak punya pasangan\n", k.ID)
			return 2
		}
		cetakBanding(keluar, rubrik.Bandingkan(k, p, nil))
		return 0
	default:
		fmt.Fprintf(galat, "perintah tidak dikenal: %q\n", arg[0])
		return 2
	}
}

func cetakSumbu(w io.Writer) {
	for _, k := range []rubrik.Kelompok{rubrik.Keterlibatan, rubrik.Ketersediaan} {
		fmt.Fprintf(w, "[%s]\n", k)
		for _, s := range rubrik.SumbuKelompok(k) {
			fmt.Fprintf(w, "  %-13s %s\n", s.Kode, s.Pertanyaan)
		}
	}
}

// garis menggambar posisi indeks pada spektrum selebar 21 kolom.
func garis(indeks float64) string {
	const lebar = 21
	pos := int((indeks+1)/2*(lebar-1) + 0.5)
	return "fokal [" + strings.Repeat("-", pos) + "o" + strings.Repeat("-", lebar-1-pos) + "] perangkat"
}

func cetakDaftar(w io.Writer, daftar []rubrik.Kasus) {
	hasil := make([]rubrik.Hasil, len(daftar))
	for i, k := range daftar {
		hasil[i] = rubrik.Nilai(k, nil)
	}
	// Urut dari paling fokal ke paling perangkat; ID memecah seri.
	sort.SliceStable(hasil, func(i, j int) bool { return hasil[i].Indeks < hasil[j].Indeks })
	fmt.Fprintf(w, "%-17s %7s  %-20s %s\n", "ID", "INDEKS", "GOLONGAN", "SPEKTRUM")
	for _, h := range hasil {
		k := h.Kasus
		fmt.Fprintf(w, "%-17s %+7.2f  %-20s %s\n", k.ID, h.Indeks, h.Golongan, garis(h.Indeks))
	}
}

func cetakNilai(w io.Writer, h rubrik.Hasil) {
	fmt.Fprintf(w, "%s (%s)\n\n%s\n\n", h.Kasus.Nama, h.Kasus.ID, h.Kasus.Kisah)
	for _, kel := range []rubrik.Kelompok{rubrik.Keterlibatan, rubrik.Ketersediaan} {
		fmt.Fprintf(w, "[%s]\n", kel)
		for _, s := range rubrik.SumbuKelompok(kel) {
			p := h.Kasus.Penilaian[s.Kode]
			fmt.Fprintf(w, "  %-13s %d/%d  %s\n", s.Kode, p.Skor, rubrik.SkorMaks, p.Alasan)
		}
	}
	fmt.Fprintf(w, "\nKeterlibatan %.2f, ketersediaan %.2f (skala 0..4)\n", h.Keterlibatan, h.Ketersediaan)
	fmt.Fprintf(w, "Indeks %+.2f: %s\n%s\n", h.Indeks, h.Golongan, garis(h.Indeks))
}

func cetakBanding(w io.Writer, p rubrik.Perbandingan) {
	fmt.Fprintf(w, "%s (%+.2f) menuju %s (%+.2f)\n", p.Fokal.Kasus.Nama, p.Fokal.Indeks, p.Lawan.Kasus.Nama, p.Lawan.Indeks)
	cetakSelisih(w, "\nYang didapat", p.Didapat)
	cetakSelisih(w, "\nYang tergeser", p.Tergeser)
	var bergerak []rubrik.Selisih
	for _, s := range p.Lainnya {
		if s.Beda != 0 {
			bergerak = append(bergerak, s)
		}
	}
	cetakSelisih(w, "\nBergerak berlawanan dengan dugaan", bergerak)
}

func cetakSelisih(w io.Writer, judul string, xs []rubrik.Selisih) {
	if len(xs) == 0 {
		return
	}
	fmt.Fprintf(w, "%s:\n", judul)
	for _, s := range xs {
		fmt.Fprintf(w, "  %s %d -> %d: %s\n", s.Sumbu.Nama, s.Fokal.Skor, s.Lawan.Skor, s.Lawan.Alasan)
	}
}
