package cars

import "strings"

func CarsWithOneOwner(data []Info, name string) []Info {

	carList := []Info{}

	for _, v := range data {

		if strings.ToLower(v.FirstName) == name || strings.ToLower(v.LastName) == name {

			carList = append(carList, v)
		}
	}

	return carList
}

func CarWithSameManufacturer(data []Info, manufacturer string) []Info {

	carList := []Info{}

	for _, v := range data {

		if strings.ToLower(v.Car_Manufactur) == manufacturer {
			carList = append(carList, v)
		}
	}

	return carList
}

func SearchByCity(data []Info, city string) []Info {
	carList := []Info{}

	for _, v := range data {
		if strings.HasSuffix(v.Address, city) {
			carList = append(carList, v)
		}
	}

	return carList
}
