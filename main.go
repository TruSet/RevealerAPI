package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
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
	var client *ethclient.Client
	var err error

	environment := flag.String("e", "development", "Specify an environment {development, docker, infura}")
	port := flag.String("p", "8080", "Specify a port for the gin server")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	log.Println(fmt.Sprintf("Starting TruSet Revealer API server in %v mode...", *environment))
	env := config.GetConfig()
  // Try IPC, if configured
  clientString = env.GetString("ethereumIpc")

  if clientString == "" {
    log.Println("Using websockets...")
    clientString = env.GetString("ethereumWs")
  }

	client, err = ethclient.Dial(clientString)

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client %v: %v", clientString, err)
	}
	defer log.Println("Shutting down TruSet Revealer API server...")

	// Open a database connection, and close it when we are terminated
	postgresUri := env.GetString("postgresUri")
  db := SetupDB(postgresUri)
	defer db.Close()
	database.InitDb(db)

	// Uncomment to create some test data
	//database.SetupTestData()

	// Read data from all past events and write them into our DB
	commitRevealVotingContractAddress := env.GetString("commitRevealVotingContractAddress")
	events.Init(commitRevealVotingContractAddress)
	events.ProcessPastEvents(client)

	if poll {
		// In parallel with serving our REST resources, we continue to process incoming blockchain logs
		// TODO: make sure we don't miss any events in the time windows between ProcessPastEvents and
		// ProcessFutureEvents. (More likely when ProcessPastEvents takes significant elapsed time?)
		go events.ProcessFutureEvents(client)
	}

	// Run a REST server to serve TruSet API requests
	r := SetupRouter()
  r.Run(":"+*port) // listen and serve on 0.0.0.0:8080 by default
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

	v1 := router.Group("/revealer/v0.1")
	{
		commitments := v1.Group("/commitments")
		{
			commitments.POST("", database.StoreCommitment) // Supports /instruments?proposalstate=xyz
		}
	}
  return router
}
