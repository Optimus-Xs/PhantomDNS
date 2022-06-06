# PhantomDNS

---

[ENGLISH README](https://github.com/Optimus-Xs/PhantomDNS/blob/main/README.md)

PhantomDNS 是一个使用Go实现的DoH（DNS over HTTPS）服务器

同时支持一定程度的请求过滤和自定义域名解析，以及DDNS over HTTPS实现（通过 API 接口，非标准协议）

## 功能
### DoH请求
若要使用 PhantomDNS 解析域名，在安装完成后请在应用程序同同目录添加 `app.yml` 

同时在其中添加你想使用的上游 DoH 服务提供商

配置示例如下
```yml
publicDNS:
    dns.google
```
此配置会让 PhantomDNS 使用
`https://dns.google/dns-query` (RFC 8484 标准)  和 `https://dns.google/resolve` (Json API请求)
作为 DoH 的上游请求接口

### 自定义域名解析
PhantomDNS 同时允许您使用此服务解析一个自定义的域名地址

注意：此功能只是让 PhantomDNS 对您自定义的域名的相关解析请求返回预先设置的IP地址，不需要您拥有并在 Whois 注册域名，
同时这也会造成 DNS 污染，建议只使用未被注册的域名作为您的自定义域名解析规则，或在局域网内部署

配置方法：
在成功部署 Phantom 后发起一下API请求：
```
https://127.0.0.1:8000/registerHost?ip=192.168.0.1&domain=example.com
```
请使用HTTP Get 请求
参数：
- ip:自定义域名指向的 IP 地址
- domain:自定义的域名
- baseAuth：此接口使用 BaseAuth 作为请求验证方案，请在发起请求时添加 BaseAuth 信息
    - username：PhantomDNS.db 中 devices 表 register_name 值
    - password：PhantomDNS.db 中 devices 表 register_password 值

然后，当使用向 PhantomDNS 发起对 `example.com` 的解析请求时会返回 `192.168.0.1`

同时您也可以在客户端设置定时脚本向 PhantomDNS 更新 IP 地址实现 DDNS over HTTPS 效果

### 解析请求过滤
由于 PhantomDNS 提供的自定义解析服务 是使用的未注册域名，同时大概率其指向 IP 为您不想暴露的 IP 地址

>以我目前对 RFC 8484 标准的研究，在标准 DoH 协议下好像不支持类似于 DNSCrypt 的权限验证机制，如果那个老哥有能解决的思路请发 Issue

所以 PhantomDNS 提供了基于请求端 IP 地址的权限验证

此方案需要客户端向 PhantomDNS 发起请求以注册自己的 IP 地址为授权的请求地址

注意：此验证方法通过读取HTTP请求的 x-forwarded-for 字段获取客户端IP，请正确配置反向代理服务器或 CDN 服务，
同时由于 NAT 技术的存在，可能会有和您共享 IP 地址的设备，此时它们也可以正常通过 PhantomDNS 解析 DNS

请求格式如下：

```
https://127.0.0.1:8000/registerClient?ip=192.168.0.1
```
请使用HTTP Post 请求
参数：
- ip:可以向 PhantomDNS 发起请求的客户端IP
- baseAuth：此接口使用 BaseAuth 作为请求验证方案，请在发起请求时添加 BaseAuth 信息
  - username：PhantomDNS.db 中 devices 表 register_name 值
  - password：PhantomDNS.db 中 devices 表 register_password 值


## 安装方法
### 手动编译安装
前置条件：
- 安装 Golang 1.18 或以上
- 安装 GCC

1. Pull 此仓库
2. 让后进入此 `/PhantomDNS` 目录
3. 运行编译命令
   ```shell
    $ go build main.go -o PhantomDNS
   ```
   其中 PhantomDNS 为输出程序文件名，可以自定义
4. 新出现的 `PhantomDNS` 文件为Phantom的可执行文件，安装完成

### 添加配置文件
在安装完成后，在 PhantomDNS 的可执行文件同目录下添加 `app.yml` 示例如下
```yaml
publicDNS:
    dns.google # 上游DoH服务提供商
SqliteStorage:
    PhantomDNS.db # PhantomDNS使用的sqlite数据库文件储存路径
DDns:
    ttl:30 # 自定义域名解析结果中的TTL（单位：秒）
```

### 在Db中添加设备信息
在添加完 `app.yml` 就可以通过
```shell
$ ./PhantomDNS #您刚刚编译时输入的可执行文件名
```
启动 PhantomDNS 了

在初次启动后会自动在 `app.yml` 中 `SqliteStorage` 属性的路径生成数据库文件，您需要手动在 devices 表中添加需要使用 PhantomDNS 然后重启 Phantom

其中 register_name 和  register_password 作为后续使用 API 更新客户端地址和自定义域名解析规则的 BaseAuth 认证信息，

目前版本此数据只能通过，手动修改数据库文件实现，后续迭代会添加使用配置文件和API接口导入

## 授权许可
本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 [LICENSE](https://github.com/Optimus-Xs/PhantomDNS/blob/main/LICENSE) 文件中。
