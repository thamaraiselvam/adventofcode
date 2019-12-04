package _2019

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findFuel(t *testing.T) {
	t.Run("Should return fuel for given mass", func(t *testing.T) {
		tests := [] struct {
			name         string
			mass         int
			expectedFuel int
		}{
			{
				name:         "Should return zero fuel",
				mass:         0,
				expectedFuel: 0,
			}, {
				name:         "Should return 2 fuel",
				mass:         12,
				expectedFuel: 2,
			}, {
				name:         "Should return 2 fuel",
				mass:         14,
				expectedFuel: 2,
			}, {
				name:         "Should return 654 fuel",
				mass:         1969,
				expectedFuel: 654,
			}, {
				name:         "Should return 33583 fuel",
				mass:         100756,
				expectedFuel: 33583,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := findFuel(test.mass, false, 0)
				assert.Equal(t, test.expectedFuel, actual)
			})
		}
	})

	t.Run("Should return additional fuel for given mass", func(t *testing.T) {
		tests := [] struct {
			name           string
			mass           int
			expectedFuel   int
			additionalFuel bool
		}{
			{
				name:           "Should return additional 2 fuel",
				mass:           12,
				expectedFuel:   2,
				additionalFuel: true,
			}, {
				name:           "Should return additional 966 fuel",
				mass:           1969,
				expectedFuel:   966,
				additionalFuel: true,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := findFuel(test.mass, test.additionalFuel, 0)
				assert.Equal(t, test.expectedFuel, actual)
			})
		}

	})
}

func Test_findFuels(t *testing.T) {
	tests := [] struct {
		name         string
		masses       []int
		expectedFuel int
	}{
		{
			name:         "Should return 2 fuel for given mass",
			masses:       []int{12, 14},
			expectedFuel: 4,
		}, {
			name:         "should return 3412207 fuel for given mass",
			masses:       []int{137139, 104321, 137149, 97698, 115133, 64329, 82531, 111730, 56831, 145953, 73388, 57230, 61935, 58542, 147631, 79366, 115484, 86997, 80362, 129109, 58568, 121969, 63696, 68116, 86668, 62059, 59485, 132507, 107823, 94467, 53032, 140962, 129499, 80599, 147570, 96463, 126169, 108575, 133312, 146929, 79826, 114449, 110908, 66107, 62171, 91677, 128188, 90483, 81045, 96006, 110366, 140765, 148360, 54565, 56664, 121547, 78839, 123739, 115408, 123245, 92419, 132564, 80022, 103370, 145366, 145211, 110360, 145897, 140817, 77978, 138064, 148134, 86208, 89472, 67117, 63423, 148536, 105835, 107783, 98601, 66702, 50459, 55127, 79808, 79628, 76264, 134392, 125547, 118186, 80947, 121669, 107315, 145093, 56296, 117226, 105409, 149238, 142651, 103286, 139215},
			expectedFuel: 3412207,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findTotalFuel(test.masses)
			assert.Equal(t, test.expectedFuel, actual)
		})
	}
}

func Test_sumOfGivenArray(t *testing.T) {

	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Should return 10 as expected",
			values:   []int{2, 3, 5},
			expected: 10,
		},
		{
			name:     "Should return 5 as expected",
			values:   []int{1, 10, -6},
			expected: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sumOfGivenArray(test.values)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsMassOrFuelBelowZero(t *testing.T) {
	type args struct {
		mass        int
		currentFuel int
	}
	tests := []struct {
		name     string
		args     args
		expected bool
	}{
		{
			name: "Should return true",
			args: args{
				mass:        0,
				currentFuel: -10,
			},
			expected: true,
		},
		{
			name: "Should return false",
			args: args{
				mass:        1,
				currentFuel: 1,
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isMassOrFuelBelowZero(test.args.mass, test.args.currentFuel)
			assert.Equal(t, test.expected, actual)
		})
	}
}
