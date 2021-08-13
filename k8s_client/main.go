package main

import (
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

func main() {
	var kubeconfig = flag.String("kubeconfig", filepath.Join("/Users/deo", ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployment, err := clientset.AppsV1().Deployments("bigwin").Get("bigwin-backend", v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	_, err = clientset.AppsV1().Deployments("bigwin").Update(deployment)
	if err != nil {
		panic("发布失败")
	}
	fmt.Println("执行完成")
}
