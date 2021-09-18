package usecase

import (
	"gaming-company-test/lib/socketio"
	"gaming-company-test/service/socket"

	socketLib "github.com/googollee/go-socket.io"
)

type Usecase struct {
	//socketRepo socket.SocketRepository
	socket socketio.SocketClient
}

func New(socket socketio.SocketClient) socket.Usecase {
	return &Usecase{
		socket: socket,
	}
}

func (u *Usecase) Socket() (*socketLib.Server, error) {
	server := u.socket.GetClient()

	server.OnConnect("/", func(conn socketLib.Conn) error {
		conn.Join(socketio.SocketRoomPrice)
		return nil
	})

	server.OnDisconnect("/", func(conn socketLib.Conn, s string) {
	})

	server.OnError("/", func(e error) {
	})

	return server, nil
}
