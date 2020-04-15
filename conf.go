package beer

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type configManager struct {
	lock sync.RWMutex
	data map[string]interface{}
}

var conf *configManager

func init() {
	conf = &configManager{
		lock: sync.RWMutex{},
		data: make(map[string]interface{}),
	}
}

func Loadini(iniPath string) {
	_, err := os.Stat(iniPath)
	if err != nil {
		panic(err)
	}
	conf.loadini(iniPath)
}

func (c *configManager) loadini(iniPath string) {
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

func (c *configManager) setConf(key, val string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = val
}

func (c *configManager) parseLine(line string) (key string, value string) {
	key = line
	index := strings.Index(line, "=")
	if index > 0 {
		key = line[0:index]
		value = line[index+1 : len(line)]
	}

	return key, value
}

func (c *configManager) GetString(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	i, ok := c.data[key]
	if !ok {
		return ""
	}

	return i.(string)
}

func (c *configManager) GetInt(key string) int {
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

func (c *configManager) Set(key string, val interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = val
}

//Config 获取config管理器对象..
func Config() *configManager {
	return conf
}
