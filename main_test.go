package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestTanpaArgumenMenampilkanBantuan(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan(nil, &keluar, &galat); kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	if !strings.Contains(keluar.String(), "Pemakaian") {
		t.Errorf("bantuan tidak tampil: %q", keluar.String())
	}
}

func TestPerintahSumbu(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan([]string{"sumbu"}, &keluar, &galat); kode != 0 {
		t.Fatalf("kode keluar = %d", kode)
	}
	for _, kata := range []string{"[keterlibatan]", "[ketersediaan]", "tubuh", "mudah"} {
		if !strings.Contains(keluar.String(), kata) {
			t.Errorf("keluaran tidak memuat %q", kata)
		}
	}
}

func TestPerintahTidakDikenal(t *testing.T) {
	var keluar, galat bytes.Buffer
	if kode := jalankan([]string{"terbang"}, &keluar, &galat); kode != 2 {
		t.Fatalf("kode keluar = %d, ingin 2", kode)
	}
	if !strings.Contains(galat.String(), "tidak dikenal") {
		t.Errorf("pesan galat keliru: %q", galat.String())
	}
}
