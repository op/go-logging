package trace

import (
	"runtime"
	"strings"
	"regexp"
)

var rx, _ = regexp.Compile("\\.func\\d+$")
func GetCallerFunctionName(backLevel int) string {

	backLevel += 2
	var pc = make([]uintptr, backLevel  + 3)
	runtime.Callers(backLevel, pc)
	tryAgain := true
	var fn *runtime.Func
	for i:=0; i < len(pc) && tryAgain; i++ {
		fn = runtime.FuncForPC(pc[i])
		tryAgain = rx.MatchString(fn.Name())
		//log.Printf("for, tryAgain=%v, name=%s", tryAgain, fn.Name())
	}
	if sp := strings.Split(fn.Name(), "."); len(sp) != 0 {
		return sp[len(sp) - 1]
	}
	return ""
}