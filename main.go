package main

import (
	"github.com/kyma-project/runtime-watcher/kcp/istio/mtls2/certificate"
	"github.com/kyma-project/runtime-watcher/kcp/istio/mtls2/server"
)

const (
	clientPublicKey  = "istio/mtls2/certs/client.pem"
	clientPrivateKey = "istio/mtls2/certs/client.key"

	serverPublicKey  = "istio/mtls2/certs/server.pem"
	serverPrivateKey = "istio/mtls2/certs/server.key"

	rootPublicKey  = "istio/mtls2/certs/ca.pem"
	rootPrivateKey = "istio/mtls2/certs/ca.key"
)

func main() {
	certificate.Verification(
		"localhost",
		clientPublicKey, clientPrivateKey,
		serverPublicKey, serverPrivateKey,
		rootPublicKey, rootPrivateKey,
		true,
	)

	srv := server.HTTPServer(serverPublicKey, serverPrivateKey, rootPublicKey)
	defer srv.Close()

	server.SendRequest(srv.URL, clientPublicKey, clientPrivateKey, rootPublicKey)
}
