package main

import (
	"flag"
	"net/http"

	"github.com/isaacgr/nmos-node-generator/capabilities/is04"
)

var configFile = flag.String(
	"config",
	"config.json",
	"Conifg file containing resource generation info",
)
var useRandomNode = flag.Bool(
	"random-node-id",
	true,
	"Pass this flag to use a non-random UUID for the node",
)
var useRandomDevice = flag.Bool(
	"random-device-id",
	true,
	"Pass this flag to use a non-random UUID for the device",
)
var useRandomResource = flag.Bool(
	"random-resource-id",
	true,
	"Pass this flag to use a non-random UUID for the device's resources",
)
var requestTimeout = flag.Int(
	"request-timeout",
	20,
	"Set the timeout for HTTP requests",
)
var noKeepalive = flag.Bool(
	"connection-keepalive",
	true,
	"Pass this flag to use a persistent HTTP connection",
)

func main() {
	flag.Parse()

	server := http.NewServeMux()
	is04.RegisterRoutes(server)

	http.ListenAndServe(":8080", server)
	// randomDeviceUUID := *useRandomDevice
	// randomResourceUUID := *useRandomResource
	// randomNodeUUID := *useRandomNode
	// httpRequestTimeout := time.Duration(*requestTimeout)
	// connectionKeepalive := *noKeepalive
	//
	// config := config.New()
	// baseUrl := config.Registry.Scheme + "://" + config.Registry.IP
	// port := config.Registry.Port
	//
	//	transport := &http.Transport{
	//		DisableKeepAlives:   false,
	//		MaxIdleConns:        0,
	//		MaxIdleConnsPerHost: math.MaxInt64,
	//		IdleConnTimeout:     300 * time.Second,
	//		TLSClientConfig: &tls.Config{
	//			InsecureSkipVerify: true,
	//		},
	//	}
	//
	//	if connectionKeepalive == false {
	//		transport = &http.Transport{
	//			TLSClientConfig: &tls.Config{
	//				InsecureSkipVerify: true,
	//			},
	//		}
	//	}
	//
	//	httpclient := &http.Client{
	//		Transport: transport,
	//		Timeout:   httpRequestTimeout * time.Second,
	//	}
}
