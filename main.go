package main

import (
	"github.com/kyma-project/verify-mtls-requests/certificate"
	"github.com/kyma-project/verify-mtls-requests/server"
)

const (
	clientPublicKey  = "certs/client.pem"
	clientPrivateKey = "certs/client.key"

	serverPublicKey  = "certs/server.pem"
	serverPrivateKey = "certs/server.key"

	rootPublicKey  = "certs/ca.pem"
	rootPrivateKey = "certs/ca.key"
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
