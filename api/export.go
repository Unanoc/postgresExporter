package api

import "sync"

// Export exports data from PostgreSQL database to .CSV files.
func Export(wg *sync.WaitGroup, OutDir, name, query, maxLines string) {
	defer wg.Done()

}
