package models

type SystemInfo struct {
	CPUUtilization    float64   `json:"cpu_utilization"`
	MemoryUtilization float64   `json:"memory_utilization"`
	DiskUsage         float64   `json:"disk_usage"`
	BandwidthUsage    Bandwidth `json:"bandwidth_usage"`
}

type Bandwidth struct {
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}
