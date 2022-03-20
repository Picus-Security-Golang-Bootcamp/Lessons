package csv_utils

import (
	"encoding/csv"
	"os"

	"github.com/h4yfans/file-operation/models"
)

func ReadCSV(filename string) (models.LocationSlice, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var locations models.LocationSlice
	for _, line := range records[1:] {
		locations = append(locations, models.Location{
			CountryCode: line[0],
			CountryName: line[1],
			CityName:    line[2],
			IATACode:    line[3],
			Latitude:    line[4],
			Longitude:   line[5],
		})
	}

	return locations, nil
}
