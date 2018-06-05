package logging

import (
	"path/filepath"
	"os"
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"strings"
)

const (
	toStdout = "STDOUT"
	toFile   = "FILE"
	toMemory = "MEMORY"
	toSyslog = "SYSLOG"
)

// DEBUG_CONFIG allows to show debugging information
var DEBUG_CONFIG = false

func LoadConfig() {
	Reset()
	configFile := "config/logging2.yaml"

	if path, ok := os.LookupEnv("logging"); ok {
		configFile = path
	} else {
		debug("'logging' not set use default")
	}
	LoadConfigFromFile(configFile)
}

func LoadConfigFromFile( configFile string) {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		yamlFile := filepath.Join(dir, configFile)
		debugf( "Read '%s'\n", yamlFile)
		if buffer, readErr := ioutil.ReadFile(yamlFile); readErr == nil {
			config := Config{}
			if yamlError := yaml.Unmarshal(buffer, &config); yamlError == nil {
				config.buildAndSetBackends()
				config.setLoggerLevel()
			} else {
				debug("logging.yaml", yamlError, "\n")
			}
		} else {
			debug("readFile", readErr,"\n")
		}
	} else {
		log.Fatalln("couldn't create abasolute path\n\n", err)
	}
}

func buildFormatter(format string) (Formatter, bool) {
	formatter, err := NewStringFormatter(format)
	if err != nil {
		debugf( "couldn't implement formatter '%s'\n\n%s", format, err)
		return nil, true
	}

	return formatter, false
}
func buildStdout(config stdoutConfig) (Backend, bool) {
	if formatter, err := buildFormatter(config.format); !err {
		backend := NewLogBackend(os.Stdout, config.prefix, 0)
		backend.Color = config.color
		return NewBackendFormatter(backend, formatter), false
	}
	return nil, true
}

