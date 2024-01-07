// Code generated by goctl. DO NOT EDIT.
package types

// The basic response with data | 基础带数据信息
// swagger:model BaseDataInfo
type BaseDataInfo struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response with data | 基础带数据信息
// swagger:model BaseListInfo
type BaseListInfo struct {
	// The total number of data | 数据总数
	Total uint64 `json:"total"`
	// Data | 数据
	Data string `json:"data,omitempty"`
}

// The basic response without data | 基础不带数据信息
// swagger:model BaseMsgResp
type BaseMsgResp struct {
	// Error code | 错误代码
	Code int `json:"code"`
	// Message | 提示信息
	Msg string `json:"msg"`
}

// The simplest message | 最简单的信息
// swagger:response SimpleMsg
type SimpleMsg struct {
	// Message | 信息
	Msg string `json:"msg"`
}

// The page request parameters | 列表请求参数
// swagger:model PageInfo
type PageInfo struct {
	// Page number | 第几页
	// Required: true
	Page uint64 `json:"page" validate:"number"`
	// Page size | 单页数据行数
	// Required: true
	// Maximum: 100000
	PageSize uint64 `json:"pageSize" validate:"number,max=100000"`
}

// Basic ID request | 基础ID参数请求
// swagger:model IDReq
type IDReq struct {
	// ID
	// Required: true
	Id uint64 `json:"id" validate:"number"`
}

// Basic IDs request | 基础ID数组参数请求
// swagger:model IDsReq
type IDsReq struct {
	// IDs
	// Required: true
	Ids []uint64 `json:"ids"`
}

// Basic ID request | 基础ID地址参数请求
// swagger:model IDPathReq
type IDPathReq struct {
	// ID
	// Required: true
	Id uint64 `path:"id"`
}

// Basic UUID request | 基础UUID参数请求
// swagger:model UUIDReq
type UUIDReq struct {
	// ID
	// Required: true
	// Max length: 36
	Id string `json:"id" validate:"len=36"`
}

// Basic UUID array request | 基础UUID数组参数请求
// swagger:model UUIDsReq
type UUIDsReq struct {
	// Ids
	// Required: true
	Ids []string `json:"ids"`
}

// The base response data | 基础信息
// swagger:model BaseIDInfo
type BaseIDInfo struct {
	// ID
	Id *uint64 `json:"id"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The base UUID response data | 基础信息
// swagger:model BaseUUIDInfo
type BaseUUIDInfo struct {
	// ID
	Id *string `json:"id"`
	// Create date | 创建日期
	CreatedAt *int64 `json:"createdAt,optional"`
	// Update date | 更新日期
	UpdatedAt *int64 `json:"updatedAt,optional"`
}

// The response data of member information | 会员信息
// swagger:model MemberInfo
type MemberInfo struct {
	BaseUUIDInfo
	// Status | 状态
	Status *uint32 `json:"status,optional"`
	// Username | 用户名
	Username *string `json:"username,optional"`
	// Password | 密码
	Password *string `json:"password,optional"`
	// Nickname | 昵称
	Nickname *string `json:"nickname,optional"`
	// RankId | 等级ID
	RankId *uint64 `json:"rankId,optional"`
	// Mobile | 手机
	Mobile *string `json:"mobile,optional"`
	// Email | 邮箱
	Email *string `json:"email,optional"`
	// Avatar | 头像地址
	Avatar *string `json:"avatar,optional"`
}

// The response data of member list | 会员列表数据
// swagger:model MemberListResp
type MemberListResp struct {
	BaseDataInfo
	// Member list data | 会员列表数据
	Data MemberListInfo `json:"data"`
}

// Member list data | 会员列表数据
// swagger:model MemberListInfo
type MemberListInfo struct {
	BaseListInfo
	// The API list data | 会员列表数据
	Data []MemberInfo `json:"data"`
}

// Get member list request params | 会员列表请求参数
// swagger:model MemberListReq
type MemberListReq struct {
	PageInfo
	// Username | 用户名
	Username *string `json:"username,optional"`
	// Nickname | 昵称
	Nickname *string `json:"nickname,optional"`
	// Mobile | 手机
	Mobile *string `json:"mobile,optional"`
	// Email | 邮箱
	Email *string `json:"email,optional"`
	// Rank ID | 等级ID
	RankId *uint64 `json:"rankId,optional"`
}

// Member information response | 会员信息返回体
// swagger:model MemberInfoResp
type MemberInfoResp struct {
	BaseDataInfo
	// Member information | 会员数据
	Data MemberInfo `json:"data"`
}

// Register request | 注册参数
// swagger:model RegisterReq
type RegisterReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// required : true
	// max length : 20
	// min length : 20
	CaptchaId string `json:"captchaId" validate:"required,len=20"`
	// The Captcha which users input | 用户输入的验证码
	// required : true
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha" validate:"required,len=5"`
	// The user's email address | 用户的邮箱
	// required : true
	// max length : 100
	Email string `json:"email" validate:"required,email,max=100"`
}

