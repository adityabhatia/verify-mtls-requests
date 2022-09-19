# verify-mtls-requests

``` shell
mkdir certs

openssl req -newkey rsa:2048 \
-nodes -x509 \
-days 365 \
-out certs/ca.pem \
-keyout certs/ca.key \
-subj "/C=US/ST=Frankfurt/L=Hessen/O=Aditya/OU=dev/CN=localhost"`
```