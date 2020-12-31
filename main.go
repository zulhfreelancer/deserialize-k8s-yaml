package main

import (
	"bytes"
	"fmt"
	"os"
	k8Yaml "k8s.io/apimachinery/pkg/util/yaml"
	appsv1 "k8s.io/api/apps/v1"
)

var deploymentManifest = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
`

func main() {
	d := &appsv1.Deployment{}
	dec := k8Yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(deploymentManifest)), 1000)

	if err := dec.Decode(&d); err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Println("kind:", d.Kind) // output: Deployment
	fmt.Println("matchLabels:", d.Spec.Selector.MatchLabels["app"]) // output: nginx
	fmt.Println("containerPort:",d.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort) // output: 80
}