func buildFile(config fileConfig) (Backend, bool) {
	if formatter, err := buildFormatter(config.format); !err {
		if target, err1 := os.OpenFile(config.target, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err1 == nil {
			backend := NewLogBackend(target, config.prefix, 0)
			return NewBackendFormatter(backend, formatter), false
		} else {
			debugf("couldn't openFile '%s'\n\n%s", config.target, err1)
		}
	}
	return nil, true
}

func buildMemory(config memoryConfig) (Backend, bool) {
	return NewMemoryBackend(config.size), false
}

type Config struct {
	Formatter map[string]interface{} `yaml:"formatter"`
	Backend   map[string]interface{} `yaml:"backend"`
	Logger    map[string]interface{} `yaml:"logger"`
}

const (
	mapOutput = "output"
	mapFormat = "format"
	mapPrefix = "prefix"
)

func (config Config) GetFormatter(name string) (ret string) {
	if formatter, ok := config.Formatter[name]; ok {
		mapping := formatter.(map[interface{}]interface{})
		if val, ok2 := mapping[mapFormat]; ok2 {
			ret = val.(string)
		} else {
			log.Fatalf("formatter '%s' is missing 'format'", name)
		}
	} else {
		log.Fatalf("undefined formatter '%s'", name)
	}

	return
}

func (config Config) buildAndSetBackends(){
	result := make([]Backend, 0)
	for _, name := range config.getBackends() {
		switch typ := config.getBackend(name).(type) {
		case stdoutConfig:
			if backend, loop := buildStdout(typ); loop {
				continue
			} else {
				result = append(result, backend)
			}
		case fileConfig:
			if backend, loop := buildFile(typ); loop {
				continue
			} else {
				result = append(result, backend)
			}
		case memoryConfig:
			if backend, loop := buildMemory(typ); loop {
				continue
			} else {
				result = append(result, backend)
			}
		case syslogConfig:
			if backend, loop := buildSyslog(typ); loop {
				continue
			} else {
				result = append(result, backend)
			}
		default:
			log.Fatalf("undefined implementation %#v", typ)
		}
	} // for _, name range config.getBackends()

	if len(result) > 0 {
		SetBackend(result...)
	}
}

func (config Config) getBackends() (ret []string) {
	ret = make([]string, len(config.Backend))
	i := 0
	for name := range config.Backend {
		ret[i] =name
		i++
	}

	debugf("%#v", ret)
	return
}

func (config Config) getBackend(name string) (ret iConfig) {
	tmp := config.Backend[name].(map[interface{}]interface{})
	useFormat := config.GetFormatter(tmp[mapFormat].(string))

	prefix := ""
	if pre, ok := tmp[mapPrefix]; ok {
		prefix = pre.(string)
	}

	switch strings.ToUpper(tmp[mapOutput].(string)) {
	case toFile:
		ret = config.buildFileConfig(tmp, name, prefix, useFormat)
		break
	case toMemory:
		ret = config.buildMemoryConfig(tmp)
		break
	case toStdout:
		ret = config.buildStdoutConfig(tmp, prefix, useFormat)
		break
	case toSyslog:
		ret = config.buildSyslogConfig(tmp, prefix)
		break
	default:
		log.Fatalf("Unkown output '%s'", tmp[mapOutput])
		break
	}
	return
}

func (config Config) setLoggerLevel() {
	for _, name := range config.getLoggers() {
		module := name
		if "root" == strings.ToLower(name) {
			module = ""
		}

		debugf( "SetLevel\tName: %s\tLevel: %#v", name, config.getLevel(name))
		SetLevel(config.getLevel(name), module)
	}
}
func (config Config) getLoggers() (ret []string){
	ret =make([]string , len(config.Logger))

	i := 0
	for name := range config.Logger {
		ret[i] = name
		i++
	}
	return
}

func (config Config) getLevel(name string) (ret Level) {
	logger := config.Logger[name].(map[interface{}]interface{})

	if val, ok := logger["level"]; ok {
		if lvl, err := LogLevel(strings.ToUpper(val.(string))); err == nil {
			ret = lvl
		}else{
			debugf("Level for '%s' ist invalid\n%s", name, err)
			ret = lvl
		}
	}
	return
}


func (config Config) buildFileConfig(yaml map[interface{}]interface{}, name, prefix, format string) fileConfig {
	saveTo, err := filepath.Abs(yaml["target"].(string))
	if err != nil {
		log.Fatalf("Backend '%s', target failed %#v", name, err)
	}
	return fileConfig{format: format, prefix: prefix, target: saveTo}
}

func (config Config) buildMemoryConfig(yaml map[interface{}]interface{}) memoryConfig {
	if val, ok := yaml["size"]; ok {
		return memoryConfig{size: val.(int)}
	} else {
		return memoryConfig{size: 1000}
	}
}

func (config Config) buildStdoutConfig(yaml map[interface{}]interface{}, prefix, format string) stdoutConfig {
	var useColor = false
	if val, ok := yaml["color"]; ok {
		useColor, _ = val.(bool)
	}
	return stdoutConfig{prefix: prefix, color: useColor, format: format}
}

func (config Config) buildSyslogConfig(yaml map[interface{}]interface{}, prefix string) syslogConfig {
	usePriority := -1
	if val, ok := yaml["priority"]; ok {
		usePriority = val.(int)
	}
	return syslogConfig{prefix: prefix, priority: usePriority}
}

type iConfig interface {
	//
}

type stdoutConfig struct {
	prefix string
	format string
	color  bool
	iConfig
}

type fileConfig struct {
	prefix string
	format string
	target string
	iConfig
}

type memoryConfig struct {
	size int
	iConfig
}

type syslogConfig struct {
	prefix   string
	priority int
	iConfig
}

func debug ( msg ...interface{}){
	if DEBUG_CONFIG {
		log.Print(msg)
	}
}

func debugf ( format string, msg ...interface{}){
	if DEBUG_CONFIG {
		log.Printf(format, msg)
	}
}
