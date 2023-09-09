package main

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

func (c *ConfigBuilder) SetDomain(domain string) *ConfigBuilder{
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
		url: c.Url,
		key: c.Key,
		domain: c.Domain,
		port: c.Port,
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
	url string
	key string
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

func createDevConfig(userId int, url string, key string ) iConfig {
	return &devConfig{
		config: config{
			userId: userId,
			url: url,
			key: key,
		},
	}
}

func createProdConfig(userId int, url string, key string) iConfig {
	return &prodConfig{
		config: config{
			userId: userId,
			url: url,
			key: key,
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