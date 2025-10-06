package fake

// DeleteProduct is an implement of repository.PostProduct interface
func (f *Fake) DeleteProduct(_ string) error {
	return nil
}
