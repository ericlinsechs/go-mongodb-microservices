package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ericlinsechs/go-mongodb-microservices/movies/pkg/models/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	movies   *mongodb.MovieModel
}

func main() {
	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 3000, "HTTP server network port")
	mongoDatabase := flag.String("mongoDatabase", "movies", "Database name")
	// enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()
	flag.PrintDefaults()

	// Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// load .env file from given path
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	errorLog.Fatalf("Error loading .env file")
	// }

	mongodbURI := os.Getenv("MONGODB_ATLAS_URI")

	opts := options.Client().ApplyURI(mongodbURI)
	// if *enableCredentials {
	// 	co.Auth = &options.Credential{
	// 		Username: os.Getenv("MONGODB_USERNAME"),
	// 		Password: os.Getenv("MONGODB_PASSWORD"),
	// 	}
	// }

	// Establish database connection
	client, err := mongo.NewClient(opts)
	if err != nil {
		errorLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	infoLog.Printf("Database connection established")

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Printf("Ping mondoDB success")

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		movies: &mongodb.MovieModel{
			Collection: client.Database(*mongoDatabase).Collection("movies"),
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	server := gin.Default()
	server = app.routes(server)
	err = server.Run(serverURI)
	errorLog.Fatal(err)
}
