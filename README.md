# PhantomDNS

---

[中文README](https://github.com/Optimus-Xs/PhantomDNS/blob/main/doc/README_zh.md)

PhantomDNS is a DoH (DNS over HTTPS) server implemented using Go

Also supports request filtering and custom domain name resolution, as well as DDNS over HTTPS implementation (via API interface, non-standard protocol)

## Features
### DoH request
To use PhantomDNS to resolve your domain name, add `app.yml` to the same directory as your application after installation

and add the upstream DoH service provider you want to use in it

A sample configuration is as follows
```yml
publicDNS:
    dns.google
```
This configuration will cause PhantomDNS to use
`https://dns.google/dns-query` (RFC 8484 standard) and `https://dns.google/resolve` (Json API request)
as the upstream request interface for DoH

### Custom Domain Resolution
PhantomDNS also allows you to use this service to resolve a custom domain name address

Note: This feature only allows PhantomDNS to return pre-set IP addresses for resolution requests related to your custom domain name, and does not require you to own and register the domain name with Whois.
It also causes DNS pollution, and it is recommended to use only unregistered domains for your custom domain name resolution rules or to deploy them on a local area network.

Configuration method.
After successfully deploying Phantom, initiate the following API request.

```
https://127.0.0.1:8000/registerHost?ip=192.168.0.1&domain=example.com
```
Please use the HTTP Get request
Parameters.
- ip:the IP address to which the custom domain name points
- domain:custom domain name
- baseAuth: This interface uses BaseAuth as the request authentication scheme, please add BaseAuth information when initiating the request
    - username: `register_name` value in the `devices` table of PhantomDNS.db
    - password: `register_password` value in the `devices` table of PhantomDNS.db

Then, when a request for `example.com` is sent to PhantomDNS, it will return `192.168.0.1`.

You can also set up a script on the client side to update the IP address to PhantomDNS to achieve the DDNS over HTTPS effect

### Resolve request filtering
Since the custom resolution service provided by PhantomDNS uses an unregistered domain name, and most likely the IP it points to is an IP address you don't want to expose

>From my research on the RFC 8484 standard, it seems that the standard DoH protocol does not support a permission validation mechanism similar to DNSCrypt, if that anyone has a solution, please New an Issue

So PhantomDNS provides permission verification based on the IP address of the requesting end

This solution requires the client to send a request to PhantomDNS to register its IP address as the authorized request address

Note: This authentication method reads the x-forwarded-for field of the HTTP request to get the client's IP, so please configure the reverse proxy server or CDN service correctly.
Also, due to the existence of NAT technology, there may be devices that share your IP address, and they can also resolve DNS through PhantomDNS normally.

The request format is as follows.

```
https://127.0.0.1:8000/registerClient?ip=192.168.0.1
```
Please use HTTP Post request
Parameters.
- ip: the IP of the client that can make the request to PhantomDNS
- baseAuth: This interface uses BaseAuth as the request authentication scheme, please add BaseAuth information when initiating the request
    - username: the `register_name` value of the `devices` table in PhantomDNS.db
    - password: `register_password` value in the `devices` table of PhantomDNS.db


## Installation method
### Manual compilation and installation
Prerequisites.
- Golang 1.18 or above installed
- GCC installed

1. pull this repository
2. get into this `/PhantomDNS` directory
3. Run the compile command
   ```shell
    $ go build main.go -o PhantomDNS
   ```
   where PhantomDNS is the name of the output program file and can be customized
4. The newly created ``PhantomDNS`` file is the PhantomDNS executable, and the installation is complete

### Adding configuration files
After the installation is complete, add `app.yml` to the same directory as the PhantomDNS executable as follows
```yaml
publicDNS:
    dns.google # Upstream DoH service providers
SqliteStorage:
    PhantomDNS.db # The path to the sqlite database file storage used by PhantomDNS
DDns:
    ttl:30 # TTL in custom domain name resolution results (unit: seconds)
```

### Adding device information to the Db
After adding `app.yml` you can add the device information via
```shell
$ ./PhantomDNS #The name of the executable file you just entered when compiling
```
and PhantomDNS is running !

After the initial startup, the database file is automatically generated in the path of the `SqliteStorage` attribute in `app.yml`, you need to manually add it to the `devices` table to use Phantom and then restart PhantomDNS

The `register_name` and `register_password` are used as BaseAuth authentication information for updating client addresses and custom domain name resolution rules using the API.

The current version of this data can only be achieved by manually modifying the database file, subsequent iterations will add the use of configuration files and API interfaces to import

## License
This project is licensed under the MIT Open Source License. The full license description is placed in the [LICENSE](https://github.com/Optimus-Xs/PhantomDNS/blob/main/LICENSE) file.