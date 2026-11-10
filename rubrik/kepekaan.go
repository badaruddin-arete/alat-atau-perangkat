package rubrik

import "math"

// PilihanBobot adalah nilai yang dicoba untuk setiap sumbu dalam uji kepekaan:
// separuh, sama, atau dua kali lipat bobot bawaan.
var PilihanBobot = []float64{0.5, 1, 2}

// SemuaBobot menghasilkan setiap kombinasi PilihanBobot untuk delapan sumbu
// (3^8 = 6561 kombinasi) dan memanggil f untuk masing-masing.
func SemuaBobot(f func(Bobot)) {
	n := len(DaftarSumbu)
	idx := make([]int, n)
	for {
		b := make(Bobot, n)
		for i, s := range DaftarSumbu {
			b[s.Kode] = PilihanBobot[idx[i]]
		}
		f(b)
		// Naikkan "angka" idx dalam basis len(PilihanBobot).
		i := 0
		for i < n {
			idx[i]++
			if idx[i] < len(PilihanBobot) {
				break
			}
			idx[i] = 0
			i++
		}
		if i == n {
			return
		}
	}
}

// RentangKasus adalah sebaran indeks dan golongan satu kasus di semua kombinasi bobot.
type RentangKasus struct {
	Kasus         Kasus
	Bawaan        float64
	Min, Maks     float64
	Golongan      map[string]int // golongan -> jumlah kombinasi
	JumlahBobot   int
	GolonganTetap int // kombinasi yang golongannya sama dengan golongan bawaan
}

// KepekaanPasangan mencatat apakah urutan sebuah pasangan bertahan.
type KepekaanPasangan struct {
	Fokal, Lawan Kasus // menurut bobot bawaan
	Bertahan     int   // kombinasi tempat lawan tetap lebih condong ke perangkat
	JumlahBobot  int
	SelisihMin   float64 // selisih indeks terkecil (lawan - fokal)
}

// Kepekaan menjalankan semua kombinasi bobot pada kasus dan pasangan.
func Kepekaan(daftar []Kasus, pasangan [][2]Kasus) ([]RentangKasus, []KepekaanPasangan) {
	rentang := make([]RentangKasus, len(daftar))
	for i, k := range daftar {
		bawaan := Indeks(k, nil)
		rentang[i] = RentangKasus{Kasus: k, Bawaan: bawaan, Min: math.Inf(1), Maks: math.Inf(-1), Golongan: map[string]int{}}
	}
	kp := make([]KepekaanPasangan, len(pasangan))
	for i, p := range pasangan {
		f, l := p[0], p[1]
		if Indeks(l, nil) < Indeks(f, nil) {
			f, l = l, f
		}
		kp[i] = KepekaanPasangan{Fokal: f, Lawan: l, SelisihMin: math.Inf(1)}
	}
	SemuaBobot(func(b Bobot) {
		for i := range rentang {
			r := &rentang[i]
			x := Indeks(r.Kasus, b)
			r.Min, r.Maks = math.Min(r.Min, x), math.Max(r.Maks, x)
			g := Golongan(x)
			r.Golongan[g]++
			if g == Golongan(r.Bawaan) {
				r.GolonganTetap++
			}
			r.JumlahBobot++
		}
		for i := range kp {
			p := &kp[i]
			d := Indeks(p.Lawan, b) - Indeks(p.Fokal, b)
			if d > 0 {
				p.Bertahan++
			}
			p.SelisihMin = math.Min(p.SelisihMin, d)
			p.JumlahBobot++
		}
	})
	return rentang, kp
}

// TitikBalik menghitung berapa perubahan skor sebesar satu poin, dengan bobot
// bawaan, yang paling sedikit diperlukan supaya sisi fokal tidak lagi lebih
// rendah indeksnya daripada lawan. Setiap poin menggeser indeks satu kasus
// sebesar 1/16 (satu dari empat sumbu kelompok, dibagi skor maksimum 4).
// Angka ini menjawab: berapa banyak ketidaksetujuan pembaca, sumbu demi sumbu,
// yang cukup untuk membalik kesimpulan pasangan? Hasil -1 berarti tidak mungkin.
func TitikBalik(fokal, lawan Kasus) int {
	const perPoin = 1.0 / (4 * SkorMaks)
	selisih := Indeks(lawan, nil) - Indeks(fokal, nil)
	if selisih <= 0 {
		return 0
	}
	perlu := int(math.Ceil(selisih/perPoin - 1e-9))
	ruang := 0
	for _, s := range DaftarSumbu {
		f, l := fokal.Penilaian[s.Kode].Skor, lawan.Penilaian[s.Kode].Skor
		if s.Kelompok == Ketersediaan {
			ruang += (SkorMaks - f) + (l - SkorMin) // fokal naik, lawan turun
		} else {
			ruang += (f - SkorMin) + (SkorMaks - l) // fokal turun, lawan naik
		}
	}
	if perlu > ruang {
		return -1
	}
	return perlu
}
