package ably

// import(
// 	"github.com/YashKumarVerma/hentry-client/internal/ably"
// )

func UserOnlinePresence(){
	ably.AblyClient.Enter("Online")
}
func UserWritingCodePresence(){
	ably.AblyClient.Update("Coding")
}
func UserLeavePresence(){
	ably.AblyClient.Leave("Offline")
}