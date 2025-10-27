package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"weekly/src/utils"
)

func FecthData(path string) ([]MenuItem, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/ariekas/koda-b4-golang-weekly-data/refs/heads/main/dataProduct.json")

	if err != nil {
		fmt.Println("Error: Failed to Fecth data")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(
		resp.Body,
	)

	var menus []MenuItem

	if err != nil {
		fmt.Println("Failed to raid body")
	}

	err = json.Unmarshal(body, &menus)
	if err != nil {
		fmt.Println("Failed to converd data")
	}

	os.WriteFile(path, body, 0644)

	fmt.Println("Data berhasil disimpan ke:", path)
	return menus, nil
}

func GetData() ([]MenuItem, error) {
	// Menuju file temporary
	tempDir := os.TempDir()
	// Membaca file temporary apakah ada file data.json
	getData := filepath.Join(tempDir, "data.json")

	// Membaca waktu file data.json dibuat pertama kali
	createAt, err := os.Stat(getData)

	// data.json nya ada
	if err == nil {
		// menghitung sudah berapa lama durasi waktu sejak file data.json
		getCreateAt := time.Since(createAt.ModTime())
		if getCreateAt >= time.Duration(utils.Time()) {
			return FecthData(getData)
		} else {
			// Membaca file data.json
			data, err := os.ReadFile(getData)
			if err != nil {
				return nil, fmt.Errorf("failed to read cached data: %v", err)
			}

			var menus []MenuItem
			// Mengubah file json menjadi slice
			json.Unmarshal(data, &menus)
			return menus, nil
		}
	}

	// data.json nya tidak ada
	return FecthData(getData)
}

type MenuItem struct {
	ID       int
	Name     string
	Price    float64
	Category string
}
type Order struct {
	Item     MenuItem
	Quantity int
}

type Transaction struct {
	OrderID   string
	Custemer  string
	Order     []Order
	Total     float64
	DateOrder time.Time
}

func (m MenuItem) PrintProduct(i int) {
	fmt.Printf("%d. %s - Rp %.0f \n", i+1, m.Name, m.Price)
}

func (o Order) PrintProduct(i int) {
	fmt.Printf("%d.\nProduct: %s\nPrice: Rp %.0f\nQuantity: %d\n\n",
		i+1, o.Item.Name, o.Item.Price, o.Quantity)
}

func SearchMenu(menus []MenuItem, keyword string) []MenuItem {

	// chanel untuk mengirim hasil pencarian antar goroutine
	dataMenu := make(chan MenuItem)
	var wg sync.WaitGroup

	// jumlah goroutine yang akan berjalan
	numWorker := 4
	// membagi rata tiap goroutine untuk handel tiap tiap data
	chunckSize := (len(menus) + numWorker - 1) / numWorker

	for i := 0; i < numWorker; i++ {
		startIdx := i * chunckSize
		endIdx := startIdx + chunckSize
		if endIdx > len(menus) {
			endIdx = len(menus)
		}
		part := menus[startIdx:endIdx]

		wg.Add(1)

		go func(part []MenuItem) {
			defer wg.Done()
			for _, menu := range part {
				if strings.Contains(strings.ToLower(menu.Name), strings.ToLower(keyword)) {
					dataMenu <- menu
				}
			}
		}(part)
	}
	go func() {
		wg.Wait()
		close(dataMenu)
	}()

	var results []MenuItem
	for item := range dataMenu {
		results = append(results, item)
	}

	return results
}