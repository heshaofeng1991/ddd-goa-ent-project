package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"ddd-goa-ent-project/ent"
	"ddd-goa-ent-project/gen/web"
	repository "ddd-goa-ent-project/infrastructure/repository/web"

	"goa.design/goa/v3/security"
)

// NewWeb returns the web service implementation.
func NewWeb(entClient *ent.Client, logger *log.Logger) web.Service {
	return &basicWebService{
		entClient: entClient,
		logger:    logger,
		web:       repository.New(entClient),
	}
}

// basicWebService implements Service interface.
type basicWebService struct {
	entClient *ent.Client
	web       repository.Web
	logger    *log.Logger
}

// JWTAuth implements the authorization logic for service "web" for the "jwt"
// security scheme.
func (b basicWebService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

func (b basicWebService) Getddd-goa-ent-projectConfigList(ctx context.Context, pagination *web.Pagination) (res *web.ddd-goa-ent-projectBoxConfigListResp, err error) {
	res, err = b.web.Getddd - goa - ent - projectConfigList(ctx, pagination)

	return res, errors.Unwrap(err)
}

func (b basicWebService) Getddd-goa-ent-projectConfigInfo(ctx context.Context, req *web.GetConfigInfoByIDReq) (res *web.ddd-goa-ent-projectBoxConfigInfoResp, err error) {
	res, err = b.web.Getddd - goa - ent - projectConfigInfo(ctx, req)

	return res, errors.Unwrap(err)
}
