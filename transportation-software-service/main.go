package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
	"github.com/sebsvt/software-prototype/transportation-software-services/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	db := client.Database("aiselena")
	return db
}

func randomName() (string, string) {
	firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis"}
	return firstNames[rand.Intn(len(firstNames))], lastNames[rand.Intn(len(lastNames))]
}

func parseCSV(filePath string) ([]aggregate.LogisticTransaction, error) {
	var transactions []aggregate.LogisticTransaction

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Assuming first row is headers, start from row 1
	for _, record := range records[1:] {
		// fmt.Println(record)
		firstName, lastName := randomName()
		transaction := aggregate.LogisticTransaction{
			TransactionID: record[1],
			Consignor: entity.Person{
				ID:          fmt.Sprintf("consignor-%s", record[1]), // Mock ID
				FirstName:   firstName,
				LastName:    lastName,
				Phone:       "123-456-7890", // Mock phone number
				AddressLine: "Mock Address", // Mock address line
				Locality:    record[3],
				City:        record[4],
				State:       record[5],
				PostalCode:  record[6],
				Country:     "Thailand",
			},
			Consignee: entity.Person{
				ID:          fmt.Sprintf("consignee-%s", record[0]), // Mock ID
				FirstName:   "Consignee",                            // Mock name
				LastName:    "Placeholder",
				Phone:       "098-765-4321", // Mock phone number
				AddressLine: "Mock 10/100 Chiang Mai Mahidon",
				Locality:    record[3],
				City:        record[4],
				State:       record[5],
				PostalCode:  record[6],
				Country:     "Thailand",
			},
			Package: entity.Item{
				ItemCode:        record[7],
				ItemName:        record[9],
				ItemDescription: record[8],
				Unit:            record[11],
				Quantity:        parseQuantity(record[10]),
			},
			Status:    "Pending",                              // Default status
			PartnerID: "7f4512af-3649-4c9d-95fa-4559329f3ef6", // Adjust based on your needs
			Timestamp: parseDate(record[2]),
			DeletedAt: nil, // Assuming no deletion info
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// Helper function to parse quantity
func parseQuantity(s string) int {
	quantity, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return quantity
}

// Helper function to parse date
func parseDate(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s) // Adjust format if needed
	return t
}

func main() {
	db := initDB()
	filePath := "./cleaned_data.csv"
	transactions, err := parseCSV(filePath)
	if err != nil {
		log.Fatalf("Error parsing CSV file: %v", err)
	}

	start := time.Now()
	repo := repository.NewLogisticTransactionRepositoryMongoDB(db.Collection("logistic_transactions"))
	// srv := service.NewLogisticTransactionService(repo)
	if err = repo.InsertMany(context.TODO(), transactions); err != nil {
		log.Fatalf("Error inserting data into MongoDB: %v", err)
	}
	end := time.Since(start)
	fmt.Println("Data successfully imported to MongoDB!")
	fmt.Println("Time: ", end)
}
