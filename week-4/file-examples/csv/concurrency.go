package csv_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"

	"github.com/h4yfans/file-operation/models"
)

func ReadLocationWithWorkerPool(path string) error {
	const numJobs = 5
	jobs := make(chan []string, numJobs)
	results := make(chan models.Location, numJobs)
	wg := sync.WaitGroup{}

	for w := 1; w <= 3; w++ {
		fmt.Println("worker starting", w)
		wg.Add(1)
		go toStruct(jobs, results, &wg)
	}

	go func() {
		fmt.Println("open file running...")
		f, _ := os.Open(path)
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		for _, line := range lines[1:] {
			fmt.Println("line", line[0])
			jobs <- line
		}

		close(jobs)
	}()

	go func() {
		fmt.Println("wait")
		wg.Wait()
		close(results)
	}()

	for v := range results {
		fmt.Println("print", v)
	}

	return nil
}

func toStruct(jobs <-chan []string, results chan<- models.Location, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println()
	for j := range jobs {
		location := models.Location{
			CountryCode: j[0],
			CountryName: j[1],
			CityName:    j[2],
			IATACode:    j[3],
			Latitude:    j[4],
			Longitude:   j[5],
		}
		//process...

		results <- location
	}
}
