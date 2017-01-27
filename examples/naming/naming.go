type Ingredient string

var defaultIngredients = []Ingredient{
	Ingredient{"tomatoes"},
	Ingredient{"sauce"},
}

type Pizza struct {
	Name        string
	ingredients []Ingredients
	price       int
}

func (p Pizza) AddBaseIngredients() {
	p.Ingredients = append(p.Ingredients, defaultIngredients...)
}

func (p Pizza) Price() int {
	return p.price
}
