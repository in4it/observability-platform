package main

import (
	"embed"
	"flag"
	"log"

	"github.com/in4it/go-devops-platform/auth/provisioning/scim"
	"github.com/in4it/go-devops-platform/licensing"
	"github.com/in4it/go-devops-platform/rest"
	localstorage "github.com/in4it/go-devops-platform/storage/local"
	"github.com/in4it/go-devops-platform/users"
	"github.com/in4it/observability-platform/pkg/observability"
)

var (
	//go:embed static
	assets embed.FS
)

func main() {
	var (
		httpPort  int
		httpsPort int
	)
	flag.IntVar(&httpPort, "http-port", 80, "http port to run server on")
	flag.IntVar(&httpsPort, "https-port", 443, "https port to run server on")
	flag.Parse()

	localStorage, err := localstorage.New()
	if err != nil {
		log.Fatalf("couldn't initialize storage: %s", err)
	}
	licenseUserCount, cloudType := licensing.GetMaxUsers(localStorage)

	userStore, err := users.NewUserStore(localStorage, licenseUserCount)
	if err != nil {
		log.Fatalf("startup failed: userstore initialization error: %s", err)
	}

	scimInstance := scim.New(localStorage, userStore, "")

	apps := map[string]rest.AppClient{
		"observability": observability.New(localStorage),
	}

	c, err := rest.NewContext(localStorage, rest.SERVER_TYPE_OBSERVABILITY, userStore, scimInstance, licenseUserCount, cloudType, apps)
	if err != nil {
		log.Fatalf("startup failed: %s", err)
	}

	rest.StartServer(httpPort, httpsPort, localStorage, c, assets)
}
