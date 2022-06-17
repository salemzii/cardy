package cars

type Info struct {
	Id             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Car_Manufactur string `json:"car_manufactur"`
	Car_Model      string `json:"car_model"`
	Car_Model_Year int    `json:"car_model_year"`
}
