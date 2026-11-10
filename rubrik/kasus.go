package rubrik

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Penilaian adalah skor satu sumbu beserta alasannya. Skor tanpa alasan ditolak.
type Penilaian struct {
	Skor   int    `json:"skor"`
	Alasan string `json:"alasan"`
}

// Kasus adalah satu praktik fiktif yang dinilai pada semua sumbu.
type Kasus struct {
	ID        string               `json:"id"`
	Nama      string               `json:"nama"`
	Pasangan  string               `json:"pasangan"`
	Kisah     string               `json:"kisah"`
	Penilaian map[string]Penilaian `json:"penilaian"`
}

// Validasi memeriksa bahwa kasus lengkap dan setiap skor dapat dibela.
// Semua masalah dikumpulkan sekaligus supaya mudah diperbaiki.
func (k Kasus) Validasi() error {
	var masalah []error
	if strings.TrimSpace(k.ID) == "" {
		masalah = append(masalah, errors.New("id kosong"))
	}
	if strings.TrimSpace(k.Nama) == "" {
		masalah = append(masalah, errors.New("nama kosong"))
	}
	if k.Pasangan != "" && k.Pasangan == k.ID {
		masalah = append(masalah, errors.New("kasus tidak boleh berpasangan dengan dirinya sendiri"))
	}
	for _, s := range DaftarSumbu {
		p, ada := k.Penilaian[s.Kode]
		if !ada {
			masalah = append(masalah, fmt.Errorf("sumbu %q belum dinilai", s.Kode))
			continue
		}
		if p.Skor < SkorMin || p.Skor > SkorMaks {
			masalah = append(masalah, fmt.Errorf("sumbu %q: skor %d di luar %d..%d", s.Kode, p.Skor, SkorMin, SkorMaks))
		}
		if strings.TrimSpace(p.Alasan) == "" {
			masalah = append(masalah, fmt.Errorf("sumbu %q: alasan kosong", s.Kode))
		}
	}
	for kode := range k.Penilaian {
		if _, dikenal := CariSumbu(kode); !dikenal {
			masalah = append(masalah, fmt.Errorf("sumbu %q tidak dikenal", kode))
		}
	}
	if len(masalah) > 0 {
		return fmt.Errorf("kasus %q tidak sah: %w", k.ID, errors.Join(masalah...))
	}
	return nil
}

// Muat membaca satu kasus dari JSON, menolak bidang yang tidak dikenal,
// lalu memvalidasinya.
func Muat(r io.Reader) (Kasus, error) {
	dek := json.NewDecoder(r)
	dek.DisallowUnknownFields()
	var k Kasus
	if err := dek.Decode(&k); err != nil {
		return Kasus{}, fmt.Errorf("membaca kasus: %w", err)
	}
	if err := k.Validasi(); err != nil {
		return Kasus{}, err
	}
	return k, nil
}
