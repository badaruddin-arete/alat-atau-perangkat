# Alat atau Perangkat: menimbang benda fokal dengan rubrik terbuka

![tests](https://github.com/badaruddin-arete/alat-atau-perangkat/actions/workflows/tests.yml/badge.svg)

Kalau sebuah alat membuat hidup lebih mudah, apa yang ikut hilang? Dan bisakah kehilangan itu dinilai secara terbuka tanpa berpura-pura objektif?

Program baris perintah dalam Go ini menilai dua belas praktik sehari-hari fiktif, misalnya kopi tubruk di teras dan kopi kapsul di meja kantor, pada delapan sumbu. Hasilnya adalah posisi tiap praktik di antara **benda fokal** dan **perangkat**, daftar "yang didapat" dan "yang tergeser" untuk setiap pasangan, dan uji seberapa tahan kesimpulan itu bila bobot atau skor dipersoalkan. Setiap angka disertai alasan tertulis.

Laporan lengkap: [docs/laporan.md](docs/laporan.md).

## Gagasan

Albert Borgmann, dalam *Technology and the Character of Contemporary Life* (1984), membedakan dua cara hidup bersama benda:

- **Benda fokal** menuntut keterlibatan. Tungku kayu harus dinyalakan dan dijaga, tetapi ia juga mengumpulkan orang dan membuat musim serta bahan bakar terasa.
- **Perangkat** menyajikan hasilnya, misalnya panas, dengan cara yang serba tersedia, sambil menyembunyikan mesin di baliknya. Kita mendapat hasilnya tanpa harus terlibat dengan dunianya.

Borgmann tidak menolak perangkat. Pertanyaannya adalah apa yang tergeser ketika pola perangkat menjadi cara utama kita menjalani hidup. Karena itu rubrik ini tidak memberi skor "baik" atau "buruk", melainkan memisahkan dua hal yang biasanya tercampur: apa yang didapat dari kemudahan dan apa yang dibayar untuknya.

## Sumbu rubrik

Setiap sumbu diberi skor 0 sampai 4, dan setiap skor wajib disertai alasan.

| Kelompok | Kode | Pertanyaan |
|---|---|---|
| Keterlibatan | `tubuh` | Seberapa banyak tangan, indra, dan tenaga yang dituntut? |
| Keterlibatan | `keterampilan` | Seberapa banyak kecakapan yang perlu dilatih dan bisa tumbuh? |
| Keterlibatan | `keterlihatan` | Seberapa terlihat bahan, asal-usul, dan cara kerjanya bagi pelaku? |
| Keterlibatan | `kebersamaan` | Seberapa jauh praktik ini mengumpulkan orang atau membentuk kebiasaan bersama? |
| Ketersediaan | `instan` | Seberapa cepat hasilnya tersedia sejak diinginkan? |
| Ketersediaan | `di_mana_saja` | Seberapa lepas hasilnya dari tempat, musim, dan keadaan tertentu? |
| Ketersediaan | `aman` | Seberapa kecil risiko gagal, kotor, atau celaka bagi pemakai? |
| Ketersediaan | `mudah` | Seberapa sedikit usaha dan perhatian yang diperlukan? |

Empat sumbu ketersediaan mengikuti empat ciri yang dipakai Borgmann untuk menggambarkan komoditas yang "tersedia": *instantaneous, ubiquitous, safe, and easy*, yaitu instan, di mana saja, aman, dan mudah. Frasa ini dikutip Brey (1997) dari halaman 41 buku Borgmann; lihat Rujukan. Empat sumbu keterlibatan adalah terjemahan penyusun atas gambaran Borgmann tentang benda fokal, bukan daftar dari bukunya.

## Kasus fiktif

Enam pasangan, semuanya dengan tokoh dan tempat rekaan. Dua di antaranya adalah **pasangan penguji**: ditulis belakangan, setelah empat pasangan pertama terbukti terlalu mendukung rubrik, dengan sisi perangkat yang justru bisa menambah kebersamaan.

| Sisi fokal | Sisi perangkat | Jenis |
|---|---|---|
| Kopi tubruk di teras | Kopi kapsul di meja kantor | biasa |
| Tungku kayu di dapur | Pemanas ruangan listrik | biasa |
| Surat tulisan tangan | Pesan instan di grup keluarga | biasa |
| Kebun sayur di pekarangan | Belanja sayur lewat aplikasi | biasa |
| Menanak nasi di dandang | Penanak nasi listrik | penguji |
| Kartu pos dari perantauan | Panggilan video keluarga | penguji |

Sebuah uji memastikan setiap pasangan penguji memang menaikkan sedikitnya satu sumbu keterlibatan di sisi perangkat, supaya labelnya tidak menyesatkan.

## Hasil

### Empat pasangan biasa

Dengan bobot bawaan, keempat sisi fokal jatuh antara −0,81 dan −0,38, dan keempat sisi perangkat antara +0,75 dan +0,88. Urutan keempatnya bertahan di semua 6561 kombinasi bobot (setiap sumbu diberi bobot 0,5, 1, atau 2). Titik baliknya, yaitu jumlah perubahan skor satu poin yang cukup untuk membalik urutan, adalah 19 sampai 27 poin. Padahal satu pasangan hanya punya 16 skor.

Hasil setegas ini **tidak** membuktikan bahwa rubriknya kuat. Dari 32 perbandingan sumbu dalam empat pasangan itu, tidak satu pun bergerak berlawanan dengan dugaan. Itu tanda bahwa kasus-kasusnya ditulis oleh orang yang sama yang merancang sumbunya.

### Dua pasangan penguji

Pasangan penguji memperlihatkan apa yang disembunyikan empat pasangan pertama:

- **Kebersamaan bisa naik.** Penanak nasi listrik membebaskan pemasak untuk ikut makan bersama (kebersamaan 1 → 2). Panggilan video mengumpulkan seluruh keluarga setiap Minggu malam (1 → 3). Rubrik melaporkan keduanya di baris "berlawanan dugaan", tidak menyembunyikannya.
- **Sisi fokal tidak selalu fokal.** Kartu pos jatuh di +0,06, "di tengah". Tidak semua cara lama menuntut keterlibatan.
- **Dari sisi keterlibatan saja, kartu pos dan panggilan video seri** (rata-rata keduanya 1,25). Selisih indeks seluruhnya datang dari sisi ketersediaan.
- **Urutan tetap bertahan**, tetapi jauh lebih tipis. Titik baliknya 8 poin (kartu pos) dan 13 poin (nasi), dibanding 19 sampai 27 poin pada pasangan biasa. Golongan penanak nasi listrik hanya tetap di 52,6% kombinasi bobot.

### Temuan utama

Rubrik ini punya bias bawaan. Empat sumbu ketersediaan hampir selalu menguntungkan sisi perangkat, karena memang begitulah perangkat didefinisikan. Akibatnya, kenaikan kebersamaan di satu atau dua sumbu jarang cukup untuk membalik indeks. Dengan kata lain, **indeks tunggal mengukur "seberapa perangkat", bukan "apa yang hilang"**. Untuk pertanyaan kedua, yang perlu dibaca adalah daftar "didapat", "tergeser", dan "berlawanan dugaan" di tiap pasangan, bukan angkanya. Karena itu laporan selalu menampilkan ketiganya.

### Uji lain

- **Buang satu sumbu.** Tidak ada satu pasangan pun yang urutannya bergantung pada satu sumbu saja.
- **Median, bukan rata-rata.** Skor 0 sampai 4 sebenarnya peringkat. Dengan median, urutan keenam pasangan tetap sama.
- **Penilai lain.** Mekanismenya sudah ada (`go run . antarpenilai`, lihat [penilaian/PETUNJUK.md](penilaian/PETUNJUK.md)), tetapi belum ada lembar dari penilai lain.

Semua angka di atas dihitung ulang oleh program di [docs/laporan.md](docs/laporan.md).

## Keputusan desain

| Keputusan | Alasan |
|---|---|
| Ditulis dalam Go, hanya pustaka standar | Pengujian sudah bawaan (`go test`), tanpa ketergantungan luar yang perlu dipercaya. |
| Skala 0 sampai 4 | Skala pendek memaksa setiap angka bisa dibela dengan satu kalimat. Skala halus seperti 0 sampai 100 memberi kesan presisi palsu. |
| Skor tanpa alasan ditolak | Penilaian harus bisa dibaca dan dibantah, bukan sekadar angka. |
| Bidang JSON yang tidak dikenal ditolak | Tidak ada bobot atau penyesuaian tersembunyi yang bisa diselipkan ke data kasus. |
| Indeks = (rata ketersediaan − rata keterlibatan) / 4 | Satu angka dari −1 (benda fokal) sampai +1 (perangkat) yang mudah dibaca, sementara kedua rata-rata tetap ditampilkan. |
| Ambang golongan −0,5, −0,15, 0,15, 0,5 | Pilihan penyusun untuk memberi nama kasar, bukan batas alami. Angka indeks selalu ditampilkan di sampingnya. |
| Kasus selalu berpasangan | Borgmann membaca perangkat sebagai pengganti sesuatu. Perbandingan pasangan menjawab "apa yang didapat dan apa yang tergeser", bukan sekadar peringkat. |
| Sumbu yang bergerak berlawanan arah dilaporkan terpisah | Supaya rubrik tidak menyembunyikan hal yang tidak sesuai dugaan. |
| Uji kepekaan bobot menyeluruh (6561 kombinasi), bukan sampel acak | Jumlahnya kecil, jadi semua kombinasi bisa dihitung dan hasilnya selalu sama. |
| Titik balik dihitung dalam poin skor | Lebih mudah dipahami daripada selisih indeks: "19 ketidaksetujuan satu poin" bisa dibayangkan. |
| `docs/laporan.md` dibuat oleh program dan diuji | Uji gagal bila laporan tidak sama persis dengan keluaran program, jadi angkanya tidak pernah basi. |
| Pasangan penguji ditambahkan dan ditandai, bukan menggantikan yang lama | Pembaca bisa melihat bagaimana hasil berubah ketika kasus sengaja dibuat sulit. Riwayat commit menunjukkan urutannya. |
| Uji buang satu sumbu dan indeks median | Memeriksa apakah kesimpulan bergantung pada satu sumbu atau pada anggapan bahwa skor adalah ukuran selang. |
| Lembar penilai tidak memuat skor penyusun | Supaya penilai lain tidak terpancing angka yang sudah ada. |
| Penilai memakai nama samaran di nama berkas | Tidak ada data pribadi yang disimpan di repo. |
| Kasus ditanam ke program (`go:embed`) | Program berjalan dari satu berkas tanpa perlu menyalin data. |
| Dua kelompok sumbu, bukan satu skor "baik" | Rubrik menggambarkan posisi, bukan menjatuhkan vonis bahwa praktik tertentu lebih baik. |

## Menjalankan

Butuh Go 1.23 atau lebih baru.

```bash
go run . daftar                # semua kasus, urut dari paling fokal
go run . nilai kopi-tubruk     # rincian skor dan alasan
go run . banding kopi-tubruk   # apa yang didapat dan apa yang tergeser
go run . kepekaan              # uji bobot dan titik balik
go run . laporan               # laporan lengkap dalam Markdown
go run . lembar                # lembar isian kosong untuk penilai lain
go run . antarpenilai          # bandingkan penilai lain dengan penyusun
go run . sumbu                 # delapan sumbu rubrik
go test ./...                  # jalankan semua uji
```

Contoh keluaran `go run . daftar`:

```
ID                 INDEKS  GOLONGAN             SPEKTRUM
tungku-kayu         -0.81  benda fokal          fokal [--o------------------] perangkat
kebun-pekarangan    -0.69  benda fokal          fokal [---o-----------------] perangkat
kopi-tubruk         -0.44  condong benda fokal  fokal [------o--------------] perangkat
surat-tangan        -0.38  condong benda fokal  fokal [------o--------------] perangkat
nasi-dandang        -0.31  condong benda fokal  fokal [-------o-------------] perangkat
kartu-pos           +0.06  di tengah            fokal [-----------o---------] perangkat
nasi-listrik        +0.50  perangkat            fokal [---------------o-----] perangkat
panggilan-video     +0.56  perangkat            fokal [----------------o----] perangkat
sayur-aplikasi      +0.75  perangkat            fokal [------------------o--] perangkat
pesan-instan        +0.81  perangkat            fokal [------------------o--] perangkat
kopi-kapsul         +0.88  perangkat            fokal [-------------------o-] perangkat
pemanas-listrik     +0.88  perangkat            fokal [-------------------o-] perangkat
```

## Susunan kode

| Paket | Isi |
|---|---|
| `rubrik` | Sumbu, bentuk data kasus dan validasinya, indeks, perbandingan pasangan, uji kepekaan, titik balik, buang satu sumbu, indeks median |
| `kasus` | Dua belas kasus JSON yang ditanam ke program, pemeriksaan pasangan |
| `penilai` | Lembar isian, pembacaan skor penilai lain, perbandingan dengan penyusun |
| `laporan` | Penyusun `docs/laporan.md` |
| `main` | Perintah baris perintah |

Uji berjalan di GitHub Actions pada Go 1.23 dan 1.24, didahului pemeriksaan format (`gofmt`) dan mutu (`go vet`).

## Keterbatasan

- **Ironi rubrik.** Mengubah keterlibatan menjadi angka adalah langkah yang sangat "perangkat": ia membuat penilaian cepat dan mudah dibaca, sambil menyembunyikan pengalaman di balik angka. Rubrik ini membuat penilaian terbuka dan dapat dibantah, tetapi tidak membuatnya objektif.
- **Penyusun kasus sama dengan perancang rubrik.** Pasangan penguji mengurangi masalah ini, tetapi tidak menghapusnya, karena pasangan penguji pun ditulis oleh penyusun yang sama. Mekanisme penilai lain sudah tersedia, tetapi belum ada yang mengisi.
- **Indeks tunggal condong ke perangkat secara bawaan.** Lihat Temuan utama. Indeks berguna untuk menempatkan kasus, bukan untuk menimbang apa yang hilang.
- **Uji kepekaan hanya menguji bobot dan skor, bukan sumbunya.** Kalau sumbu yang dipilih keliru, misalnya ada aspek penting yang tidak tercakup, kedua uji tidak akan menangkapnya.
- **Skala dianggap selang.** Rata-rata skor 0 sampai 4 mengandaikan jarak 1 ke 2 sama dengan 3 ke 4. Indeks median disertakan sebagai pembanding, dan untuk kasus-kasus ini urutannya sama.
- **Kasus berlatar Indonesia, rubrik dari pemikir Amerika.** Contoh Borgmann sendiri (perapian, makan bersama keluarga) berasal dari kehidupan Amerika. Pemindahannya ke tungku dan kopi tubruk adalah tafsiran penyusun.
- **Rujukan belum dicocokkan langsung ke buku.** Frasa empat ciri ketersediaan dan halaman 41 diambil dari kutipan Brey (1997) dalam jurnal ilmiah, bukan dari buku Borgmann sendiri. Gambaran benda fokal disusun dari pemahaman atas buku dan pembahasan sekunder, jadi bagian Gagasan sebaiknya dibaca sebagai ringkasan.
- Semua kasus fiktif dan tidak menggambarkan produk, merek, atau orang tertentu. Program ini bukan alat untuk menilai kehidupan orang sungguhan.

## Rujukan

- Albert Borgmann. 1984. *Technology and the Character of Contemporary Life: A Philosophical Inquiry*. Chicago: University of Chicago Press. (Belum dicocokkan langsung; lihat Keterbatasan.)
- Philip Brey. 1997. "New Media and the Quality of Life." *Society for Philosophy and Technology* 3 (1). https://scholar.lib.vt.edu/ejournals/SPT/v3n1/brey.html. Sumber kutipan "instantaneous, ubiquitous, safe, and easy" (Borgmann 1984, 41).

## Transparansi bantuan AI

Kode, uji, kasus fiktif, dan teks disusun bersama Claude (Anthropic) dalam sesi Claude Code. Gagasan, pilihan bahasa, peta karya, dan keputusan akhir ditinjau dan disetujui oleh penyusun. Setiap commit yang dibantu memuat baris `Co-Authored-By`.

## Lisensi

Hak cipta © 2026 badaruddin-arete. All rights reserved. Kode boleh dibaca dan dipelajari, tetapi tidak boleh disalin, diubah, atau disebarkan tanpa izin tertulis. Lihat [LICENSE](LICENSE).
