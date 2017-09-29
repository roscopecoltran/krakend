// Package gin provides some basic implementations for building routers based on gin-gonic/gin
package gin

import (
	"context"
	"fmt"
	"net/http"
	"time"
	// "go.uber.org/zap"

	"github.com/gin-contrib/gzip"
	// "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	// "github.com/gregjones/httpcache"
	// "github.com/gin-contrib/pprof"
	// "github.com/casbin/casbin"
	// "github.com/gin-contrib/authz"
	// "github.com/gin-contrib/size"
	// "github.com/vanng822/gin-csrf"

	"github.com/roscopecoltran/krakend/config"
	"github.com/roscopecoltran/krakend/logging"
	"github.com/roscopecoltran/krakend/proxy"
	"github.com/roscopecoltran/krakend/router"
)

// Config is the struct that collects the parts the router should be builded from
type Config struct {
	Engine         *gin.Engine
	Middlewares    []gin.HandlerFunc
	HandlerFactory HandlerFactory
	ProxyFactory   proxy.Factory
	Logger         logging.Logger
}

/*
// ref. https://github.com/stevokk/netflixplusplus/blob/master/api/main.go
var reqClient = &http.Client{
	Timeout:   10 * time.Second,
	Transport: httpcache.NewMemoryCacheTransport(),
}
*/

// DefaultFactory returns a gin router factory with the injected proxy factory and logger.
// It also uses a default gin router and the default HandlerFactory
func DefaultFactory(proxyFactory proxy.Factory, logger logging.Logger) router.Factory {
	return NewFactory(
		Config{
			Engine:         gin.Default(),
			Middlewares:    []gin.HandlerFunc{},
			HandlerFactory: EndpointHandler,
			ProxyFactory:   proxyFactory,
			Logger:         logger,
		},
	)
}

// NewFactory returns a gin router factory with the injected configuration
func NewFactory(cfg Config) router.Factory {
	return factory{cfg}
}

type factory struct {
	cfg Config
}

// New implements the factory interface
func (rf factory) New() router.Router {
	return ginRouter{rf.cfg, context.Background()}
}

func (rf factory) NewWithContext(ctx context.Context) router.Router {
	return ginRouter{rf.cfg, ctx}
}

type ginRouter struct {
	cfg Config
	ctx context.Context
}

// Run implements the router interface
func (r ginRouter) Run(cfg config.ServiceConfig) {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	} else {
		r.cfg.Logger.Debug("Debug enabled")
	}

	r.cfg.Engine.RedirectTrailingSlash = true
	r.cfg.Engine.RedirectFixedPath = true
	r.cfg.Engine.HandleMethodNotAllowed = true

	r.cfg.Engine.Use(gzip.Gzip(gzip.DefaultCompression))
	r.cfg.Engine.Use(r.cfg.Middlewares...)

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	// r.cfg.Engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.cfg.Engine.Use(gin.Recovery())

	r.cfg.Engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// options := csrf.DefaultOptions()
	// options.MaxUsage = 10
	// options.MaxAge = 15 * 60
	// r.cfg.Engine.Use(csrf.Csrf(options))

	/*
		pprof.Register(r.cfg.Engine, &pprof.Options{
			// default is "debug/pprof"
			RoutePrefix: "debug/pprof",
		})
	*/

	// load the casbin model and policy from files, database is also supported.
	// e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	// r.cfg.Engine.Use(middleware.CORS())
	// r.cfg.Engine.Use(middleware.Recovery())
	// r.cfg.Engine.Use(authz.NewAuthorizer(e))

	/*
		if cfg.Admin {
			r.registerAdminEndpoints()
		}
	*/

	fmt.Printf("r.registerDebugEndpoints? %t \n", cfg.Debug)
	if !cfg.Debug {
		r.registerDebugEndpoints()
	}

	fmt.Printf("r.registerAdminEndpoints? %t \n", cfg.Admin)
	if !cfg.Admin {
		r.registerAdminEndpoints()
	}

	r.registerKrakendEndpoints(cfg.Endpoints)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r.cfg.Engine,
	}

	go func() {
		r.cfg.Logger.Critical(s.ListenAndServe())
	}()

	<-r.ctx.Done()
	r.cfg.Logger.Error(s.Shutdown(context.Background()))
}

func (r ginRouter) registerDebugEndpoints() {
	handler := DebugHandler(r.cfg.Logger)
	r.cfg.Engine.GET("/__debug/*param", handler)
	r.cfg.Engine.POST("/__debug/*param", handler)
	r.cfg.Engine.PUT("/__debug/*param", handler)
}

func (r ginRouter) registerAdminEndpoints() {
	handler := AdminHandler(r.cfg.Logger)
	r.cfg.Engine.Any("/admin/*w", handler)
}

func (r ginRouter) registerAuthEndpoints() {
	handler := AuthRequiredHandler(r.cfg.Logger)
	authorized := r.cfg.Engine.Group("/api")
	authorized.Use(handler)
	{
		authorized.GET("/*w", handler)
		authorized.POST("/*w", handler)
		authorized.PUT("/*w", handler)
	}
}

func (r ginRouter) registerKrakendEndpoints(endpoints []*config.EndpointConfig) {
	for _, c := range endpoints {
		proxyStack, err := r.cfg.ProxyFactory.New(c)
		if err != nil {
			r.cfg.Logger.Error("calling the ProxyFactory", err.Error())
			continue
		}

		r.registerKrakendEndpoint(c.Method, c.Endpoint, r.cfg.HandlerFactory(c, proxyStack), len(c.Backend))
	}
}

func (r ginRouter) registerKrakendEndpoint(method, path string, handler gin.HandlerFunc, totBackends int) {
	if method != "GET" && totBackends > 1 {
		r.cfg.Logger.Error(method, "endpoints must have a single backend! Ignoring", path)
		return
	}
	switch method {
	case "GET":
		r.cfg.Engine.GET(path, handler)
	case "POST":
		r.cfg.Engine.POST(path, handler)
	case "PUT":
		r.cfg.Engine.PUT(path, handler)
	case "PATCH":
		r.cfg.Engine.PATCH(path, handler)
	case "DELETE":
		r.cfg.Engine.DELETE(path, handler)
	default:
		r.cfg.Logger.Error("Unsupported method", method)
	}
}

/*

// https://github.com/lebedev-yury/cities-api/blob/master/router.go
func registerGeoHelpersEndpoints(db *bolt.DB, options *config.Options, c *cache.Cache) *gin.Engine {
	// router := gin.Default()
	// router.Use(middleware.CommonHeaders(options.CORSOrigins))

	handler := GeoHelperHandler(r.cfg.Logger)
	geoHelpers := r.cfg.Engine.Group("/api")

	v1 := geoHelpers.Group("/geo")
	{
		v1.GET("/app/status", MakeApplicationStatusEndpoint(db))
		v1.GET("/cities/:id", MakeCityEndpoint(db, options))
		v1.GET("/cities", MakeCitiesEndpoint(db, options, c))
	}

	return router
}
*/
