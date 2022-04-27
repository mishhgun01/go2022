package maxage

import "sort"

//employee структура сотрудника
type employee struct {
	name string
	age  int
}

//customer - структура покупателя
type customer struct {
	name string
	age  int
}

//Методы получения возраста у покупателя и сотрудника
func (e *employee) getAge() int {
	return e.age
}

func (c *customer) getAge() int {
	return c.age
}

//getter реализует контракт
type getter interface {
	getAge() int
}

//MaxAge принимает на вход неопределённое количество данных любого типа и возвращает возраст старшего человека
func MaxAge(p ...getter) int {
	sort.Slice(p, func(i, j int) bool {
		return p[i].getAge() > p[j].getAge()
	})
	return p[0].getAge()
}
