package builder

type Grouper interface {
	// Returns the name of the group by being used.
	Name() string

	// Validates that the contents of the group by.
	Validate() error
}
