package ip

import "encoding/binary"

type IPv4 struct {
	Header V4Header
	Body   []byte
}

type IPv6 struct {
	Header V6Header
	Body   []byte
}

type V4Header struct {
	Version       uint8   // 版本（4位）和头部长度（4位）
	TOS           uint8   // 服务类型
	TotalLength   uint16  // 总长度
	ID            uint16  // 标识
	Flags         uint8   // 标志（3位）和片偏移（13位）
	TTL           uint8   // 存活时间
	Protocol      uint8   // 协议
	Checksum      uint16  // 头部校验和
	SourceIP      [4]byte // 源IP地址
	DestinationIP [4]byte // 目的IP地址
	Options       []byte  // 可选字段
	Padding       []byte  // 填充字段
}

type V6Header struct {
	Version       uint8    // 版本（4位），Traffic Class（8位），Flow Label（20位）
	PayloadLength uint16   // 载荷长度
	NextHeader    uint8    // 下一个头部
	HopLimit      uint8    // 跳限制
	SourceIP      [16]byte // 源IPv6地址
	DestinationIP [16]byte // 目的IPv6地址
}

func (ipv4 *IPv4) CheckSum() {
	ipv4.Header.Checksum = 0

	// Convert IPv4 header to a byte slice
	headerBytes := make([]byte, 20)
	binary.BigEndian.PutUint16(headerBytes[0:2], uint16(ipv4.Header.Version<<4|(20/4))) // Version and Header Length
	headerBytes[1] = ipv4.Header.TOS
	binary.BigEndian.PutUint16(headerBytes[2:4], ipv4.Header.TotalLength)
	binary.BigEndian.PutUint16(headerBytes[4:6], ipv4.Header.ID)
	binary.BigEndian.PutUint16(headerBytes[6:8], uint16(ipv4.Header.Flags<<13))
	headerBytes[8] = ipv4.Header.TTL
	headerBytes[9] = ipv4.Header.Protocol
	binary.BigEndian.PutUint16(headerBytes[10:12], 0) // Placeholder for Checksum
	copy(headerBytes[12:16], ipv4.Header.SourceIP[:])
	copy(headerBytes[16:20], ipv4.Header.DestinationIP[:])
	// Calculate checksum
	checksum := uint32(0)
	for i := 0; i < 10; i += 2 {
		checksum += uint32(headerBytes[i])<<8 | uint32(headerBytes[i+1])
	}
	for checksum > 0xffff {
		checksum = (checksum >> 16) + (checksum & 0xffff)
	}
	ipv4.Header.Checksum = ^uint16(checksum)
}
