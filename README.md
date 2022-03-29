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
    