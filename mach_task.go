package mach

// #include <mach/mach_init.h>
import "C"

type Task C.mach_port_t

func TaskSelf() Task {
	port := C.mach_task_self();
	return Task(port)
}
