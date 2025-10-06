package fake

type Fake struct{}

func NewFakeProductRepository() *Fake {
	return &Fake{}
}

// Close implements the repository.ProductRepository interface.
func (f *Fake) Close() error {
	return nil
}
