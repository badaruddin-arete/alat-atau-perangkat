package rubrik

import "testing"

func TestSemuaBobotMenghasilkan6561KombinasiUnik(t *testing.T) {
	dilihat := map[[8]float64]bool{}
	SemuaBobot(func(b Bobot) {
		var kunci [8]float64
		for i, s := range DaftarSumbu {
			kunci[i] = b[s.Kode]
		}
		dilihat[kunci] = true
		if err := b.Validasi(); err != nil {
			t.Fatal(err)
		}
	})
	if len(dilihat) != 6561 {
		t.Fatalf("kombinasi unik = %d, ingin 6561", len(dilihat))
	}
}

func TestKepekaanKasusSeragamTidakBergerak(t *testing.T) {
	// Skor seragam per kelompok: bobot apa pun memberi rata-rata yang sama.
	a, b := kasusDua("a", 4, 0), kasusDua("b", 0, 4)
	rentang, kp := Kepekaan([]Kasus{a, b}, [][2]Kasus{{b, a}})
	for _, r := range rentang {
		if r.Min != r.Maks || r.GolonganTetap != r.JumlahBobot || r.JumlahBobot != 6561 {
			t.Errorf("%s: rentang [%v, %v], tetap %d", r.Kasus.ID, r.Min, r.Maks, r.GolonganTetap)
		}
	}
	if kp[0].Fokal.ID != "a" || kp[0].Bertahan != 6561 || !dekat(kp[0].SelisihMin, 2) {
		t.Errorf("pasangan: %+v", kp[0])
	}
}

func TestKepekaanMenangkapUrutanYangGoyah(t *testing.T) {
	a, b := kasusDua("a", 2, 2), kasusDua("b", 2, 2)
	a.Penilaian["tubuh"] = Penilaian{3, "uji"}  // a sedikit lebih fokal
	b.Penilaian["mudah"] = Penilaian{1, "uji"}  // tetapi b kurang mudah
	b.Penilaian["instan"] = Penilaian{4, "uji"} // dan lebih instan
	_, kp := Kepekaan([]Kasus{a, b}, [][2]Kasus{{a, b}})
	if kp[0].Bertahan == 0 || kp[0].Bertahan == kp[0].JumlahBobot {
		t.Fatalf("urutan seharusnya bergantung pada bobot, bertahan %d/%d", kp[0].Bertahan, kp[0].JumlahBobot)
	}
}

func TestTitikBalik(t *testing.T) {
	uji := []struct {
		nama         string
		fokal, lawan Kasus
		ingin        int
	}{
		{"sudah terbalik", kasusDua("a", 0, 4), kasusDua("b", 4, 0), 0},
		{"sama", kasusDua("a", 2, 2), kasusDua("b", 2, 2), 0},
		// selisih indeks 0,25 = 4 poin.
		{"selisih kecil", kasusDua("a", 2, 2), kasusDua("b", 1, 2), 4},
		// selisih 2 = 32 poin, ruang tepat 32.
		{"ujung", kasusDua("a", 4, 0), kasusDua("b", 0, 4), 32},
	}
	for _, u := range uji {
		if got := TitikBalik(u.fokal, u.lawan); got != u.ingin {
			t.Errorf("%s: TitikBalik = %d, ingin %d", u.nama, got, u.ingin)
		}
	}
}
