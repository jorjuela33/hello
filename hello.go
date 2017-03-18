package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		db, err := gorm.Open("mssql", "server=181.49.12.194;database=BD_TEMP;user id=sa;password=Qu4l1ty;port=1433")
		if err == nil {
			fmt.Println("connected")
		} else {
			fmt.Println(err)
		}

		var products []Product
		db.Table("alm_insumos").Scan(&products)
		json.NewEncoder(w).Encode(products)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(""+port, nil)
}

type Product struct {
	gorm.Model
	Name string `gorm:"column:Nomins"`
}

func (Product) TableName() string {
	return "alm_insumos"
}
