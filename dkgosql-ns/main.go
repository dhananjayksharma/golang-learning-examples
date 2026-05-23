package main

import (
	"dkgosql-ns/k8s"
	"log"
)

func main() {
	client, err := k8s.GetKubeClient()
	if err != nil {
		log.Fatal(err)
	}

	err = k8s.CreateNamespace(client, "payment-service-namespace-v2")
	if err != nil {
		log.Fatal(err)
	}
}
