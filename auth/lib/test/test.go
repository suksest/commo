package test

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

// SetupDBTests for mocked gorm package
// WIP
func SetupDBTests() (*gorm.DB, error) {
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	// GORM
	db, err := gorm.Open("mysql", "shop:shop@/shop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	return db, nil
}
