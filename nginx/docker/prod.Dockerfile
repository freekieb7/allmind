FROM nginx:latest

RUN rm /etc/nginx/conf.d/default.conf

COPY ./templates/ /etc/nginx/templates/