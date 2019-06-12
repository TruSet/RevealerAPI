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

	"github.com/TruSet/RevealerAPI/database"
	"github.com/TruSet/RevealerAPI/events"
)

var localMode bool
var poll bool
var rinkeby bool

func main() {
	service := flag.String("s", "reveal", "Service should be 'api' or 'reveal' to indicate whether this is a REST api that accepts delegated reveals, or a 'revealer' service that reveals votes made to a voting contract")
	ethereumRpc := flag.String("rpc", os.Getenv("ETHEREUM_RPC"), "The rpc endpoint of your ethereum client (defaults to ETHEREUM_RPC env var)")
	postgresUri := flag.String("db", os.Getenv("DATABASE_URL"), "The postgres database endpoint (defaults to DATABASE_URL env var)")
	commitRevealVotingContractAddress := flag.String("crv", os.Getenv("CRV_ADDRESS"), "The address of the CommitRevealVoting contract (defaults to CRV_ADDRESS env var)")

	flag.Usage = func() {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	log.Printf("Starting TruSet Revealer %v service...", *service)

	// ethereum rpc path - ipc, websockets
	if *ethereumRpc == "" {
		log.Fatal("No -rpc flag or ETHEREUM_RPC specified (wss://..., *.ipc)")
	}

	defer log.Printf("Shutting down TruSet Revealer %v service", *service)

	// Open a database connection, and close it when we are terminated
	if *postgresUri == "" {
		log.Fatal("No -db flag or DATABASE_URL specified (postgresql://...)")
	}
	db := SetupDB(*postgresUri)
	defer db.Close()
	database.InitDb(db)

	// Uncomment to create some test data
	//database.SetupTestData()

	events.Init(*ethereumRpc, *commitRevealVotingContractAddress)
	log.Printf("Listening to CRV contract at %v", *commitRevealVotingContractAddress)

	switch *service {
	case "reveal":
		// Read data from all past events and reveal any we have not already revealed.
		// This allows us to catch up automatically on lost logs whenever we restart.
		// (If we never restart there's no extra work to do. But if we do, this is safer than ignoring the gap.)
		// There are more efficient ways of doing this but they are more work. We would need to track
		// logs that have been processed successfully.

		// poll for contract events that result in reveals
		var systemCreatedAtBlock uint64 = 3456213 // TODO: get from env var?
		events.ProcessPastEvents(systemCreatedAtBlock)
		events.ProcessFutureEvents()
	case "api":
		port := os.Getenv("PORT")
		// Run a REST server to serve TruSet API requests
		r := SetupRouter()
		r.Run(":" + port) // listen and serve on 0.0.0.0:8080 by default
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
