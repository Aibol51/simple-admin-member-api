package memberrank

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetMemberRankByIdLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetMemberRankByIdLogic {
	return &GetMemberRankByIdLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetMemberRankByIdLogic) GetMemberRankById(req *types.IDReq) (resp *types.MemberRankInfoResp, err error) {
	data, err := l.svcCtx.MmsRpc.GetMemberRankById(l.ctx, &mms.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MemberRankInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
		},
		Data: types.MemberRankInfo{
			BaseInfo: types.BaseInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:        data.Name,
			Description: data.Description,
			Remark:      data.Remark,
			Code:        data.Code,
		},
	}, nil
}
