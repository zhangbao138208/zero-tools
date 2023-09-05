package makefile

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"strings"

	"github.com/pkg/errors"
)

// extractInfo extracts the information to context
func extractInfo(g *GenContext) error {
	makefileData, err := fileutil.ReadFileToString(g.TargetPath)
	if err != nil {
		return err
	}

	if strings.Contains(makefileData, "gen-api") {
		if strings.Contains(makefileData, "gen-ent") {
			g.UseEnt = true
			g.IsSingle = true
		} else {
			g.IsApi = true
		}
	}

	if strings.Contains(makefileData, "gen-rpc") {
		g.IsRpc = true
		if strings.Contains(makefileData, "gen-ent") {
			g.UseEnt = true
		}
	}

	dataSplit := strings.Split(makefileData, "\n")

	if g.Style == "" && !strings.Contains(makefileData, "PROJECT_STYLE=") {
		return errors.New("style not set, use -s to set style")
	} else if g.Style == "" && strings.Contains(makefileData, "PROJECT_STYLE=") {
		style := findDefined("PROJECT_STYLE=", dataSplit)
		if style == "" {
			return errors.New("failed to find style definition, please set it manually by -s")
		}
		g.Style = style
	}

	if val := findDefined("PROJECT_I18N", dataSplit); val != "" {
		if val == "true" {
			g.UseI18n = true
		}
	}

	if g.ServiceName == "" {
		g.ServiceName = findDefined("SERVICE", dataSplit)
	}

	g.EntFeature = findDefined("ENT_FEATURE", dataSplit)

	if g.EntFeature == "" {
		g.EntFeature = "sql/execquery"
	}

	return err
}

func findDefined(target string, data []string) string {
	for _, v := range data {
		if strings.Contains(v, target) {
			dataSplit := strings.Split(v, "=")
			if len(dataSplit) == 2 {
				return strings.TrimSpace(dataSplit[1])
			} else {
				return ""
			}
		}
	}

	return ""
}
