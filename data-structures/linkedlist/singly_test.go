package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

func generateTestList(data []int) SLinkedList {
	if len(data) == 0 {
		return SLinkedList{}
	}

	res := SLinkedList{}
	tmp := newSNode(data[0])
	res.head = tmp

	for i := 1; i < len(data); i++ {
		cur := newSNode(data[i])
		tmp.next = cur
		tmp = cur
	}
	return res
}

func generateStringForm(listData []int) string {
	builder := strings.Builder{}
	builder.WriteString("head -> ")
	for i := 0; i < len(listData); i++ {
		builder.WriteString(fmt.Sprintf("%d -> ", listData[i]))
	}
	builder.WriteString("nil")
	return builder.String()
}

func TestSLListBasic(t *testing.T) {

	// test: String()
	a, b := generateTestList(nil), generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	t.Run("Test linked list's stringer function", func(t *testing.T) {
		as := a.String()
		if as != "head -> nil" {
			t.Errorf("Expected empty list to be \"head -> nil\", got \"%s\"", as)
		}

		bs := b.String()
		if bs != "head -> 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> nil" {
			t.Errorf("Expected %s, got %s",
				"head -> 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> nil",
				bs,
			)
		}
	})

	// Test: IsEmpty()
	t.Run("Test IsEmpty() receiver function", func(t *testing.T) {
		ae, be := a.IsEmpty(), b.IsEmpty()

		if ae != true {
			t.Errorf("First linked list is supposed to be empty, got %t from IsEmpty() call", ae)
		}

		if be != false {
			t.Errorf("Second linked list is not supposed to be empty, got %t from IsEmpty() call", be)
		}
	})

	// Test: Length()
	t.Run("Test Length() receiver function", func(t *testing.T) {
		if al := a.Length(); al != 0 {
			t.Errorf("First linked list is supposed to be empty, got %d from Length() call", al)
		}

		if bl := b.Length(); bl != 11 {
			t.Errorf("Second linked list's length is 11, got %d from Length() call", bl)
		}
	})
}

func TestSLListFind(t *testing.T) {
	t.Run("Find first node", func(t *testing.T) {
		a, b := generateTestList(nil), generateTestList([]int{1, 2, 3})

		if af := a.First(); af != nil {
			t.Errorf("First() should return nil when called on empty list, received %p", af)
		}

		bf := b.First()
		if bf == nil {
			t.Error("First() should return a valid address when called on a non-empty list, got nil")
		} else {
			if bf.data != 1 {
				t.Errorf("The head value of the second list should be 1, got %d", bf.data)
			}
		}
	})

	t.Run("Find last node", func(t *testing.T) {
		a, b, c := generateTestList(nil), generateTestList([]int{1}), generateTestList([]int{1, 2})

		if al := a.Last(); al != nil {
			t.Errorf("Last() should return nil when called on an empty list, got %p", al)
		}

		bl := b.Last()
		if bl == nil {
			t.Error("Calling Last() on a non-empty array should return a valid address, got nil")
		} else {
			if bl.data != 1 {
				t.Errorf("The last element on the second list should have the value 1, got %d", bl.data)
			}
		}

		cl := c.Last()
		if cl == nil {
			t.Error("Calling Last() on a non-empty array should return a valid address, got nil")
		} else {
			if cl.data != 2 {
				t.Errorf("The last element on the third list should have the value 2, got %d", cl.data)
			}
		}
	})

	t.Run("Find node by index", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			a := generateTestList(nil)

			if neg := a.GetAt(-1); neg != nil {
				t.Errorf("On an empty list, get element at -1 should return nil, got %p", neg)
			}

			if zero := a.GetAt(0); zero != nil {
				t.Errorf("On an empty list, get element at 0 should return nil, got %p", zero)
			}

			if pos := a.GetAt(1); pos != nil {
				t.Errorf("On an empty list, get element at 1 should return nil, got %p", pos)
			}
		})

		t.Run("on non-empty list", func(t *testing.T) {
			a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

			t.Run("with negative index", func(t *testing.T) {
				idx := -1
				e := a.GetAt(idx)
				if e != nil {
					t.Errorf("Getting an element with negative index should return nil, got %p", e)
				}
			})

			t.Run("with valid index", func(t *testing.T) {
				idx := 0
				e := a.GetAt(idx)
				if e == nil {
					t.Errorf("Getting an element with valid index should return a valid pointer, got nil")
				} else {
					if e.data != idx {
						t.Errorf("The element with index %d should be %d, got %d", idx, idx, e.data)
					}
				}

				idx = 5
				e = a.GetAt(idx)
				if e == nil {
					t.Errorf("Getting an element with valid index should return a valid pointer, got nil")
				} else {
					if e.data != idx {
						t.Errorf("The element with index %d should be %d, got %d", idx, idx, e.data)
					}
				}

				idx = 9
				e = a.GetAt(idx)
				if e == nil {
					t.Errorf("Getting an element with valid index should return a valid pointer, got nil")
				} else {
					if e.data != idx {
						t.Errorf("The element with index %d should be %d, got %d", idx, idx, e.data)
					}
				}
			})

			t.Run("with out of range index", func(t *testing.T) {
				oor := a.GetAt(10)
				if oor != nil {
					t.Errorf("List doesn't have element with index 10, got %p pointing to %d", oor, oor.data)
				}
			})
		})
	})

	t.Run("Find node by value", func(t *testing.T) {
		a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		t.Run("with existing value", func(t *testing.T) {
			e := a.GetWhere(5)

			if e == nil {
				t.Errorf("GetWhere() call returns nil, expected a valid pointer")
			} else {
				if e.data != 5 {
					t.Errorf("Expected the value 5 to be found at index 5, got %d", e.data)
				}
			}
		})

		t.Run("with non existing value", func(t *testing.T) {
			e := a.GetWhere(10)

			if e != nil {
				t.Errorf("Expected element with value 10 not to be found in list, got %p which points to node with value %d", e, e.data)
			}
		})
	})
}

