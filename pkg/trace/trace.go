package trace

// returns full package and function name
// 0 for the caller name, 1 for a up level, etc...
import (
	"runtime"
	"strings"
	"regexp"
)

func GetCallerName(backLevel int) string {

	pc := make([]uintptr, 10)  // at least 1 entry needed
	runtime.Callers(backLevel + 2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name();

}

// returns only function name
func GetCallerFunctionName(backLevel int) string {
	caller := GetCallerName(backLevel + 1)
	sp := strings.Split(caller, ".")

	if (len(sp) == 0) {
		return ""
	}

	return sp[len(sp) - 1]
}


// give the name of functions in stack
// backlevel
// 0 - returns the direct caller name
// 1 - returns the second to last function caller name and go on
var rx, _ = regexp.Compile("^func\\d+")
func GetCallerFunctionNameSkippingAnnonymous(backlevel int) string {

	var name string = "";
	counter := 0
	for tryAgain := true; tryAgain; counter++ {
		name = GetCallerFunctionName(backlevel + 1 + counter)
		tryAgain = rx.MatchString(name)
	}

	return name

}