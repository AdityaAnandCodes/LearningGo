package main

import (
	   "context"
	   "encoding/json"
	   "fmt"
	   "log"
	   "net/http"
	   "os"
	   "time"

	   "go.mongodb.org/mongo-driver/bson"
	   "go.mongodb.org/mongo-driver/mongo"
	   "go.mongodb.org/mongo-driver/mongo/options"
	   "github.com/joho/godotenv"
)

// User structure
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var collection *mongo.Collection

func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Replace with your MongoDB URI
	mongoURI := os.Getenv("MONGODB_URI")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ MongoDB connection failed:", err)
	}

	// Ping DB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	// Choose database and collection
	db := client.Database("Database")
	collection = db.Collection("GolangUsers")

	// Set up API
	http.HandleFunc("/user", userHandler)

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode incoming JSON
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" || user.Email == "" {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	// Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{
		"name":  user.Name,
		"email": user.Email,
		"ts":    time.Now(),
	})
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":    "User added successfully",
		"insertedId": res.InsertedID,
	})
}
