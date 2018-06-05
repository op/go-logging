//+build !windows,!plan9

package logging

func buildSyslog(config syslogConfig) (Backend, bool) {
	if ret, err := NewSyslogBackend(config.prefix); err == nil {
		return ret, false
	}
	return nil, true
}
