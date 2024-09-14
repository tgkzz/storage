package dto

type UpdateItem struct {
	Id       int64
	Name     *string
	Desc     *string
	Quantity *int32
	Price    *float32
	Currency *string
}

func (u UpdateItem) HaveName() bool {
	return u.Name != nil
}

func (u UpdateItem) SetName(str string) {
	u.Name = &str
}

func (u UpdateItem) HaveDesc() bool {
	return u.Desc != nil
}

func (u UpdateItem) SetDesc(str string) {
	u.Desc = &str
}

func (u UpdateItem) HaveQuantity() bool {
	return u.Quantity != nil
}

func (u UpdateItem) SetQuantity(q int32) {
	u.Quantity = &q
}

func (u UpdateItem) HavePrice() bool {
	return u.Price != nil
}

func (u UpdateItem) SetPrice(p float32) {
	u.Price = &p
}

func (u UpdateItem) HaveCurrency() bool {
	return u.Currency != nil
}

func (u UpdateItem) SetCurrency(s string) {
	u.Currency = &s
}
