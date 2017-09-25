// Package config defines the config structs and some config parser interfaces and implementations
package config

import (
	"errors"
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/cep21/xdgbasedir"
	"github.com/joho/godotenv"
	// "github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend/encoding"
	// "github.com/roscopecoltran/krakend/logging"
	// "github.com/roscopecoltran/krakend/utils"
	// "github.com/yudppp/structs"
)

/*
	refs:
	- https://github.com/daichirata/yaml2env
	- https://yudppp.github.io/json2struct/
	- github.com/yudppp/structs

    user := structs.NewExample(User{}).(User)
    fmt.Println(user.Name) # -> ichiro

    user = structs.NewDefault(User{}).(User)
    fmt.Println(user.Name) # -> suzuki

*/
var Config = struct {
	Env struct {
		Files []string          `json:"files" yaml:"files"`
		Keys  map[string]string `json:"kvs" yaml:"kvs"`
	} `json:"envs" yaml:"envs" file:".env"`
	Server    ServerConfig    `json:"server" yaml:"server" file:"providers"`
	Datastore DatabaseConfig  `json:"datastores" yaml:"datastores" file:"providers"`
	Gates     []GatewayConfig `json:"gateways" yaml:"gateways" file:"providers"`
	Providers []Provider      `json:"providers" yaml:"providers" file:"providers"`
	// Tasks     []dog.Dogfile   `json:"tasks" yaml:"tasks"`
}{}

const (
	// BracketsRouterPatternBuilder uses brackets as route params delimiter
	BracketsRouterPatternBuilder = iota
	// ColonRouterPatternBuilder use a colon as route param delimiter
	ColonRouterPatternBuilder
	// Time822 formt time for RFC 822
	Time822 = "02 Jan 2006 15:04:05 -0700" // "02 Jan 06 15:04 -0700"
)

// RoutingPattern to use during route conversion. By default, use the colon router pattern
var RoutingPattern = ColonRouterPatternBuilder

type ExtraConfig map[string]interface{}

// ConfigGetter is a function for parsing ExtraConfig into a previously know type
type ConfigGetter func(ExtraConfig) interface{}

// DefaultConfigGetter is the Default implementation for ConfigGetter, it just returns the ExtraConfig map.
func DefaultConfigGetter(extra ExtraConfig) interface{} { return extra }

const defaultNamespace = "github.com/roscopecoltran/krakend/config"

var configDirectoryPath string

// ConfigGetters map than match namespaces and ConfigGetter so the components knows which type to expect returned by the
// ConfigGetter ie: if we look for the defaultNamespace in the map, we will get the DefaultConfigGetter implementation
// which will return a ExtraConfig when called
var ConfigGetters = map[string]ConfigGetter{defaultNamespace: DefaultConfigGetter}

var (
	simpleURLKeysPattern = regexp.MustCompile(`\{([a-zA-Z\-_0-9]+)\}`)
	debugPattern         = "^[^/]|/__debug(/.*)?$"
	errInvalidHost       = errors.New("invalid host")
	defaultPort          = 8080
)

var (
	DefaultGatewayModels = []interface{}{&ServiceConfig{}, &EndpointConfig{}, &Backend{}} // &Group{}, &Node{}, &Topic{}, &Plugin{},
)

func init() {
	baseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		log.Fatal("config/config.go > Can't find XDG BaseDirectory")
	} else {
		configDirectoryPath = path.Join(baseDir, "krakend")
	}
}

// refs
// - https://github.com/gooops/env_strings

