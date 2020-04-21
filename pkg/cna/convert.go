package cna

import (
	"log"

	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/kubernetes/kompose/pkg/loader"
	"github.com/kubernetes/kompose/pkg/transformer"
	"github.com/kubernetes/kompose/pkg/transformer/kubernetes"
)

// Convert transforms and sanitizes docker compose or dab file to k8s objects
func Convert(opt kobject.ConvertOptions, objType string) error {
	komposeObject, err := getKomposeObject(opt)
	if err != nil {
		return nil
	}
	// Modify the Kompose Object
	newKompose := sanitize(komposeObject, objType)
	// Get a transformer that maps komposeObject to provider's primitives
	t := getTransformer(opt)

	// Do the transformation
	objects, err := t.Transform(newKompose, opt)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Print output
	err = kubernetes.PrintList(objects, opt)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return nil
}

func getKomposeObject(opt kobject.ConvertOptions) (*kobject.KomposeObject, error) {
	komposeObject := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
	}

	l, err := loader.GetLoader("compose")
	if err != nil {
		log.Fatal(err)
	}

	komposeObject, err = l.LoadFile(opt.InputFiles)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &komposeObject, nil
}

// Convenience method to return the appropriate Transformer based on
// what provider we are using.
func getTransformer(opt kobject.ConvertOptions) transformer.Transformer {
	// Create/Init new Kubernetes object with CLI opts
	return &kubernetes.Kubernetes{Opt: opt}
}
