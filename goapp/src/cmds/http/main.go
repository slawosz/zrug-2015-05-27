package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-gorp/gorp"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name  string  `db:"name"`
	Price float64 `db:"price"`

	Attributes []Attribute `json:",omitempty"`
}

type Attribute struct {
	Key       string `db:"key"`
	Value     string `db:"value"`
	ProductId int64  `db:"product_id"`
}

var dbmap *gorp.DbMap

func main() {
	// setup db
	db, err := sql.Open("mysql", "root:@/railsapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()

	// setup http
	r := mux.NewRouter()
	r.HandleFunc("/products/{product_id}", show)
	http.Handle("/", r)

	http.HandleFunc("/products", index)

	// start http
	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	var products []Product
	_, err := dbmap.Select(&products, "select p.name, p.price from products p order by p.id")
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b))
}

func show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["product_id"]

	var attributes []Attribute
	_, err := dbmap.Select(&attributes, "select a.key, a.value, a.product_id from products p inner join attributes a on a.product_id = p.id where p.id = ?", product_id)
	if err != nil {
		panic(err)
	}

	var product *Product
	err = dbmap.SelectOne(&product, "select p.name, p.price from products p where p.id = ?", product_id)
	if err != nil {
		panic(err)
	}
	product.Attributes = attributes
	b, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b))
}
