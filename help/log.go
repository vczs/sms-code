package help

import (
	"fmt"
	"runtime"
)

func VczLog(desc string, err error) {
	p, _, line, _ := runtime.Caller(1)
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("vczs_error_log: [%s(%d)](%s):%v\n", name, line, desc, err)
}
