package protocol

import (
	"github.com/mdlayher/ethernet"
	"net"
)

// The only EtherType used here
const EtherType ethernet.EtherType = 0x1234

var Broadcast net.HardwareAddr = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

// TODO ADDR
type Addr net.HardwareAddr // TODO whats with this Addr ??

func (Addr) Network() string { return "protocol" }
func (Addr) String() string { return string(Broadcast) }
