package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member/create",
		Description: "apiDesc.createMember",
		ApiGroup:    "member",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member/update",
		Description: "apiDesc.updateMember",
		ApiGroup:    "member",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member/delete",
		Description: "apiDesc.deleteMember",
		ApiGroup:    "member",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member/list",
		Description: "apiDesc.getMemberList",
		ApiGroup:    "member",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member",
		Description: "apiDesc.getMemberById",
		ApiGroup:    "member",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	// MEMBER RANK

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member_rank/create",
		Description: "apiDesc.createMemberRank",
		ApiGroup:    "member_rank",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member_rank/update",
		Description: "apiDesc.updateMemberRank",
		ApiGroup:    "member_rank",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member_rank/delete",
		Description: "apiDesc.deleteMemberRank",
		ApiGroup:    "member_rank",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member_rank/list",
		Description: "apiDesc.getMemberRankList",
		ApiGroup:    "member_rank",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/member_rank",
		Description: "apiDesc.getMemberRankById",
		ApiGroup:    "member_rank",
		Method:      "Post",
	})

	if err != nil {
		return err
	}

	return nil
}

func (l *InitDatabaseLogic) insertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  common.DefaultParentId,
		Path:      "",
		Name:      "MemberManagementDirectory",
		Redirect:  "",
		Component: "LAYOUT",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.memberManagement",
			Icon:               "ic:round-person-outline",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 0,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/member",
		Name:      "MemberManagement",
		Redirect:  "",
		Component: "/mms/member/index",
		Sort:      1,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.memberManagement",
			Icon:               "ic:round-person-outline",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Id:        0,
		CreatedAt: 0,
		UpdatedAt: 0,
		Level:     2,
		ParentId:  menuData.Id,
		Path:      "/member_rank",
		Name:      "MemberRankManagement",
		Redirect:  "",
		Component: "/mms/memberRank/index",
		Sort:      2,
		Disabled:  false,
		Meta: &core.Meta{
			Title:              "route.memberRankManagement",
			Icon:               "ic:round-person-outline",
			HideMenu:           false,
			HideBreadcrumb:     false,
			IgnoreKeepAlive:    false,
			HideTab:            false,
			FrameSrc:           "",
			CarryParam:         false,
			HideChildrenInMenu: false,
			Affix:              false,
			DynamicLevel:       0,
			RealPath:           "",
		},
		MenuType: 1,
	})

	if err != nil {
		return err
	}

	return err
}
