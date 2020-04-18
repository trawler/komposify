package cna

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kubernetes/kompose/pkg/kobject"
)

func sanitize(kObject *kobject.KomposeObject) kobject.KomposeObject {
	deleteDockerLbLabels(kObject)
	deleteSecrets(kObject)
	return *kObject
}

func deleteSecrets(opt *kobject.KomposeObject) {
	//for secret := range opt.Secrets {
	//}
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
