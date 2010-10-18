package mach

// #include <mach/mach_host.h>
// #include <mach/processor_info.h>
import "C"

import (
	"unsafe"
)

type ProcessorCpuLoadInfo C.processor_cpu_load_info_t

func (h Host) ProcessorCpuLoadInfo() {
 	var pcli *[]ProcessorCpuLoadInfo
	var numcpu C.natural_t
	var nummsg C.mach_msg_type_number_t

	if C.host_processor_info(C.host_t(h), C.PROCESSOR_CPU_LOAD_INFO, &numcpu, (*C.processor_info_array_t)(unsafe.Pointer(pcli)), &nummsg) != 0 {
		//
	}
}