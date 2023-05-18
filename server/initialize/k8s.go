package initialize

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/pddzl/kubefish/server/global"
)

func K8sConnect() (*kubernetes.Clientset, *rest.Config) {
	// client in pod
	config, err := rest.InClusterConfig()

	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", global.KF_CONFIG.K8s.KubeConfig)
		if err != nil {
			global.KF_LOG.Error("k8s connect failed", zap.Error(err))
			return nil, nil
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.KF_LOG.Error("k8s clientSet", zap.Error(err))
	}

	return clientSet, config
}
