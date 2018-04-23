# HTTP-Facter

http-facter is a simple to use REST API that allows
applications to easily retrieve facts from remote
systems using HTTP GET.

For example, if you want the OS version of a remote system:
```bash
$ curl -XGET http://testserver:4023/os/version
```
The result would be:
```json
{
    "os" : {
        "version": "RHEL7"
    }
}
```

