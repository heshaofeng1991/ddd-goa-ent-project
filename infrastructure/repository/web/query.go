package web

import (
	"ddd-goa-ent-project/ent"
ddd-goa-ent-projectConf "ddd-goa-ent-project/ent/cbddd-goa-ent-projectboxconfig"
)
type ddd
-goa-ent-projectConfig interface {
Getddd-goa-ent-projectConfigList(ctx context.Context, entClient *ent.Client,
page, pageSize int64) ([]*ent.Cbddd-goa-ent-projectBoxConfig, error)

Getddd-goa-ent-projectConfigInfoByID(ctx context.Context, entClient *ent.Client, id int64) (*ent.Cbddd-goa-ent-projectBoxConfig, error)
}
type ddd
-goa-ent-project struct {
}

func (ddd -goa-ent-project) Getddd-goa-ent-projectConfigList(ctx context.Context, entClient *ent.Client,
page, pageSize int) ([]*ent.Cbddd-goa-ent-projectBoxConfig, error) {
	return entClient.Cbddd - goa - ent - projectBoxConfig.Query().
		Order(ent.Desc(ddd-goa-ent-projectConf.FieldID)).Limit(pageSize).
		Offset((page-1)*pageSize).All(ctx)
}

func (ddd -goa-ent-project) Getddd-goa-ent-projectConfigInfoByID(ctx context.Context,
entClient *ent.Client, id int) (*ent.Cbddd-goa-ent-projectBoxConfig, error) {
	return entClient.Cbddd - goa - ent - projectBoxConfig.Query().Where(ddd-goa-ent-projectConf.IDEQ(int64(id))).First(ctx)
}
