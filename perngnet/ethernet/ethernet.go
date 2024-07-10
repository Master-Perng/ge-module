package ethernet

type Ethernet struct {
	Header Header
	Body   []byte
}

type Header struct {
	DestinationMAC [6]byte
	SourceMAC      [6]byte
	TypeLength     uint16
}
