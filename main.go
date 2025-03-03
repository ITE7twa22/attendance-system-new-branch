package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/rs/cors"
	"google.golang.org/api/option"
)

// Firestore collection name
const outOfOrg_FirestoreCollection = "outOfOrg"
const volunteers_FirestoreCollection = "Volunteers"

// A global variable to store the number of volunteers
var volunteerCount int

// RequestBody represents the structure of the incoming JSON request body
type RequestBody struct {
	VolunteerID int `json:"volunteerID"`
	NationalId  int `json:"nationalId"`
}

// ResponseData represents the structure of the JSON response body
type ResponseData struct {
	VolunteerID            int       `json:"Code"`
	NationalId             int       `json:"newVolunteerId"`
	RelationsVolunteerName string    `json:"RelationsVolunteerName"`
	Gender                 string    `json:"Gender"`
	LoginDateTime          time.Time `json:"LoginDateTime"`
	LogoutDateTime         time.Time `json:"LogoutDateTime"`
}

// Volunteer represents the structure of a volunteer document
type Volunteer struct {
	NationalId int    `json:"NationalId"`
	Name       string `json:"Name"`
	Phone      string `json:"Phone"`
	Gender     string `json:"Gender"`
}

// Function to enable CORS for the response
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// Function to handle preflight requests
func handlePreflight(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

// Function to decode the JSON request body into the RequestBody struct
func decodeRequestBody(r *http.Request) (RequestBody, error) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	return requestBody, err
}

// Function to decode the JSON request body into the responseData struct
func decodeResponsetBody(r *http.Request) (ResponseData, error) {
	var requestBody ResponseData
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	return requestBody, err
}

// Function to initialize the Firebase app and Firestore client
func initFirebase() (*firestore.Client, error) {
	opt := option.WithCredentialsFile("volunteersdata-cf17b-firebase-adminsdk-fbsvc-f1d035e292.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Function to increment the volunteer count
func incrementVolunteerCount() {
	volunteerCount++
}

// Function to decrement the volunteer count
func decrementVolunteerCount() {
	volunteerCount--
}

// Handler function to return the current volunteer count
func getVolunteerCountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responseData := map[string]int{"VolunteerCount": volunteerCount}
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func decodeVolunteer(r *http.Request) (Volunteer, error) {
	var newVolunteer Volunteer
	err := json.NewDecoder(r.Body).Decode(&newVolunteer)
	return newVolunteer, err
}

// This function adds new volunteer's by National ID to the outOfOrg Collection -- Done by Nada
func addNewVolunteer(w http.ResponseWriter, r *http.Request) {
	log.Println("1- In Add New Volunteer Function")

	// Enable CORS
	enableCORS(w)

	// Handle preflight OPTIONS request
	if handlePreflight(w, r) {
		return
	}

	// Decode the request body
	newVolunteer, err := decodeVolunteer(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Initialize Firebase client
	client, err := initFirebase()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Reference to the Firestore collections for volunteers
	log.Println("2- Referencing the Firestore collections to volunteers")

	collection1 := client.Collection(outOfOrg_FirestoreCollection)
	collection2 := client.Collection(volunteers_FirestoreCollection)

	// Check if the volunteer already exists in either collection
	if exists, err := checkIfVolunteerExists(collection1, "NationalId", newVolunteer.NationalId); err != nil || exists {
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Error(w, "Volunteer with the provided NationalId already exists", http.StatusBadRequest)
			log.Println("Volunteer with the provided NationalId already exists in collection 1")

		}
		return
	}

	if exists, err := checkIfVolunteerExists(collection2, "NationalID", newVolunteer.NationalId); err != nil || exists {
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			http.Error(w, "Volunteer with the provided NationalId already exists", http.StatusBadRequest)
			log.Println("Volunteer with the provided NationalId already exists in collection 2")
		}
		return
	}

	// Create a new document in the collection with the volunteer data
	log.Println("4- Creating a new document with the volunteer data")

	_, _, err = collection1.Add(context.Background(), map[string]interface{}{
		"NationalId": newVolunteer.NationalId,
		"Name":       newVolunteer.Name,
		"Phone":      newVolunteer.Phone,
		"Gender":     newVolunteer.Gender,
		"fromE7twaa": false,
		// Add other fields as needed...
	})
	if err != nil {
		http.Error(w, "Error adding volunteer to collection", http.StatusInternalServerError)
		return
	}
	log.Printf("New Volunteer: %+v", newVolunteer)

	// Return success response if the volunteer is added successfully
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Volunteer added successfully"))
}

// Function to search collections if volunteer already exists -- Done by Nada
func checkIfVolunteerExists(collection *firestore.CollectionRef, field string, nationalId int) (bool, error) {

	log.Println("3- searching collections if volunteer already exists")

	query := collection.Where(field, "==", nationalId).Limit(1)
	snapshot, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return false, err
	}
	return len(snapshot) > 0, nil
}

