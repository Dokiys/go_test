genrsa:
	 openssl genrsa -out ca.key 2048

gen_default_ca:
	 openssl req -x509 -new -nodes -key ca.key -subj "/CN=example.ca.com" -days 5000 -out ca.crt

gen_ca:
	openssl req -new -sha256 \
		-key ca.key \
		-subj "/C=CN/ST=Beijing/L=Beijing/O=UnitedStack/OU=Devops/CN=www.zchd.ltd" \
		-reqexts SAN \
		-config <(cat /System/Library/OpenSSL/openssl.cnf \
			<(printf "[SAN]\nsubjectAltName=DNS:www.zchd.ltd,DNS:www.test.zchd.ltd")) \
		-out zchd.csr

sign_ca:
	openssl x509 -req -days 365000 \
        -in zchd.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
        -extfile <(printf "subjectAltName=DNS:www.zchd.ltd,DNS:www.test.zchd.ltd") \
        -out zchd.crt

