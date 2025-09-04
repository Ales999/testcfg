# Test using encore local config

```bash
$ encore version
encore version v1.49.3
```

Build and run with docker

```bash
$ encore build docker tst-config:tstconf
$ docker run -it --rm -e PORT=8081 -p 8081:8081 tst-config:tstconf
```

Check:

```bash
$ curl 'localhost:8081/hello/tstuser'
{
  "code": "internal",
  "message": "configuration not initialized",
  "details": null
}
```

If running local as ```encore run``` and using the default config, it should return:

```bash
$ curl 'localhost:4000/hello/tstuser'
{
  "Message": "Config work FALSE, tstuser!"
}
```
