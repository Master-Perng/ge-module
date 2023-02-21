package win_api

type MIB_IPFORWARDTABLE struct {
	DwNumEntries uint32
	Table        [4096]MIB_IPFORWARDROW
}

type MIB_IPFORWARDROW struct {
	DwForwardDest      [4]byte
	DwForwardMask      [4]byte
	DwForwardPolicy    [4]byte
	DwForwardNextHop   [4]byte
	DwForwardIfIndex   [4]byte
	DwForwardType      [4]byte
	DwForwardProto     [4]byte
	DwForwardAge       [4]byte
	DwForwardNextHopAS [4]byte
	DwForwardMetric1   [4]byte
	DwForwardMetric2   [4]byte
	DwForwardMetric3   [4]byte
	DwForwardMetric4   [4]byte
	DwForwardMetric5   [4]byte
}
