package ip

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