func TestAdd(t *testing.T) {
	t.Run("using Init() call", func(t *testing.T) {
		a := SLinkedList{}
		a.Init(0)
		expected := generateStringForm([]int{0})
		if as := a.String(); as != expected {
			t.Errorf("Expected initialized list to be \"%s\", got \"%s\"", expected, as)
		}
	})

	t.Run("using AddHead() call", func(t *testing.T) {
		a := SLinkedList{}
		a.AddHead(1)
		a.AddHead(2)
		a.AddHead(3)
		expected := generateStringForm([]int{3, 2, 1})
		if as := a.String(); as != expected {
			t.Errorf("Expected final list to be \"%s\", got \"%s\"", expected, as)
		}
	})

	t.Run("using AddTail() call", func(t *testing.T) {
		a := SLinkedList{}
		a.AddTail(1)
		a.AddTail(2)
		a.AddTail(3)
		expected := generateStringForm([]int{1, 2, 3})
		if as := a.String(); as != expected {
			t.Errorf("Expected final list to be \"%s\", got \"%s\"", expected, as)
		}
	})

	t.Run("using AddAt() call", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			a := SLinkedList{}
			t.Run("with negative index", func(t *testing.T) {
				index, val := -1, 100
				neg := a.AddAt(index, val)
				if neg != nil {
					t.Errorf("Expected nil, got %p pointing to %d", neg, neg.data)
				}
			})

			t.Run("with zero index", func(t *testing.T) {
				index, val := 0, 100
				zero := a.AddAt(index, val)
				if zero != nil {
					t.Errorf("Expected nil, got %p pointing to %d", zero, zero.data)
				}
			})

			t.Run("with positive index", func(t *testing.T) {
				index, val := 1, 100
				pos := a.AddAt(index, val)
				if pos != nil {
					t.Errorf("Expected nil, got %p pointing to %d", pos, pos.data)
				}
			})
		})

		t.Run("on non-empty list", func(t *testing.T) {
			t.Run("with negative index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index, val := -1, 100
				res := a.AddAt(index, val)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d. List: %s", res, res.data, a.String())
				}
			})

			t.Run("with valid index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index, val := 5, 100
				res := a.AddAt(index, val)
				if res == nil {
					t.Errorf("Expected a pointer of inserted element, got nil")
				} else {
					if res.data != 100 {
						t.Errorf("Expected to receive a pointer to element with value 100, got %d", res.data)
					}
					if as, es := a.String(), generateStringForm([]int{0, 1, 2, 3, 4, 5, 100, 6, 7, 8, 9}); as != es {
						t.Errorf("Expected final list to be \"%s\", got \"%s\"", es, as)
					}
				}
			})

			t.Run("with out of range index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index, val := 50, 100
				res := a.AddAt(index, val)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d. List: %s", res, res.data, a.String())
				}
			})
		})
	})
}

