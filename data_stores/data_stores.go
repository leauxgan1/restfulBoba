package data_stores

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Location struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Menu []MenuItem `json:"menu"`
}

var Locations = []Location{
	{
		ID:   "01",
		Name: "Boba Fiend",
		Menu: []MenuItem{
			{
				Name: "Green Milk Tea", Price: 3.50,
			},
			{
				Name: "Passionfruit Green Tea", Price: 5.00,
			},
		},
	},
}
