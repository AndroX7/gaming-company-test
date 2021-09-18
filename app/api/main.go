package main

import (
	"context"
	"fmt"
	"gaming-company-test/app/api/middleware"
	"gaming-company-test/app/api/server"
	"gaming-company-test/events"
	"gaming-company-test/events/subscriber"
	"gaming-company-test/lib/database_transaction"
	"gaming-company-test/lib/redis"
	"gaming-company-test/lib/s3"
	"gaming-company-test/lib/socketio"
	"gaming-company-test/lib/validators"
	"gaming-company-test/listener/database_transaction_callbacks"
	"gaming-company-test/listener/event_listeners"
	"gaming-company-test/service/artist"
	artistHTTP "gaming-company-test/service/artist/delivery/http"
	"gaming-company-test/service/response_cache"
	socketHTTP "gaming-company-test/service/socket/delivery/http"
	"log"
	"net/http"
	"os"
	"time"

	artistModule "gaming-company-test/service/artist/module"
	responCacheModule "gaming-company-test/service/response_cache/module"

	socketModule "gaming-company-test/service/socket/module"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gookit/event"
	"github.com/subosito/gotenv"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type libs struct {
	fx.Out

	S3                 s3.S3Client
	Redis              redis.Client
	Socket             socketio.SocketClient
	TransactionManager database_transaction.Client
}

type appEvents struct {
	fx.Out

	Subscriber subscriber.Client
	Event      events.Client
}

type appListeners struct {
	fx.Out

	DatabaseTransactionCallback database_transaction_callbacks.Client
	EventListener               event_listeners.Client
}

type handlers struct {
	fx.In

	ArtistHandler *artistHTTP.Handler
	SocketHandler *socketHTTP.Handler
}

func main() {
	log.Println("server is starting")

	loadEnv()

	// set log to file
	if os.Getenv("APP_ENV") != "development" {
		log.Println("running in ", os.Getenv("APP_ENV"), " environment")
		f, err := os.OpenFile("error-log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()

		//set output of logs to f
		log.SetOutput(f)
	}

	app := fx.New(
		fx.Provide(
			setupDatabase,
			initLibs,
			registerEvents,
		),
		artistModule.Module,
		responCacheModule.Module,
		socketModule.Module,
		fx.Invoke(
			validators.NewValidator,
			// startCron,
			startListeners,
			startServer,
		),
	)

	app.Run()
}

// func startCron(lc fx.Lifecycle, c *cron.Cron) {
// 	c.Start()

// 	lc.Append(fx.Hook{
// 		OnStop: func(ctx context.Context) error {
// 			c.Stop()
// 			return nil
// 		},
// 	})
// }

func startServer(lc fx.Lifecycle, handlers handlers) {
	m := middleware.New(
		middleware.Config{
			JwtToken: os.Getenv("JWT_SECRET"),
		},
		redis.NewClient(redis.Credentials{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}, os.Getenv("APP_ENV")),
	)

	h := server.BuildHandler(m,
		handlers.ArtistHandler,
		handlers.SocketHandler,
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(s *http.Server) {
				log.Printf("api is available at %s\n", s.Addr)
				if err := s.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}(s)
			return nil
		},
		OnStop: func(c context.Context) error {
			_ = s.Shutdown(c)
			log.Println("api gracefully stopped")
			return nil
		},
	})
}

func initLibs(lc fx.Lifecycle, db *gorm.DB) libs {
	l := libs{
		S3: s3.NewS3Client(),
		Redis: redis.NewClient(redis.Credentials{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		}, os.Getenv("APP_ENV")),
		Socket:             socketio.NewSocketClient(),
		TransactionManager: database_transaction.New(db),
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			_ = l.Redis.Close()

			return nil
		},
	})

	return l
}

func registerEvents(
	lc fx.Lifecycle,
) appEvents {
	e := appEvents{
		Event: events.Register(),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
	})

	return e
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		log.Println("failed to load from .env")
	}
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return db
}

func startListeners(
	lc fx.Lifecycle,
	db *gorm.DB,
	events events.Client,
	responseCacheUsecase response_cache.Usecase,
	artistUsecase artist.Usecase,
) appListeners {
	l := appListeners{
		DatabaseTransactionCallback: database_transaction_callbacks.New(db, events, responseCacheUsecase, artistUsecase),
		EventListener:               event_listeners.New(responseCacheUsecase),
	}

	ignoreCache := false
	if ignoreCache || os.Getenv("RESPONSE_CACHE_ENABLED") == "false" || os.Getenv("RESPONSE_CACHE_ENABLED") == "" {
		ignoreCache = true
	}

	if !ignoreCache {
		//after gorm:create
		db.Callback().Create().Register("after_create_flush_response_cache", l.DatabaseTransactionCallback.FlushResponseCache)

		// after gorm:update
		db.Callback().Update().Register("after_update_flush_response_cache", l.DatabaseTransactionCallback.FlushResponseCache)

		// after gorm:delete
		db.Callback().Delete().Register("after_delete_flush_response_cache", l.DatabaseTransactionCallback.FlushResponseCache)
	}

	//after gorm:create
	db.Callback().Create().Register("after_create_flush_redis_key", l.DatabaseTransactionCallback.FlushRedisKey)

	// after gorm:update
	db.Callback().Update().Register("after_update_flush_redis_key", l.DatabaseTransactionCallback.FlushRedisKey)

	// after gorm:delete
	db.Callback().Delete().Register("after_delete_flush_redis_key", l.DatabaseTransactionCallback.FlushRedisKey)
	e := appEvents{
		Subscriber: subscriber.New(l.EventListener),
	}

	event.AddSubscriber(e.Subscriber)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
	})

	return l
}