func TestRemove(t *testing.T) {
	t.Run("using RemoveHead() call", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			a := SLinkedList{}
			res := a.RemoveHead()
			if res != nil {
				t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
			}
		})

		t.Run("on non-empty list", func(t *testing.T) {
			a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			a.RemoveHead()
			a.RemoveHead()
			a.RemoveHead()
			if as, exp := a.String(), generateStringForm([]int{3, 4, 5, 6, 7, 8, 9}); as != exp {
				t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, as)
			}
		})
	})

	t.Run("using RemoveTail() call", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			a := SLinkedList{}
			res := a.RemoveTail()
			if res != nil {
				t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
			}
		})

		t.Run("on non-empty list", func(t *testing.T) {
			a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			a.RemoveTail()
			a.RemoveTail()
			a.RemoveTail()
			if res, exp := a.String(), generateStringForm([]int{0, 1, 2, 3, 4, 5, 6}); res != exp {
				t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, res)
			}
		})
	})

	t.Run("using RemoveAt() call", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			t.Run("with negative index", func(t *testing.T) {
				a := SLinkedList{}
				index := -1
				res := a.RemoveAt(index)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
				}
			})

			t.Run("with zero index", func(t *testing.T) {
				a := SLinkedList{}
				index := 0
				res := a.RemoveAt(index)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
				}
			})

			t.Run("with positive index", func(t *testing.T) {
				a := SLinkedList{}
				index := 1
				res := a.RemoveAt(index)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
				}
			})
		})

		t.Run("on non-empty list", func(t *testing.T) {
			t.Run("with negative index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index := -5
				res := a.RemoveAt(index)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
				}
			})

			t.Run("with valid index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index := 5
				res := a.RemoveAt(index)
				if res == nil {
					t.Errorf("Expected a pointer to removed element, got nil. List: %s", a.String())
				} else {
					if as, exp := a.String(), generateStringForm([]int{0, 1, 2, 3, 4, 6, 7, 8, 9}); as != exp {
						t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, as)
					}
				}
			})

			t.Run("with out of range index", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				index := 10
				res := a.RemoveAt(index)
				if res != nil {
					t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
				}
			})
		})
	})

	t.Run("using RemoveWhere() call", func(t *testing.T) {
		t.Run("on empty list", func(t *testing.T) {
			a := SLinkedList{}
			res := a.RemoveWhere(100)
			if res != nil {
				t.Errorf("Expected nil, got %p pointing to %d", res, res.data)
			}
		})

		t.Run("on non-empty list", func(t *testing.T) {
			t.Run("with existing element", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				a.RemoveWhere(5)
				if as, exp := a.String(), generateStringForm([]int{0, 1, 2, 3, 4, 6, 7, 8, 9}); as != exp {
					t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, as)
				}
			})

			t.Run("with non-existing element", func(t *testing.T) {
				a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
				a.RemoveWhere(10)
				if as, exp := a.String(), generateStringForm([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}); as != exp {
					t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, as)
				}
			})
		})
	})

	t.Run("using Clear() call", func(t *testing.T) {
		a := generateTestList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		a.Clear()
		if as, exp := a.String(), "head -> nil"; as != exp {
			t.Errorf("Expected list to be \"%s\", got \"%s\"", exp, as)
		}
	})
}

func TestReverse(t *testing.T) {
	t.Run("on empty list", func(t *testing.T) {
		a := SLinkedList{}
		a.Reverse()
		if a.String() != "head -> nil" {
			t.Errorf("Expected list to be \"head -> nil\", got \"%s\"", a.String())
		}
	})

	t.Run("on non-empty list", func(t *testing.T) {
		inputs := [][]int{
			{0},
			{0, 1},
			{0, 1, 2},
			{0, 1, 2, 3},
			{0, 1, 2, 3, 4},
			{0, 1, 2, 3, 4, 5},
		}
		expectedOutputs := [][]int{
			{0},
			{1, 0},
			{2, 1, 0},
			{3, 2, 1, 0},
			{4, 3, 2, 1, 0},
			{5, 4, 3, 2, 1, 0},
		}

		for i := 0; i < len(inputs); i++ {
			inputList := generateTestList(inputs[i])
			initial := inputList.String()
			expectedString := generateStringForm(expectedOutputs[i])
			inputList.Reverse()
			if res := inputList.String(); res != expectedString {
				t.Errorf("Expected the reverse of \"%s\" to be \"%s\", got \"%s\"", initial, expectedString, res)
			}
		}
	})
}
