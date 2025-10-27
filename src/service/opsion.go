package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Opsion() {
	var input string
	tempDir := os.TempDir()
	getData := filepath.Join(tempDir, "data.json")

	fmt.Printf("Apakah kamu yakin ingin menghapus data product ? (y/n): ")

	input = strings.TrimSpace(strings.ToLower(input))

	fmt.Scan(&input)
	if input != "y" {
		fmt.Println("Dibatalkan, file tidak dihapus.")
		fmt.Scanln()
		return
	}

	err := os.Remove(getData)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File tidak ditemukan.")
		} else {
			fmt.Println("Gagal menghapus file:", err)
		}
	} else {
		fmt.Println("File data.json berhasil dihapus.")
	}

	fmt.Println("Tekan Enter untuk kembali ke menu...")
	fmt.Scanln()
}
