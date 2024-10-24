package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Dharmadurai95/channelModule/helpers"
)

type FoodData struct {
	Hotel_name  string   `json:"hotel_name"`
	Place       string   `json:"place"`
	Image_url   string   `json:"image_url"`
	Description string   `json:"description"`
	Rating      float64  `json:"rating"`
	Tags        []string `json:"tags"`
	Vegetarian  bool     `json:"vegetarian"`
	Price       float64  `json:"price"`
}

func randomNumChan(c chan int) {
	number := helpers.GenerateRandomNum(10)
	c <- number
}

func main() {

	intChan := make(chan int)
	defer close(intChan)
	go randomNumChan(intChan)
	fmt.Println(<-intChan)

	data, _ := jsonData()
	log.Println(data, "dharma")

}

func jsonData() ([]FoodData, error) {
	data := `[
		{
			"hotel_name": "Hotel Bella Vista",
			"place": "New York",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/qhyrbmiouhwzkngbqpie",
			"description": "Delicious Margherita Pizza with fresh tomatoes and basil",
			"rating": 4.5,
			"tags": ["pizza", "Italian", "Margherita"],
			"vegetarian": true,
			"price": 10.99
		},
		{
			"hotel_name": "The Wing Stop",
			"place": "Chicago",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/ed4ab9e836ffbea562154d5c086d39d6",
			"description": "Spicy Chicken Wings with BBQ sauce",
			"rating": 4.2,
			"tags": ["wings", "chicken", "spicy"],
			"vegetarian": false,
			"price": 8.49
		},
		{
			"hotel_name": "Green Leaf Cafe",
			"place": "San Francisco",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/e416a07fb12739b91164126d1dc247f8",
			"description": "Classic Caesar Salad with fresh romaine lettuce and dressing",
			"rating": 4.7,
			"tags": ["salad", "Caesar", "healthy"],
			"vegetarian": true,
			"price": 6.99
		},
		{
			"hotel_name": "Burger Palace",
			"place": "Los Angeles",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/418e7dc325cf65d7c1d9faab2e646421",
			"description": "Authentic Beef Burger with cheese and pickles",
			"rating": 4.4,
			"tags": ["burger", "beef", "cheese"],
			"vegetarian": false,
			"price": 9.99
		},
		{
			"hotel_name": "Sushi Cove",
			"place": "Seattle",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/csao7mzsg8trda8ede9w",
			"description": "Fresh Sushi Roll with salmon, avocado, and rice",
			"rating": 4.9,
			"tags": ["sushi", "Japanese", "salmon"],
			"vegetarian": false,
			"price": 12.99
		},
		{
			"hotel_name": "Green Leaf Cafe",
			"place": "San Francisco",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/55de0860b7ddbc2d79f8e5c8fa94e899",
			"description": "Vegetable Stir-Fry with tofu and mixed vegetables",
			"rating": 4.3,
			"tags": ["stir-fry", "vegetarian", "tofu"],
			"vegetarian": true,
			"price": 7.99
		},
		{
			"hotel_name": "Pasta Paradise",
			"place": "Miami",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/3aade67d243274c32f49c8b52d127ab2",
			"description": "Creamy Penne Alfredo with parmesan cheese",
			"rating": 4.6,
			"tags": ["pasta", "Alfredo", "creamy"],
			"vegetarian": true,
			"price": 8.99
		},
		{
			"hotel_name": "Steak House",
			"place": "New York",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/9de2675e3aa5c9dc295a2db60dcdcba9",
			"description": "Grilled Steak with roasted potatoes and asparagus",
			"rating": 4.8,
			"tags": ["steak", "grilled", "meat"],
			"vegetarian": false,
			"price": 15.99
		},
		{
			"hotel_name": "Green Leaf Cafe",
			"place": "San Francisco",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/c2ac2ddde82f7f83fc92064f98ba8b43",
			"description": "Refreshing Watermelon Salad with feta cheese",
			"rating": 4.5,
			"tags": ["salad", "watermelon", "cheese"],
			"vegetarian": true,
			"price": 5.99
		},
		{
			"hotel_name": "Wok n' Roll",
			"place": "Los Angeles",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/csao7mzsg8trda8ede9w",
			"description": "Fried Rice with shrimp and vegetables",
			"rating": 4.2,
			"tags": ["rice", "fried", "shrimp"],
			"vegetarian": false,
			"price": 9.49
		},
		{
			"hotel_name": "Pizza Haven",
			"place": "Chicago",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/423c5dde006571290e983f695ab2c8aa",
			"description": "Cheesy Garlic Breadsticks with marinara sauce",
			"rating": 4.7,
			"tags": ["breadsticks", "garlic", "cheesy"],
			"vegetarian": true,
			"price": 4.99
		},
		{
			"hotel_name": "Taco Time",
			"place": "Seattle",
			"image_url": "https://media-assets.swiggy.com/swiggy/image/upload/fl_lossy,f_auto,q_auto,w_1024/49987621bc4a080044b2c9bd22937a51",
			"description": "Spicy Tofu Tacos with avocado and salsa",
			"rating": 4.3,
			"tags": ["tacos", "tofu", "spicy"],
			"vegetarian": true,
			"price": 7.49
		}
	]
	   
	`

	var unmarshall []FoodData
	err := json.Unmarshal([]byte(data), &unmarshall)
	if err != nil {
		return unmarshall, errors.New("Unable to unmarshall data")
		// log.Println("Unmashalling data error", err)
	}
	// log.Println(unmarshall)

	//marshal
	var newFootdata FoodData
	var newFootdata1 FoodData
	var jsonSlice []FoodData
	newFootdata.Description = "test description data"
	newFootdata.Hotel_name = "belal"
	newFootdata.Image_url = "https:google.com"
	newFootdata.Place = "chennai"
	newFootdata.Price = 123.321
	newFootdata.Tags = []string{"dall", "rice", "lemon-rice"}
	newFootdata.Vegetarian = true
	newFootdata.Rating = 7.5

	newFootdata1.Description = "test one description data"
	newFootdata1.Hotel_name = "thalappakatti"
	newFootdata1.Image_url = "https:google.com/api/v1"
	newFootdata1.Place = "trichy"
	newFootdata1.Price = 183.321
	newFootdata1.Tags = []string{"briyani", "chicken rice", "motto fry"}
	newFootdata1.Vegetarian = false
	newFootdata1.Rating = 6.5
	jsonSlice = append(jsonSlice, newFootdata)
	jsonSlice = append(jsonSlice, newFootdata1)

	//stringify the json object
	stringifyJson, err := json.MarshalIndent(jsonSlice, "", "")
	if err != nil {
		log.Println("unable to marshal data", err)
	}
	log.Println(string(stringifyJson))
	return unmarshall, nil
}
