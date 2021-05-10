package watchman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
)

// NotifyBackend makes call to server to notify about project entropy
func NotifyBackend(entropy uint64) bool {
	postBody, _ := json.Marshal(map[string]string{
		"id": sensor.Load.MachineID,
		"v":  strconv.FormatUint(entropy, 10),
	})
	responseBody := bytes.NewBuffer(postBody)

	// make the request
	resp, err := http.Post(config.Load.Feeder+"/data", "application/json", responseBody)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)

	}

	fmt.Println("Publisher API Call : ")
	fmt.Println("Publisher API Call : Target : ", config.Load.Feeder+"/data")
	fmt.Println("Publisher API Call : Params : ", strconv.FormatUint(entropy, 10))
	fmt.Println("Publisher API Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	fmt.Println("Publisher API Call : Response : ", string(body))

	return true
}
