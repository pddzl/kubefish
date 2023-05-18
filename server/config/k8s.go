package config

type K8s struct {
	KubeConfig string `mapstructure:"kube-config" json:"kube-config" yaml:"kube-config"` // k8s密钥
}