// Search for volunteer in all collections -- Done by Nada
func handleFirestoreQueries(w http.ResponseWriter, query1, query2, query3 firestore.Query) {
	ctx := context.Background()
	log.Println("Enabled Handle Firestore Queries Function")
	// querySnapshot4, err := query4.Documents(ctx).GetAll()
	// if len(querySnapshot4) > 0 {
	// 	http.Error(w, fmt.Sprintf("Code not active"), http.StatusRequestTimeout)
	// 	return
	// }

	// Query the third collection by (id)
	querySnapshot3, err := query3.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error querying Firestore (third collection): %v", err) // Detailed log
		http.Error(w, fmt.Sprintf("Error querying Firestore (third collection): %v", err), http.StatusInternalServerError)
		return
	}

	if len(querySnapshot3) > 0 {
		var data map[string]interface{}
		if err := querySnapshot3[0].DataTo(&data); err != nil {
			log.Printf("Error converting Firestore data (third collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error converting Firestore data (third collection): %v", err), http.StatusInternalServerError)
			return
		}

		relationsVolunteerName, ok := data["ArabicName"].(string)
		if !ok || relationsVolunteerName == "" {
			log.Printf("Invalid or missing ArabicName field (third collection): %#v", data)
			http.Error(w, "Invalid or missing ArabicName field", http.StatusInternalServerError)
			return
		}

		log.Println("Name from third collection:", relationsVolunteerName)
		responseData := ResponseData{
			RelationsVolunteerName: relationsVolunteerName,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			log.Printf("Error encoding JSON response (third collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
			return
		}
		return
	}

	// Query the second collection by (NationalID)
	querySnapshot2, err := query2.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error querying Firestore (second collection): %v", err) // Detailed log
		http.Error(w, fmt.Sprintf("Error querying Firestore (second collection): %v", err), http.StatusInternalServerError)
		return
	}

	if len(querySnapshot2) > 0 {
		var data map[string]interface{}
		if err := querySnapshot2[0].DataTo(&data); err != nil {
			log.Printf("Error converting Firestore data (second collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error converting Firestore data (second collection): %v", err), http.StatusInternalServerError)
			return
		}

		relationsVolunteerName, ok := data["ArabicName"].(string)
		if !ok || relationsVolunteerName == "" {
			log.Printf("Invalid or missing ArabicName field (second collection): %#v", data)
			http.Error(w, "Invalid or missing ArabicName field", http.StatusInternalServerError)
			return
		}

		volunteerIDValue, idExists := data["Code"]
		if !idExists {
			log.Printf("Missing id field in Firestore document (second collection): %#v", data) // Detailed log
			http.Error(w, "Invalid or missing id field", http.StatusInternalServerError)
			return
		}

		var volunteerID int
		switch v := volunteerIDValue.(type) {
		case int:
			volunteerID = v
		case int64:
			volunteerID = int(v)
		case float64:
			volunteerID = int(v)
		default:
			log.Printf("Invalid type for id field (second collection): %#v", volunteerIDValue) // Detailed log
			http.Error(w, "Invalid type for id field", http.StatusInternalServerError)
			return
		}

		log.Println("Name from second collection:", relationsVolunteerName)
		log.Println("ID from second collection:", volunteerID)

		responseData := ResponseData{
			VolunteerID:            volunteerID,
			RelationsVolunteerName: relationsVolunteerName,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			log.Printf("Error encoding JSON response (second collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
			return
		}
		return
	}

	// Query the first collection by (NationalId)
	querySnapshot1, err := query1.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error querying Firestore (first collection): %v", err) // Detailed log
		http.Error(w, fmt.Sprintf("Error querying Firestore (first collection): %v", err), http.StatusInternalServerError)
		return
	}

	if len(querySnapshot1) > 0 {
		var data map[string]interface{}
		if err := querySnapshot1[0].DataTo(&data); err != nil {
			log.Printf("Error converting Firestore data (first collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error converting Firestore data (first collection): %v", err), http.StatusInternalServerError)
			return
		}

		relationsVolunteerName, ok := data["Name"].(string)
		if !ok || relationsVolunteerName == "" {
			log.Printf("Invalid or missing Name field (first collection): %#v", data) // Detailed log
			http.Error(w, "Invalid or missing Name field", http.StatusInternalServerError)
			return
		}

		log.Println("Name from first collection:", relationsVolunteerName)
		responseData := ResponseData{
			RelationsVolunteerName: relationsVolunteerName,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(responseData); err != nil {
			log.Printf("Error encoding JSON response (first collection): %v", err) // Detailed log
			http.Error(w, fmt.Sprintf("Error encoding JSON response: %v", err), http.StatusInternalServerError)
			return
		}
		return
	}

	// If no data is found
	http.Error(w, "No data found", http.StatusNotFound)
}

