package conf

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type ConfigManager struct {
	lock sync.RWMutex
	data map[string]interface{}
}

var Manager *ConfigManager

func init() {
	conf := &ConfigManager{
		lock: sync.RWMutex{},
		data: make(map[string]interface{}),
	}

	Manager = conf
}

func (c *ConfigManager) Loadini(iniPath string) {
	_, err := os.Stat(iniPath)
	if err != nil {
		panic(err)
	}
	c.loadini(iniPath)
}

func (c *ConfigManager) loadini(iniPath string) {
	b, err := ioutil.ReadFile(iniPath)
	if err != nil {
		panic(err)
	}
	confLine := strings.Split(string(b), "\n")

	for _, line := range confLine {
		//替换掉当前行里面所有的空格.
		line = strings.Replace(line, " ", "", -1)

		//判断是否为注释.
		if strings.HasPrefix(line, "#") {
			continue
		}

		//去掉换行符.
		if len(line) < 1 {
			continue
		}

		key, value := c.parseLine(line)
		c.setConf(key, value)
	}
}

func (c *ConfigManager) setConf(key, val string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = val
}

func (c *ConfigManager) parseLine(line string) (key string, value string) {
	key = line
	index := strings.Index(line, "=")
	if index > 0 {
		key = line[0:index]
		value = line[index+1 : len(line)]
	}

	return key, value
}

func (c *ConfigManager) GetString(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	i, ok := c.data[key]
	if !ok {
		return ""
	}

	return i.(string)
}

func (c *ConfigManager) GetInt(key string) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	i, ok := c.data[key]
	if !ok {
		return 0
	}

	n, err := strconv.Atoi(i.(string))
	if err != nil {
		log.Printf("err:%+v\n", err)
		return 0
	}
	return n
}

func (c *ConfigManager) Set(key string, val interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = val
}