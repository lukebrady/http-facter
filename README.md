# HTTP-Facter

http-facter is a simple to use REST API that allows
applications to easily retrieve facts from remote
systems using HTTP GET.

For example, if you want the OS version of a remote system:
```bash
$ curl -XGET http://testserver:4023/fact/os
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
You can also use http-facter with PowerShell. For example:
```powershell
Invoke-RestMethod -Method Get http://testserver:4023/fact
```
The result would be:
```powershell
aio_agent_version : 5.3.4
augeas            : @{version=1.8.1}
disks             : @{fd0=; sda=; sr0=}
dmi               : @{bios=; board=; chassis=; manufacturer=VMware, Inc.; product=}
facterversion     : 3.9.4
filesystems       : ext2,ext3,ext4,xfs
hypervisors       : @{vmware=}
identity          : @{gid=0; group=root; privileged=True; uid=0; user=root}
is_virtual        : True
kernel            : Linux
kernelmajversion  : 3.10
kernelrelease     : 3.10.0-693.21.1.el7.x86_64
kernelversion     : 3.10.0
load_averages     : @{15m=0.06; 1m=0.22; 5m=0.08}
memory            : @{swap=; system=}
mountpoints       : @{/=; /boot=; /dev/shm=; /home=; /opt=; /run=; /run/user/1000=; /run/user/1002=; /run/user/996=;
                    /sys/fs/cgroup=; /var=}
networking        : @{testserver network information}
os                : @{architecture=x86_64; family=RedHat; hardware=x86_64; name=RedHat; release=; selinux=}
partitions        : @{/dev/sda1=; /dev/sda2=; /dev/sda3=; /dev/sda4=; /dev/sda5=; /dev/sda6=; /dev/sda7=}
path              : /sbin:/bin:/usr/sbin:/usr/bin
pe_cert_change    : testserver
processors        : @{count=4; isa=x86_64; models=System.Object[]; physicalcount=4}
ruby              : @{platform=x86_64-linux; sitedir=/opt/puppetlabs/puppet/lib/ruby/site_ruby/2.4.0; version=2.4.3}
ssh               : @{ecdsa=; ed25519=; rsa=}
system_uptime     : @{days=23; hours=905; seconds=3260802; uptime=23 days}
timezone          : EDT
virtual           : vmware
```

