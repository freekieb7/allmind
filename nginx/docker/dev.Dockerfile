FROM nginx:latest

RUN rm /etc/nginx/conf.d/default.conf

COPY --chown=nginx:nginx ./private/freekkuijpers.nl.key /etc/ssl/freekkuijpers.nl/private/freekkuijpers.nl.key
COPY --chown=nginx:nginx ./certs/freekkuijpers.nl.pem /etc/ssl/freekkuijpers.nl/certs/freekkuijpers.nl.pem

COPY ./templates/ /etc/nginx/templates/