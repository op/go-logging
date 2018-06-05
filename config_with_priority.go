//+build windows plan9

package logging

func buildSyslog(config syslogConfig) (Backend, bool) {
	if config.priority == -1 {
		if ret, err := NewSyslogBackend(config.prefix); err == nil {
			return ret, false
		}
	} else {

		if ret, err := NewSyslogBackendPriority(config.prefix, Priority(config.priority)); err == nil {
			return ret, false
		}
	}
	return nil, true
}