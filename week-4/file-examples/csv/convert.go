package csv_utils

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/h4yfans/file-operation/models"
)

func JSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var locations models.LocationSlice
	err = json.NewDecoder(sourceFile).Decode(&locations)
	if err != nil {
		return err
	}

	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := csv.NewWriter(output)
	headers := []string{"CountryISO", "CountryNameEN", "CityNameEN", "CityIATA", "CityLatitude", "CityLongitude"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, loc := range locations {
		var row []string
		row = append(row, loc.CountryCode, loc.CountryName, loc.CityName, loc.IATACode, loc.Latitude, loc.Longitude)
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	log.Println("JSON to CSV process completed")
	return nil
}
