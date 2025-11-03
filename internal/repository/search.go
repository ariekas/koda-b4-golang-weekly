package repository

import (
	"strings"
	"sync"
	"weekly/internal/model"
)

func SearchMenu(menus []model.MenuItem, keyword string) []model.MenuItem {

	// chanel untuk mengirim hasil pencarian antar goroutine
	dataMenu := make(chan model.MenuItem)
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

		go func(part []model.MenuItem) {
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

	var results []model.MenuItem
	for item := range dataMenu {
		results = append(results, item)
	}

	return results
}
