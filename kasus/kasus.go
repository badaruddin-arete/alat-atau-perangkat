// Paket kasus memuat delapan kasus fiktif yang ditanam ke dalam program.
package kasus

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/badaruddin-arete/alat-atau-perangkat/rubrik"
)

//go:embed data/*.json
var berkas embed.FS

// Semua memuat kasus bawaan, urut menurut ID.
func Semua() ([]rubrik.Kasus, error) {
	return MuatDari(berkas, "data")
}

// MuatDari memuat semua berkas .json di satu direktori, memvalidasi tiap kasus,
// dan memastikan ID unik serta setiap pasangan saling menunjuk.
func MuatDari(fsys fs.FS, dir string) ([]rubrik.Kasus, error) {
	entri, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	var daftar []rubrik.Kasus
	for _, e := range entri {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		f, err := fsys.Open(path.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		k, err := rubrik.Muat(f)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		if want := strings.TrimSuffix(e.Name(), ".json"); k.ID != want {
			return nil, fmt.Errorf("%s: id %q harus sama dengan nama berkas", e.Name(), k.ID)
		}
		daftar = append(daftar, k)
	}
	sort.Slice(daftar, func(i, j int) bool { return daftar[i].ID < daftar[j].ID })
	if err := periksaPasangan(daftar); err != nil {
		return nil, err
	}
	return daftar, nil
}

func periksaPasangan(daftar []rubrik.Kasus) error {
	menurutID := map[string]rubrik.Kasus{}
	for _, k := range daftar {
		if _, ganda := menurutID[k.ID]; ganda {
			return fmt.Errorf("id ganda: %q", k.ID)
		}
		menurutID[k.ID] = k
	}
	for _, k := range daftar {
		if k.Pasangan == "" {
			continue
		}
		p, ada := menurutID[k.Pasangan]
		if !ada {
			return fmt.Errorf("kasus %q berpasangan dengan %q yang tidak ada", k.ID, k.Pasangan)
		}
		if p.Pasangan != k.ID {
			return fmt.Errorf("pasangan %q dan %q tidak saling menunjuk", k.ID, p.ID)
		}
	}
	return nil
}

// Cari mengembalikan kasus dengan ID tertentu.
func Cari(daftar []rubrik.Kasus, id string) (rubrik.Kasus, bool) {
	for _, k := range daftar {
		if k.ID == id {
			return k, true
		}
	}
	return rubrik.Kasus{}, false
}

// Pasangan mengembalikan setiap pasangan satu kali, sisi pertama menurut urutan ID.
func Pasangan(daftar []rubrik.Kasus) [][2]rubrik.Kasus {
	var hasil [][2]rubrik.Kasus
	for _, k := range daftar {
		if k.Pasangan == "" || k.ID > k.Pasangan {
			continue
		}
		if p, ada := Cari(daftar, k.Pasangan); ada {
			hasil = append(hasil, [2]rubrik.Kasus{k, p})
		}
	}
	return hasil
}
