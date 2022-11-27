package query

import (
	structs2 "atc-runner/src/data/structs"
	"atc-runner/src/helpers/database/utils"
	"fmt"
	"log"
	"sync"
)


func GetArbPairRoutes(ArbPair structs2.ArbPair, ArbPairRoutesWaitGroup *sync.WaitGroup, ArbPairRoutesChannel chan structs2.ArbPair) {

	// Schedule The Call To WaitGroup's Done To Tell GoRoutine Is Completed.
	defer ArbPairRoutesWaitGroup.Done()

	// Query
	GetArbPairRoutesQuery := fmt.Sprintf("SELECT routes.* FROM routes WHERE routes.pair_id = %d", ArbPair.PairDbId)

	// Create Connection To DB
	DBConnection := utils.CreateDatabaseConnection()

	// Create List Of Pair
	var Routes []structs2.Route

	// Execute DB Query
	QueryError := DBConnection.Select(&Routes, GetArbPairRoutesQuery)

	// Catch Any Errors When Querying
	if QueryError != nil {
		log.Fatal(QueryError)
	}

	// Close Connection
	DBConnectionCloseError := DBConnection.Close()
	if DBConnectionCloseError != nil {
		log.Fatal(DBConnectionCloseError)
	}

	ArbPair.PairRoutes = Routes

	// Send Return Value Back In Channel
	ArbPairRoutesChannel <- ArbPair

}