package macro

import "C"
import (
	"fmt"
	"sync"

	"github.com/Rob-Fletcher/go-gdal"
)

// Mosaic accepts a list of files (including VSI files) and mosaics them together and outputs them in the provided
// projection
func Mosaic(urls []string, epsg string) {

	if len(urls) == 0 {
		return
	}

	datasets := make([]gdal.Dataset, 0)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			ds, err := gdal.Open(url, gdal.ReadOnly)
			if err != nil {
				println(fmt.Sprintf("Error opening file %s -- skipping", url))
			} else {
				println(fmt.Sprintf("Opened %s", url))
				datasets = append(datasets, ds)
			}
		}(url)
	}
	wg.Wait()

	//todo: combine and close datasets

}
