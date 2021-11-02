package types

type Money int64 //в минимальных единицах(копейки , дирамы , центы и тд)

type Category string //категории в которых был совершен платеж(авто , рестораны,аптеки итд.)

type Payment struct { // информация о платяже
	ID       int
	Amount   Money
	Category Category
}