// Get Volunteers who have a code -- Done by Nada
func getVolunteerByNumber(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCORS(w)

	// Handle preflight OPTIONS request
	if handlePreflight(w, r) {
		return
	}

	// Decode the request body
	requestBody, err := decodeRequestBody(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Firebase Client
	client, err := initFirebase()
	if err != nil {
		return
	}

	// fmt.Print("Request Body: ", requestBody.VolunteerID)
	// Process the requestBody, to get the volunteer data
	query1 := client.Collection(outOfOrg_FirestoreCollection).Where("NationalId", "==", requestBody.VolunteerID).Limit(1)
	query2 := client.Collection(volunteers_FirestoreCollection).Where("NationalID", "==", requestBody.VolunteerID).Limit(1)
	query3 := client.Collection(volunteers_FirestoreCollection).
		Where("Code", "==", requestBody.VolunteerID).
		Where("IsActive", "==", "YES").
		Limit(1)
	// query4 := client.Collection(volunteers_FirestoreCollection).
	// 	Where("Code", "==", requestBody.VolunteerID).
	// 	Where("IsActive", "==", "NO").
	// 	Limit(1)

	handleFirestoreQueries(w, query1, query2, query3)

}

func getRelationsInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS by setting the appropriate headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Decode the JSON request body
	var requestBody ResponseData
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Use the downloaded JSON key file to authenticate the Firebase app
	opt := option.WithCredentialsFile("volunteersdata-cf17b-firebase-adminsdk-fbsvc-f1d035e292.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Initialize Firestore client
	client, err := app.Firestore(context.Background())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Reference to the Firestore collection
	collection := client.Collection(volunteers_FirestoreCollection)

	// Query to get RelationsVolunteerName by VolunteerID
	query := collection.Where("Code", "==", requestBody.VolunteerID).Limit(1)
	querySnapshot, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if document with the provided VolunteerID exists
	if len(querySnapshot) == 0 {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	// Get RelationsVolunteerName from Firestore document
	var data map[string]interface{}
	if err := querySnapshot[0].DataTo(&data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	relationsVolunteerName, ok := data["ArabicName"].(string)
	if !ok || relationsVolunteerName == "" {
		log.Printf("Raw Firestore Data: %#v", data)
		http.Error(w, "Invalid or missing RelationsVolunteerName field", http.StatusInternalServerError)
		return
	}

	// Encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	responseData := ResponseData{RelationsVolunteerName: relationsVolunteerName}
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
	log.Printf("Raw Firestore Data: %#v", data)

}

// Store the Volunteer's log in data and updating the unlogged volunteer log out field to null -- Dalal/Nada
func addRelationsVolunteerHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCORS(w)

	// Handle preflight requests
	if handlePreflight(w, r) {
		return
	}

	// Decode the JSON request body
	requestBody, err := decodeResponsetBody(r)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Now you can access VolunteerID, RelationsVolunteerName, and DateTime from requestBody
	// fmt.Println("VolunteerID:", requestBody.VolunteerID)
	// fmt.Println("RelationsVolunteerName:", requestBody.RelationsVolunteerName)
	// fmt.Println("LoginDateTime:", requestBody.LoginDateTime)
	// fmt.Println("LoginDateTime:", time.Now())

	// Initialize Firestore client
	client, err := initFirebase()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()
	// Firestore setup
	newCollection := client.Collection("Attendance")

	// Set a custom document ID
	currentTime := time.Now()

	// Extract the date from the current time
	currentDate := currentTime.Format("2006-01-02")

	volunteerIDString := strconv.Itoa(requestBody.VolunteerID)

	customDocumentID := volunteerIDString + "-" + currentDate
	docRef := newCollection.Doc(customDocumentID)

	// Create a new document in the collection with the retrieved data
	_, err = docRef.Set(context.Background(), map[string]interface{}{
		"RelationsVolunteerName": requestBody.RelationsVolunteerName,
		"VolunteerID":            requestBody.VolunteerID,
		"LoginDateTime":          time.Now(),
		"LogoutDateTime":         nil,
	})
	if err != nil {
		http.Error(w, "Error adding data to new collection", http.StatusInternalServerError)
		return
	}

	incrementVolunteerCount()
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Data added to new collection successfully"))
} //dalal done

// ----------------------------------------------------------------------------

// القدرة على تسجيل الدخول بعد الساعة 12
func updateLogoutDateTimeHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handlePreflight(w, r) {
		return
	}

	requestBody, err := decodeRequestBody(r)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := initFirebase()
	if err != nil {
		log.Println("Error creating Firebase app:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Reference to the Firestore collection
	collection := client.Collection("Attendance")

	// Query to check if the volunteer has an attendance record for today
	// Get current date and time
	now := time.Now()

	// Calculate start and end times for the query
	var start, end time.Time

	//start 8 end 6 AM
	// Case 1: After midnight
	if now.Hour() >= 6 && now.Hour() <= 8 {
		start = time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
		end = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
	} else {
		// Case 2: From 8 AM to 6 AM the next day
		if now.Hour() < 6 {
			// If the current time is before 6 AM, we are in the "next day" period
			start = time.Date(now.Year(), now.Month(), now.Day()-1, 8, 0, 0, 0, time.Local)
			end = time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
		} else {
			// If the current time is after 8 AM
			start = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
			end = time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, time.Local)
		}
	}

	// Perform the query
	query := collection.Where("VolunteerID", "==", requestBody.VolunteerID).
		Where("LoginDateTime", ">=", start).
		Where("LoginDateTime", "<=", end).
		Where("LoginDateTime", ">", time.Time{})

	querySnapshot, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		log.Println("Error querying Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check if a document with the provided VolunteerID exists for today
	if len(querySnapshot) == 0 {
		log.Println("No attendance record found for VolunteerID:", requestBody.VolunteerID)
		w.Write([]byte("Volunteer did not come today"))
		return
	}

	// Calculate the duration between LogoutDateTime and LoginDateTime
	logoutTime := time.Now()
	loginTime := querySnapshot[0].Data()["LoginDateTime"].(time.Time)

	duration := logoutTime.Sub(loginTime)

	// Calculate the total seconds from the duration
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	// Construct a formatted duration string (HH:MM:SS)
	formattedDuration := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

	// Volunteer has an attendance record for today, update LogoutDateTime
	documentRef := querySnapshot[0].Ref
	_, err = documentRef.Update(context.Background(), []firestore.Update{
		{Path: "LogoutDateTime", Value: time.Now()},
		{Path: "Hours", Value: formattedDuration},
	})
	if err != nil {
		log.Println("Error updating LogoutDateTime:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	decrementVolunteerCount()
	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("LogoutDateTime updated successfully"))
}

// تحديد فترة البحث بناءً على الوقت الحالي
func determineSearchPeriod() (time.Time, time.Time) {
	now := time.Now()
	var start, end time.Time

	if now.Hour() >= 6 && now.Hour() <= 8 {
		start = time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
		end = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
	} else {
		// Case 2: From 8 AM to 6 AM the next day
		if now.Hour() < 6 {
			// If the current time is before 6 AM, we are in the "next day" period
			start = time.Date(now.Year(), now.Month(), now.Day()-1, 8, 0, 0, 0, time.Local)
			end = time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.Local)
		} else {
			// If the current time is after 8 AM
			start = time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)
			end = time.Date(now.Year(), now.Month(), now.Day()+1, 6, 0, 0, 0, time.Local)
		}
	}

	return start, end
}

// تنفيذ استعلام Firestore والتحقق من وجود سجل الحضور
func performQueryAndCheckAttendance(client *firestore.Client, volunteerID int, start, end time.Time) (bool, error) {
	collection := client.Collection("Attendance")
	query := collection.Where("VolunteerID", "==", volunteerID).
		Where("LoginDateTime", ">=", start).
		Where("LoginDateTime", "<=", end)

	querySnapshot, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		return false, err
	}

	return len(querySnapshot) > 0, nil
}

// إرسال الاستجابة اعتمادًا على وجود سجل الحضور
func sendResponseBasedOnAttendance(w http.ResponseWriter, volunteerID int, attended bool) {
	if !attended {
		log.Println("No attendance record found for VolunteerID:", volunteerID)
		w.Write([]byte("Volunteer did not come today"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Volunteer came today"))
}

// المعالج الرئيسي للبحث عن الحضور
func searchAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	if handlePreflight(w, r) {
		return
	}

	requestBody, err := decodeRequestBody(r)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := initFirebase()
	if err != nil {
		log.Println("Error creating Firebase app:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	start, end := determineSearchPeriod()
	attended, err := performQueryAndCheckAttendance(client, requestBody.VolunteerID, start, end)
	if err != nil {
		log.Println("Error querying Firestore:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	sendResponseBasedOnAttendance(w, requestBody.VolunteerID, attended)
}

//----------------------------------------------------------------------------

// Add a new route for the new handler
func main() {
	http.HandleFunc("/getRelationsInfo", getRelationsInfoHandler)
	http.HandleFunc("/getNumberHandler", getVolunteerByNumber)
	http.HandleFunc("/addNewVolunteer", addNewVolunteer)
	http.HandleFunc("/addRelationsVolunteer", addRelationsVolunteerHandler)
	http.HandleFunc("/searchAttendanceHandler", searchAttendanceHandler)
	http.HandleFunc("/updateLogoutDateTime", updateLogoutDateTimeHandler)
	http.HandleFunc("/getVolunteerCount", getVolunteerCountHandler)

	// Initialize CORS middleware
	handler := cors.Default().Handler(http.DefaultServeMux)

	// Start the server on port 8080
	log.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
