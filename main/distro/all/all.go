package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/Github-Aiko/Aiko-Core/app/dispatcher"
	_ "github.com/Github-Aiko/Aiko-Core/app/proxyman/inbound"
	_ "github.com/Github-Aiko/Aiko-Core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/Github-Aiko/Aiko-Core/app/commander"
	_ "github.com/Github-Aiko/Aiko-Core/app/log/command"
	_ "github.com/Github-Aiko/Aiko-Core/app/proxyman/command"
	_ "github.com/Github-Aiko/Aiko-Core/app/stats/command"

	// Developer preview services
	_ "github.com/Github-Aiko/Aiko-Core/app/observatory/command"

	// Other optional features.
	_ "github.com/Github-Aiko/Aiko-Core/app/dns"
	_ "github.com/Github-Aiko/Aiko-Core/app/dns/fakedns"
	_ "github.com/Github-Aiko/Aiko-Core/app/log"
	_ "github.com/Github-Aiko/Aiko-Core/app/metrics"
	_ "github.com/Github-Aiko/Aiko-Core/app/policy"
	_ "github.com/Github-Aiko/Aiko-Core/app/reverse"
	_ "github.com/Github-Aiko/Aiko-Core/app/router"
	_ "github.com/Github-Aiko/Aiko-Core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/Github-Aiko/Aiko-Core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/Github-Aiko/Aiko-Core/proxy/blackhole"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/dns"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/dokodemo"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/freedom"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/http"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/loopback"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/mtproto"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/shadowsocks"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/socks"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/trojan"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/vless/inbound"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/vless/outbound"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/vmess/inbound"
	_ "github.com/Github-Aiko/Aiko-Core/proxy/vmess/outbound"

	// Transports
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/domainsocket"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/grpc"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/http"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/kcp"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/quic"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/tcp"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/tls"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/udp"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/websocket"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/xtls"

	// Transport headers
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/http"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/noop"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/srtp"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/tls"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/utp"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/wechat"
	_ "github.com/Github-Aiko/Aiko-Core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/Github-Aiko/Aiko-Core/main/json"
	_ "github.com/Github-Aiko/Aiko-Core/main/toml"
	_ "github.com/Github-Aiko/Aiko-Core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/Github-Aiko/Aiko-Core/main/confloader/external"

	// Commands
	_ "github.com/Github-Aiko/Aiko-Core/main/commands/all"
)
