package eldest_person

import (
	"go2022/hw9/pkg/models"
	"testing"
)

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name   string
		people []interface{}
		want   interface{}
	}{
		{
			name: "Among Employees",
			people: []interface{}{
				models.Employee{Age: 54, Name: "John"},
				models.Employee{Age: 25, Name: "Mike"},
				models.Employee{Age: 30, Name: "Keith"},
			},
			want: models.Employee{Age: 54, Name: "John"},
		},
		{
			name: "Among Customers",
			people: []interface{}{
				models.Customer{Age: 54, Name: "John"},
				models.Customer{Age: 25, Name: "Mike"},
				models.Customer{Age: 67, Name: "Bob"},
			},
			want: models.Customer{Age: 67, Name: "Bob"},
		},
		{
			name: "Among Everybody",
			people: []interface{}{
				models.Employee{Age: 56, Name: "Mike"},
				models.Employee{Age: 30, Name: "Keith"},
				models.Customer{Age: 54, Name: "John"},
				models.Customer{Age: 25, Name: "Mike"},
			},
			want: models.Employee{Age: 56, Name: "Mike"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.people...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
