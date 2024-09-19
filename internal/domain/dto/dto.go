package dto

type UpdateItem struct {
	Id       int64
	Name     *string
	Desc     *string
	Quantity *int32
	Price    *float32
	Currency *string
}

func (u *UpdateItem) HaveName() bool {
	return u.Name != nil
}

func (u *UpdateItem) SetName(str string) *UpdateItem {
	u.Name = &str

	return u
}

func (u *UpdateItem) HaveDesc() bool {
	return u.Desc != nil
}

func (u *UpdateItem) SetDesc(str string) *UpdateItem {
	u.Desc = &str

	return u
}

func (u *UpdateItem) HaveQuantity() bool {
	return u.Quantity != nil
}

func (u *UpdateItem) SetQuantity(q int32) *UpdateItem {
	u.Quantity = &q

	return u
}

func (u *UpdateItem) HavePrice() bool {
	return u.Price != nil
}

func (u *UpdateItem) SetPrice(p float32) *UpdateItem {
	u.Price = &p

	return u
}

func (u *UpdateItem) HaveCurrency() bool {
	return u.Currency != nil
}

func (u *UpdateItem) SetCurrency(s string) *UpdateItem {
	u.Currency = &s

	return u
}

func NewUpdateItem(id int64) *UpdateItem {
	return &UpdateItem{
		Id: id,
	}
}
