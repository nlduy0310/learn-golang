package queue

import (
	"fmt"
	"strings"
	"testing"
)

func generateTestQueue(values []int) *DQueue {
	res := DQueue{}
	for _, val := range values {
		res.Enqueue(val)
	}
	return &res
}

func generateStringForm(values []int) string {
	builder := strings.Builder{}
	builder.WriteString("first -> ")
	for _, val := range values {
		builder.WriteString(fmt.Sprintf("%d -> ", val))
	}
	builder.WriteString("last")
	return builder.String()
}

func TestDQueueBasics(t *testing.T) {
	t.Run("IsEmpty() receiver function", func(t *testing.T) {
		testCases := []struct {
			name     string
			vals     []int
			expected bool
		}{
			{
				name:     "on non-empty queue",
				vals:     []int{1, 2, 3, 4, 5},
				expected: false,
			},
			{
				name:     "on empty queue",
				vals:     nil,
				expected: true,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				q := generateTestQueue(testCase.vals)
				actual := q.IsEmpty()
				if actual != testCase.expected {
					t.Errorf("Expected %t, got %t", testCase.expected, actual)
				}
			})
		}
	})

	t.Run("Size() receiver function", func(t *testing.T) {
		testCases := []struct {
			name     string
			vals     []int
			expected int
		}{
			{
				name:     "on an empty queue",
				vals:     nil,
				expected: 0,
			},
			{
				name:     "on a queue with 1 elements",
				vals:     []int{1},
				expected: 1,
			},
			{
				name:     "on a queue with 10 elements",
				vals:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				expected: 10,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				q := generateTestQueue(testCase.vals)
				actual := q.Size()
				if actual != testCase.expected {
					t.Errorf("Expected %d, got %d", testCase.expected, actual)
				}
			})
		}
	})

	t.Run("String() stringer", func(t *testing.T) {
		testCases := []struct {
			name     string
			vals     []int
			expected string
		}{
			{
				name: "on an empty queue",
				vals: nil,
			},
			{
				name: "on a queue with 1 element",
				vals: []int{1},
			},
			{
				name: "on a queue with 10 elements",
				vals: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				q := generateTestQueue(testCase.vals)
				actual := q.String()
				expected := generateStringForm(testCase.vals)
				if actual != expected {
					t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
				}
			})
		}
	})
}

func TestDQueueUpdateOperations(t *testing.T) {
	t.Run("Enqueue() operation", func(t *testing.T) {
		testCases := []struct {
			name           string
			initialValues  []int
			enqueuedValues []int
		}{
			{
				name:           "on an empty queue",
				initialValues:  nil,
				enqueuedValues: []int{1, 2, 3},
			},
			{
				name:           "on a queue with 1 element",
				initialValues:  []int{0},
				enqueuedValues: []int{1, 2, 3},
			},
			{
				name:           "on a queue with 5 elements",
				initialValues:  []int{6, 7, 8, 9, 10},
				enqueuedValues: []int{1, 2, 3},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				q := generateTestQueue(testCase.initialValues)
				for _, val := range testCase.enqueuedValues {
					q.Enqueue(val)
				}
				expected := generateStringForm(append(testCase.initialValues, testCase.enqueuedValues...))
				actual := q.String()
				if actual != expected {
					t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
				}
			})
		}
	})

	t.Run("Dequeue() operation", func(t *testing.T) {
		testCases := []struct {
			name          string
			initialValues []int
			times         []struct {
				nTimes   int
				expected string
			}
		}{
			{
				name:          "on an empty queue",
				initialValues: nil,
				times: []struct {
					nTimes   int
					expected string
				}{
					{
						nTimes:   1,
						expected: generateStringForm(nil),
					},
					{
						nTimes:   5,
						expected: generateStringForm(nil),
					},
				},
			},
			{
				name:          "on a queue with 5 elements",
				initialValues: []int{1, 2, 3, 4, 5},
				times: []struct {
					nTimes   int
					expected string
				}{
					{
						nTimes:   1,
						expected: generateStringForm([]int{2, 3, 4, 5}),
					},
					{
						nTimes:   5,
						expected: generateStringForm(nil),
					},
					{
						nTimes:   8,
						expected: generateStringForm(nil),
					},
				},
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				for _, nt := range testCase.times {
					t.Run(fmt.Sprintf("%d time(s)", nt.nTimes), func(t *testing.T) {
						q := generateTestQueue(testCase.initialValues)
						for i := 0; i < nt.nTimes; i++ {
							q.Dequeue()
						}
						actual := q.String()
						if actual != nt.expected {
							t.Errorf("Expected \"%s\", got \"%s\"", nt.expected, actual)
						}
					})
				}
			})
		}
	})
}

func TestDQueueAccessOperations(t *testing.T) {
	t.Run("First() accessor", func(t *testing.T) {
		t.Run("on empty queue", func(t *testing.T) {
			q := generateTestQueue(nil)
			first := q.First()
			if first != nil {
				t.Errorf("Expected nil, got %p pointing to %d", first, first.data)
			}
		})

		t.Run("on non-empty queue", func(t *testing.T) {
			testCases := []struct {
				name     string
				vals     []int
				expected int
			}{
				{
					name:     "on a queue with 1 element",
					vals:     []int{1},
					expected: 1,
				},
				{
					name:     "on a queue with 5 elements",
					vals:     []int{0, 2, 3, 4, 5},
					expected: 0,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					q := generateTestQueue(testCase.vals)
					first := q.First()
					if first == nil {
						t.Errorf("Expected pointer to %d, got nil", testCase.expected)
					} else {
						if first.data != testCase.expected {
							t.Errorf("Expected pointer to %d, got pointer to %d", testCase.expected, first.data)
						}
					}
				})
			}
		})
	})

	t.Run("Last() accessor", func(t *testing.T) {
		t.Run("on empty queue", func(t *testing.T) {
			q := generateTestQueue(nil)
			first := q.First()
			if first != nil {
				t.Errorf("Expected nil, got %p pointing to %d", first, first.data)
			}
		})

		t.Run("on non-empty queue", func(t *testing.T) {
			testCases := []struct {
				name     string
				vals     []int
				expected int
			}{
				{
					name:     "on a queue with 1 element",
					vals:     []int{1},
					expected: 1,
				},
				{
					name:     "on a queue with 5 elements",
					vals:     []int{0, 2, 3, 4, 5},
					expected: 5,
				},
			}

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					q := generateTestQueue(testCase.vals)
					last := q.Last()
					if last == nil {
						t.Errorf("Expected pointer to %d, got nil", testCase.expected)
					} else {
						if last.data != testCase.expected {
							t.Errorf("Expected pointer to %d, got pointer to %d", testCase.expected, last.data)
						}
					}
				})
			}
		})
	})
}
