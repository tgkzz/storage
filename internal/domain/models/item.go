package models

type Item struct {
	ID       int64
	Name     string
	Desc     string
	Quantity int32
	Price    float32
	Currency string
}

type OptionFunc func(*Item)

func SetId(id int64) OptionFunc {
	return func(item *Item) {
		item.ID = id
	}
}

func SetName(name string) OptionFunc {
	return func(item *Item) {
		item.Name = name
	}
}

func SetDesc(desc string) OptionFunc {
	return func(item *Item) {
		item.Desc = desc
	}
}

func SetQuantity(n int32) OptionFunc {
	return func(item *Item) {
		item.Quantity = n
	}
}

func SetPrice(price float32) OptionFunc {
	return func(item *Item) {
		item.Price = price
	}
}

func SetCurrency(curr string) OptionFunc {
	return func(item *Item) {
		item.Currency = curr
	}
}

func NewItem(opts ...OptionFunc) *Item {
	item := &Item{}

	for _, opt := range opts {
		opt(item)
	}
	return item
}
