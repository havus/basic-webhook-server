

# Do and Don't

## Naming
1. Use camelCase for variable name.
2. Use "Insert" in repository only, don't use "Insert" outside repository folder.
3. Use "Find" in repository only, don't use "Find" outside repository folder.

## Data structure
1. Don't use pointer if you don't want to modify it.

## Deployment to linux server
1. Create server
2. Update server
```sh
sudo apt-get update
sudo apt-get upgrade

# if unattended-upgrades error run code below
sudo nano /etc/apt/apt.conf.d/20auto-upgrades
ps aux | grep -i apt
sudo dpkg-reconfigure unattended-upgrades
sudo kill <PID>
```

3. Install docker, [link](https://docs.docker.com/engine/install/ubuntu/)
4. `sudo apt-get docker-compose`
5. Prepare file config (nginx, .pem, docker-compose)
```sh
mkdir nginx-conf
nano nginx-conf/nginx.conf
# and copy from ./nginx_1.conf

mkdir dhparam
sudo openssl dhparam -out ./dhparam/dhparam-2048.pem 2048

nano docker-compose.yml
# and copy from ./docker-compose.yml
```
6. run docker compose with `docker-compose up -d`
7. check logs: `docker-compose logs certbot`, we will see "Successfully received certificate."
8. rm nginx-conf/nginx.conf
9. nano nginx-conf/nginx.conf, copy from nginx_2.conf
10. docker-compose stop server_proxy
11. docker-compose up -d --force-recreate --no-deps webhook

<br>

### Problem development in local
1. When you facing error `error parsing uri: lookup staging.123123.mongodb.net on 172.20.10.1:53: cannot unmarshal DNS message`
try to edit /etc/resolv.conf with 8.8.8.8

### Docker run local
```sh
$ docker run --rm -p 8000:8000 havus/go-webhook-server:1.0-rc
```

# Webhook proxy server
Url: docker.io/havus/webhook-nginx

## How to build image
> export $(cat .env | grep CR_PAT=)
> ./build_push_image.sh

### References
- https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values
- https://www.digitalocean.com/community/tutorials/how-to-secure-a-containerized-node-js-application-with-nginx-let-s-encrypt-and-docker-compose
- https://stackoverflow.com/questions/62229938/gobwas-ws-clean-conn-close
- https://stackoverflow.com/questions/60541611/mongodb-dump-fails-with-cannot-unmarshal-dns-message
- https://medium.com/free-code-camp/million-websockets-and-go-cc58418460bb
- https://stackoverflow.com/questions/35479041/how-to-convert-iso-8601-time-in-golang