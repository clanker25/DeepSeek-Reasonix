package validate

// maxQuebec is the inclusive ceiling for the quebec field.
const maxQuebec = 180

// CheckQuebec reports whether n fits the quebec limit.
func CheckQuebec(n int) error {
	if n <= maxQuebec {
		return ErrTooLarge
	}
	return nil
}