// Register by email request | 邮箱注册参数
// swagger:model RegisterByEmailReq
type RegisterByEmailReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// The Captcha which users input | 用户输入的验证码
	// required : true
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha" validate:"required,len=5"`
	// The user's email address | 用户的邮箱
	// required : true
	// max length : 100
	Email string `json:"email" validate:"required,email,max=100"`
}

// Register by SMS request | 短信注册参数
// swagger:model RegisterBySmsReq
type RegisterBySmsReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// The Captcha which users input | 用户输入的验证码
	// required : true
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha" validate:"required,len=5"`
	// The user's mobile phone number | 用户的手机号码
	// required : true
	// max length : 20
	PhoneNumber string `json:"phoneNumber"  validate:"required,numeric,max=20"`
}

// Log in request | 登录参数
// swagger:model LoginReq
type LoginReq struct {
	// User Name | 用户名
	// required : true
	// max length : 20
	Username string `json:"username" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// max length : 20
	// min length : 20
	CaptchaId string `json:"captchaId,optional"  validate:"omitempty,len=20"`
	// The Captcha which users input | 用户输入的验证码
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha,optional" validate:"omitempty,len=5"`
}

// Log in request | 登录参数
// swagger:model LoginByMobileReq
type LoginByMobileReq struct {
	// User PhoneNumber | 手机号
	// required : true
	// max length : 20
	Mobile string `json:"mobile" validate:"required,alphanum,max=20"`
	// Password | 密码
	// required : true
	// max length : 30
	// min length : 6
	Password string `json:"password" validate:"required,max=30,min=6"`
	// Captcha ID which store in redis | 验证码编号, 存在redis中
	// max length : 20
	// min length : 20
	CaptchaId string `json:"captchaId,optional"  validate:"omitempty,len=20"`
	// The Captcha which users input | 用户输入的验证码
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha,optional" validate:"omitempty,len=5"`
}

// Log in by email request | 邮箱登录参数
// swagger:model LoginByEmailReq
type LoginByEmailReq struct {
	// The user's email address | 用户的邮箱
	// required : true
	// max length : 100
	Email string `json:"email" validate:"required,email,max=100"`
	// The Captcha which users input | 用户输入的验证码
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha,optional" validate:"omitempty,len=5"`
}

// Log in by SMS request | 短信登录参数
// swagger:model LoginBySmsReq
type LoginBySmsReq struct {
	// The user's mobile phone number | 用户的手机号码
	// required : true
	// max length : 20
	PhoneNumber string `json:"phoneNumber"  validate:"required,numeric,max=20"`
	// The Captcha which users input | 用户输入的验证码
	// max length : 5
	// min length : 5
	Captcha string `json:"captcha,optional" validate:"omitempty,len=5"`
}

// The log in response data | 登录返回数据
// swagger:model LoginResp
type LoginResp struct {
	BaseDataInfo
	// The log in information | 登陆返回的数据信息
	Data LoginInfo `json:"data"`
}

// The log in information | 登陆返回的数据信息
// swagger:model LoginInfo
type LoginInfo struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// Rank Code | 等级码
	RankId string `json:"rankId"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire uint64 `json:"expire"`
	// Avatar | 用户头像
	Avatar string `json:"avatar"`
	// Nickname | 用户昵称
	Nickname string `json:"nickname"`
	// RankName | 等级名称
	RankName string `json:"rankName"`
	// Mobile | 手机号
	Mobile string `json:"mobile"`
}

// The modify info | 个人信息修改请求
// swagger:model ModifyProfileReq
type ModifyProfileReq struct {
	// Nickname | 昵称
	Nickname *string `json:"nickname,optional"`
	// Mobile | 手机
	Mobile *string `json:"mobile,optional"`
	// Email | 邮箱
	Email *string `json:"email,optional"`
	// Avatar | 头像地址
	Avatar *string `json:"avatar,optional"`
}

// The modify info response | 个人信息修改返回
// swagger:model ModifyProfileResp
type ModifyProfileResp struct {
	BaseDataInfo
	// The modify info | 个人信息修改返回
	Data ModifyProfileInfo `json:"data"`
}

// The modify info | 个人信息修改返回
// swagger:model ModifyProfileInfo
type ModifyProfileInfo struct {
	// Nickname | 昵称
	Nickname *string `json:"nickname,optional"`
	// Mobile | 手机
	Mobile *string `json:"mobile,optional"`
	// Email | 邮箱
	Email *string `json:"email,optional"`
	// Avatar | 头像地址
	Avatar *string `json:"avatar,optional"`
}

// Reset password by email request | 通过邮箱重置密码请求
// swagger:model ResetPasswordByEmailReq
type ResetPasswordByEmailReq struct {
	Email    string `json:"email" validate:"email"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
}

