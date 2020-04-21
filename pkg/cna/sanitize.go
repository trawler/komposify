package cna

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kubernetes/kompose/pkg/kobject"
)

func sanitize(k *kobject.KomposeObject, t string) kobject.KomposeObject {
	switch t {
	case "services":
		deleteDockerLbLabels(k)
		deleteSecrets(k)
	case "secrets":
		deleteServiceConfigs(k)
	}

	return *k
}

func deleteSecrets(opt *kobject.KomposeObject) {
	for secret := range opt.Secrets {
		delete(opt.Secrets, secret)
	}
}

func deleteServiceConfigs(opt *kobject.KomposeObject) {
	for svcConfig := range opt.ServiceConfigs {
		delete(opt.ServiceConfigs, svcConfig)
	}
}

func deleteDockerLbLabels(opt *kobject.KomposeObject) {
	for svc := range opt.ServiceConfigs {
		deployLabels := opt.ServiceConfigs[svc].DeployLabels
		for key := range deployLabels {
			if strings.HasPrefix(key, "com.docker.lb") {
				delete(deployLabels, key)
			}
		}
	}
}

// PrettyPrint prints data structs for debugging
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return nil
}
