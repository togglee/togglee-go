package drivers

// ReleaseDriver -- release toggles driver resolver
type ReleaseDriver struct{ }

// isActive -- release toggles driver resolver function
func (h ReleaseDriver) IsActive(name string) *bool {
	panic("Not implemented")
	return nil
}
