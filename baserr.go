// Package baserr provides a simple way to create inherited errors and check if an error is inherited from another error
// like it is done in OOP languages in the "catch" blocks for handling different types of exceptions.
package baserr

type (
	// IBase represents an interface for inherited errors.
	IBase interface {
		Error() string
		Parent() error
	}
	// Base represents an error inherited from another error of type T.
	Base[T error] struct {
		message string
		parent  T
	}
	// BaseError represents a first level error. It is used as parent for other errors.
	BaseError struct {
		Base[error]
	}
)

// Error returns an error message.
func (e Base[T]) Error() string {
	return e.message
}

// Parent returns an error of type T.
func (e Base[T]) Parent() error {
	return e.parent
}

// NewError returns an error of type T with a message.
func NewError[T error](err string) IBase {
	return &Base[T]{
		message: err,
	}
}

// InheritedFrom checks if an error is inherited from an error of type T.
func InheritedFrom[T any](err any) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(T); ok {
		return true
	}
	if val, ok := err.(IBase); ok {
		parent := val.Parent()
		return InheritedFrom[T](parent)
	}

	return false
}
