package tcp

/*
*  A echo server to test whether the server is functioning normally
*/

import(
	"bufio"                              
	"context"
	"github.com/hdt3213/godis/lib/logger"
	"github.com/hdt3213/godis/lib/sync/atomic"
	"github.com/hdt3213/godis/lib/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

//EchoHandler echos received line to client,using for test
type EchoHandler struct{

}
