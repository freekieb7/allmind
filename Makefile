up:
	docker compose up -d --remove-orphans

build:
	docker compose build

down:
	docker compose down

gen-cert:
	sudo openssl req -subj "/countryName=EN/stateOrProvinceName=Somewhere/organizationName=Personal/localityName=Somewhere/commonName=freekkuijpers.nl/organizationalUnitName=IT/emailAddress=freek@freekkuijpers.nl/" -x509 -nodes -days 3650 -newkey rsa:2048 -keyout ./nginx/private/freekkuijpers.nl.key -out ./nginx/certs/freekkuijpers.nl.pem