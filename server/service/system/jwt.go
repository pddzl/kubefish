package system

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

	"github.com/pddzl/kubefish/server/global"
	"github.com/pddzl/kubefish/server/model/system"
)

type JwtService struct{}

// GetRedisJWT 获取jwt
func (jwtService *JwtService) GetRedisJWT(username string) (redisJWT string, err error) {
	redisJWT, err = global.KF_REDIS.Get(context.Background(), username).Result()
	return redisJWT, err
}

// SetRedisJWT jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(username string, jwt string) (err error) {
	// 此处过期时间等于jwt过期时间
	err = global.KF_REDIS.Set(context.Background(), username, jwt, time.Duration(global.KF_CONFIG.JWT.ExpiresTime)*time.Second).Err()
	return err
}

// JoinInBlacklist 拉黑jwt
func (jwtService *JwtService) JoinInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.KF_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	return
}

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.KF_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}
