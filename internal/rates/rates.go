package rates

import (
	"github.com/majorchork/tech-crib-africa/config"
	"net/http"
)

type Client struct {
	Http   http.Client
	Config config.Config
}
