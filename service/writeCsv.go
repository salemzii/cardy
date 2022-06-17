package service

import (
    "encoding/csv"
    "log"
    "os"
    "strconv"

    "github.com/salemzii/cardy/cars"
)

var data [][]string


func WriteFile(dataStore []cars.Info){
	file, err := os.Create("data.csv")
	
	if err != nil {
	    log.Fatalln("failed to open file", err)
	}

	defer file.Close()

    for _, record := range dataStore {
        row := []string{strconv.Itoa(record.Id), record.FirstName, record.LastName,
        record.Email, record.Address, record.Car_Manufactur, record.Car_Model,
        strconv.Itoa(record.Car_Model_Year)}
        data = append(data, row)
    }

	writer := csv.NewWriter(file)
	defer writer.Flush()
    writer.WriteAll(data)
}