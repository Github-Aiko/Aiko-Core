package api

import (
	"github.com/Github-Aiko/Aiko-Core/main/commands/base"
)

// CmdAPI calls an API in an Aiko process
var CmdAPI = &base.Command{
	UsageLine: "{{.Exec}} api",
	Short:     "Call an API in an Aiko process",
	Long: `{{.Exec}} {{.LongName}} provides tools to manipulate Aiko via its API.
`,
	Commands: []*base.Command{
		cmdRestartLogger,
		cmdGetStats,
		cmdQueryStats,
		cmdSysStats,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds,
	},
}
