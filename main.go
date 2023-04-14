package main

import (
	"net/http"

	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"systementor.se/yagolangapi/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

var theRandom *rand.Rand

func start(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", &PageView{Title: "test", Rubrik: "Hej Golang"})
}

// HTML
// JSON
/*
func patrik(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "Patrik Iden", "city": "Stockholm"})
}
*/
func userName(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "patrik":
		c.JSON(http.StatusOK, gin.H{
			"name": "Patrik Iden",
			"city": "Stockholm",
		})
	case "stefan":
		c.JSON(http.StatusOK, gin.H{
			"name": "Stefan Holmberg",
			"city": "Stockholm",
		})
		// Add more name alternatives here!
	default:
		c.String(http.StatusNotFound, "Error! Could not found name: %s", name)
	}
}

func employeesJson(c *gin.Context) {
	var employees []data.Employee
	data.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func addEmployee(c *gin.Context) {

	data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})

}

func addManyEmployees(c *gin.Context) {
	//Here we create 10 Employees
	for i := 0; i < 10; i++ {
		data.DB.Create(&data.Employee{Age: theRandom.Intn(50) + 18, Namn: randomdata.FirstName(randomdata.RandomGender), City: randomdata.City()})
	}

}

var config Config

func main() {
	theRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	readConfig(&config)

	data.InitDatabase(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := gin.Default()

	router.LoadHTMLGlob("templates/**")
	router.GET("/", start)

	api := router.Group("/api")
	api.GET("/employees", employeesJson)
	api.GET("/addemployee", addEmployee)
	api.GET("/addmanyemployees", addManyEmployees)
	// api.GET("/patrik", patrik)
	api.GET("/test", func(c *gin.Context) { c.String(http.StatusOK, "Hej") })
	api.GET("/:name", userName)

	router.Run(":8080")

	// e := data.Employee{
	// 	Age:  1,
	// 	City: "Strefabn",
	// 	Namn: "wddsa",
	// }

	// if e.IsCool() {
	// 	fmt.Printf("Namn is cool:%s\n", e.Namn)
	// } else {
	// 	fmt.Printf("Namn:%s\n", e.Namn)
	// }

	// fmt.Println("Hello")
	// t := tabby.New()
	// t.AddHeader("Namn", "Age", "City")
	// t.AddLine("Stefan", "50", "Stockholm")
	// t.AddLine("Oliver", "14", "Stockholm")
	// t.Print()
}
