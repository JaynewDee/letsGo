package patterns

import (
	"fmt"
)

/*
BUILDER
*/
type Config struct {
	userId int
	url    string
	key    string
	domain string
	port   int
}

type ConfigBuilder struct {
	UserId int
	Url    string
	Key    string
	Domain string
	Port   int
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (c *ConfigBuilder) SetUserId(id int) *ConfigBuilder {
	c.UserId = id
	return c
}

func (c *ConfigBuilder) SetUrl(url string) *ConfigBuilder {
	c.Url = url
	return c
}

func (c *ConfigBuilder) SetKey(key string) *ConfigBuilder {
	c.Key = key
	return c
}

func (c *ConfigBuilder) SetDomain(domain string) *ConfigBuilder {
	c.Domain = domain
	return c
}

func (c *ConfigBuilder) SetPort(port int) *ConfigBuilder {
	c.Port = port
	return c
}

func (c *ConfigBuilder) Build() *Config {
	return &Config{
		userId: c.UserId,
		url:    c.Url,
		key:    c.Key,
		domain: c.Domain,
		port:   c.Port,
	}
}

/*
FACTORY
*/
type iConfig interface {
	setUserId(id int) *config
	setUrl(url string) *config
	setKey(key string) *config
	getUserId() int
	getUrl() string
	getKey() string
}

type config struct {
	userId int
	url    string
	key    string
}

type devConfig struct {
	config
}

type prodConfig struct {
	config
}

func (c *config) setUserId(id int) *config {
	c.userId = id
	return c
}
func (c *config) setUrl(url string) *config {
	c.url = url
	return c
}
func (c *config) setKey(key string) *config {
	c.key = key
	return c
}
func (c *config) getUserId() int {
	return c.userId
}
func (c *config) getUrl() string {
	return c.url
}
func (c *config) getKey() string {
	return c.key
}

func createDevConfig(userId int, url string, key string) iConfig {
	return &devConfig{
		config: config{
			userId: userId,
			url:    url,
			key:    key,
		},
	}
}

func createProdConfig(userId int, url string, key string) iConfig {
	return &prodConfig{
		config: config{
			userId: userId,
			url:    url,
			key:    key,
		},
	}
}

func NewFactoryConfig(configType string, userId int, url string, key string) iConfig {
	if configType == "dev" {
		return createDevConfig(userId, url, key)
	}
	if configType == "prod" {
		return createProdConfig(userId, url, key)
	}

	return nil
}

/*
SINGLETON
*/
// type Logger struct {
// 	logLevel int
// }

// func (l *Logger) Log(s string) {
// 	fmt.Println(l.logLevel, ":", s)
// }

// func (l *Logger) SetLogLevel(lvl int) {
// 	l.logLevel = lvl
// }

// var logger *Logger
// var once sync.Once

// func getLoggerInstance() *Logger {
// 	// Wrapping instantiation in once.Do() provides concurrency safety to singleton!
// 	once.Do(func() {
// 		if logger == nil {
// 			fmt.Println("Creating Logger Singleton Instance ... ")
// 			logger = &Logger{}
// 		}
// 	})

// 	fmt.Println("Returning Logger Singleton ... ")

// 	return logger
// }

// func LogConcurrently() {
// 	// Demonstrates that Logger is only created a single time
// 	for i := 1; i < 10; i++ {
// 		go getLoggerInstance()
// 	}

//		fmt.Scanln()
//	}

type Log = func(s string)
type Level = func() int
type SetLevel = func(lvl int)

func Logger(logLevel int) (Log, Level, SetLevel) {
	return func(s string) {
			fmt.Println(logLevel, ":", s)
		},

		func() int {
			return logLevel
		},

		func(lvl int) {
			logLevel = lvl
		}

}

// type ILogger interface {
// 	Log(s string)
// 	Level() int
// 	SetLevel(lvl int)
// }

// type Logger struct {
// 	logLevel int
// }

// func (l *Logger) Log(s string) {
// 	fmt.Println(l.logLevel, ":", s)
// }

// func (l *Logger) Level() int {
// 	return l.logLevel
// }

// func (l *Logger) SetLevel(lvl int) {
// 	l.logLevel = lvl
// }
