package consul_client

import (
	"time"
	"log"

	consul "github.com/hashicorp/consul/api"
)

// type Service struct {
// 	Name        string
// 	URL        string
// 	ConsulAgent *consul.Agent
// }

// register (or update) the service
func RegisterService(AppName string, URL string) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

    pair := &api.KVPair{Key: AppName, Value: []byte(URL)}

    if _, err := c.Put(pair, nil); err != nil {
        log.Fatal(err)
    } else {
    	log.Println("register (or update) service '", pair.Key, "' successfully")
    }
}

// deregister the service
func DeRegisterService(AppName string, URL string) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

    pair := &api.KVPair{Key: AppName}

    if _, err := c.Delete(pair.Key, nil); err != nil {
        log.Fatal(err)
    } else {
        log.Println("deregister service '", pair.Key, "' successfully")
    }
}

// list all services
func ListAllService() {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

    // TODO
}

func GetURLFromAppName(AppName string) {
	c, err := consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return nil, err
	}

	URL, _, err := consul.KV().Get(AppName, nil)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	if URL.Value == nil {
		fmt.Fprintf(w, "Configuration empty")
		return
	}
	val := string(URL.Value)

	return val
}