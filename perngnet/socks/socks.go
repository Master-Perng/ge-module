package socks

// Socks4AHeader represents the header of SOCKS4A protocol
type Socks4AHeader struct {
	Version  uint8   // Version number (always 4)
	Command  uint8   // Command code
	Port     uint16  // Destination port
	IPAddr   [4]byte // IPv4 address
	UserID   []byte  // User ID
	NullByte uint8   // Null byte (always 0)
	Domain   []byte  // Domain name
}

// Socks5Header represents the header of SOCKS5 protocol
type Socks5Header struct {
	Version     uint8  // Version number (always 5)
	Command     uint8  // Command code
	Reserved    uint8  // Reserved (must be 0x00)
	AddressType uint8  // Address type
	Destination string // Destination address (IPv4, IPv6, or domain name)
	Port        uint16 // Destination port
}
