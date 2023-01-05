package web

type Web interface {
	Getddd
-goa-ent-projectConfigList(ctx context.Context, p *web.Pagination) (res *web.ddd-goa-ent-projectBoxConfigListResp, err error)
Getddd-goa-ent-projectConfigInfo(ctx context.Context, p *web.GetConfigInfoByIDReq) (res *web.ddd-goa-ent-projectBoxConfigInfoResp, err error)
}