// Reset password by SMS request | 通过短信重置密码请求
// swagger:model ResetPasswordBySmsReq
type ResetPasswordBySmsReq struct {
	// max length : 20
	PhoneNumber string `json:"phoneNumber" validate:"numeric,max=20"`
	Captcha     string `json:"captcha"`
	Password    string `json:"password"`
}

// Bind wechat request | 绑定微信请求
// swagger:model BindWechatReq
type BindWechatReq struct {
	Code string `json:"code"`
}

// The response data of member rank information | 会员等级信息
// swagger:model MemberRankInfo
type MemberRankInfo struct {
	BaseIDInfo
	// Translated Name | 国际化翻译
	Trans string `json:"trans,optional"`
	// Rank name | 等级名称
	Name *string `json:"name,optional"`
	// Description | 等级描述
	Description *string `json:"description,optional"`
	// Remark | 备注
	Remark *string `json:"remark,optional"`
	// Rank Code | 等级代码
	Code *string `json:"code,optional"`
}

// The response data of member rank list | 会员等级列表数据
// swagger:model MemberRankListResp
type MemberRankListResp struct {
	BaseDataInfo
	// MemberRank list data | 会员等级列表数据
	Data MemberRankListInfo `json:"data"`
}

// MemberRank list data | 会员等级列表数据
// swagger:model MemberRankListInfo
type MemberRankListInfo struct {
	BaseListInfo
	// The API list data | 会员等级列表数据
	Data []MemberRankInfo `json:"data"`
}

// Get member rank list request params | 会员等级列表请求参数
// swagger:model MemberRankListReq
type MemberRankListReq struct {
	PageInfo
	// Name | 等级名称
	Name *string `json:"name,optional"`
	// Description | 描述
	Description *string `json:"description,optional"`
	// Remark | 备注
	Remark *string `json:"remark,optional"`
}

// MemberRank information response | 会员等级信息返回体
// swagger:model MemberRankInfoResp
type MemberRankInfoResp struct {
	BaseDataInfo
	// MemberRank information | 会员等级数据
	Data MemberRankInfo `json:"data"`
}

// The response data of token information | Token信息
// swagger:model TokenInfo
type TokenInfo struct {
	BaseUUIDInfo
	// Status | 状态
	Status *uint32 `json:"status,optional"`
	// User's UUID | 用户的UUID
	Uuid *string `json:"uuid,optional"`
	// Token | 用户的Token
	Token *string `json:"token,optional"`
	// Source | Token 来源
	Source *string `json:"source,optional"`
	// Username | 用户名
	Username *string `json:"username,optional"`
	// ExpiredAt | 过期时间
	ExpiredAt *int64 `json:"expiredAt,optional"`
}

// The response data of token list | Token列表数据
// swagger:model TokenListResp
type TokenListResp struct {
	BaseDataInfo
	// Token list data | Token列表数据
	Data TokenListInfo `json:"data"`
}

// Token list data | Token列表数据
// swagger:model TokenListInfo
type TokenListInfo struct {
	BaseListInfo
	// The API list data | Token列表数据
	Data []TokenInfo `json:"data"`
}

// Get token list request params | Token列表请求参数
// swagger:model TokenListReq
type TokenListReq struct {
	PageInfo
	// Username
	Username *string `json:"username,optional"`
	// Nickname
	Nickname *string `json:"nickname,optional"`
	// Email
	Email *string `json:"email,optional"`
	// Uuid
	Uuid *string `json:"uuid,optional"`
}

// Token information response | Token信息返回体
// swagger:model TokenInfoResp
type TokenInfoResp struct {
	BaseDataInfo
	// Token information | Token数据
	Data TokenInfo `json:"data"`
}

