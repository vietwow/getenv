package consul_client

import (
	"log"
    "fmt"
	consul "github.com/hashicorp/consul/api"
)

// type Service struct {
// 	Name        string
// 	URL        string
// 	ConsulAgent *consul.Agent
// }

// register (or update) the service
func RegisterService(AppName string, URL string) {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

    pair := &consul.KVPair{Key: AppName, Value: []byte(URL)}

    if _, err := c.KV().Put(pair, nil); err != nil {
        log.Fatal(err)
    } else {
    	log.Println("register (or update) service '", pair.Key, "' successfully")
    }
}

// deregister the service
func DeRegisterService(AppName string) {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

    pair := &consul.KVPair{Key: AppName}

    if _, err := c.KV().Delete(pair.Key, nil); err != nil {
        log.Fatal(err)
    } else {
        log.Println("deregister service '", pair.Key, "' successfully")
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

func GetURLFromAppName(AppName string) string {
	config := consul.DefaultConfig()
	config.Address = "consul:8500"
	c, err := consul.NewClient(config)
	
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("hahaha")
	URL, _, err := c.KV().Get(AppName, nil)
	fmt.Println(URL.Value)
	fmt.Println("hihihi")
	if err != nil {
		log.Fatal(err)
	}
	if URL.Value == nil {
		log.Println("Configuration empty")
	}
	val := string(URL.Value)

	return val
}