# HTTP-Facter

http-facter is a simple to use REST API that allows
applications to easily retrieve facts from remote
systems using HTTP GET.

[![BCH compliance](https://bettercodehub.com/edge/badge/lukebrains/http-facter?branch=master)](https://bettercodehub.com/)

For example, if you want the OS version of a remote system:
```bash
$ curl -XGET http://testserver:4023/os
```
The result would be:
```json
{
  "architecture" : "x86_64",
  "family" : "RedHat",
  "hardware" : "x86_64",
  "name" : "RedHat",
  "release" : {
    "full" : "7.4",
    "major" : "7",
    "minor" : "4"
  },
  "selinux" : {
    "config_mode" : "enforcing",
    "current_mode" : "enforcing",
    "enabled" : true,
    "enforced" : true,
    "policy_version" : "28"
  }
}
```

