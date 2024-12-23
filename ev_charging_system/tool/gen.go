package tool

import (
	"log"

	"fmt"

	"ev_charging_system/model"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenDao() {
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dao", // output directory, default value is ./query
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	format := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(format, "lsh", "lsh666hh", "zhoupb.com:33060", "db_charging")
	// Initialize a *gorm.DB instance
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(dbs)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(model.FeeRule{}, model.Gun{}, model.Parameter{}, model.Pile{}, model.RepairRequest{}, model.Repairman{}, model.Station{})

	// Generate default DAO interface for those generated structs from database
	// companyGenerator := g.GenerateModelAs("company", "MyCompany")
	// g.ApplyBasic(
	// 	g.GenerateModel("users"),
	// 	companyGenerator,
	// 	g.GenerateModelAs("people", "Person",
	// 		gen.FieldIgnore("deleted_at"),
	// 		gen.FieldNewTag("age", `json:"-"`),
	// 	),
	// )

	// Execute the generator
	g.Execute()
}
