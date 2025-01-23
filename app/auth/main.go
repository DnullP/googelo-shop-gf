package main

import (
	"fmt"
	"googelo-shop-gf/app/auth/internal/controller/auth"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/contrib/trace/otlpgrpc/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 配置检查
	var ctx = gctx.New()
	test, err := g.Cfg().Get(ctx, "test", "no")
	if err != nil {
	}
	fmt.Println(test)

	// 链路追踪
	serviceName := "auth-service"
	endpoint, err := g.Cfg().Get(ctx, "Jaeger.address", "127.0.0.1:4317")
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	g.Log().Debugf(ctx, "Jaeger at %s", endpoint.String())
	traceToken := "******_******"
	shutdown, err := otlpgrpc.Init(serviceName, endpoint.String(), traceToken)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	defer shutdown(ctx)

	// Database Test
	testDB := g.Model("test")
	_, err = testDB.Data(g.Map{
		"name": "jack",
		"age":  123,
	}).Insert()
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// Redis test
	ExampleCache_SetAdapter()

	g.Log().Info(ctx)

	s := grpcx.Server.New()
	auth.Register(s)
	s.Run()
}

func ExampleCache_SetAdapter() {
	var (
		err         error
		ctx         = gctx.New()
		cache       = gcache.New()
		redisConfig = &gredis.Config{
			Address: "redis-service.shop.svc.cluster.local",
			Db:      9,
		}
		cacheKey   = `key`
		cacheValue = `value`
	)
	// Create redis client object.
	redis, err := gredis.New(redisConfig)
	if err != nil {
		panic(err)
	}
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	// Set and Get using cache object.
	err = cache.Set(ctx, cacheKey, cacheValue, time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// Get using redis client.
	fmt.Println(redis.MustDo(ctx, "GET", cacheKey).String())

	// Output:
	// value
	// value
}
