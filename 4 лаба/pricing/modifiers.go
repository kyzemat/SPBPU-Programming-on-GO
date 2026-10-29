package pricing

import (
	"math"
)

// PriceModifier — тип-функция, описывающая любое преобразование цены.
type PriceModifier func(float64) float64

// CombineModifiers — функция высшего порядка, которая объединяет
// произвольное количество модификаторов в одну функцию.
// Модификаторы применяются в порядке их передачи.
//
// Если модификаторов нет, возвращается функция "тождество",
// которая возвращает исходную цену без изменений.
func CombineModifiers(modifiers ...PriceModifier) PriceModifier {
	// Замыкание: захватывает слайс modifiers из внешней области видимости.
	return func(price float64) float64 {
		result := price
		for _, mod := range modifiers {
			if mod != nil {
				result = mod(result)
			}
		}
		return result
	}
}

// WithDiscount возвращает модификатор, применяющий скидку в процентах.
// Пример: WithDiscount(10) уменьшает цену на 10%.
func WithDiscount(percent float64) PriceModifier {
	// Замыкание: захватывает переменную percent.
	return func(price float64) float64 {
		return price * (1.0 - percent/100.0)
	}
}

// WithTax возвращает модификатор, добавляющий налог (в процентах).
// Пример: WithTax(20) увеличивает цену на 20%.
func WithTax(rate float64) PriceModifier {
	return func(price float64) float64 {
		return price * (1.0 + rate/100.0)
	}
}

// WithRounding возвращает модификатор, который округляет цену
// до заданного количества знаков после запятой (precision).
// Используется математическое округление (до ближайшего).
func WithRounding(precision int) PriceModifier {
	// Замыкание: захватывает precision.
	// Предвычисляем множитель, чтобы не делать math.Pow на каждом вызове.
	factor := math.Pow(10, float64(precision))
	return func(price float64) float64 {
		return math.Round(price*factor) / factor
	}
}

// WithMinimumPrice возвращает модификатор, который гарантирует,
// что итоговая цена не будет ниже заданного минимума.
func WithMinimumPrice(minPrice float64) PriceModifier {
	return func(price float64) float64 {
		if price < minPrice {
			return minPrice
		}
		return price
	}
}

// WithMaximumPrice возвращает модификатор, ограничивающий цену сверху.
// (Дополнительный пример, чтобы показать гибкость подхода).
func WithMaximumPrice(maxPrice float64) PriceModifier {
	return func(price float64) float64 {
		if price > maxPrice {
			return maxPrice
		}
		return price
	}
}
