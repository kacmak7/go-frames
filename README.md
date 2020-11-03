# go-frames

Chat with other ethernet machines using only the data link layer.

## Usage on Linux
### Receive messages
Start receiver process in order to intercept proper frames (0x1234 EtherType).

### Send messages
Message any physical machine in the same network. Don't provide any destination MAC address in order to Broadcast FF:FF:FF:FF:FF:FF
