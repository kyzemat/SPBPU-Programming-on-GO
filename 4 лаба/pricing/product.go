package pricing

import "fmt"

// Product описывает книгу, которую можно положить в заказ.
type Product struct {
	Title  string  // название книги
	Author string  // автор
	Price  float64 // базовая цена в рублях
}

// NewProduct создаёт книгу с заданными названием, автором и ценой.
func NewProduct(title, author string, price float64) Product {
	return Product{Title: title, Author: author, Price: price}
}

// String реализует интерфейс fmt.Stringer.
// Пример: «Война и мир» (Л. Н. Толстой) — 450.00 ₽
func (p Product) String() string {
	return fmt.Sprintf("«%s» (%s) — %.2f ₽", p.Title, p.Author, p.Price)
}

// PriceWith возвращает цену книги после применения модификатора.
// Если модификатор nil, возвращается базовая цена.
func (p Product) PriceWith(mod PriceModifier) float64 {
	if mod == nil {
		return p.Price
	}
	return mod(p.Price)
}
