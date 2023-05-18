package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/pddzl/kubefish/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	KF_VP                  *viper.Viper
	KF_CONFIG              config.Server
	KF_LOG                 *zap.Logger
	KF_DB                  *gorm.DB
	KF_REDIS               *redis.Client
	KF_Concurrency_Control = &singleflight.Group{}
	KF_K8S_Client          *kubernetes.Clientset
	KF_K8S_CONFIG          *rest.Config
)
