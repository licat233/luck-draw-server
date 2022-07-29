package listen

import (
	"context"
	"luck-draw/service/v1/luckdraw/mq/internal/config"
	"luck-draw/service/v1/luckdraw/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config) []service.Service {
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq: pub sub
	services = append(services, KqMqs(c, ctx, svcCtx)...)
	return services
}
