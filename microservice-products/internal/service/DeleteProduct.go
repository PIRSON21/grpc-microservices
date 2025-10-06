package service

import "github.com/sirupsen/logrus"

// DeleteProduct deletes a product by its ID.
func (s *ProductService) DeleteProduct(id string) error {
	log := logrus.WithFields(logrus.Fields{
		"op": "service.ProductService.DeleteProduct",
		"id": id,
	})

	log.Debug("Deleting product")

	// Delete the product
	err := s.repo.DeleteProduct(id)
	if err != nil {
		log.Errorf("Error deleting product: %v", err)
		return err
	}

	log.Infof("Product with ID %s deleted successfully", id)
	return nil
}
