package order

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error

type OrderDB struct {
	db *gorm.DB
}

var orderrepository *OrderDB

func OrderDBInit() {
	orderrepository = &OrderDB{}
	orderrepository.db, err = gorm.Open(sqlite.Open("Order_table.db"), &gorm.Config{})

	if err != nil {
		panic("DB Connection Error")
	}
	orderrepository.db.AutoMigrate(&Order{})

}

func OrderRepository() *OrderDB {
	return orderrepository
}

func (self *OrderDB) save(entity interface{}) {

	err := self.db.Create(entity)

	if err != nil {
		log.Print(err)
	}
}

func (self *OrderDB) GetList() []Order {

	entities := []Order{}
	self.db.Find(&entities)

	return entities
}

func (self *OrderDB) GetID(id int) *Order {
	entity := &Order{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *OrderDB) Delete(entity *Order) error {
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *OrderDB) Update(id int, params map[string]string) error {
	entity := &Order{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	} else {
		update := &Order{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}
