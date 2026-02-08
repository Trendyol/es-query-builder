package condition

// Branch represents a conditional branch used in IfElse chains.
// Create branches using ElseIf or Else functions.
type Branch struct {
	item      any
	condition bool
	isElse    bool
}

// ElseIf creates a conditional branch for use in IfElse.
// The item is selected only if no previous condition was true and this condition is true.
//
// Parameters:
//   - condition: A boolean that determines if `item` should be selected.
//   - item: The value to return if this branch is selected.
//
// Returns:
//
//	A Branch to be passed to IfElse.
//
// Example usage:
//
//	condition.IfElse(x > 10, es.Term("foo", "bar"),
//	    condition.ElseIf(y < 20, es.Exists("fizz")),
//	)
func ElseIf(condition bool, item any) Branch {
	return Branch{condition: condition, item: item}
}

// Else creates a default branch for use in IfElse.
// The item is selected when no previous condition was true.
//
// Parameters:
//   - item: The value to return if no other condition matched.
//
// Returns:
//
//	A Branch to be passed as the last argument to IfElse.
//
// Example usage:
//
//	condition.IfElse(x > 10, es.Term("foo", "bar"),
//	    condition.ElseIf(y < 20, es.Exists("fizz")),
//	    condition.Else(es.Range("date").LesserThan("now-2h")),
//	)
func Else(item any) Branch {
	return Branch{item: item, isElse: true}
}

// IfElse evaluates conditions in order and returns the item associated with the
// first true condition. If no condition is true and no Else branch is provided,
// it returns nil.
//
// Parameters:
//   - condition: A boolean that determines if `item` should be returned.
//   - item: The map or slice to return if the condition is true.
//   - branches: Optional ElseIf and Else branches to evaluate in order.
//
// Returns:
//
//	The item from the first matching condition, the Else item, or nil.
//
// Example usage:
//
//	// Simple if
//	condition.IfElse(x > 10, es.Term("foo", "bar"))
//
//	// if + elseIf
//	condition.IfElse(x > 10, es.Term("foo", "bar"),
//	    condition.ElseIf(y < 20, es.Exists("fizz")),
//	)
//
//	// if + elseIf + else
//	condition.IfElse(x > 10, es.Term("foo", "bar"),
//	    condition.ElseIf(y < 20, es.Exists("fizz")),
//	    condition.Else(es.Range("date").LesserThan("now-2h")),
//	)
func IfElse[T ~map[string]any | ~[]any](condition bool, item T, branches ...Branch) any {
	if condition {
		return item
	}
	for _, branch := range branches {
		if branch.isElse {
			return branch.item
		}
		if branch.condition {
			return branch.item
		}
	}
	return nil
}
