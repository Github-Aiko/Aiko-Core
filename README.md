# Project X

[Project X](https://github.com/XTLS) originates from XTLS protocol, provides a set of network tools such as [Aiko-core](https://github.com/Github-Aiko/Aiko-Core).

## License

[Mozilla Public License Version 2.0](https://github.com/Github-Aiko/Aiko-Core/blob/main/LICENSE)

## Installation

- Linux Script
  - [Aiko-install](https://github.com/XTLS/Aiko-install)
  - [Aiko-script](https://github.com/kirin10000/Aiko-script)
- Docker
  - [teddysun/Aiko](https://hub.docker.com/r/teddysun/Aiko)
- One Click
  - [ProxySU](https://github.com/proxysu/ProxySU)
  - [v2ray-agent](https://github.com/mack-a/v2ray-agent)
  - [Aiko-yes](https://github.com/jiuqi9997/Aiko-yes)
  - [Aiko_onekey](https://github.com/wulabing/Aiko_onekey)
- Magisk
  - [Aiko4Magisk](https://github.com/CerteKim/Aiko4Magisk)
  - [Aiko_For_Magisk](https://github.com/E7KMbb/Aiko_For_Magisk)
- Homebrew
  - `brew install Aiko`
  - [(Tap) Repository 0](https://github.com/N4FA/homebrew-Aiko)
  - [(Tap) Repository 1](https://github.com/xiruizhao/homebrew-Aiko)

## Contributing
[Code Of Conduct](https://github.com/Github-Aiko/Aiko-Core/blob/main/CODE_OF_CONDUCT.md)

## Usage

[Aiko-examples](https://github.com/XTLS/Aiko-examples) / [VLESS-TCP-XTLS-WHATEVER](https://github.com/XTLS/Aiko-examples/tree/main/VLESS-TCP-XTLS-WHATEVER)

## GUI Clients

- OpenWrt
  - [PassWall](https://github.com/xiaorouji/openwrt-passwall)
  - [Hello World](https://github.com/jerrykuku/luci-app-vssr)
  - [ShadowSocksR Plus+](https://github.com/fw876/helloworld)
  - [luci-app-Aiko](https://github.com/yichya/luci-app-Aiko) ([openwrt-Aiko](https://github.com/yichya/openwrt-Aiko))
- Windows
  - [v2rayN](https://github.com/2dust/v2rayN)
  - [Qv2ray](https://github.com/Qv2ray/Qv2ray) (This project had been archived and currently inactive)
  - [Netch (NetFilter & TUN/TAP)](https://github.com/NetchX/Netch) (This project had been archived and currently inactive)
- Android
  - [v2rayNG](https://github.com/2dust/v2rayNG)
  - [Kitsunebi](https://github.com/rurirei/Kitsunebi/tree/release_xtls)
- iOS & macOS (with M1 chip)
  - [Shadowrocket](https://apps.apple.com/app/shadowrocket/id932747118)
  - [Stash](https://apps.apple.com/app/stash/id1596063349)
- macOS (Intel chip & M1 chip)
  - [Qv2ray](https://github.com/Qv2ray/Qv2ray) (This project had been archived and currently inactive)
  - [V2RayXS](https://github.com/tzmax/V2RayXS)

## Credits

This repo relies on the following third-party projects:

- Special thanks:
  - [v2fly/v2ray-core](https://github.com/v2fly/v2ray-core)
- In production:
  - [ghodss/yaml](https://github.com/ghodss/yaml)
  - [gorilla/websocket](https://github.com/gorilla/websocket)
  - [lucas-clemente/quic-go](https://github.com/lucas-clemente/quic-go)
  - [pelletier/go-toml](https://github.com/pelletier/go-toml)
  - [pires/go-proxyproto](https://github.com/pires/go-proxyproto)
  - [refraction-networking/utls](https://github.com/refraction-networking/utls)
  - [seiflotfy/cuckoofilter](https://github.com/seiflotfy/cuckoofilter)
  - [google/starlark-go](https://github.com/google/starlark-go)
- For testing only:
  - [miekg/dns](https://github.com/miekg/dns)
  - [stretchr/testify](https://github.com/stretchr/testify)
  - [h12w/socks](https://github.com/h12w/socks)

## Compilation

### Windows

```bash
go build -o Aiko.exe -trimpath -ldflags "-s -w -buildid=" ./main
```

### Linux / macOS

```bash
go build -o Aiko -trimpath -ldflags "-s -w -buildid=" ./main
```

## Telegram

[Project X](https://t.me/projectAiko)

[Project X Channel](https://t.me/projectXtls)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/Github-Aiko/Aiko-Core.svg)](https://starchart.cc/Github-Aiko/Aiko-Core)
