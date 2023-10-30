package watchman

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"strconv"
	"time"
	"github.com/YashKumarVerma/hentry-client/internal/config"
	"github.com/YashKumarVerma/hentry-client/internal/sensor"
	"github.com/YashKumarVerma/hentry-client/internal/ably"
	
)

// NotifyBackend makes call to server to notify about project entropy
func NotifyBackend(entropy uint64) bool {
	channel := ably.AblyClient.Channels.Get("History_"+sensor.Load.MachineID)
	channel.PublishAsync("entropy", string(strconv.FormatUint(entropy, 10)), func(err error) {
		if err != nil {
			fmt.Println("failed to publish", err)
		} else {
			fmt.Println("publish ok" +sensor.Load.MachineID )
		}
	})


	
	// postBody, _ := json.Marshal(map[string]string{
	// 	"id": sensor.Load.MachineID,
	// 	"v":  string(strconv.FormatUint(entropy, 10)),
	// })
	// responseBody := bytes.NewBuffer(postBody)

	// // make the request
	// resp, err := http.Post(config.Load.Server+"/timeseries", "application/json", responseBody)

	// // check if non 200 response
	// if err != nil {
	// 	log.Fatalf("An Error Encountered %v", err)
	// }
	// defer resp.Body.Close()

	// // read response body as required
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	if config.Load.Debug == true {
		fmt.Println("Publisher Entropy Call : Target : ", config.Load.Server+"/data")
		fmt.Println("Publisher Entropy Call : Params : ", strconv.FormatUint(entropy, 10))
		fmt.Println("Publisher Entropy Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	} else {
		fmt.Println("Entropy Transmitted.")
	}

	return true
}

// UpdateSnapshot makes call to server to tell about current project snapshot
func UpdateSnapshot(data uint64) bool {

	channel := ably.AblyClient.Channels.Get("History_"+sensor.Load.MachineID)
	channel.PublishAsync("snapshot",string(strconv.FormatUint(data, 10)), func(err error) {
		if err != nil {
			fmt.Println("failed to publish", err)
		} else {
			fmt.Println("publish ok" +sensor.Load.MachineID )
		}
	})

	// postBody, _ := json.Marshal(map[string]string{
	// 	"id": sensor.Load.MachineID,
	// 	"v":  string(strconv.FormatUint(data, 10)),
	// })
	// responseBody := bytes.NewBuffer(postBody)

	// // make the request
	// resp, err := http.Post(config.Load.Server+"/snapshot", "application/json", responseBody)

	// // check if non 200 response
	// if err != nil {
	// 	log.Fatalf("An Error Encountered %v", err)
	// }
	// defer resp.Body.Close()

	// // read response body as required
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)

	// }

	if config.Load.Debug == true {
		fmt.Println("Publisher Snapshot Call : Target : ", config.Load.Server+"/data")
		fmt.Println("Publisher Snapshot Call : Params : ", strconv.FormatUint(data, 10))
		fmt.Println("Publisher Snapshot Call : Stamp : ", time.Unix(time.Now().Unix(), 0).Format(time.RFC1123Z))
	} else {
		fmt.Println("Snapshot transmitted.")
	}

	return true
}
