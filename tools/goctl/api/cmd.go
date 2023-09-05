package api

import (
	"github.com/spf13/cobra"

	"github.com/zeromicro/go-zero/tools/goctl/api/apigen"
	"github.com/zeromicro/go-zero/tools/goctl/api/docgen"
	"github.com/zeromicro/go-zero/tools/goctl/api/format"
	"github.com/zeromicro/go-zero/tools/goctl/api/gogen"
	"github.com/zeromicro/go-zero/tools/goctl/api/new"
	"github.com/zeromicro/go-zero/tools/goctl/api/validate"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/internal/cobrax"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

var (
	// Cmd describes an api command.
	Cmd       = cobrax.NewCommand("api", cobrax.WithRunE(apigen.CreateApiTemplate))
	docCmd    = cobrax.NewCommand("doc", cobrax.WithRunE(docgen.DocCommand))
	formatCmd = cobrax.NewCommand("format", cobrax.WithRunE(format.GoFormatApi))
	goCmd     = cobrax.NewCommand("go", cobrax.WithRunE(gogen.GoCommand))
	newCmd    = cobrax.NewCommand("new", cobrax.WithRunE(new.CreateServiceCommand),
		cobrax.WithArgs(cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs)))
	validateCmd = cobrax.NewCommand("validate", cobrax.WithRunE(validate.GoValidateApi))
	pluginCmd   = cobrax.NewCommand("plugin", cobrax.WithRunE(plugin.PluginCommand))

	protoCmd = cobrax.NewCommand("proto", cobrax.WithRunE(gogen.GenCRUDLogicByProto))

	entCmd = cobrax.NewCommand("ent", cobrax.WithRunE(gogen.GenCRUDLogicByEnt))
)

