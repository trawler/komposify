package cna

import (
	"reflect"
	"testing"

	"github.com/kubernetes/kompose/pkg/kobject"
)

func TestDeleteDockerLBLabels(t *testing.T) {
	service := kobject.ServiceConfig{
		ContainerName: "test-container-name",
		Image:         "test-image",
		DeployLabels: map[string]string{
			"com.docker.lb.hosts":   "test.example.com",
			"com.docker.lb.port":    "8080",
			"com.docker.lb.network": "test-tier",
			"kompose.service.type":  "clusterip",
		},
		Annotations: map[string]string{"test-annotations": "my-test-service"},
		Privileged:  true,
		Restart:     "always",
	}

	k := kobject.KomposeObject{
		ServiceConfigs: map[string]kobject.ServiceConfig{"test-app": service},
	}

	deployLabels := k.ServiceConfigs["test-app"].DeployLabels

	expected := map[string]string{
		"kompose.service.type": "clusterip",
	}
	deleteDockerLbLabels(&k)
	if !reflect.DeepEqual(expected, deployLabels) {
		t.Errorf("Deploy labels not sanitized. Expected:\n%q\nGot:\n%q", expected, deployLabels)
	}
}

/*
func TestDeleteServiceConfigs(t *testing.T) {
	kobject := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
		Secrets:        make(map[string]dockerCliTypes.SecretConfig),
	}
}*/
