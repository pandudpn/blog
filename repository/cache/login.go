package cache

import (
	"context"
	"os"
	"strings"
	"time"
	
	"github.com/google/uuid"
	"github.com/pandudpn/blog/model"
	"github.com/pandudpn/blog/utils/hash"
)

func (c *cacheRepository) GetDataFromAccessToken(ctx context.Context, accessToken string) (int, error) {
	res := c.redis.Get(ctx, accessToken)
	if res.Err() != nil {
		return 0, res.Err()
	}
	
	val := res.Val()
	h := hash.NewHash()
	userId, err := h.Decode(val)
	if err != nil {
		return 0, err
	}
	
	return userId, nil
}

func (c *cacheRepository) SetLoginAccessToken(ctx context.Context, user *model.User) (model.ResponseLogin, error) {
	var res model.ResponseLogin
	
	uid := uuid.New().String()
	uid = strings.ReplaceAll(uid, "-", "")
	
	expired, err := time.ParseDuration(os.Getenv("EXPIRED_ACCESS_TOKEN"))
	if err != nil {
		expired = time.Duration(48) * time.Hour
	}
	
	h := hash.NewHash()
	userId, err := h.Encode(user.Id)
	if err != nil {
		return res, err
	}
	
	err = c.redis.Set(ctx, uid, userId, expired).Err()
	if err != nil {
		return res, err
	}
	
	now := time.Now().In(c.timezone).Add(expired)
	res.TokenType = "Bearer"
	res.AccessToken = uid
	res.ExpiredAt = now.Unix()
	
	return res, nil
}
