package dbrepository

import (
	"mongorestaurantsample-master/domain"
	//"mongorestaurantsample-master/utils"
)

type Factory struct {
}

func (f *Factory) NewUser(name string, address string, addressline2 string, url string, outcode string, postcode string, rating float32, type_of_food string) *domain.Restaurant {
	return &domain.Restaurant{
		Name:         name,
		Address:      address,
		AddressLine2: addressline2,
		URL:          url,
		Outcode:      outcode,
		Postcode:     postcode,
		Rating:       rating,
		TypeOfFood:   type_of_food,
	}

}
