package ably

import (
    "github.com/ably/ably-go/ably"
    "fmt"
)


func Init() {
	fmt.Println("Ably Client Init called")
    var err error
    AblyClient, err = ably.NewRealtime(ably.WithKey("J9x8VQ.Lw-2eg:Vgo3n8N5OLbPAX1CSu16seIVC3TZcz8FjKdCUuz9SAU"))
    if err != nil {
        panic(err)
    }

    AblyClient.Connect()

}
var AblyClient *ably.Realtime