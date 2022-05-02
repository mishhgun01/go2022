package maxage

import "testing"

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name   string
		people []getter
		want   int
	}{
		{
			name: "Among Employees",
			people: []getter{
				&Employee{Age: 54, Name: "John"},
				&Employee{Age: 25, Name: "Mike"},
				&Employee{Age: 30, Name: "Keith"},
			},
			want: 54,
		},
		{
			name: "Among Customers",
			people: []getter{
				&Customer{Age: 54, Name: "John"},
				&Customer{Age: 25, Name: "Mike"},
				&Customer{Age: 67, Name: "Bob"},
			},
			want: 67,
		},
		{
			name: "Among Everybody",
			people: []getter{
				&Employee{Age: 56, Name: "Mike"},
				&Employee{Age: 30, Name: "Keith"},
				&Customer{Age: 54, Name: "John"},
				&Customer{Age: 25, Name: "Mike"},
			},
			want: 56,
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
