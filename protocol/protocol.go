package protocol

import (
	"github.com/mdlayher/ethernet"
	"net"
)

// The only EtherType used here
const EtherType ethernet.EtherType = 0x1234

var Broadcast net.HardwareAddr = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

// Address implementation
type Addr struct {
	HardwareAddr net.HardwareAddr
}

func (a *Addr) Network() string {return "frames"}

func (a *Addr) String() string {return a.HardwareAddr.String()}

