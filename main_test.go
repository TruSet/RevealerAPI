package main
import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
   "fmt"
   "github.com/gin-gonic/gin"
   "github.com/stretchr/testify/assert"
   "github.com/miguelmota/go-solidity-sha3"
   "github.com/TruSet/RevealerAPI/database"
   "bytes"
   "encoding/hex"
)

var postgresUri = "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
var router = SetupRouter()

func requestBodyBuffer(jsonStr string) (*bytes.Buffer) {
   return bytes.NewBuffer([]byte(jsonStr))
}

func TestInvalidCommitment(t *testing.T) {
   db := SetupDB(postgresUri)
   defer db.Close()
   database.InitDb(db)

   test_voteOption := 1
   test_salt := 666
   test_pollID := "1234"
   test_voterAddress := "0x1234"

   var jsonStr = fmt.Sprintf(
     "{\"pollID\":\"%s\", \"voterAddress\": \"%s\", \"commitHash\": \"%s\", \"voteOption\": %d, \"salt\": %d}",
     test_pollID,
     test_voterAddress,
     "NotAValidCommitHash",
     test_voteOption,
     test_salt,
   )

   req, _ := http.NewRequest("POST", "/revealer/v0.1/commitments", requestBodyBuffer(jsonStr))
   req.Header.Set("Content-Type", "application/json")

   w := httptest.NewRecorder()
   router.ServeHTTP(w, req)

   // the request gives a 406
   assert.Equal(t, http.StatusNotAcceptable, w.Code)

}

func TestValidCommitment(t *testing.T) {
   db := SetupDB(postgresUri)
   defer db.Close()
   database.InitDb(db)

  
   // Build our expected body
   body := gin.H{
      "status": http.StatusCreated,
      "message": "vote will be revealed when voting closes",
   }

   test_voteOption := 1
   test_salt := 666
   test_pollID := "1234"
   test_voterAddress := "0x1234"

   test_commitHash := solsha3.SoliditySHA3(
    solsha3.Uint256(test_voteOption),
    solsha3.Uint256(test_salt),
  )

   var jsonStr = fmt.Sprintf(
     "{\"pollID\":\"%s\", \"voterAddress\": \"%s\", \"commitHash\": \"%s\", \"voteOption\": %d, \"salt\": %d}",
     test_pollID,
     test_voterAddress,
     hex.EncodeToString(test_commitHash),
     test_voteOption,
     test_salt,
   )

   req, _ := http.NewRequest("POST", "/revealer/v0.1/commitments", requestBodyBuffer(jsonStr))
   req.Header.Set("Content-Type", "application/json")

   w := httptest.NewRecorder()
   router.ServeHTTP(w, req)

   // Assert we encoded correctly,
   // the request gives a 200
   assert.Equal(t, http.StatusCreated, w.Code)
   // Convert the JSON response to a map
   var response map[string]string
   unmarshallErr := json.Unmarshal([]byte(w.Body.String()), &response)

   // Grab the value & check whether or not it exists
   value, exists := response["message"]

   // Make some assertions on the correctness of the response.
   assert.Nil(t, unmarshallErr)
   assert.True(t, exists)
   assert.Equal(t, body["message"], value)


   var commitment database.Commitment;
   database.Db.Where(
     "poll_id = ? and voter_address = ? and commit_hash = ?",
     test_pollID,
     test_voterAddress,
     hex.EncodeToString(test_commitHash),
   ).Last(&commitment)

   assert.True(t, commitment.CommitHash == hex.EncodeToString(test_commitHash))

}
