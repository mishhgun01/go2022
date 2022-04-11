package hw6

import (
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func (geom *Geom) CalculateDistance() float64 {

	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {

		return -1
	}
	var distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	// возврат расстояния между точками
	return distance
}

//--------изменения-------//
//в декларации метода определил аргумент как указатель, а не значение
//избавился от else
//оставил только тип возврщаемого значения, без имени и определил distance явно
//убрал fmt.Println - ошибки должны обрабатываться в вызывющем коде

//плюсом написал простенькие тесты
