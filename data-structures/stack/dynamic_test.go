package stack

import (
	"fmt"
	"strings"
	"testing"
)

func generateTestStack(values []int) *DStack {
	res := DStack{}
	for i := len(values) - 1; i >= 0; i-- {
		res.Push(values[i])
	}
	return &res
}

func generateStringForm(values []int) string {
	builder := strings.Builder{}
	builder.WriteString("top -> ")
	for _, v := range values {
		builder.WriteString(fmt.Sprintf("%d -> ", v))
	}
	builder.WriteString("bottom")
	return builder.String()
}

func TestStackBasics(t *testing.T) {
	t.Run("IsEmpty() receiver function", func(t *testing.T) {
		t.Run("on empty stack", func(t *testing.T) {
			s := generateTestStack(nil)
			if se := s.IsEmpty(); se != true {
				t.Errorf("Expected %t, got %t", true, se)
			}
		})

		t.Run("on non-empty stack", func(t *testing.T) {
			s := generateTestStack([]int{0, 1, 2, 3, 4, 5})
			if se := s.IsEmpty(); se != false {
				t.Errorf("Expected %t, got %t", false, se)
			}
		})
	})

	t.Run("Size() receiver function", func(t *testing.T) {
		testCases := []struct {
			name     string
			stack    DStack
			expected int
		}{
			{
				name:     "on empty stack",
				stack:    *generateTestStack(nil),
				expected: 0,
			},
			{
				name:     "on non-empty stack",
				stack:    *generateTestStack([]int{1, 2, 3, 4, 5}),
				expected: 5,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				if s := testCase.stack.Size(); s != testCase.expected {
					t.Errorf("Expected %d, got %d", testCase.expected, s)
				}
			})
		}
	})

	t.Run("String() stringer", func(t *testing.T) {
		testCases := []struct {
			name     string
			stack    DStack
			expected string
		}{
			{
				name:     "on empty stack",
				stack:    *generateTestStack(nil),
				expected: generateStringForm(nil),
			},
			{
				name:     "on non-empty stack",
				stack:    *generateTestStack([]int{1, 2, 3, 4, 5}),
				expected: generateStringForm([]int{1, 2, 3, 4, 5}),
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				if str := testCase.stack.String(); str != testCase.expected {
					t.Errorf("Expected %s, got %s", testCase.expected, str)
				}
			})
		}
	})
}

func TestStackOperations(t *testing.T) {
	t.Run("Push() operation", func(t *testing.T) {
		testCases := []struct {
			name     string
			stack    DStack
			values   []int
			expected string
		}{
			{
				name:     "on empty stack",
				stack:    *generateTestStack(nil),
				values:   []int{3, 4, 5},
				expected: generateStringForm([]int{5, 4, 3}),
			},
			{
				name:     "on non-empty stack",
				stack:    *generateTestStack([]int{7, 8, 9}),
				values:   []int{3, 2, 1},
				expected: generateStringForm([]int{1, 2, 3, 7, 8, 9}),
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				for _, v := range testCase.values {
					testCase.stack.Push(v)
				}
				if str := testCase.stack.String(); str != testCase.expected {
					t.Errorf("Expected \"%s\", got \"%s\"", testCase.expected, str)
				}
			})
		}
	})

	t.Run("Pop() operation", func(t *testing.T) {
		t.Run("on empty stack", func(t *testing.T) {
			stack := generateTestStack(nil)
			if res := stack.Pop(); res != nil {
				t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
			}
		})

		t.Run("on non-empty stack", func(t *testing.T) {
			stack := generateTestStack([]int{6, 8, 10})
			res := stack.Pop()
			if str := stack.String(); str != "top -> 8 -> 10 -> bottom" {
				t.Errorf("Expected stack to be \"top -> 8 -> 10 -> bottom\", got \"%s\"", str)
			}
			if res == nil {
				t.Errorf("Expected a valid pointer to 6, got nil")
			} else {
				if res.data != 6 {
					t.Errorf("Expected value 6, got %d", res.data)
				}
			}
		})
	})

	t.Run("Peek() operation", func(t *testing.T) {
		t.Run("on empty stack", func(t *testing.T) {
			stack := generateTestStack(nil)
			if res := stack.Peek(); res != nil {
				t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
			}
		})

		t.Run("on non-empty stack", func(t *testing.T) {
			stack := generateTestStack([]int{6, 8, 10})
			res := stack.Peek()
			if str := stack.String(); str != "top -> 6 -> 8 -> 10 -> bottom" {
				t.Errorf("Expected stack to be \"top -> 6 -> 8 -> 10 -> bottom\", got \"%s\"", str)
			}
			if res == nil {
				t.Errorf("Expected a pointer to 6, got nil")
			} else {
				if res.data != 6 {
					t.Errorf("Expected pointer to value 6, got pointer to %d", res.data)
				}
			}
		})
	})

	t.Run("Clear() operation", func(t *testing.T) {
		testCases := []struct {
			name     string
			stack    DStack
			expected string
		}{
			{
				name:     "on empty stack",
				stack:    *generateTestStack(nil),
				expected: generateStringForm(nil),
			},
			{
				name:     "on non-empty stack",
				stack:    *generateTestStack([]int{0, 1, 2, 3, 4, 5}),
				expected: generateStringForm(nil),
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				testCase.stack.Clear()
				if str := testCase.stack.String(); str != testCase.expected {
					t.Errorf("Expected stack to be \"%s\", got \"%s\"", testCase.expected, str)
				}
			})
		}
	})
}
