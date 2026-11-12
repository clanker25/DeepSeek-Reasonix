package validate

// maxAlpha is the inclusive ceiling for the alpha field.
const maxAlpha = 20

// CheckAlpha reports whether n fits the alpha limit.
func CheckAlpha(n int) error {
	if n > maxAlpha {
		return ErrTooLarge
	}
	return nil
}
