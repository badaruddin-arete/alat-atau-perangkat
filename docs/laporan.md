# Laporan perbandingan

Dibuat otomatis dengan `go run . laporan > docs/laporan.md`. Jangan disunting dengan tangan; uji akan gagal bila laporan ini tidak sama dengan keluaran program.

Semua kasus fiktif. Skor 0 sampai 4, bobot bawaan 1 untuk setiap sumbu. Indeks −1 berarti benda fokal penuh, +1 perangkat penuh.

## Spektrum

| Kasus | Keterlibatan | Ketersediaan | Indeks | Golongan |
|---|---:|---:|---:|---|
| Tungku kayu di dapur | 3,75 | 0,50 | -0,81 | benda fokal |
| Kebun sayur di pekarangan | 3,50 | 0,75 | -0,69 | benda fokal |
| Kopi tubruk di teras | 3,25 | 1,50 | -0,44 | condong benda fokal |
| Surat tulisan tangan | 2,75 | 1,25 | -0,38 | condong benda fokal |
| Belanja sayur lewat aplikasi | 0,50 | 3,50 | +0,75 | perangkat |
| Pesan instan di grup keluarga | 0,75 | 4,00 | +0,81 | perangkat |
| Kopi kapsul di meja kantor | 0,25 | 3,75 | +0,88 | perangkat |
| Pemanas ruangan listrik | 0,00 | 3,50 | +0,88 | perangkat |

## Pasangan: apa yang didapat, apa yang tergeser

### Kebun sayur di pekarangan → Belanja sayur lewat aplikasi

Indeks -0,69 → +0,75. Titik balik: 23 perubahan skor satu poin.

| Arah | Sumbu | Skor | Alasan di sisi perangkat |
|---|---|---|---|
| didapat | Di mana saja | 0 → 4 | Sayur apa pun tersedia kapan saja, lepas dari musim dan tempat. |
| didapat | Instan | 0 → 3 | Sampai dalam sekitar satu jam. |
| didapat | Mudah | 1 → 4 | Hampir tanpa usaha. |
| didapat | Aman | 2 → 3 | Risiko kecil: sayur kadang layu atau tertukar. |
| tergeser | Keterlibatan tubuh | 4 → 0 | Hanya menggulir dan mengetuk layar; kurir yang berjalan. |
| tergeser | Keterlihatan cara kerja | 4 → 1 | Ada label asal kebun di aplikasi, tetapi rantai dari kebun sampai pintu tidak terlihat. |
| tergeser | Kebersamaan dan ritual | 3 → 0 | Dipesan dan diterima sendiri; tidak ada kebiasaan bersama yang terbentuk. |
| tergeser | Keterampilan | 3 → 1 | Memilih sayur dari foto butuh sedikit pengalaman, tetapi tidak ada yang perlu dilatih. |

### Kopi tubruk di teras → Kopi kapsul di meja kantor

Indeks -0,44 → +0,88. Titik balik: 21 perubahan skor satu poin.

| Arah | Sumbu | Skor | Alasan di sisi perangkat |
|---|---|---|---|
| didapat | Instan | 1 → 4 | Cangkir terisi dalam sekitar tiga puluh detik. |
| didapat | Mudah | 1 → 4 | Hampir tanpa usaha dan bisa dilakukan sambil memikirkan hal lain. |
| didapat | Aman | 2 → 4 | Air panas tertutup di dalam mesin; hampir tidak ada risiko gagal atau kotor. |
| didapat | Di mana saja | 2 → 3 | Ada di banyak kantor dan rumah, tetapi butuh listrik dan kapsul yang cocok dengan mesinnya. |
| tergeser | Keterlibatan tubuh | 3 → 0 | Hanya memasukkan kapsul dan menekan satu tombol. |
| tergeser | Keterampilan | 3 → 0 | Tidak ada kecakapan yang perlu dilatih; hasilnya sama siapa pun yang menekan. |
| tergeser | Keterlihatan cara kerja | 3 → 0 | Biji, gilingan, dan suhu tersembunyi di dalam kapsul dan mesin; asal kopi tidak diketahui pemakai. |
| tergeser | Kebersamaan dan ritual | 4 → 1 | Kadang ada obrolan singkat di dekat mesin, tetapi praktiknya dirancang untuk sendiri dan cepat. |

### Tungku kayu di dapur → Pemanas ruangan listrik

Indeks -0,81 → +0,88. Titik balik: 27 perubahan skor satu poin.

