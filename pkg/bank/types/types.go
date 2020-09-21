package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы т.д)
type Money int64

//Category представляет собой категорию, в который был совершон платёж(авто, аптеки, ресторан и т.д)
type Category string

//Payment представляет информатцию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
