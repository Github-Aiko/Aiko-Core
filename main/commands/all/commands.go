package all

import (
	"github.com/Github-Aiko/Aiko-Core/main/commands/all/api"
	"github.com/Github-Aiko/Aiko-Core/main/commands/all/tls"
	"github.com/Github-Aiko/Aiko-Core/main/commands/base"
)

// go:generate go run github.com/Github-Aiko/Aiko-Core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
