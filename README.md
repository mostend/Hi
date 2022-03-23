# Hi
A base gui tool for xray/v2ray/hysteria/trojan-go/brook without system proxy 

Support windows and macOS

## build
go env
  - MacOS 
    - chmod +x build.sh
    - ./build.sh
  - Windows 
    - go build  -ldflags="-s -w -H windowsgui"


about brook
```json
{
    "server": "client",
    "serverType": "--server",
    "addr": "0.0.0.0:123456",
    "password": "xxx",
    "http": "0.0.0.0:8010",
    "socks5": "0.0.0.0:1087",
    "noBrook": ""
}
```
```json
{
    "server": "wssclient",
    "serverType": "--wssserver",
    "addr": "wss://0.0.0.0:10002",
    "password": "xxx",
    "http": "0.0.0.0:8010",
    "socks5": "0.0.0.0:1087",
    "noBrook": "--withoutBrookProtocol"
}
```