package main

vehicule interface {}

car []struct {
	model string, make string, typeVehicle string}


truck []struct {
	model string, make string, typeVehicle string}

bike []struct {
	model string, make string, typeVehicle string}

// Values array for the feedback.json file We will create two more structs to capture the feedback from the JSON file. For the first struct, find the comment // Values array for the feedback.json file and, on the very next line define a new struct named Values.
Values []struct {
	// VehicleType is the type of vehicle
	VehicleType string `json:"vehicleType"`
	// VehicleModel is the model of the vehicle
	VehicleModel string `json:"vehicleModel"`
	// VehicleMake is the make of the vehicle
	VehicleMake string `json:"vehicleMake"`
	// Rating is the rating of the vehicle
	Rating int `json:"rating"`
} `json:"values"`
}

// Model array for the feedback.json file As part of Values,
//define one field named Models which is of type []Model and has the JSON tag json:"values"
//As part of Model, define two fields: one named Name which is of type string and has the JSON tag json:"model",
//and another named Feedback which is of type []string with JSON tag json:"feedback".

Model []struct {
	// Name is the name of the model
	Name string `json:"model"`
	// Feedback is the feedback for the model
	Feedback []string `json:"feedback"`
} `json:"values"`
}
//Define feedback struct and fields
//Create another struct named feedbackResult and assign it four fields: feedbackTotal, feedbackPositive, feedbackNegative, and feedbackNeutral. All field type are int.

feedbackResult struct {
	feedbackTotal int
	feedbackPositive int
	feedbackNegative int
	feedbackNeutral int

}

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
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

	// Print ratings for the different vehicles
}

/*
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
*/
