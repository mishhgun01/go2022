package eldest_person

import "go2022/hw9/pkg/models"

//MaxAge принимает неопределённое количество данных любого типа и возвращает самого старшего человека
func MaxAge(p ...interface{}) interface{} {
	var max int
	var res interface{}
	for i := range p {
		switch p[i].(type) {
		case models.Employee:
			per := p[i].(models.Employee)
			if per.Age > max {
				res = per
				max = per.Age
			}
		case models.Customer:
			per := p[i].(models.Customer)
			if per.Age > max {
				res = per
				max = per.Age
			}
		}
	}
	return res
}
