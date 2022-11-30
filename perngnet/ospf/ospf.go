package ospf

// Header is the header of OSPF packet
type Header struct {
	Version  uint8
	Type     uint8
	Length   uint16
	Router   [4]byte
	Area     uint32
	Checksum uint16
	AuthType uint16
	AuthData uint64
}

// Hello is the hello packet of OSPF
type Hello struct {
	NetMask         [4]byte
	HelloInterval   uint16
	Options         uint8
	RtrPri          uint8
	RtrDeadInterval uint32
	Descr           [4]byte
	BackupDR        [4]byte
	Neighbor        [4]byte
}

type DBD struct {
	InterfaceMtu uint16
	Options      uint8
	Flags        [5]byte
	I            [1]byte
	MS           [1]byte
	Sequence     uint32
	LSAHeader    []byte
}
