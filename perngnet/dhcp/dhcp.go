package dhcp

// Header is the header of DHCP packet
type Header struct {
	Op      uint8
	HType   uint8
	HLen    uint8
	Hops    uint8
	XID     uint32
	Secs    uint16
	Flags   uint16
	CIAddr  [4]byte
	YIAddr  [4]byte
	SIAddr  [4]byte
	GIAddr  [4]byte
	CHAddr  [16]byte
	SName   [64]byte
	File    [128]byte
	Magic   [4]byte
	Options []byte
}
