// Paket rubrik memuat sumbu penilaian dan bentuk data kasus untuk menimbang
// sebuah praktik di antara "benda fokal" dan "perangkat" menurut Albert Borgmann.
package rubrik

// Kelompok membagi sumbu menjadi dua sisi yang ditarik Borgmann.
type Kelompok string

const (
	// Keterlibatan: seberapa jauh praktik menuntut tubuh, keterampilan,
	// pemahaman, dan kebersamaan. Ciri benda fokal.
	Keterlibatan Kelompok = "keterlibatan"
	// Ketersediaan: seberapa jauh hasilnya disajikan instan, di mana saja,
	// aman, dan mudah. Ciri perangkat.
	Ketersediaan Kelompok = "ketersediaan"
)

// Batas skor tiap sumbu. Skala pendek dipilih supaya setiap angka
// harus dapat dibela dengan satu kalimat alasan.
const (
	SkorMin  = 0
	SkorMaks = 4
)

// Sumbu adalah satu pertanyaan penilaian.
type Sumbu struct {
	Kode       string
	Nama       string
	Kelompok   Kelompok
	Pertanyaan string
}

// DaftarSumbu berisi delapan sumbu, empat di tiap kelompok, dalam urutan tetap.
var DaftarSumbu = []Sumbu{
	{"tubuh", "Keterlibatan tubuh", Keterlibatan,
		"Seberapa banyak tangan, indra, dan tenaga yang dituntut?"},
	{"keterampilan", "Keterampilan", Keterlibatan,
		"Seberapa banyak kecakapan yang perlu dilatih dan bisa tumbuh?"},
	{"keterlihatan", "Keterlihatan cara kerja", Keterlibatan,
		"Seberapa terlihat bahan, asal-usul, dan cara kerjanya bagi pelaku?"},
	{"kebersamaan", "Kebersamaan dan ritual", Keterlibatan,
		"Seberapa jauh praktik ini mengumpulkan orang atau membentuk kebiasaan bersama?"},
	{"instan", "Instan", Ketersediaan,
		"Seberapa cepat hasilnya tersedia sejak diinginkan?"},
	{"di_mana_saja", "Di mana saja", Ketersediaan,
		"Seberapa lepas hasilnya dari tempat, musim, dan keadaan tertentu?"},
	{"aman", "Aman", Ketersediaan,
		"Seberapa kecil risiko gagal, kotor, atau celaka bagi pemakai?"},
	{"mudah", "Mudah", Ketersediaan,
		"Seberapa sedikit usaha dan perhatian yang diperlukan?"},
}

// CariSumbu mengembalikan sumbu dengan kode tertentu.
func CariSumbu(kode string) (Sumbu, bool) {
	for _, s := range DaftarSumbu {
		if s.Kode == kode {
			return s, true
		}
	}
	return Sumbu{}, false
}

// SumbuKelompok mengembalikan sumbu milik satu kelompok, dalam urutan tetap.
func SumbuKelompok(k Kelompok) []Sumbu {
	var hasil []Sumbu
	for _, s := range DaftarSumbu {
		if s.Kelompok == k {
			hasil = append(hasil, s)
		}
	}
	return hasil
}
