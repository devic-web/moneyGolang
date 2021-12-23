package main

type Product struct {
	Name string `json:"name"`
	Date struct {
		Mount int `json:"mount"`
		Year  int `json:"year"`
	} `json:"date"`
	Datails struct {
		Interes  float64 `json:"interes"`
		Precing  float64 `json:"precing"`
		Mensualy int     `json:"mensualy"`
	} `json:"datails"`
}

func (p Product) PriceMount() float64 {
	var price float64
	if p.Datails.Interes > 0 && p.Datails.Interes < 1 {
		price = p.Datails.Precing * (p.Datails.Interes + 1)
	} else {
		price = p.Datails.Precing
	}

	if p.Datails.Mensualy == 0 {
		return 0
	} else {
		return price / float64(p.Datails.Mensualy)
	}

}

func (p Product) UltimeMount() int {
	return (p.Date.Mount + p.Datails.Mensualy) % 12
}

func (p Product) UltimeYear() int {
	return p.Date.Year + (p.Date.Mount/12 + p.Datails.Mensualy/12)
}