// Init initializes the configuration struct and its defined endpoints and backends.
// Init also sanitizes the values, applies the default ones whenever necessary and
// normalizes all the things.
func (s *ServiceConfig) Init() error {

	if len(Config.Env.Files) == 0 {
		Config.Env.Files = append(Config.Env.Files, ".env")
		// pp.Println(Config.Env.Files)
	}

	Config.Env.Keys, _ = godotenv.Read(Config.Env.Files...)
	// fmt.Println("config/config.go > Config.Env.Keys: ")
	// pp.Println(Config.Env.Keys)

	s.uriParser = NewURIParser()

	if s.Version != 1 {
		return fmt.Errorf("config/config.go > Unsupported version: %d\n", s.Version)
	}

	if s.Port == 0 {
		s.Port = Config.Server.Port
	}

	s.Host = s.uriParser.CleanHosts(s.Host)

	for i, e := range s.Endpoints {

		e.Endpoint = s.uriParser.CleanPath(e.Endpoint)

		if err := e.validate(); err != nil {
			return err
		}

		inputParams := s.extractPlaceHoldersFromURLTemplate(e.Endpoint, endpointURLKeysPattern)
		inputSet := map[string]interface{}{}

		for ip := range inputParams {
			inputSet[inputParams[ip]] = nil
		}

		e.Endpoint = s.uriParser.GetEndpointPath(e.Endpoint, inputParams)

		s.initEndpointDefaults(i)

		for j, b := range e.Backend {

			s.initBackendDefaults(i, j)

			// fmt.Println("initBackendDefaults: ")
			// pp.Print(b)

			b.Method = strings.ToTitle(b.Method)

			if err := s.initBackendURLMappings(i, j, inputSet); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *ServiceConfig) extractPlaceHoldersFromURLTemplate(subject string, pattern *regexp.Regexp) []string {
	matches := pattern.FindAllStringSubmatch(subject, -1)
	keys := make([]string, len(matches))
	for k, v := range matches {
		keys[k] = v[1]
	}
	return keys
}

func (s *ServiceConfig) initEndpointDefaults(e int) {
	endpoint := s.Endpoints[e]

	if endpoint.Method == "" {
		endpoint.Method = "GET"
	} else {
		endpoint.Method = strings.ToTitle(endpoint.Method)
	}
	if s.CacheTTL != 0 && endpoint.CacheTTL == 0 {
		endpoint.CacheTTL = s.CacheTTL
	}
	if s.Timeout != 0 && endpoint.Timeout == 0 {
		endpoint.Timeout = s.Timeout
	}
	if endpoint.ConcurrentCalls == 0 {
		endpoint.ConcurrentCalls = 1
	}
}

func (s *ServiceConfig) initBackendDefaults(e, b int) {
	endpoint := s.Endpoints[e]
	backend := endpoint.Backend[b]
	for n, h := range backend.Header {
		for k, v := range Config.Env.Keys {
			holderKey := fmt.Sprintf("{%s}", strings.Replace(k, "\"", "", -1))
			h = strings.Replace(h, holderKey, v, -1) // backend.Header[n]
			s.Endpoints[e].Backend[b].Header[n] = strings.Trim(h, " ")
			backend.Header[n] = strings.Trim(h, " ")
		}
	}
	if len(backend.Host) == 0 {
		backend.Host = s.Host
	} else if !backend.HostSanitizationDisabled {
		backend.Host = s.uriParser.CleanHosts(backend.Host)
	}
	if backend.Method == "" {
		backend.Method = endpoint.Method
	}
	backend.Timeout = endpoint.Timeout
	backend.ConcurrentCalls = endpoint.ConcurrentCalls

	switch strings.ToLower(backend.Encoding) {
	case encoding.XML:
		backend.Decoder = encoding.NewXMLDecoder(backend.IsCollection)
	//case encoding.XPATH:
	//	backend.Decoder = encoding.NewXPATHDecoder(backend.IsCollection)
	case encoding.RSS:
		backend.Decoder = encoding.NewRSSDecoder()
	default:
		backend.Decoder = encoding.NewJSONDecoder(backend.IsCollection)
	}
}

func (s *ServiceConfig) initBackendURLMappings(e, b int, inputParams map[string]interface{}) error {
	backend := s.Endpoints[e].Backend[b]

	backend.URLPattern = s.uriParser.CleanPath(backend.URLPattern)

	outputParams := s.extractPlaceHoldersFromURLTemplate(backend.URLPattern, simpleURLKeysPattern)

	outputSet := map[string]interface{}{}
	for op := range outputParams {
		outputSet[outputParams[op]] = nil
	}

	if len(outputSet) > len(inputParams) {
		return fmt.Errorf("config/config.go > initBackendURLMappings > Too many output params! input: %v, output: %v\n", outputSet, outputParams)
	}

	tmp := backend.URLPattern
	backend.URLKeys = make([]string, len(outputParams))
	for o := range outputParams {
		if _, ok := inputParams[outputParams[o]]; !ok {
			return fmt.Errorf("config/config.go > initBackendURLMappings > Undefined output param [%s]! input: %v, output: %v\n", outputParams[o], inputParams, outputParams)
		}
		tmp = strings.Replace(tmp, "{"+outputParams[o]+"}", "{{."+strings.Title(outputParams[o])+"}}", -1)
		backend.URLKeys = append(backend.URLKeys, strings.Title(outputParams[o]))
	}
	backend.URLPattern = tmp
	return nil
}

func (e *EndpointConfig) validate() error {
	matched, err := regexp.MatchString(debugPattern, e.Endpoint)
	if err != nil {
		log.Printf("config/config.go > validate > ERROR: parsing the endpoint url [%s]: %s. Ignoring\n", e.Endpoint, err.Error())
		return err
	}
	if matched {
		return fmt.Errorf("config/config.go validate > > ERROR: the endpoint url path [%s] is not a valid one!!! Ignoring\n", e.Endpoint)
	}

	if len(e.Backend) == 0 {
		return fmt.Errorf("config/config.go validate > > WARNING: the [%s] endpoint has 0 backends defined! Ignoring\n", e.Endpoint)
	}
	return nil
}
