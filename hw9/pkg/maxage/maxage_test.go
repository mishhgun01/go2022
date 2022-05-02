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
				&employee{age: 54, name: "John"},
				&employee{age: 25, name: "Mike"},
				&employee{age: 30, name: "Keith"},
			},
			want: 54,
		},
		{
			name: "Among Customers",
			people: []getter{
				&customer{age: 54, name: "John"},
				&customer{age: 25, name: "Mike"},
				&customer{age: 67, name: "Bob"},
			},
			want: 67,
		},
		{
			name: "Among Everybody",
			people: []getter{
				&employee{age: 56, name: "Mike"},
				&employee{age: 30, name: "Keith"},
				&customer{age: 54, name: "John"},
				&customer{age: 25, name: "Mike"},
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
