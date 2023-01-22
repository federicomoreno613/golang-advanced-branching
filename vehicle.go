package main

import (
	"encoding/json"
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
					s := strings.Trim(strings.ToLower(word), " ,.,!,?,\t,\n,\r")
					// Create a switch statement for rating
					switch s {
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
