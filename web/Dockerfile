FROM node:16

WORKDIR /gva_web/
COPY . .

RUN #npm install -g cnpm && cnpm install && cnpm run build

FROM nginx:alpine
LABEL MAINTAINER="w17681264762@163.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /gva_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