func init() {
	var (
		apiCmdFlags      = Cmd.Flags()
		docCmdFlags      = docCmd.Flags()
		formatCmdFlags   = formatCmd.Flags()
		goCmdFlags       = goCmd.Flags()
		newCmdFlags      = newCmd.Flags()
		pluginCmdFlags   = pluginCmd.Flags()
		validateCmdFlags = validateCmd.Flags()
		protoCmdFlags    = protoCmd.Flags()
		entCmdFlags      = entCmd.Flags()
	)

	apiCmdFlags.StringVar(&apigen.VarStringOutput, "o")
	apiCmdFlags.StringVar(&apigen.VarStringHome, "home")
	apiCmdFlags.StringVar(&apigen.VarStringRemote, "remote")
	apiCmdFlags.StringVar(&apigen.VarStringBranch, "branch")

	docCmdFlags.StringVar(&docgen.VarStringDir, "dir")
	docCmdFlags.StringVar(&docgen.VarStringOutput, "o")

	formatCmdFlags.StringVar(&format.VarStringDir, "dir")
	formatCmdFlags.BoolVar(&format.VarBoolIgnore, "iu")
	formatCmdFlags.BoolVar(&format.VarBoolUseStdin, "stdin")
	formatCmdFlags.BoolVar(&format.VarBoolSkipCheckDeclare, "declare")

	goCmdFlags.StringVarP(&gogen.VarStringDir, "dir", "d")
	goCmdFlags.StringVarP(&gogen.VarStringAPI, "api", "a")
	goCmdFlags.StringVar(&gogen.VarStringHome, "home")
	goCmdFlags.StringVar(&gogen.VarStringRemote, "remote")
	goCmdFlags.StringVar(&gogen.VarStringBranch, "branch")
	goCmdFlags.StringVarPWithDefaultValue(&gogen.VarStringStyle, "style", "s", config.DefaultFormat)
	goCmdFlags.BoolVarP(&gogen.VarBoolErrorTranslate, "trans_err", "t")
	goCmdFlags.BoolVarP(&gogen.VarBoolUseCasbin, "casbin", "c")
	goCmdFlags.StringVarP(&gogen.VarStringExtraField, "extra_field", "e")
	goCmdFlags.BoolVarP(&gogen.VarBoolUseI18n, "i18n", "i")

	newCmdFlags.StringVar(&new.VarStringHome, "home")
	newCmdFlags.StringVar(&new.VarStringRemote, "remote")
	newCmdFlags.StringVar(&new.VarStringBranch, "branch")
	newCmdFlags.StringVarPWithDefaultValue(&new.VarStringStyle, "style", "s", config.DefaultFormat)
	newCmdFlags.BoolVarP(&new.VarBoolUseCasbin, "casbin", "c")
	newCmdFlags.BoolVarP(&new.VarBoolUseI18n, "i18n", "i")
	newCmdFlags.StringVarPWithDefaultValue(&new.VarStringGoZeroVersion, "go_zero_version", "z", config.DefaultGoZeroVersion)
	newCmdFlags.StringVarPWithDefaultValue(&new.VarStringToolVersion, "tool_version", "t", config.DefaultToolVersion)
	newCmdFlags.StringVarP(&new.VarModuleName, "module_name", "m")
	newCmdFlags.BoolVarP(&new.VarBoolErrorTranslate, "trans_err", "a")
	newCmdFlags.IntVarPWithDefaultValue(&new.VarIntServicePort, "port", "p", 9100)
	newCmdFlags.BoolVarP(&new.VarBoolGitlab, "gitlab", "g")
	newCmdFlags.BoolVarP(&new.VarBoolEnt, "ent", "e")

	pluginCmdFlags.StringVarP(&plugin.VarStringPlugin, "plugin", "p")
	pluginCmdFlags.StringVar(&plugin.VarStringDir, "dir")
	pluginCmdFlags.StringVar(&plugin.VarStringAPI, "api")
	pluginCmdFlags.StringVar(&plugin.VarStringStyle, "style")

	validateCmdFlags.StringVar(&validate.VarStringAPI, "api")

	protoCmdFlags.StringVarP(&gogen.VarStringProto, "proto", "p")
	protoCmdFlags.StringVarP(&gogen.VarStringOutput, "output", "o")
	protoCmdFlags.StringVarP(&gogen.VarStringAPIServiceName, "api_service_name", "a")
	protoCmdFlags.StringVarP(&gogen.VarStringRPCServiceName, "rpc_service_name", "r")
	protoCmdFlags.StringVarPWithDefaultValue(&gogen.VarStringStyle, "style", "s", config.DefaultFormat)
	protoCmdFlags.StringVarP(&gogen.VarStringModelName, "model", "m")
	protoCmdFlags.IntVarPWithDefaultValue(&gogen.VarIntSearchKeyNum, "search_key_num", "k", 3)
	protoCmdFlags.StringVarP(&gogen.VarStringRpcName, "rpc_name", "n")
	protoCmdFlags.StringVarP(&gogen.VarStringGrpcPbPackage, "grpc_package", "g")
	protoCmdFlags.BoolVar(&gogen.VarBoolMultiple, "multiple")
	protoCmdFlags.BoolVarP(&gogen.VarBoolUseI18n, "i18n", "i")
	protoCmdFlags.StringVarP(&gogen.VarStringImportPrefix, "import_prefix", "x")
	protoCmdFlags.StringVarPWithDefaultValue(&gogen.VarStringJSONStyle, "json_style", "j", "goZero")
	protoCmdFlags.BoolVarP(&gogen.VarBoolOverwrite, "overwrite", "w")

	entCmdFlags.StringVarP(&gogen.VarStringSchema, "schema", "c")
	entCmdFlags.StringVarP(&gogen.VarStringOutput, "output", "o")
	entCmdFlags.StringVarP(&gogen.VarStringAPIServiceName, "api_service_name", "a")
	entCmdFlags.StringVarPWithDefaultValue(&gogen.VarStringStyle, "style", "s", config.DefaultFormat)
	entCmdFlags.StringVarP(&gogen.VarStringModelName, "model", "m")
	entCmdFlags.IntVarPWithDefaultValue(&gogen.VarIntSearchKeyNum, "search_key_num", "k", 3)
	entCmdFlags.StringVarP(&gogen.VarStringGroupName, "group", "g")
	entCmdFlags.BoolVarP(&gogen.VarBoolOverwrite, "overwrite", "w")
	entCmdFlags.BoolVarP(&gogen.VarBoolUseI18n, "i18n", "i")
	entCmdFlags.StringVarP(&gogen.VarStringImportPrefix, "import_prefix", "x")
	entCmdFlags.StringVarPWithDefaultValue(&gogen.VarStringJSONStyle, "json_style", "j", "goZero")

	// Add sub-commands
	Cmd.AddCommand(docCmd)
	Cmd.AddCommand(formatCmd)
	Cmd.AddCommand(goCmd)
	Cmd.AddCommand(newCmd)
	Cmd.AddCommand(pluginCmd)
	Cmd.AddCommand(validateCmd)
	Cmd.AddCommand(protoCmd)
	Cmd.AddCommand(entCmd)
}
