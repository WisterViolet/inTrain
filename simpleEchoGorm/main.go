// go:build ignore
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id" gorm:"column:id;primarykey;not null;autoIncrement"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	Age       int       `json:"age" gorm:"column:age;not null"`
	Height    int       `json:"height" gorm:"column:height;not null"`
	Weight    float64   `json:"weight" gorm:"column:weight;not null"`
	BirthDate time.Time `json:"birthdate" gorm:"column:birthdate;not null;type:date"`
	Color     string    `json:"color" gorm:"column:color;not null"`
}

func main() {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := os.Getenv("DBHOST")
	dbName := os.Getenv("DBNAME")
	dbOptions := os.Getenv("DBOPTIONS")
	dsn := fmt.Sprintf("%s:%s@%s/%s?%s", dbUser, dbPass, dbHost, dbName, dbOptions)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	DB.AutoMigrate(&User{})
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		var u User
		DB.Where("id = ?", id).First(&u)
		userJSON, err := json.Marshal(u)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error")
		}
		return c.String(http.StatusOK, string(userJSON))
	})
	e.Logger.Fatal(e.Start(":8080"))

}
