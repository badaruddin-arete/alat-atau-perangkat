# Alat atau Perangkat: menimbang benda fokal dengan rubrik terbuka

![tests](https://github.com/badaruddin-arete/alat-atau-perangkat/actions/workflows/tests.yml/badge.svg)

Kalau sebuah alat membuat hidup lebih mudah, apa yang ikut hilang? Dan bisakah kehilangan itu dinilai secara terbuka tanpa berpura-pura objektif?

Program baris perintah dalam Go ini menilai delapan praktik sehari-hari fiktif, misalnya kopi tubruk di teras dan kopi kapsul di meja kantor, pada delapan sumbu. Hasilnya adalah posisi tiap praktik di antara **benda fokal** dan **perangkat**, daftar "yang didapat" dan "yang tergeser" untuk setiap pasangan, dan uji seberapa tahan kesimpulan itu bila bobot atau skor dipersoalkan. Setiap angka disertai alasan tertulis.

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

Empat sumbu ketersediaan mengikuti empat ciri yang dipakai Borgmann untuk menggambarkan komoditas yang "tersedia": instan, di mana saja (*ubiquitous*), aman, dan mudah. Empat sumbu keterlibatan adalah terjemahan penyusun atas gambaran Borgmann tentang benda fokal, bukan daftar dari bukunya.

## Kasus fiktif

Empat pasangan, semuanya dengan tokoh dan tempat rekaan.

| Sisi fokal | Sisi perangkat |
|---|---|
| Kopi tubruk di teras | Kopi kapsul di meja kantor |
| Tungku kayu di dapur | Pemanas ruangan listrik |
| Surat tulisan tangan | Pesan instan di grup keluarga |
| Kebun sayur di pekarangan | Belanja sayur lewat aplikasi |

## Hasil

Dengan bobot bawaan, keempat sisi fokal jatuh antara −0,81 dan −0,38, dan keempat sisi perangkat antara +0,75 dan +0,88. Dua uji kepekaan dijalankan:

- **Bobot.** Setiap sumbu diberi bobot 0,5, 1, atau 2, sehingga ada 6561 kombinasi. Urutan keempat pasangan bertahan di semua kombinasi. Yang goyah adalah nama golongan untuk kasus di tengah: kopi tubruk berpindah dari "condong benda fokal" ke "benda fokal" pada 9,8% kombinasi, dan surat tulisan tangan pada 5,2% kombinasi.
- **Skor (titik balik).** Berapa perubahan skor satu poin yang cukup untuk membalik urutan sebuah pasangan? Jawabannya 19 poin (surat dan pesan instan) sampai 27 poin (tungku dan pemanas). Satu pasangan hanya punya 16 skor, jadi rata-rata setiap skor harus digeser lebih dari satu poin ke arah yang berlawanan sebelum kesimpulannya berubah.

### Bacaan yang jujur

Hasil yang setegas ini **tidak** membuktikan bahwa rubriknya kuat. Ia terutama memperlihatkan bahwa kasus-kasusnya ditulis oleh orang yang sama yang merancang sumbunya. Tanda paling jelas: dari 32 perbandingan sumbu dalam empat pasangan, tidak satu pun bergerak berlawanan dengan dugaan. Program sanggup melaporkan arah yang berlawanan (dan diuji untuk itu), tetapi data ini tidak memberinya kesempatan. Pesan instan sempat mendekati dengan skor kebersamaan 2, karena grup keluarga memang mengumpulkan orang.

Jadi temuan utamanya bukan "perangkat mengurangi keterlibatan", sebab itu sudah dimasukkan lewat cara kasus ditulis. Temuannya adalah bahwa rubrik terbuka mengubah perdebatan. Daripada bertanya "apakah kopi kapsul buruk?", pembaca bisa menunjuk satu baris, misalnya *keterlihatan 3 → 0*, lalu membantah alasannya. Titik balik memberi ukuran berapa banyak bantahan seperti itu yang diperlukan sebelum kesimpulan berubah.

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
sayur-aplikasi      +0.75  perangkat            fokal [------------------o--] perangkat
pesan-instan        +0.81  perangkat            fokal [------------------o--] perangkat
kopi-kapsul         +0.88  perangkat            fokal [-------------------o-] perangkat
pemanas-listrik     +0.88  perangkat            fokal [-------------------o-] perangkat
```

## Susunan kode

| Paket | Isi |
|---|---|
| `rubrik` | Sumbu, bentuk data kasus dan validasinya, indeks, perbandingan pasangan, uji kepekaan, titik balik |
| `kasus` | Delapan kasus JSON yang ditanam ke program, pemeriksaan pasangan |
| `laporan` | Penyusun `docs/laporan.md` |
| `main` | Perintah baris perintah |

Uji berjalan di GitHub Actions pada Go 1.23 dan 1.24, didahului pemeriksaan format (`gofmt`) dan mutu (`go vet`).

## Keterbatasan

- **Ironi rubrik.** Mengubah keterlibatan menjadi angka adalah langkah yang sangat "perangkat": ia membuat penilaian cepat dan mudah dibaca, sambil menyembunyikan pengalaman di balik angka. Rubrik ini membuat penilaian terbuka dan dapat dibantah, tetapi tidak membuatnya objektif.
- **Penyusun kasus sama dengan perancang rubrik.** Seperti dijelaskan di bagian Hasil, ketegasan hasil sebagian besar berasal dari cara kasus ditulis. Langkah yang lebih jujur adalah meminta beberapa orang lain menilai kasus yang sama lalu membandingkan skornya. Ini belum dilakukan.
- **Uji kepekaan hanya menguji bobot dan skor, bukan sumbunya.** Kalau sumbu yang dipilih keliru, misalnya ada aspek penting yang tidak tercakup, kedua uji tidak akan menangkapnya.
- **Skala dianggap selang.** Rata-rata skor 0 sampai 4 mengandaikan jarak 1 ke 2 sama dengan 3 ke 4. Itu penyederhanaan.
- **Kasus berlatar Indonesia, rubrik dari pemikir Amerika.** Contoh Borgmann sendiri (perapian, makan bersama keluarga) berasal dari kehidupan Amerika. Pemindahannya ke tungku dan kopi tubruk adalah tafsiran penyusun.
- **Rujukan belum dicocokkan halaman demi halaman.** Empat ciri ketersediaan dan gambaran benda fokal disusun dari pemahaman atas buku Borgmann dan pembahasan sekunder. Kutipan langsung dan nomor halamannya belum diperiksa ke teks buku, jadi bagian Gagasan sebaiknya dibaca sebagai ringkasan, bukan kutipan.
- Semua kasus fiktif dan tidak menggambarkan produk, merek, atau orang tertentu. Program ini bukan alat untuk menilai kehidupan orang sungguhan.

## Rujukan

- Albert Borgmann. 1984. *Technology and the Character of Contemporary Life: A Philosophical Inquiry*. Chicago: University of Chicago Press. (Belum dicocokkan halaman demi halaman; lihat Keterbatasan.)

## Transparansi bantuan AI

Kode, uji, kasus fiktif, dan teks disusun bersama Claude (Anthropic) dalam sesi Claude Code. Gagasan, pilihan bahasa, peta karya, dan keputusan akhir ditinjau dan disetujui oleh penyusun. Setiap commit yang dibantu memuat baris `Co-Authored-By`.

## Lisensi

Hak cipta © 2026 badaruddin-arete. All rights reserved. Kode boleh dibaca dan dipelajari, tetapi tidak boleh disalin, diubah, atau disebarkan tanpa izin tertulis. Lihat [LICENSE](LICENSE).
