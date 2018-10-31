package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"github.com/TruSet/RevealerAPI/config"
	"github.com/TruSet/RevealerAPI/database"
	"github.com/TruSet/RevealerAPI/events"
)

var localMode bool
var poll bool
var rinkeby bool

func init() {
	// Log subscriptions are not supported over RPC
	flag.BoolVar(&poll, "poll", false, "Poll future events. If false (default) we will populate past logs but not future logs")
}

func main() {

	var clientString string
	var err error

	environment := flag.String("e", "development", "Specify an environment {development, docker, infura}")
	port := flag.String("p", "8080", "Specify a port for the gin server")
	service := flag.String("s", "reveal", "Mode should be 'rest' or 'reveal' to indicate whether this is a REST api that accepts delegated reveals, or a 'revealer' service that reveals votes made to a voting contract")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	log.Printf("Starting TruSet Revealer %v service in %v mode...", *service, *environment)
	env := config.GetConfig()
	// Try IPC, if configured
	clientString = env.GetString("ethereumIpc")

	if clientString == "" {
		log.Println("Using websockets...")
		clientString = env.GetString("ethereumWs")
	}

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client %v: %v", clientString, err)
	}
	defer log.Printf("Shutting down TruSet Revealer %v service", *service)

	// Open a database connection, and close it when we are terminated
	postgresUri := env.GetString("postgresUri")
	if postgresUri == "" {
		postgresUri = os.Getenv("DATABASE_URL")
	}
	db := SetupDB(postgresUri)
	defer db.Close()
	database.InitDb(db)

	// Uncomment to create some test data
	//database.SetupTestData()
	// Run a REST server to serve TruSet API requests
	// Do this before processing our backlog, otherwise Heroku thinks we are taking too long to start
	r := SetupRouter()
	go r.Run(":" + *port) // listen and serve on 0.0.0.0:8080 by default

	// Read data from all past events and reveal any we have not already revealed.
	// This allows us to catch up automatically on lost logs whenever we restart.
	// (If we never restart there's no extra work to do. But if we do, this is safer than ignoring the gap.)
	// There are more efficient ways of doing this but they are more work. We would need to track
	// logs that have been processed successfully.
	commitRevealVotingContractAddress := env.GetString("commitRevealVotingContractAddress")
	events.Init(clientString, commitRevealVotingContractAddress)
	log.Printf("Listening to CRV contract at %v", commitRevealVotingContractAddress)

	switch *service {
	case "reveal":
		// poll for contract events that result in reveals
		events.ProcessPastEvents()
		events.ProcessFutureEvents()
	case "api":
		// Run a REST server to serve TruSet API requests
		r := SetupRouter()
		r.Run(":" + *port) // listen and serve on 0.0.0.0:8080 by default
	default:
		log.Fatal("-s option is not one of the supported service types - should be 'reveal' or 'api'")
	}
}

func SetupDB(postgresUri string) *gorm.DB {
	db, err := gorm.Open("postgres", postgresUri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// TODO set cors
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}

	//router.Use(cors.New(config))

	router.Use(cors.Default())
	v1 := router.Group("/revealer/v0.1")
	{
		commitments := v1.Group("/commitments")
		{
			commitments.POST("", database.StoreCommitment) // Supports /instruments?proposalstate=xyz
		}
	}
	return router
}
