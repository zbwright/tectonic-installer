package testutils

import (
	"k8s.io/client-go/kubernetes"
	"os"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

const (
	// kubeconfigEnv is the environment variable that is checked for a the kubeconfig path to be loaded.
	kubeconfigEnv = "TEST_KUBECONFIG"
)

func NewClient() *kubernetes.Clientset {

	kcfgPath := os.Getenv(kubeconfigEnv)

	if len(kcfgPath) == 0 {
		log.Fatalf("no kubeconfig path in environment variable %s", kubeconfigEnv)
	}

	rules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kcfgPath}
	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	restConfig, err := cfg.ClientConfig()
	if err != nil {
		log.Fatalf("could not create client config: %v", err)
	}

	cs, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("could not create clientset: %v", err)
	}
	return cs
}
