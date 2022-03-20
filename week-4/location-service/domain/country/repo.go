package country

import "gorm.io/gorm"

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (c *CountryRepository) GetAllCountriesWithCityInformation() ([]Country, error) {
	var countries []Country
	result := c.db.Preload("Cities").Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}

func (c *CountryRepository) GetCountryWithName(name string) (*Country, error) {
	var country *Country
	result := c.db.
		Where(Country{Name: name}).
		Attrs(Country{Code: "NULL", Name: "NULL"}).
		FirstOrInit(&country) // Eğer sorgu sonucunda veri bulunursa Attrs kısmında yazılanlar ignore edilir.

	if result.Error != nil {
		return nil, result.Error
	}

	return country, nil
}

func (c *CountryRepository) GetAllCountriesWithoutCityInformation() ([]Country, error) {
	var countries []Country
	result := c.db.Find(&countries)

	if result.Error != nil {
		return nil, result.Error
	}

	return countries, nil
}

func (c *CountryRepository) Migration() {
	c.db.AutoMigrate(&Country{})
}

func (c *CountryRepository) InsertSampleData() {
	cities := []Country{
		{Name: "Türkiye", Code: "TR"},
		{Name: "Amerika", Code: "US"},
	}

	for _, city := range cities {
		c.db.Create(&city)
	}
}
