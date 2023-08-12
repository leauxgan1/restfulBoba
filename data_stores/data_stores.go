package data_stores

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
  Ingredients []string `json:"ingredients`
}

type Location struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Menu []MenuItem `json:"menu"`
}

var Locations = []Location{
	{
		ID:   "01",
		Name: "Guppy House",
		Menu: []MenuItem{
			{
        Name: "Green Milk Tea", 
        Price: 4.50, 
        Ingredients: []string{"Non-dairy milk powder", "Jasmine Green Tea"},
      },
      {
        Name: "Passionfruit Green Tea", 
        Price: 5.00, 
        Ingredients: []string{"Jasmine Green Tea", "Passionfruit Syrup", "High Fructose Corn Syrup"},
      },
      {
        Name: "Brown Sugar Boba", 
        Price: 0.50, 
        Ingredients: []string{"Tapioca starch", "Brown Sugar"},
      },
      {
        Name: "Lychee Jelly", 
        Price: 0.50, 
        Ingredients: []string{"Lychee fruit", "Gelatin"},
      },
		},
  },
}
