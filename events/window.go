package events

import (
	"context"
	"github.com/getsentry/raven-go"
	"log"
	"time"
)

const MinSecondsBetweenBlockNumberRequests = 15
const MinLogConfirmations = 10

var (
	currentBlockNumberLastQueriedAt time.Time
	latestBlockNumberResult         uint64
)

func CurrentBlockNumber() uint64 {
	if time.Since(currentBlockNumberLastQueriedAt).Seconds() >= MinSecondsBetweenBlockNumberRequests {
		// Update the block number
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			raven.CaptureError(err, nil)
			// If we cannot get the block number, just log and return the last known block number
			log.Printf("Failed to get latest block number from client: %v", err)
		} else {
			latestBlockNumberResult = header.Number.Uint64()
		}
		currentBlockNumberLastQueriedAt = time.Now()
	}

	return latestBlockNumberResult
}

func FinalizedBlockNumber() uint64 {
	return CurrentBlockNumber() - MinLogConfirmations
}
