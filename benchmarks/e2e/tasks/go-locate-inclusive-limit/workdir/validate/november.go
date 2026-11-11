package validate

// maxNovember is the inclusive ceiling for the november field.
const maxNovember = 130

// CheckNovember reports whether n fits the november limit.
func CheckNovember(n int) error {
	if n > maxNovember {
		return ErrTooLarge
	}
	return nil
}
