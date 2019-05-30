package consul_client

import (
	"log"
	// "fmt"
	consul "github.com/hashicorp/consul/api"
)

// type Service struct {
// 	Name        string
// 	URL        string
// 	ConsulAgent *consul.Agent
// }

func StoreKeyValue(key string, value string) {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

    pair := &consul.KVPair{Key: key, Value: []byte(value)}

    if _, err := c.KV().Put(pair, nil); err != nil {
        log.Fatal(err)
    } else {
    	log.Println("store (or update) key '", pair.Key, "' successfully")
    }
}

func DeleteKeyValue(key string) {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

    pair := &consul.KVPair{Key: key}

    if _, err := c.KV().Delete(pair.Key, nil); err != nil {
        log.Fatal(err)
    } else {
        log.Println("remove key '", pair.Key, "' successfully")
    }
}

// list all services
// func ListAllService() error {
// 	// c, err := consul.NewClient(consul.DefaultConfig())
// 	// if err != nil {
// 	// 	return err
// 	// }

//     // TODO
//     return err
// }

func GetKeyValue(key string) string {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	
	if err != nil {
		log.Fatal(err)
	}

	kv, _, err := c.KV().Get(key, nil)
	if err != nil {
		log.Fatal(err)
	}

	if kv == nil {
		log.Println("Configuration empty")
		return ""
	}

	value := string(kv.Value)

	return value
}