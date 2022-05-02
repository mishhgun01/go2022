package maxage

import "sort"

// Employee структура сотрудника
type Employee struct {
	Name string
	Age  int
}

// Customer - структура покупателя
type Customer struct {
	Name string
	Age  int
}

// Методы получения возраста у покупателя и сотрудника
func (e *Employee) getAge() int {
	return e.Age
}

func (c *Customer) getAge() int {
	return c.Age
}

// getter реализует контракт
type getter interface {
	getAge() int
}

// MaxAge принимает на вход неопределённое количество данных любого типа и возвращает возраст старшего человека
func MaxAge(p ...getter) int {
	sort.Slice(p, func(i, j int) bool {
		return p[i].getAge() > p[j].getAge()
	})
	return p[0].getAge()
}
