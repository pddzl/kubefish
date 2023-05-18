package system

import (
	"github.com/pddzl/kubefish/server/global"
)

type JwtBlacklist struct {
	global.KF_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
