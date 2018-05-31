# README
[![Build Status](https://travis-ci.org/ildarusmanov/adfsuserinfo.svg?branch=master)](https://travis-ci.org/ildarusmanov/adfsuserinfo)
[![Maintainability](https://api.codeclimate.com/v1/badges/88f99d4d7a7cb9fd7b23/maintainability)](https://codeclimate.com/github/ildarusmanov/adfsuserinfo/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/88f99d4d7a7cb9fd7b23/test_coverage)](https://codeclimate.com/github/ildarusmanov/adfsuserinfo/test_coverage)


# Setup

## Create config from config.example.com
```
cp config.example.yml config.yml
vim config.yml
```
## Setup dependencies
```
dep ensure
```

## Run tests


```
go test ./...
```

## Try some requests like this

Get Userinfo
```
GET http://0.0.0.0:8001/api/v1/me?access_token=[access_token]
```
```
POST http://0.0.0.0:8001/api/v1/me
Content-Type: application/x-www-form-urlencoded

access_token=[access_token]
```

Status check
```
GET http://0.0.0.0:8002/api/v1/status/check
```

## Run with docker
```
cd [project path]

sudo docker build -t adfsuserinfo .

// prod
sudo docker run --restart=always -d -p 8002:8002 --network host adfsuserinfo

// or dev
sudo docker run -d -p 8002:8002 --network host adfsuserinfo
// list containers
sudo docker ps
```