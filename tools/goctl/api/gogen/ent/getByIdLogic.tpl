package {{.packageName}}

import (
	"context"

	"{{.projectPath}}{{.importPrefix}}/internal/svc"
	"{{.projectPath}}{{.importPrefix}}/internal/types"
	"{{.projectPath}}{{.importPrefix}}/internal/utils/dberrorhandler"

{{if .useI18n}}    "github.com/suyuan32/simple-admin-common/i18n"
{{else}}    "github.com/suyuan32/simple-admin-common/msg/errormsg"
{{end}}{{if .useUUID}}    "github.com/suyuan32/simple-admin-common/utils/uuidx"
{{end}}
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type Get{{.modelName}}ByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGet{{.modelName}}ByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get{{.modelName}}ByIdLogic {
	return &Get{{.modelName}}ByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Get{{.modelName}}ByIdLogic) Get{{.modelName}}ById(req *types.{{if .useUUID}}UU{{end}}IDReq) (*types.{{.modelName}}InfoResp, error) {
	data, err := l.svcCtx.DB.{{.modelName}}.Get(l.ctx, {{if .useUUID}}uuidx.ParseUUIDString({{end}}req.Id{{if .useUUID}}){{end}})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.{{.modelName}}InfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  {{if .useI18n}}l.svcCtx.Trans.Trans(l.ctx, i18n.Success){{else}}errormsg.Success{{end}},
        },
        Data: types.{{.modelName}}Info{
            Base{{if .useUUID}}UU{{end}}IDInfo:    types.Base{{if .useUUID}}UU{{end}}IDInfo{
				Id:          {{if .useUUID}}pointy.GetPointer(data.ID.String()){{else}}&data.ID{{end}},
				CreatedAt:    pointy.GetPointer(data.CreatedAt.Unix()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.Unix()),
            },
{{.listData}}
        },
	}, nil
}

