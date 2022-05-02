package eldest_person

import (
	"go2022/hw9/pkg/maxage"
)

// MaxAge принимает неопределённое количество данных любого типа и возвращает самого старшего человека
func MaxAge(p ...interface{}) interface{} {
	var max int
	var res interface{}
	for i := range p {
		switch p[i].(type) {
		case maxage.Employee:
			per := p[i].(maxage.Employee)
			if per.Age > max {
				res = per
				max = per.Age
			}
		case maxage.Customer:
			per := p[i].(maxage.Customer)
			if per.Age > max {
				res = per
				max = per.Age
			}
		}
	}
	return res
}
