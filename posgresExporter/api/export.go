package api

import "sync"

// Export exports data from PostgreSQL database to .CSV files.
func Export(wg *sync.WaitGroup, stopCh <-chan struct{}, outDir, name, query, maxLines string) {
	defer wg.Done()

}
