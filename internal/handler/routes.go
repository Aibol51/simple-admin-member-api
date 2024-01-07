// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "github.com/suyuan32/simple-admin-member-api/internal/handler/base"
	captcha "github.com/suyuan32/simple-admin-member-api/internal/handler/captcha"
	member "github.com/suyuan32/simple-admin-member-api/internal/handler/member"
	memberrank "github.com/suyuan32/simple-admin-member-api/internal/handler/memberrank"
	oauthprovider "github.com/suyuan32/simple-admin-member-api/internal/handler/oauthprovider"
	publicmember "github.com/suyuan32/simple-admin-member-api/internal/handler/publicmember"
	publicoauth "github.com/suyuan32/simple-admin-member-api/internal/handler/publicoauth"
	token "github.com/suyuan32/simple-admin-member-api/internal/handler/token"
	vipdemo "github.com/suyuan32/simple-admin-member-api/internal/handler/vipdemo"
	"github.com/suyuan32/simple-admin-member-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/member/create",
					Handler: member.CreateMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/update",
					Handler: member.UpdateMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/delete",
					Handler: member.DeleteMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/list",
					Handler: member.GetMemberListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member",
					Handler: member.GetMemberByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/member/profile",
				Handler: member.ModifyProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/member/logout",
				Handler: member.LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/bind/wechat",
				Handler: member.BindWechatHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/member/login",
				Handler: publicmember.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/login_by_phone",
				Handler: publicmember.LoginByPhoneHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/login_by_email",
				Handler: publicmember.LoginByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/login_by_sms",
				Handler: publicmember.LoginBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/register",
				Handler: publicmember.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/register_by_email",
				Handler: publicmember.RegisterByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/register_by_sms",
				Handler: publicmember.RegisterBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/reset_password_by_email",
				Handler: publicmember.ResetPasswordByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/reset_password_by_sms",
				Handler: publicmember.ResetPasswordBySmsHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Vip},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/member/vip",
					Handler: vipdemo.TestVipHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/create",
					Handler: memberrank.CreateMemberRankHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/update",
					Handler: memberrank.UpdateMemberRankHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/delete",
					Handler: memberrank.DeleteMemberRankHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/list",
					Handler: memberrank.GetMemberRankListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank",
					Handler: memberrank.GetMemberRankByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token/create",
					Handler: token.CreateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/update",
					Handler: token.UpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/delete",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: token.GetTokenByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: publicoauth.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login/callback",
				Handler: publicoauth.OauthCallbackHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login/wechat/mini_program",
				Handler: publicoauth.WechatMiniProgramLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/create",
					Handler: oauthprovider.CreateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/update",
					Handler: oauthprovider.UpdateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/delete",
					Handler: oauthprovider.DeleteOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/list",
					Handler: oauthprovider.GetOauthProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider",
					Handler: oauthprovider.GetOauthProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha/email",
				Handler: captcha.GetEmailCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha/sms",
				Handler: captcha.GetSmsCaptchaHandler(serverCtx),
			},
		},
	)
}
