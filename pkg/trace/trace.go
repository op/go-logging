package trace

import (
	"runtime"
	"strings"
	"regexp"
)


//
// Return the caller function name
// 0 -> returns the current caller function
// 1 -> returns the current caller parent
// etc.
//
var rx, _ = regexp.Compile("\\.func\\d+$")
func GetCallerFunctionName(backLevel int) string {

	backLevel += 2
	var pc, tryAgain = make([]uintptr, backLevel  + 3), true
	runtime.Callers(backLevel, pc)
	var fn *runtime.Func
	for i:=0; i < len(pc) && tryAgain; i++ {
		fn = runtime.FuncForPC(pc[i])
		tryAgain = rx.MatchString(fn.Name())
	}
	if index := strings.LastIndex(fn.Name(), "."); index != -1 {
		return fn.Name()[index + 1:]
	}
	return ""
}