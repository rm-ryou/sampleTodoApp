FROM nginx:1.27.0

COPY ./templates /etc/nginx/templates

COPY ./certs /etc/nginx/certs

# Running nginx in debug mode
# https://hub.docker.com/_/nginx
CMD [ "nginx-debug", "-g", "daemon off;" ]
