package xdr

import "fmt"

// ValidateExact checks that v holds exactly one well-formed value: it is
// ValidateFull, and it also fails if any bytes follow the value. It walks v
// once.
func ValidateExact[T View](v T) error {
	n, err := v.valid(0)
	if err != nil {
		return err
	}
	if n != len(v) {
		return fmt.Errorf("xdr view: %d bytes after the value", len(v)-n)
	}
	return nil
}
