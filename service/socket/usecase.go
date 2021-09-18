package socket

import (
	socketio "github.com/googollee/go-socket.io"
)

type Usecase interface {
	Socket() (*socketio.Server, error)
}
