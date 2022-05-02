package eldest_person

import (
	"go2022/hw9/pkg/maxage"
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
				maxage.Employee{Age: 54, Name: "John"},
				maxage.Employee{Age: 25, Name: "Mike"},
				maxage.Employee{Age: 30, Name: "Keith"},
			},
			want: maxage.Employee{Age: 54, Name: "John"},
		},
		{
			name: "Among Customers",
			people: []interface{}{
				maxage.Customer{Age: 54, Name: "John"},
				maxage.Customer{Age: 25, Name: "Mike"},
				maxage.Customer{Age: 67, Name: "Bob"},
			},
			want: maxage.Customer{Age: 67, Name: "Bob"},
		},
		{
			name: "Among Everybody",
			people: []interface{}{
				maxage.Employee{Age: 56, Name: "Mike"},
				maxage.Employee{Age: 30, Name: "Keith"},
				maxage.Customer{Age: 54, Name: "John"},
				maxage.Customer{Age: 25, Name: "Mike"},
			},
			want: maxage.Employee{Age: 56, Name: "Mike"},
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
