package main
import (
   "encoding/json"
   "net/http"
   "net/http/httptest"
   "testing"
   "fmt"
   "github.com/gin-gonic/gin"
   "github.com/stretchr/testify/assert"
   "github.com/TruSet/RevealerAPI/database"
   //"github.com/TruSet/RevealerAPI/events"
   "bytes"
   //"net/url"
)

func TestRevealerAPI(t *testing.T) {
	 postgresUri := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
   db := SetupDB(postgresUri)
   defer db.Close()
   database.InitDb(db)

  
   // Build our expected body
   body := gin.H{
      "status": http.StatusCreated,
      "message": "vote will be revealed when voting closes",
   }
   // Grab our router
   router := SetupRouter()

   test_pollID := "1234"
   test_voterAddress := "0x1234"
   test_commitHash := "HASH"
   test_voteOption := "2"
   test_salt := "666"

   var jsonStr = fmt.Sprintf(
     "{\"pollID\":\"%s\", \"voterAddress\": \"%s\", \"commitHash\": \"%s\", \"voteOption\": %s, \"salt\": %s}",
     test_pollID,
     test_voterAddress,
     test_commitHash,
     test_voteOption,
     test_salt,
   )
   var requestBody = bytes.NewBuffer([]byte(jsonStr)) 

   req, _ := http.NewRequest("POST", "/revealer/v0.1/commitments", requestBody)
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
   database.Db.Where("poll_id = ? and voter_address = ?",
   test_pollID, test_voterAddress).First(&commitment)

   assert.True(t, true)

}
