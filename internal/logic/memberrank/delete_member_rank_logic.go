package memberrank

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-member-rpc/mms"

	"github.com/suyuan32/simple-admin-member-api/internal/svc"
	"github.com/suyuan32/simple-admin-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberRankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewDeleteMemberRankLogic(r *http.Request, svcCtx *svc.ServiceContext) *DeleteMemberRankLogic {
	return &DeleteMemberRankLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *DeleteMemberRankLogic) DeleteMemberRank(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.MmsRpc.DeleteMemberRank(l.ctx, &mms.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