| Arah | Sumbu | Skor | Alasan di sisi perangkat |
|---|---|---|---|
| didapat | Instan | 0 → 4 | Kamar hangat dalam beberapa menit. |
| didapat | Mudah | 0 → 4 | Cukup sekali putar, lalu dilupakan. |
| didapat | Di mana saja | 1 → 3 | Bisa di kamar mana pun, asal ada listrik yang cukup. |
| didapat | Aman | 1 → 3 | Tanpa asap dan api terbuka, meski tetap ada risiko korsleting bila kabel rusak. |
| tergeser | Keterlibatan tubuh | 4 → 0 | Hanya memutar tombol atau menekan remote. |
| tergeser | Keterlihatan cara kerja | 4 → 0 | Listrik, pembangkit, dan elemen pemanas tidak terlihat; panas datang tanpa tanda asal-usulnya. |
| tergeser | Kebersamaan dan ritual | 4 → 0 | Setiap kamar punya pemanasnya sendiri sehingga tidak ada alasan untuk berkumpul. |
| tergeser | Keterampilan | 3 → 0 | Tidak ada kecakapan yang perlu dilatih selain mengatur suhu. |

### Surat tulisan tangan → Pesan instan di grup keluarga

Indeks -0,38 → +0,81. Titik balik: 19 perubahan skor satu poin.

| Arah | Sumbu | Skor | Alasan di sisi perangkat |
|---|---|---|---|
| didapat | Instan | 0 → 4 | Terkirim dan bisa dibalas dalam hitungan detik. |
| didapat | Di mana saja | 1 → 4 | Bisa dari mana saja selama ada sinyal. |
| didapat | Mudah | 1 → 4 | Hampir tanpa usaha dan bisa dilakukan sambil mengerjakan hal lain. |
| didapat | Aman | 3 → 4 | Tidak ada risiko fisik; pesan jarang hilang. |
| tergeser | Keterlihatan cara kerja | 3 → 0 | Jaringan, server, dan cara pesan sampai sepenuhnya tersembunyi. |
| tergeser | Keterlibatan tubuh | 2 → 0 | Hanya mengetik atau menekan tombol rekam di layar. |
| tergeser | Keterampilan | 3 → 1 | Ada sedikit kecakapan memilih kata dan waktu, tetapi praktiknya tidak menuntut latihan. |
| tergeser | Kebersamaan dan ritual | 3 → 2 | Grup keluarga mengumpulkan banyak orang sekaligus; kebersamaan ini nyata, meski tipis dan terputus-putus. |

## Kepekaan terhadap bobot

Setiap sumbu diberi bobot 0,5, 1, atau 2, sehingga ada 6561 kombinasi bobot.

| Pasangan | Urutan bertahan | Selisih indeks terkecil |
|---|---:|---:|
| Kebun sayur di pekarangan → Belanja sayur lewat aplikasi | 100,0% | +1,14 |
| Kopi tubruk di teras → Kopi kapsul di meja kantor | 100,0% | +1,18 |
| Tungku kayu di dapur → Pemanas ruangan listrik | 100,0% | +1,46 |
| Surat tulisan tangan → Pesan instan di grup keluarga | 100,0% | +0,89 |

| Kasus | Indeks bawaan | Rentang indeks | Golongan tetap | Sebaran golongan |
|---|---:|---|---:|---|
| Tungku kayu di dapur | -0,81 | -0,93 sampai -0,66 | 100,0% | benda fokal 100,0% |
| Kebun sayur di pekarangan | -0,69 | -0,88 sampai -0,48 | 99,9% | benda fokal 99,9%; condong benda fokal 0,1% |
| Kopi tubruk di teras | -0,44 | -0,59 sampai -0,32 | 90,2% | benda fokal 9,8%; condong benda fokal 90,2% |
| Surat tulisan tangan | -0,38 | -0,55 sampai -0,11 | 94,8% | benda fokal 5,1%; condong benda fokal 94,8%; di tengah 0,1% |
| Belanja sayur lewat aplikasi | +0,75 | +0,60 sampai +0,90 | 100,0% | perangkat 100,0% |
| Pesan instan di grup keluarga | +0,81 | +0,68 sampai +0,93 | 100,0% | perangkat 100,0% |
| Kopi kapsul di meja kantor | +0,88 | +0,71 sampai +0,96 | 100,0% | perangkat 100,0% |
| Pemanas ruangan listrik | +0,88 | +0,80 sampai +0,95 | 100,0% | perangkat 100,0% |
