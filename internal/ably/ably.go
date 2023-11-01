package ably

import (
    "github.com/ably/ably-go/ably"
    "fmt"
)

var ClientChannel *ably.RealtimeChannel 

func Init(machineID string) {
    fmt.Println("Ably Client Init called")
    var err error
    AblyClient, err := ably.NewRealtime(ably.WithKey("J9x8VQ.hxM02Q:cuc5WSlCcdoR67M26mrJPBgsliG0zum2DJfIwsZNvn8"))
    if err != nil {
        panic(err)
    }

    AblyClient.Connect()
    ClientChannel = AblyClient.Channels.Get("History_" + machineID)
}
