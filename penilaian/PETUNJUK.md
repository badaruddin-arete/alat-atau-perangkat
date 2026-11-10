# Petunjuk untuk penilai

Folder ini dipakai untuk menguji celah terbesar karya ini: semua kasus dinilai oleh orang yang juga merancang rubriknya. Penilai lain diminta menilai kasus fiktif yang sama tanpa melihat skor penyusun.

## Cara mengisi

1. Baca [kisah.md](kisah.md). Semua tokoh dan tempat di dalamnya fiktif.
2. Salin [lembar-kosong.csv](lembar-kosong.csv) menjadi berkas baru dengan **nama samaran**, misalnya `penilai-a.csv`. Jangan memakai nama asli.
3. Isi kolom `skor` (bilangan 0 sampai 4) dan `alasan` (satu kalimat) untuk setiap sumbu. Kolom `pertanyaan` hanya bantuan dan tidak perlu diubah.
4. Boleh mengisi sebagian kasus saja, asal setiap kasus yang diisi lengkap delapan sumbu. Baris yang skornya kosong dilewati.
5. Jangan membuka `kasus/data/` atau `docs/laporan.md` sebelum selesai mengisi, supaya tidak terpancing skor penyusun.

Lembar bisa dibuka dengan aplikasi lembar kerja apa pun, lalu disimpan kembali sebagai CSV.

## Membandingkan

```bash
go run . antarpenilai
```

Program membaca semua berkas `.csv` di folder ini kecuali `lembar-kosong.csv`, lalu melaporkan untuk setiap penilai: rata-rata selisih skor dengan penyusun, berapa sumbu yang berbeda 2 poin atau lebih, dan berapa pasangan yang urutannya sama.

Lembar ini tidak meminta data pribadi apa pun. Yang dinilai adalah kasus fiktif, bukan pengisinya.
