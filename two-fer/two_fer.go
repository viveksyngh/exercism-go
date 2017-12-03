package twofer

import "fmt"

// ShareWith retruns a 2-fer string with input name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
