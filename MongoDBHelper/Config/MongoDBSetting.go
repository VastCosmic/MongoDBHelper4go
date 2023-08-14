package Config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type MongoDBSetting struct {
	Host string `json:"host"`
	Port string `json:"port"`
	//DbName string `json:"dbName"`
}

type Config struct {
	Setting MongoDBSetting
	mutex   sync.RWMutex
}

var instance *Config

// GetInstance 返回Config的单例实例
func GetInstance() *Config {
	if instance == nil {
		instance = &Config{}
		instance.loadSetting()
	}
	return instance
}

func (c *Config) loadSetting() error {
	file, err := os.Open("MongoDBHelper/Config/MongoDBSetting.json")
	if err != nil {
		fmt.Println("Open MongoDBSetting.json failed, err: ", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Close MongoDBSetting.json failed, err: ", err)
		}
	}(file)

	var setting MongoDBSetting
	err = json.NewDecoder(file).Decode(&setting)
	if err != nil {
		fmt.Println("Decode MongoDBSetting.json failed, err: ", err)
		return err
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Setting = setting

	return nil
}

func (c *Config) GetHost() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.Setting.Host
}

func (c *Config) GetPort() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.Setting.Port
}

//func (c *Config) GetDbName() string {
//	c.mutex.RLock()
//	defer c.mutex.RUnlock()
//	return c.Setting.DbName
//}

func (c *Config) SetHost(host string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Setting.Host = host
}

func (c *Config) SetPort(port string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Setting.Port = port
}

//func (c *Config) SetDbName(dbName string) {
//	c.mutex.Lock()
//	defer c.mutex.Unlock()
//	c.Setting.DbName = dbName
//}
