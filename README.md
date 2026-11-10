# Alat atau Perangkat: menimbang benda fokal dengan rubrik terbuka

![tests](https://github.com/badaruddin-arete/alat-atau-perangkat/actions/workflows/tests.yml/badge.svg)

Kalau sebuah alat membuat hidup lebih mudah, apa yang ikut hilang? Dan bisakah kehilangan itu dinilai secara terbuka tanpa berpura-pura objektif?

Program baris perintah dalam Go ini menilai praktik sehari-hari fiktif, misalnya kopi tubruk dan kopi kapsul, pada delapan sumbu. Hasilnya adalah posisi tiap praktik di antara **benda fokal** dan **perangkat**, dengan alasan tertulis untuk setiap angka.

> Status: dalam pengerjaan. Mesin rubrik dan delapan kasus fiktif sudah berjalan; uji kepekaan bobot dan laporan perbandingan menyusul.

## Gagasan

Albert Borgmann, dalam *Technology and the Character of Contemporary Life* (1984), membedakan dua cara hidup bersama benda:

- **Benda fokal** menuntut keterlibatan. Tungku kayu harus dinyalakan dan dijaga, tetapi ia juga mengumpulkan orang dan membuat musim serta bahan bakar terasa.
- **Perangkat** menyajikan hasilnya, misalnya panas, dengan cara yang serba tersedia, sambil menyembunyikan mesin di baliknya. Kita mendapat hasilnya tanpa harus terlibat dengan dunianya.

Borgmann tidak menyebut perangkat buruk dengan sendirinya. Pertanyaannya adalah apa yang tergeser ketika pola perangkat menjadi cara utama kita menjalani hidup.

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

Empat sumbu ketersediaan mengikuti empat ciri yang dipakai Borgmann untuk menggambarkan komoditas yang "tersedia": instan, di mana saja, aman, dan mudah. Kutipan dan halamannya belum dicocokkan ke teks buku dan akan ditandai di sini setelah diperiksa.

## Keputusan desain

| Keputusan | Alasan |
|---|---|
| Ditulis dalam Go, hanya pustaka standar | Pengujian sudah bawaan (`go test`), tanpa ketergantungan luar yang perlu dipercaya. |
| Skala 0 sampai 4 | Skala pendek memaksa setiap angka bisa dibela dengan satu kalimat. Skala halus seperti 0 sampai 100 memberi kesan presisi palsu. |
| Skor tanpa alasan ditolak | Penilaian harus bisa dibaca dan dibantah, bukan sekadar angka. |
| Bidang JSON yang tidak dikenal ditolak | Tidak ada bobot atau penyesuaian tersembunyi yang bisa diselipkan ke data kasus. |
| Validasi melaporkan semua masalah sekaligus | Penyusun kasus memperbaiki semuanya dalam satu kali baca. |
| Indeks = (rata ketersediaan − rata keterlibatan) / 4 | Satu angka dari −1 (benda fokal) sampai +1 (perangkat) yang mudah dibaca, sementara kedua rata-rata tetap ditampilkan. |
| Ambang golongan −0,5, −0,15, 0,15, 0,5 | Pilihan penyusun untuk memberi nama kasar, bukan batas alami. Angka indeks selalu ditampilkan di sampingnya. |
| Kasus selalu berpasangan | Borgmann membaca perangkat sebagai pengganti sesuatu. Perbandingan pasangan menjawab "apa yang didapat dan apa yang tergeser", bukan sekadar peringkat. |
| Sumbu yang bergerak berlawanan arah dilaporkan terpisah | Supaya rubrik tidak menyembunyikan hal yang tidak sesuai dugaan. |
| Kasus ditanam ke program (`go:embed`) | Program berjalan dari satu berkas tanpa perlu menyalin data. |
| Dua kelompok sumbu, bukan satu skor "baik" | Rubrik menggambarkan posisi, bukan menjatuhkan vonis bahwa praktik tertentu lebih baik. |

## Menjalankan

```bash
go run . daftar                # semua kasus, urut dari paling fokal
go run . nilai kopi-tubruk     # rincian skor dan alasan
go run . banding kopi-tubruk   # apa yang didapat dan apa yang tergeser
go run . sumbu                 # delapan sumbu rubrik
go test ./...                  # jalankan semua uji
```

## Keterbatasan

- **Ironi rubrik.** Mengubah keterlibatan menjadi angka adalah langkah yang sangat "perangkat". Rubrik ini membuat penilaian terbuka dan dapat dibantah, tetapi tidak membuatnya objektif.
- Skor dan alasan disusun oleh penyusun karya ini, bukan hasil survei atau pengamatan.
- Semua kasus fiktif dan tidak menggambarkan produk, merek, atau orang tertentu.

## Transparansi bantuan AI

Kode, uji, dan teks disusun bersama Claude (Anthropic) dalam sesi Claude Code. Gagasan, pilihan kasus, dan keputusan akhir ditinjau oleh penyusun. Setiap commit yang dibantu memuat baris `Co-Authored-By`.

## Lisensi

Hak cipta © 2026 badaruddin-arete. All rights reserved. Kode boleh dibaca dan dipelajari, tetapi tidak boleh disalin, diubah, atau disebarkan tanpa izin tertulis. Lihat [LICENSE](LICENSE).
