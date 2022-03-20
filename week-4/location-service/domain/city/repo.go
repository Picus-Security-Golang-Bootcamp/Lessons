package city

import (
	"errors"

	"gorm.io/gorm"
)

type CityRepository struct {
	db *gorm.DB
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{db: db}
}

func (c *CityRepository) FindAll() []City {
	var cities []City
	c.db.Find(&cities)

	return cities
}

func (c *CityRepository) FindByCountryCode(countryCode string) []City {
	var cities []City
	//c.db.Where(`"CountryCode" = ?`, countryCode).Order("Id desc,name").Find(&cities)

	//Struct
	//c.db.Where(&City{CountryCode: countryCode}).First(&cities)
	//c.db.Where(map[string]interface{}{"CountryCode": countryCode, "code": "01"}).First(&cities)
	//c.db.Where([]int64{20,18,17}).Find(&cities) // ID IN(20,18,17)
	return cities
}

func (c *CityRepository) FindByCountryCodeOrCityCode(code string) []City {
	var cities []City
	c.db.Where(`"CountryCode = ?"`, code).Or("code = ?", code).Find(&cities)

	return cities
}

func (c *CityRepository) FindByName(name string) []City {
	var cities []City
	c.db.Where("name LIKE ? ", "%"+name+"%").Find(&cities)

	return cities
}

func (c *CityRepository) FindByNameWithRawSQL(name string) []City {
	var cities []City
	c.db.Raw("SELECT * FROM City WHERE name LIKE ?", "%"+name+"%").Scan(&cities)

	return cities
}

func (c *CityRepository) GetByID(id int) (*City, error) {
	var city City
	result := c.db.First(&city, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &city, nil
}

func (c *CityRepository) Create(city City) error {
	result := c.db.Create(city)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CityRepository) Update(city City) error {
	result := c.db.Save(city)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CityRepository) Delete(city City) error {
	result := c.db.Delete(city)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CityRepository) DeleteById(id int) error {
	result := c.db.Delete(&City{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CityRepository) Migrations() {
	c.db.AutoMigrate(&City{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

func (c *CityRepository) InsertSampleData() {
	cities := []City{
		{Name: "Adana", Code: "01", CountryCode: "TR"},
		{Name: "Adıyaman", Code: "02", CountryCode: "TR"},
		{Name: "Ankara", Code: "06", CountryCode: "TR"},
		{Name: "İstanbul", Code: "34", CountryCode: "TR"},
		{Name: "İzmir", Code: "35", CountryCode: "TR"},
	}

	for _, city := range cities {
		c.db.Where(City{Code: city.Code}).
			Attrs(City{Code: city.Code, Name: city.Name}).
			FirstOrCreate(&city)
	}
}
