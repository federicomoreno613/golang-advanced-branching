package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type vehicle interface {
}

type car struct {
	model       string
	make        string
	typeVehicle string
}

type truck struct {
	model       string
	make        string
	typeVehicle string
}

type bike struct {
	model string
	make  string
}

// Values array for the feedback.json file
type Values struct {
	Models []Model `json:"values"`
}

// Model array for the feedback.json file
type Model struct {
	Name     string   `json:"model"`
	Feedback []string `json:"feedback"`
}

// Create another struct named feedbackResult and assign it four fields: feedbackTotal, feedbackPositive, feedbackNegative, and feedbackNeutral. All field type are int.
type feedbackResult struct {
	feedbackTotal    int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral  int
}

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

// We will define two variables. The first one is a map named vehicleResult whose key-type is string, and whose value-type is feedbackResult. The second variable is a slice named inventory of type vehicle
var (
	vehicleResult map[string]feedbackResult
	inventory     []vehicle
)

func init() {

	inventory = []vehicle{
		bike{"FTR 1200", "Indian"},
		bike{"Iron 1200", "Harley"},
		car{"Sonata", "Hyundai", "Sedan"},
		car{"SantaFe", "Hyundai", "SUV"},
		car{"Civic", "Honda", "Hatchback"},
		car{"A5", "Audi", "Coupe"},
		car{"Mazda6", "Mazda", "Sedan"},
		car{"CRV", "Honda", "SUV"},
		car{"Camry", "Toyota", "Sedan"},
		truck{"F-150", "Ford", "Truck"},
		truck{"RAM1500", "Dodge", "Truck"}}

	vehicleResult = make(map[string]feedbackResult)

}

func main() {
	// Generate ratings for the different vehicles
	generateRating()
	// Print ratings for the different vehicles
	//for statement with no key, a value named veh, and range of inventory.
	for _, veh := range inventory {
		// Create a switch statement for the vehicle type
		switch v := veh.(type) {
		case car:
			v.carDetails()
		case bike:
			v.bikeDetails()
		case truck:
			v.truckDetails()
		default:
			fmt.Printf("Are you sure this Vehicle Type exists")
		}
	}
}

func readJSONFile() Values {
	jsonFile, err := os.Open("feedback.json")

	if err != nil {
		log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var content Values
	json.Unmarshal(byteValue, &content)

	return content
}

func generateRating() {
	// Call the function `readJSONfile()` and assign it to a variable `f`
	f := readJSONFile()
	// Check `for _, v := range f.Models`
	for _, v := range f.Models {
		// Declare variables within for stmt vehRating and vehResult
		var vehRating rating
		var vehResult feedbackResult
		// Check `for _, msg := range v.Feedback`
		for _, msg := range v.Feedback {
			// Check `if text := strings.Split(msg, " ") ; len(text) >= 5`vehRating = 5.0
			//				vehResult.feedbackTotal++
			if text := strings.Split(msg, " "); len(text) >= 5 {
				vehRating = 5.0
				vehResult.feedbackTotal++
				// Check for `for _, word := range text`
				for _, word := range text {
					switch s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r"); s {
					case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
						vehRating += extraPositive
					case "help", "helpful", "thanks", "thank you", "happy":
						vehRating += positive
					case "not helpful", "sad", "angry", "improve", "annoy":
						vehRating += negative
					case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
						vehRating += extraNegative
					}
				}

				// Create a switch statement for feedback
				switch {
				case vehRating > 8.0:
					vehResult.feedbackPositive++
				case vehRating >= 4.0 && vehRating <= 8.0:
					vehResult.feedbackNeutral++
				case vehRating < 4.0:
					vehResult.feedbackNegative++
				}
			}
		}
		// Add rating to the vehicle model
		vehicleResult[v.Name] = vehResult
	}
}

//The first task of this module is to create a function showRating() with one parameter model of type string, and no return value.

func showRating(model string) {
	ratingFound = false
	for m, r := range vehicleResult {
		if m == model {
			fmt.Printf("Total Ratings:%v\tPositive:%v\tNegative:%v\tNeutral:%v", r.feedbackTotal, r.feedbackPositive, r.feedbackNegative, r.feedbackNeutral)
			//The second statement will assign the bool value true to the variable ratingFound.
			ratingFound = true
		}
		if !ratingFound {
			fmt.Println("No rating for this vehicle")
		}

	}
}

// create carDetails() method which has a receiver of type *car named c with no parameters and has no return value. We will write two statements within this method
func (c *car) carDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Car", c.make, c.model)
	showRating(c.model)
}

func (b *bike) bikeDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Bike", b.make, b.model)
	showRating(b.model)
}

func (t *truck) truckDetails() {
	fmt.Printf("\n%-5v: %-8v: %-12v ", "Truck", t.make, t.model)
	showRating(t.model)
}
