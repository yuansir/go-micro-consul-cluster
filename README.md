# GO Micro 搭建 Consul服务发现集群实例

博客文章 => [Go实践微服务 -- 服务发现](https://yuanxuxu.com/2018/07/10/go-microservice-service-discovery/)

### 依赖

* Docker
* Protobuf v3
    * $ brew install protobuf
* protoc-gen libraries and other dependencies
    * $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    * $ go get -u github.com/micro/protoc-gen-micro
    * $ go get -u github.com/micro/go-micro
    * $ go get -u github.com/hailocab/go-geoindex

### 克隆项目

git clone 该项目到你的GOPATH/src 下

### 运行

```
$ ./build.sh

$ docker-compose build

$ docker-compose up

```

### 请求
Curl 请求未认证token API:

```
$ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer INVALID_TOKEN" \
           -d '{"inDate": "2015-04-09"}' \
            http://localhost:8080/hotel/rates

    {"id":"api.hotel.rates","code":401,"detail":"Unauthorized","status":"Unauthorized"}
```
    
Curl 请求认证token API:

```
 $ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer VALID_TOKEN" \
           -d '{"inDate": "2015-04-09"}' \
            http://localhost:8080/hotel/rates

    {"id":"api.hotel.rates","code":400,"detail":"Please specify inDate/outDate params","status":"Bad Request"}

```
   
Curl 请求认证token API:

```$xslt
 $ curl -H 'Content-Type: application/json' \
           -H "Authorization: Bearer VALID_TOKEN" \
           -d '{"inDate": "2015-04-09", "outDate": "2015-04-10"}' \
            http://localhost:8080/hotel/rates
```

The JSON response:

```json
{
    "hotels": [
        {
            "id": 1,
            "name": "Clift Hotel",
            "phoneNumber": "(415) 775-4700",
            "description": "A 6-minute walk from Union Square and 4 minutes from a Muni Metro station, this luxury hotel designed by Philippe Starck features an artsy furniture collection in the lobby, including work by Salvador Dali.",
            "address": {
                "streetNumber": "495",
                "streetName": "Geary St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94102"
            }
        }
    ],
    "ratePlans": [
        {
            "hotelId": 1,
            "code": "RACK",
            "inDate": "2015-04-09",
            "outDate": "2015-04-10",
            "roomType": {
                "bookableRate": 109,
                "totalRate": 109,
                "totalRateInclusive": 123.17,
                "code": "KNG"
            }
        }
    ]
}
```