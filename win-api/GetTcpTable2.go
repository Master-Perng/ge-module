package win_api

type MIB_TCPROW_OWNER_PID struct {
	dwState        uint32
	dwLocalAddr    uint32
	dwLocalPort    uint32
	dwRemoteAddr   uint32
	dwRemotePort   uint32
	dwOwningPid    uint32
	dwOffloadState uint32
}

type MIB_TCPTABLE2 struct {
	dwNumEntries uint32
	table        [4096]MIB_TCPROW_OWNER_PID
}
