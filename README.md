# Test using encore local config


```bash
$ encore version
encore version v1.46.20
```

```bash
$ encore build docker tst-config:tstconf
$ docker run -e PORT=8081 -p 8081:8081 tst-config:tstconf
```

Check:

```bash
$ curl 'localhost:8081/hello/user'
{
  "code": "internal",
  "message": "panic handling request: runtime error: invalid memory address or nil pointer dereference",
  "details": null
}
```

