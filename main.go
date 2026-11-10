// Program alat-atau-perangkat menilai praktik sehari-hari fiktif di antara
// benda fokal dan perangkat. Tahap ini baru menampilkan sumbu rubrik.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

func main() {
	os.Exit(jalankan(os.Args[1:], os.Stdout, os.Stderr))
}

// jalankan dipisah dari main supaya perintah bisa diuji tanpa keluar dari proses.
func jalankan(arg []string, keluar, galat io.Writer) int {
	if len(arg) == 0 || arg[0] == "bantuan" {
		fmt.Fprintln(keluar, "Pemakaian: alat-atau-perangkat <perintah>")
		fmt.Fprintln(keluar, "  sumbu     tampilkan delapan sumbu rubrik")
		fmt.Fprintln(keluar, "  bantuan   tampilkan pesan ini")
		return 0
	}
	switch arg[0] {
	case "sumbu":
		cetakSumbu(keluar)
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
