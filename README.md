# lpmng-fw-agent

Web service used to add user ips to the `authorized_table` of the firewall, only one url is currently exposed : `/event/session`.

It accept only `POST` method and the data must be of the form an event, for instance :

```json
{
  "action": "created",
  "name": "session",
  "param": {
    "mac": "fd:df:ad:ad:2a:ff",
    "ip4": "212.126.212.141",
    "user": 14,
    "internet": true
  }
}
```

To install the service on the firewall, clone the repo in `/opt/` build the executable with `go build fw.go`.
Then link the `lpmng-fw-agent` service to `/usr/local/etc/rc.d` and enable it by setting `lpmng_fw_agent_enable="YES"` in the `/etc/rc.conf`.
