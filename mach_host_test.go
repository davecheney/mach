package mach

import (
	"fmt"
	"testing"
)

func TestMachHostSelf(*testing.T) {
	host := HostSelf()
	fmt.Sprintf("%#v\n", host)
}

func TestHostBasicInfo(t *testing.T) {
	host := HostSelf()
	info := host.BasicInfo()
	if info == nil {
		t.Error("Unable to retrieve basic_info")
	}
}

func TestHostVmInfo(t *testing.T) {
	host := HostSelf()
	info, err := host.VmInfo()
	if err != nil {
		t.Error("Unable to retrieve vm_info")
	}
	t.Logf("Pagein: %d Pageout: %d\n", info.Pageins, info.Pageouts)
}


func TestHostLoadInfo(t *testing.T) {
	host := HostSelf()
	info := host.LoadInfo()
	if info == nil {
		t.Error("Unable to retrieve host_load_info")
	}
	t.Logf("1min loadavg: %f, 5min loadavg: %f, 15min loadavg: %f\n", info.OneMin(), info.FiveMin(), info.FifteenMin())
}

func TestPageSize(t *testing.T) {
	host := HostSelf()
	t.Logf("host_page_size: %d", host.Pagesize())
}



