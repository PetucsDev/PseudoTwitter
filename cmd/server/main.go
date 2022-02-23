package main

import(

	_ "github.com/go-sql-driver/mysql"
		"database/sql"
		"PseudoTwitter/cmd/server/routes"
		"github.com/gin-gonic/gin"
	
)


func main(){

	db, _ := sql.Open("mysql", "root:peti96_cnc@tcp(127.0.0.1:3306)/twitter")
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(":3001"); err != nil {
		panic(err)
	}
}