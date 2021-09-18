package socketio

import (
	"log"
	"net/http"

	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
)

const SocketRoomPrice = "price"

// const SocketEventUpdateBnibLowestSellPrice = "UPDATE_BNIB_LOWEST_SELL_PRICE"
// const SocketEventUpdateBnibHighestBidPrice = "UPDATE_BNIB_HIGHEST_BID_PRICE"
// const SocketEventUpdatePreOrderLowestSellPrice = "UPDATE_PRE_ORDER_LOWEST_SELL_PRICE"
// const SocketEventUpdatePreOrderHighestBidPrice = "UPDATE_PRE_ORDER_HIGHEST_BID_PRICE"

type SocketClient interface {
	GetClient() *socketio.Server
}

type Socket struct {
	socketServer *socketio.Server
}

func NewSocketClient() SocketClient {
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	pt := polling.Default

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})
	if err != nil {
		log.Println("error creating socket server")
	}

	go server.Serve()
	return &Socket{
		socketServer: server,
	}
}

func (s Socket) GetClient() *socketio.Server {
	return s.socketServer
}
