package win_api

type MIB_TCPROW_OWNER_PID struct {
	DwState        uint32
	DwLocalAddr    uint32
	DwLocalPort    uint32
	DwRemoteAddr   uint32
	DwRemotePort   uint32
	DwOwningPid    uint32
	DwOffloadState uint32
}

type MIB_TCPTABLE2 struct {
	dwNumEntries uint32
	table        [4096]MIB_TCPROW_OWNER_PID
}
