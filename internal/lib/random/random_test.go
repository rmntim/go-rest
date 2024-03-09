package random

import "testing"

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size(1)",
			size: 1,
		},
		{
			name: "size(5)",
			size: 5,
		},
		{
			name: "size(10)",
			size: 10,
		},
		{
			name: "size(50)",
			size: 50,
		},
		{
			name: "size(100)",
			size: 100,
		},
		{
			name: "size(1000)",
			size: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str1 := NewRandomString(tt.size)
			str2 := NewRandomString(tt.size)

			if len(str1) != tt.size || len(str2) != tt.size {
				t.Fatal("fail: wrong size")
			}

			if str1 == str2 {
				t.Fatal("fail: equal strings")
			}
		})
	}
}
