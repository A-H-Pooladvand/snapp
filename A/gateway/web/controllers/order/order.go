package order

type Controller struct{}

func NewOrder() *Controller {
	return new(Controller)
}

func (o Controller) Queue() string {
	return "order"
}

func (o Controller) Exchange() string {
	return ""
}
