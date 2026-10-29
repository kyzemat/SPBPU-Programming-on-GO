package pricing

import (
	"fmt"
	"strings"
)

// Order — заказ, состоящий из нескольких книг.
type Order struct {
	ID       int       // номер заказа
	Customer string    // имя покупателя
	Items    []Product // книги в заказе
}

// NewOrder создаёт пустой заказ.
func NewOrder(id int, customer string) *Order {
	return &Order{ID: id, Customer: customer}
}

// Add добавляет книги в заказ и возвращает сам заказ,
// что позволяет вызывать методы цепочкой.
func (o *Order) Add(products ...Product) *Order {
	o.Items = append(o.Items, products...)
	return o
}

// BaseTotal возвращает сумму заказа без каких-либо модификаторов.
func (o *Order) BaseTotal() float64 {
	return o.TotalWith(nil)
}

// TotalWith возвращает сумму заказа, когда модификатор применён
// к цене каждой книги по отдельности.
func (o *Order) TotalWith(mod PriceModifier) float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.PriceWith(mod)
	}
	return total
}

// Receipt формирует текстовый чек: цена каждой книги до и после
// применения модификатора, затем итог.
func (o *Order) Receipt(mod PriceModifier) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Заказ №%d, покупатель: %s\n", o.ID, o.Customer)
	for i, item := range o.Items {
		fmt.Fprintf(&sb, "  %d. «%s» — %.2f ₽ → %.2f ₽\n",
			i+1, item.Title, item.Price, item.PriceWith(mod))
	}
	fmt.Fprintf(&sb, "Без модификаторов: %.2f ₽\n", o.BaseTotal())
	fmt.Fprintf(&sb, "К оплате:          %.2f ₽", o.TotalWith(mod))
	return sb.String()
}

// String реализует интерфейс fmt.Stringer: краткое описание заказа.
func (o *Order) String() string {
	return fmt.Sprintf("Заказ №%d (%s): %d шт., базовая сумма %.2f ₽",
		o.ID, o.Customer, len(o.Items), o.BaseTotal())
}
