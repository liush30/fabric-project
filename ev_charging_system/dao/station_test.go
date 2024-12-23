package dao

import (
	"context"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestStation(t *testing.T) {
	format := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(format, "lsh", "lsh666hh", "zhoupb.com:33060", "db_charging")
	// Initialize a *gorm.DB instance
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	q := Use(dbs)
	d, err := q.Gun.WithContext(context.Background()).Find()
	if err != nil {
		t.Error(err)
	}
	t.Log(d)
}
