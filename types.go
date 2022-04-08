package socketio

import (
	"github.com/googollee/go-socket.io/v2/parser"
	"reflect"
)

// namespace
const (
	aliasRootNamespace = "/"
	rootNamespace      = ""
)

// message
const (
	clientDisconnectMsg = "client namespace disconnect"
)

type readHandler func(c *conn, header parser.Header) error

var (
	defaultHeaderType = []reflect.Type{reflect.TypeOf("")}
)