package expenses

type Percentile struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"Percentage"`
	Spent      float64 `json:"spent"`
}

func (p *Percentile) CalculatorResumen(salary float64) *Resumen {
	var r = &Resumen{
		Name:        p.Name,
		Type:        "Porcentil",
		MountInit:   1,
		MountsToPay: 12,
	}

	r.PriceYear = ((DAYS_MOUNTH * MOUNTHS_YEAR) * p.Percentage) * salary

	r.Porcentile = (r.PriceYear / salary) / (DAYS_MOUNTH * MOUNTHS_YEAR)

	r.Complete = p.Spent / r.PriceYear

	return r
}
