package web

import (
	"ddd-goa-ent-project/gen/web"
	"ddd-goa-ent-project/infrastructure/repository"

	"errors"
)

var (
	ddd-goa-ent-projectBoxConf ddd-goa-ent-project
)

func (r repo) Getddd-goa-ent-projectConfigList(ctx context.Context, p *web.Pagination) (res *web.ddd-goa-ent-projectBoxConfigListResp, err error) {
	res = &web.ddd - goa - ent - projectBoxConfigListResp{}

	pageSize, page := repository.GetPageSizeAndCurrent(&p.Page, &p.PageSize)

	result, err := ddd - goa - ent - projectBoxConf.Getddd - goa - ent - projectConfigList(ctx, r.entClient, page, pageSize)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	for _, val := range result {
		res.List = append(res.List, &web.ddd-goa-ent-projectBoxConfigList{
			AdminUserID: &val.AdminUserID,
			// ObjectID:    val.ObjectID,
			Name:  val.Name,
			Desc:  val.Desc,
			Img:   val.Img,
			ImgBg: val.ImgBg,
			// Index:       val.Index,
			SellPrice: &val.SellPrice,
			// Status:      val.Status,
			BoxType:    nil,
			Revision:   nil,
			ID:         nil,
			CreateTime: nil,
			UpdateTime: nil,
			CreateBy:   nil,
			UpdateBy:   nil,
		})
	}

	total, err := r.entClient.Cbddd - goa - ent - projectBoxConfig.Query().Count(ctx)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	newTotal := int64(total)

	res.Total = &newTotal

	return
}

func (r repo) Getddd-goa-ent-projectConfigInfo(ctx context.Context, p *web.GetConfigInfoByIDReq) (res *web.ddd-goa-ent-projectBoxConfigInfoResp, err error) {
	res = &web.ddd - goa - ent - projectBoxConfigInfoResp{}

	result, err := ddd - goa - ent - projectBoxConf.Getddd - goa - ent - projectConfigInfoByID(ctx, r.entClient, p.ID)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	res.Name = result.Name

	return
}
