package mach

// #include <mach/mach_error.h>
// #include <mach/host_info.h>
// #include <mach/mach_host.h>
import "C"

import (
	"unsafe"
	"os"
	"fmt"
)

const (
	LOAD_SCALE = 1000.
)

type Host C.host_t

func HostSelf() Host {
	host := C.mach_host_self()
	return Host(host)
}

// type host_basic_info C.host_basic_info

type BasicInfo C.host_basic_info_data_t

func (i *BasicInfo) MaxCpus() uint32 {
	return uint32(i.max_cpus)
}

func (h Host) BasicInfo() *BasicInfo {
	host_info := new(BasicInfo)
	count := C.mach_msg_type_number_t(C.HOST_BASIC_INFO_COUNT)
	if C.host_info(C.host_t(h), C.HOST_BASIC_INFO, (*C.integer_t)(unsafe.Pointer(host_info)), &count) != 0 {
		return nil
	}
	
	return host_info
}

type VmInfo struct {
	Pageins uint64
	Pageouts uint64
}

func (h Host) VmInfo() (*VmInfo, os.Error) {
	vmstat := new(C.vm_statistics_data_t)
	nummsg := C.mach_msg_type_number_t(C.HOST_VM_INFO_COUNT)
	ret := C.host_statistics(C.host_t(h), C.HOST_VM_INFO, (*C.integer_t)(unsafe.Pointer(vmstat)), &nummsg)
	if ret != 0 {
		return nil, fmt.Errorf("host_statistics: %s", C.mach_error_string(C.mach_error_t(ret)))
	}
	return &VmInfo { uint64(vmstat.pageins), uint64(vmstat.pageouts) }, nil
}

type LoadInfo C.host_load_info_data_t

func (l *LoadInfo) OneMin() float {
	return float(l.avenrun[0]) / LOAD_SCALE
}
 
func (l *LoadInfo) FiveMin() float {
	return float(l.avenrun[1]) / LOAD_SCALE
}

func (l *LoadInfo) FifteenMin() float {
	return float(l.avenrun[2]) / LOAD_SCALE
}

func (h Host) LoadInfo() *LoadInfo {
	host_load_info := new(LoadInfo)
	count := C.mach_msg_type_number_t(C.HOST_LOAD_INFO_COUNT)
	if C.host_statistics(C.host_t(h), C.HOST_LOAD_INFO, (*C.integer_t)(unsafe.Pointer(host_load_info)), &count) != 0 {
		return nil
	}
	
	return host_load_info
}

func (h Host) Pagesize() int {
	var page_size C.vm_size_t
	_ = C.host_page_size(C.host_t(h), &page_size)
	return int(page_size)
}
