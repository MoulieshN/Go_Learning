package decoratorp

type Decorator interface {
	GetPrice() int
	GetIngredients() string
}

type Coffee struct {
	price       int
	ingredients string
}

func NewCoffee(p int, ing string) Decorator {
	return &Coffee{
		price:       p,
		ingredients: ing,
	}
}

func (c *Coffee) GetPrice() int {
	return c.price
}

func (c *Coffee) GetIngredients() string {
	return c.ingredients
}

type Milk struct {
	decorator Decorator
}

func NewMilk(decorator Decorator) Decorator {
	return &Milk{
		decorator: decorator,
	}
}

func (m *Milk) GetPrice() int {
	return m.decorator.GetPrice() + 10
}

func (m *Milk) GetIngredients() string {
	return m.decorator.GetIngredients() + " milk"
}

type Sugar struct {
	decorator Decorator
}

func NewSugar(decorator Decorator) Decorator {
	return &Sugar{
		decorator: decorator,
	}
}

func (s *Sugar) GetPrice() int {
	return s.decorator.GetPrice() + 15
}

func (s *Sugar) GetIngredients() string {
	return s.decorator.GetIngredients() + " Sugar"
}
