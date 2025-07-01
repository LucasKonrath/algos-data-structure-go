package fractionalKnapsack

import "testing"

func TestFractionalKnapsack(t *testing.T) {
	tests := []struct {
		capacity float64
		items    []Item
		expected float64
	}{
		{
			capacity: 50,
			items: []Item{
				{value: 60, weight: 10},
				{value: 100, weight: 20},
				{value: 120, weight: 30},
			},
			expected: 240.0, // Take all of first two, 2/3 of last
		},
		{
			capacity: 10,
			items: []Item{
				{value: 500, weight: 30},
				{value: 400, weight: 20},
				{value: 200, weight: 10},
			},
			expected: 200.0, // Only the last item fits
		},
		{
			capacity: 0,
			items: []Item{
				{value: 100, weight: 10},
			},
			expected: 0.0,
		},
		{
			capacity: 5,
			items: []Item{
				{value: 50, weight: 10},
			},
			expected: 25.0, // Take half of the only item
		},
		{
			capacity: 100,
			items:    []Item{},
			expected: 0.0,
		},
	}

	for _, test := range tests {
		result := fractionalKnapsack(test.capacity, test.items)
		if (result-test.expected) > 1e-6 || (test.expected-result) > 1e-6 {
			t.Errorf("fractionalKnapsack(%v, %v) = %v; want %v", test.capacity, test.items, result, test.expected)
		}
	}
}