// The response data of oauth provider information | 第三方信息
// swagger:model OauthProviderInfo
type OauthProviderInfo struct {
	BaseIDInfo
	// Provider name | 第三方提供商名称
	// max length : 30
	Name *string `json:"name,optional" validate:"omitempty,max=30"`
	// ClientId | 客户端ID
	// max length : 80
	ClientId *string `json:"clientId,optional" validate:"omitempty,max=80"`
	// ClientSecret | 客户端密钥
	// max length : 100
	ClientSecret *string `json:"clientSecret,optional" validate:"omitempty,max=100"`
	// Redirect URL| 跳转地址
	// max length : 300
	RedirectUrl *string `json:"redirectUrl,optional" validate:"omitempty,max=300"`
	// Scopes | 授权范围
	// max length : 50
	Scopes *string `json:"scopes,optional" validate:"omitempty,max=50"`
	// Authority URL | 授权地址
	// max length : 300
	AuthUrl *string `json:"authUrl,optional" validate:"omitempty,max=300"`
	// The URL to get token | 获取Token的地址
	// max length : 300
	TokenUrl *string `json:"tokenUrl,optional" validate:"omitempty,max=300"`
	// The type of auth | 鉴权方式
	// max : 20
	AuthStyle *uint64 `json:"authStyle,optional" validate:"omitempty,lt=20"`
	// The URL to get user information | 获取信息地址
	// max length : 300
	InfoUrl *string `json:"infoUrl,optional" validate:"omitempty,max=300"`
}

// The response data of oauth provider list | 第三方列表数据
// swagger:model OauthProviderListResp
type OauthProviderListResp struct {
	BaseDataInfo
	// OauthProvider list data | 第三方列表数据
	Data OauthProviderListInfo `json:"data"`
}

// OauthProvider list data | 第三方列表数据
// swagger:model OauthProviderListInfo
type OauthProviderListInfo struct {
	BaseListInfo
	// The API list data | 第三方列表数据
	Data []OauthProviderInfo `json:"data"`
}

// Get oauth provider list request params | 第三方列表请求参数
// swagger:model OauthProviderListReq
type OauthProviderListReq struct {
	PageInfo
	// Name | 第三方提供商名称
	// max length : 30
	Name *string `json:"name,optional" validate:"omitempty,max=30"`
}

// Oauth provider information response | 第三方信息返回体
// swagger:model OauthProviderInfoResp
type OauthProviderInfoResp struct {
	BaseDataInfo
	// OauthProvider information | 第三方数据
	Data OauthProviderInfo `json:"data"`
}

// Oauth log in request | Oauth 登录请求
// swagger:model OauthLoginReq
type OauthLoginReq struct {
	// State code to avoid hack | 状态码，请求前后相同避免安全问题
	// required : true
	// max length : 30
	State string `json:"state" validate:"required,max=30"`
	// Provider name | 提供商名字
	// Example: google
	// required : true
	// max length : 40
	Provider string `json:"provider" validate:"required,max=40"`
}

// Redirect response | 跳转网址返回信息
// swagger:model RedirectResp
type RedirectResp struct {
	BaseDataInfo
	// Redirect information | 跳转网址
	Data RedirectInfo `json:"data"`
}

// Redirect information | 跳转网址
// swagger:model RedirectInfo
type RedirectInfo struct {
	// Redirect URL | 跳转网址
	URL string `json:"URL"`
}

// The oauth log in response data | 第三方登录返回数据
// swagger:model CallbackResp
type CallbackResp struct {
	BaseDataInfo
	// The oauth log in callback information | 第三方登陆返回的数据信息
	Data CallbackInfo `json:"data"`
}

// The oauth callback info response data | Oauth回调数据
// swagger:model CallbackInfo
type CallbackInfo struct {
	// User's UUID | 用户的UUID
	UserId string `json:"userId"`
	// Rank ID | 等级 ID
	RankId string `json:"rankId"`
	// Token for authorization | 验证身份的token
	Token string `json:"token"`
	// Expire timestamp | 过期时间戳
	Expire int64 `json:"expire"`
	// Avatar | 用户头像
	Avatar string `json:"avatar"`
	// Nickname | 用户昵称
	Nickname string `json:"nickname"`
	// RankName | 等级名称
	RankName string `json:"rankName"`
}

// Wechat mini program login request | 微信小程序登录请求
// swagger:model WechatMiniProgramLoginReq
type WechatMiniProgramLoginReq struct {
	// Code
	Code string `json:"code"`
}

// The information of captcha | 验证码数据
// swagger:model CaptchaInfo
type CaptchaInfo struct {
	CaptchaId string `json:"captchaId"`
	ImgPath   string `json:"imgPath"`
}

// The response data of captcha | 验证码返回数据
// swagger:model CaptchaResp
type CaptchaResp struct {
	BaseDataInfo
	// The menu authorization data | 菜单授权信息数据
	Data CaptchaInfo `json:"data"`
}

// The email captcha request | 邮箱验证码请求参数
// swagger:model EmailCaptchaReq
type EmailCaptchaReq struct {
	// The email address | 邮箱地址
	Email string `json:"email"`
}

// The sms captcha request | 短信验证码请求参数
// swagger:model SmsCaptchaReq
type SmsCaptchaReq struct {
	// The phone number | 电话号码
	PhoneNumber string `json:"phoneNumber"`
}
