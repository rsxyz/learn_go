openssl req -new -newkey rsa:2048  \
    -nodes -x509 -subj '/CN=self-signed.ignore' -days 1800 \
    -keyout tls.key \
    -out tls.crt