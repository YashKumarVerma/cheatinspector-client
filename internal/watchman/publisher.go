package watchman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
)

// NotifyBackend makes call to server to notify about project entropy
func NotifyBackend(entropy uint64) bool {
	postBody, _ := json.Marshal(map[string]string{
		"id": sensor.Load.MachineID,
		"v":  string(strconv.FormatUint(entropy, 10)),
	})
	responseBody := bytes.NewBuffer(postBody)

	// make the request
	resp, err := http.Post(config.Load.Feeder+"/timeseries", "application/json", responseBody)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Entropy Transmitted.")
	// fmt.Println("Publisher Entropy Call : Target : ", config.Load.Feeder+"/data")
	// fmt.Println("Publisher Entropy Call : Params : ", strconv.FormatUint(entropy, 10))
	// fmt.Println("Publisher Entropy Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	// fmt.Println("Publisher Entropy Call : Response : ", string(body))

	return true
}

// UpdateSnapshot makes call to server to tell about current project snapshot
func UpdateSnapshot(data uint64) bool {
	postBody, _ := json.Marshal(map[string]string{
		"id": sensor.Load.MachineID,
		"v":  string(strconv.FormatUint(data, 10)),
	})
	responseBody := bytes.NewBuffer(postBody)

	// make the request
	resp, err := http.Post(config.Load.Feeder+"/snapshot", "application/json", responseBody)

	// check if non 200 response
	if err != nil {
		log.Fatalf("An Error Encountered %v", err)
	}
	defer resp.Body.Close()

	// read response body as required
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)

	}

	fmt.Println("Snapshot transmitted.")
	// fmt.Println("Publisher Snapshot Call : Target : ", config.Load.Feeder+"/data")
	// fmt.Println("Publisher Snapshot Call : Params : ", strconv.FormatUint(data, 10))
	// fmt.Println("Publisher Snapshot Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	// fmt.Println("Publisher Snapshot Call : Response : ", string(body))

	return true
}
