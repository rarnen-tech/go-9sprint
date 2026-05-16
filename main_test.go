package main

// Пишите тесты в этом файле
import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "Zero size",
			size: 0,
		},
		{
			name: "Negative size",
			size: -10,
		},
		{
			name: "One element",
			size: 1,
		},
		{
			name: "Many elements",
			size: 100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := generateRandomElements(test.size)

			if test.size <= 0 {
				if len(got) != 0 {
					t.Fatalf("Expected empty slice, got ;ength%v", len(got))
				}

				return

			}

			if len(got) != test.size {
				t.Fatalf("Expected slice of length %d, got %d", test.size, len(got))
			}

		})
	}
}
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{5},
			want: 5,
		},
		{
			name: "several elements",
			data: []int{1, 7, 3, 10, 2},
			want: 10,
		},
		{
			name: "maximum first",
			data: []int{100, 1, 2, 3},
			want: 100,
		},
		{
			name: "same elements",
			data: []int{4, 4, 4, 4},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)

			if got != test.want {
				t.Fatalf("expected %d, got %d", test.want, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "less than chunks",
			data: []int{1, 9, 3},
			want: 9,
		},
		{
			name: "exact chunks",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "more than chunks",
			data: []int{1, 20, 3, 4, 5, 6, 7, 8, 100, 9, 10},
			want: 100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxChunks(test.data)

			if got != test.want {
				t.Fatalf("expected %d, got %d", test.want, got)
			}
		})
	}
}
