package wework

import (
	"encoding/json"
	"net/url"
	"os"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/go-resty/resty/v2"
	"github.com/haiyiyun/weworksdk/internal"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IWeWork interface {
	getProviderToken() string

	GetCorpId() string
	GetSuiteId() string
	GetSuiteToken() string
	GetSuiteEncodingAesKey() string
	Logger() *zap.Logger
	SetAppSecretFunc(f func(corpId string) (corpid string, secret string, customizedApp bool))
	SetAgentIdFunc(f func(corpId string) (agentId int))

	GetLoginInfo(authCode string) (resp GetLoginInfoResponse)
	GetUserInfo3rd(code string) (resp GetUserInfo3rdResponse)
	GetUserInfoDetail3rd(userTicket string) (resp GetUserInfoDetail3rdResponse)
	GetUserInfo(corpId string, code string) (resp GetUserInfoResponse)
	GetUserDetail(corpId string, userTicket string) (resp GetUserDetailResponse)

	AgentGet(corpId string, agentId int) (resp AgentGetResponse)
	AgentList(corpId string) (resp AgentListResponse)

	UpdateSuiteTicket(ticket string)
	getSuiteAccessToken() string
	GetPreAuthCode() (resp GetPreAuthCodeResponse)
	GetPermanentCode(authCode string) (resp GetPermanentCodeResponse)
	GetAuthInfo(authCorpId, permanentCode string) (resp GetAuthInfoResponse)
	GetAppQrCode(request GetAppQrCodeRequest) (resp GetAppQrCodeResponse)

	UserCreate(corpId string, user User) (resp internal.BizResponse)
	UserUpdate(corpId string, user User) (resp internal.BizResponse)
	UserDelete(corpId string, userId string) (resp internal.BizResponse)
	UserGet(corpId string, userId string) (resp UserGetResponse)
	UserSimpleList(corpId string, departId int32, fetchChild int) (resp UserSimpleListResponse)
	UserList(corpId string, departId int32, fetchChild int) (resp UserListResponse)
	UserId2OpenId(corpId string, userId string) (resp UserId2OpenIdResponse)
	OpenId2UserId(corpId string, openId string) (resp OpenId2UserIdResponse)
	ListMemberAuth(corpId string, cursor string, limit int) (resp ListMemberAuthResponse)
	CheckMemberAuth(corpId string, openUserId string) (resp CheckMemberAuthResponse)
	GetUserId(corpId string, mobile string) (resp GetUserIdResponse)
	ListSelectedTicketUser(corpId string, ticket string) (resp ListSelectedTicketUserResponse)
	UserListId(corpId string, cursor string, limit int) (resp UserListIdResponse)

	CorpTagList(corpId string, tagIds, groupIds []string) (resp CorpTagListResponse)
	CorpTagAdd(corpId string, tagGroup CorpTagGroup) (resp CorpTagAddResponse)
	CorpTagUpdate(corpId string, tag CorpTag) (resp internal.BizResponse)
	CorpTagDelete(corpId string, tagIds, groupIds []string) (resp internal.BizResponse)
	MarkTag(corpId string, userId string, externalUserId string, addTag []int, removeTag []int) (resp internal.BizResponse)

	DepartmentCreate(corpId string, department Department) (resp DepartmentCreateResponse)
	DepartmentUpdate(corpId string, department Department) (resp internal.BizResponse)
	DepartmentDelete(corpId string, id int32) (resp internal.BizResponse)
	DepartmentList(corpId string, id uint) (resp DepartmentListResponse)
	DepartmentSimpleList(corpId string, id int32) (resp DepartmentSimpleListResponse)
	DepartmentGet(corpId string, id int32) (resp DepartmentGetResponse)

	ExternalContactGetFollowUserList(corpId string) (resp ExternalContactGetFollowUserListResponse)
	ExternalContactList(corpId string, userId string) (resp ExternalContactListResponse)
	ExternalContactGet(corpId string, externalUserId, cursor string) (resp ExternalContactGetResponse)
	ExternalContactBatchGetByUser(corpId string, userIds []string, cursor string, limit int) (resp ExternalContactBatchGetByUserResponse)
	ExternalContactRemark(corpId string, remark ExternalContactRemarkRequest) (resp internal.BizResponse)
	UnionId2ExternalUserId(corpId string, unionid, openid string) (resp UnionId2ExternalUserIdResponse)
	ToServiceExternalUserid(corpId string, externalUserId string) (resp ToServiceExternalUseridResponse)

	ExternalAddContactWay(corpId string, me ContactMe) (resp ContactMeAddResponse)
	ExternalUpdateContactWay(corpId string, me ContactMe) (resp internal.BizResponse)
	ExternalGetContactWay(corpId string, configId string) (resp ContactMeGetResponse)
	ExternalListContactWay(corpId string, startTime, endTime int64, cursor string, limit int) (resp ContactMeListResponse)
	ExternalDeleteContactWay(corpId string, configId string) (resp internal.BizResponse)
	ExternalCloseTempChat(corpId string, userId, externalUserId string) (resp internal.BizResponse)

	AddMsgTemplate(corpId string, msg ExternalMsg) (resp AddMsgTemplateResponse)
	GetGroupMsgListV2(corpId string, filter GroupMsgListFilter) (resp GetGroupMsgListV2Response)
	GetGroupMsgTask(corpId string, filter GroupMsgTaskFilter) (resp GetGroupMsgTaskResponse)
	GetGroupMsgSendResult(corpId string, filter GroupMsgSendResultFilter) (resp GetGroupMsgSendResultResponse)
	SendWelcomeMsg(corpId string, msg ExternalMsg) (resp internal.BizResponse)

	GetUserBehaviorData(corpId string, filter GetUserBehaviorFilter) (resp GetUserBehaviorDataResponse)
	GroupChatStatistic(corpId string, filter GroupChatStatisticFilter) (resp GroupChatStatisticResponse)
	GroupChatStatisticGroupByDay(corpId string, filter GroupChatStatisticGroupByDayFilter) (resp GroupChatStatisticResponse)

	AddProductAlbum(corpId string, product Product) (resp AddProductAlbumResponse)
	GetProductAlbum(corpId string, productId string) (resp GetProductAlbumResponse)
	GetProductAlbumList(corpId string, limit int, cursor string) (resp GetProductAlbumListResponse)
	UpdateProductAlbum(corpId string, request ProductUpdateRequest) (resp internal.BizResponse)
	DeleteProductAlbum(corpId string, productId string) (resp internal.BizResponse)

	AddInterceptRule(corpId string, interceptRule InterceptRule) (resp AddInterceptRuleResponse)
	GetInterceptRuleList(corpId string) (resp GetInterceptRuleListResponse)
	GetInterceptRule(corpId string, ruleId string) (resp GetInterceptRuleResponse)
	UpdateInterceptRule(corpId string, request UpdateInterceptRuleRequest) (resp internal.BizResponse)
	DeleteInterceptRule(corpId string, ruleId string) (resp internal.BizResponse)

	GroupChatList(corpId string, filter GroupChatListFilter) (resp GroupChatListResponse)
	GroupChat(corpId string, request GroupChatRequest) (resp GroupChatResponse)
	GroupOpengId2ChatId(corpId string, opengid string) (resp GroupOpengId2ChatIdResponse)

	MediaUploadAttachment(corpId string, attrs Media) (resp MediaUploadResponse)
	MediaUpload(corpId string, fileType MediaType, filePath string) (resp MediaUploadResponse)
	MediaUploadImg(corpId string, filePath string) (resp MediaUploadImgResponse)
	MediaGet(corpId string, mediaId string) (resp MediaGetResponse)

	GetBillList(corpId string, req GetBillListRequest) (resp GetBillListResponse)

	MessageSend(corpId string, msg interface{}) (resp MessageSendResponse)
	MessageReCall(corpId string, msgId string) (resp internal.BizResponse)

	MessageUpdateTemplateCard(corpId string, msg TemplateCardUpdateMessage) (resp MessageUpdateTemplateCardResponse)

	AddMomentTask(corpId string, task MomentTask) (resp AddMomentTaskResponse)
	GetMomentTaskResult(corpId string, jobId string) (resp GetMomentTaskResultResponse)
	GetMomentList(corpId string, filter MomentListFilter) (resp GetMomentListResponse)
	GetMomentTask(corpId string, filter MomentTaskFilter) (resp GetMomentTaskResponse)
	GetMomentCustomerList(corpId string, filter MomentCustomerFilter) (resp GetMomentCustomerListResponse)
	GetMomentSendResult(corpId string, filter MomentCustomerFilter) (resp GetMomentSendResultResponse)
	GetMomentComments(corpId string, momentId string, userId string) (resp GetMomentCommentsResponse)

	TagCreate(corpId string, tag Tag) (resp TagCreateResponse)
	TagUpdate(corpId string, tag Tag) (resp internal.BizResponse)
	TagDelete(corpId string, id int) (resp internal.BizResponse)
	TagList(corpId string) (resp TagListResponse)
	TagUserList(corpId string, id int) (resp TagUserListResponse)
	TagAddUsers(corpId string, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse)
	TagDelUsers(corpId string, tagId int, userIds []string, partyIds []int32) (resp TagAddOrDelUsersResponse)

	TransferCustomer(corpId string, request TransferCustomerRequest) (resp TransferCustomerResponse)
	TransferResult(corpId string, request TransferResultRequest) (resp TransferResultResponse)
	GetUnassignedList(corpId string, request UnAssignedRequest) (resp UnAssignedResponse)
	TransferCustomerResigned(corpId string, request TransferCustomerRequest) (resp TransferCustomerResponse)
	TransferResultResigned(corpId string, request TransferResultRequest) (resp TransferResultResponse)
	TransferGroupChat(corpId string, request GroupChatTransferRequest) (resp GroupChatTransferResponse)

	GetInvoiceInfo(corpId string, query InvoiceInfoQuery) (resp GetInvoiceInfoResponse)
	GetInvoiceInfoBatch(corpId string, query InvoiceInfoQueryBatch) (resp GetInvoiceInfoBatchResponse)
	UpdateInvoiceStatus(corpId string, request UpdateInvoiceStatusRequest) (resp internal.BizResponse)
	UpdateInvoiceStatusBatch(corpId string, request UpdateInvoiceStatusBatchRequest) (resp internal.BizResponse)

	CreateStudent(corpId string, student Student) (resp internal.BizResponse)
	BatchCreateStudent(corpId string, students []Student) (resp BatchStudentResponse)
	DeleteStudent(corpId string, userId string) (resp internal.BizResponse)
	BatchDeleteStudent(corpId string, userIdList []string) (resp BatchStudentResponse)
	UpdateStudent(corpId string, student Student) (resp internal.BizResponse)
	BatchUpdateStudent(corpId string, students []Student) (resp BatchStudentResponse)

	CreateParent(corpId string, parent Parent) (resp internal.BizResponse)
	BatchCreateParent(corpId string, parents []Parent) (resp BatchParentResponse)
	DeleteParent(corpId string, userId string) (resp internal.BizResponse)
	BatchDeleteParent(corpId string, userIdList []string) (resp BatchParentResponse)
	UpdateParent(corpId string, parent Parent) (resp internal.BizResponse)
	BatchUpdateParent(corpId string, parents []Parent) (resp BatchParentResponse)
	ListParentWithDepartmentId(corpId string, departmentId int32) (resp ListParentWithDepartmentIdResponse)

	SchoolUserGet(corpId string, userId string) (resp SchoolUserGetResponse)
	SchoolUserList(corpId string, departmentId uint32, fetchChild int) (resp SchoolUserListResponse)
	SetArchSyncMode(corpId string, mode int) (resp internal.BizResponse)
	GetSubScribeQrCode(corpId string) (resp GetSubScribeQrCodeResponse)
	SetSubScribeMode(corpId string, mode int) (resp internal.BizResponse)
	GetSubScribeMode(corpId string) (resp GetSubScribeModeResponse)
	BatchToExternalUserId(corpId string, mobiles []string) (resp BatchToExternalUserIdResponse)
	SetTeacherViewMode(corpId string, mode int) (resp internal.BizResponse)
	GetTeacherViewMode(corpId string) (resp GetTeacherViewModeResponse)
	GetAllowScope(corpId string, agentId int) (resp GetAllowScopeResponse)
	SetUpgradeInfo(corpId string, request UpgradeRequest) (resp UpgradeInfoResponse)

	SchoolDepartmentCreate(corpId string, department SchoolDepartment) (resp SchoolDepartmentCreateResponse)
	SchoolDepartmentUpdate(corpId string, department SchoolDepartment) (resp internal.BizResponse)
	SchoolDepartmentDelete(corpId string, departmentId int32) (resp internal.BizResponse)
	SchoolDepartmentList(corpId string, departmentId int32) (resp SchoolDepartmentListResponse)

	GetUserAllLivingId(corpId string, request GetUserAllLivingIdRequest) (resp GetUserAllLivingIdResponse)
	GetLivingInfo(corpId string, liveId string) (resp GetLivingInfoResponse)
	GetWatchStat(corpId string, request GetWatchStatRequest) (resp GetWatchStatResponse)
	GetUnWatchStat(corpId string, request GetWatchStatRequest) (resp GetUnWatchStatResponse)
	DeleteReplayData(corpId string, livingId string) (resp internal.BizResponse)

	GetPaymentResult(corpId string, paymentId string) (resp GetPaymentResultResponse)
	GetTrade(corpId string, request GetTradeRequest) (resp GetTradeResponse)

	GetJsApiTicket(corpId string) (resp TicketResponse)
	GetConfigSignature(corpId string, referer string) (resp JsTicketSignatureResponse)
	GetJsApiAgentTicket(corpId string, agentId int) (resp TicketResponse)
	GetAgentConfigSignature(corpId string, agentId int, referer string) (resp JsTicketSignatureResponse)

	KfAccountAdd(corpId string, account KfAccount) (resp KfAccountAddResponse)
	KfAccountDel(corpId string, kfId string) (resp internal.BizResponse)
	KfAccountUpdate(corpId string, account KfAccount) (resp internal.BizResponse)
	KfAccountList(corpId string, request KfAccountListRequest) (resp KfAccountListResponse)
	KfAddContactWay(corpId string, kfId string, scene string) (resp KfAccContactWayResponse)
	KfServicerAdd(corpId string, request KfServicerRequest) (resp KfServicerResponse)
	KfServicerDel(corpId string, request KfServicerRequest) (resp KfServicerResponse)
	KfServicerList(corpId string, kfId string) (resp KfServicerListResponse)
	KfServiceStateGet(corpId string, request KfServiceStateGetRequest) (resp KfServiceStateGetResponse)
	KfServiceStateTrans(corpId string, request KfServiceStateTransRequest) (resp KfServiceStateTransResponse)
	KfSyncMsg(corpId string, request KfSyncMsgRequest) (resp KfSyncMsgResponse)
	KfSendMsg(corpId string, request SendMsgRequest) (resp SendMsgResponse)
	KfSendMsgOnEvent(corpId string, request SendMsgOnEventRequest) (resp SendMsgResponse)
	KfCustomerBatchGet(corpId string, userList []string, needEnterSessionContext int) (resp KfCustomerBatchGetResponse)
	// KfGetCorpQualification 仅支持第三方应用，且需具有“微信客服->获取基础信息”权限
	KfGetCorpQualification(corpId string) (resp KfGetCorpQualificationResponse)
	KfGetUpgradeServiceConfig(corpId string) (resp KfGetUpgradeServiceConfigResponse)
	KfUpgradeService(corpId string, request UpgradeServiceRequest) (resp internal.BizResponse)
	KfCancelUpgradeService(corpId string, request CancelUpgradeServiceRequest) (resp internal.BizResponse)
	// KfGetCorpStatistic
	// 查询时间区间[start_time, end_time]为闭区间，最大查询跨度为31天，用户最多可获取最近180天内的数据。
	// 当天的数据需要等到第二天才能获取，建议在第二天早上六点以后再调用此接口获取前一天的数据
	KfGetCorpStatistic(corpId string, filter KfGetCorpStatisticFilter) (resp KfGetCorpStatisticResponse)
	// KfGetServicerStatistic
	// 查询时间区间[start_time, end_time]为闭区间，最大查询跨度为31天，用户最多可获取最近180天内的数据。
	// 当天的数据需要等到第二天才能获取，建议在第二天早上六点以后再调用此接口获取前一天的数据
	KfGetServicerStatistic(corpId string, filter KfGetServicerStatisticFilter) (resp KfGetServicerStatisticResponse)
	KfKnowLedgeAddGroup(corpId string, name string) (resp KfKnowLedgeAddGroupResponse)
	KfKnowLedgeDelGroup(corpId string, groupId string) (resp internal.BizResponse)
	KfKnowLedgeModGroup(corpId string, name string, groupId string) (resp internal.BizResponse)
	KfKnowLedgeListGroup(corpId string, filter KfKnowLedgeListGroupFilter) (resp KfKnowLedgeListGroupResponse)

	CreateNewOrder(request CreateOrderRequest) (resp OrderResponse)
	CreateReNewOrderJob(request CreateReNewOrderJobRequest) (resp CreateReNewOrderJobResponse)
	SubmitOrderJob(request SubmitOrderJobRequest) (resp OrderResponse)
	ListOrder(request ListOrderRequest) (resp ListOrderResponse)
	GetOrder(request GetOrderRequest) (resp GetOrderResponse)
	ListOrderAccount(request ListOrderAccountRequest) (resp ListOrderAccountResponse)
	ActiveAccount(request ActiveAccountRequest) (resp internal.BizResponse)
	BatchActiveAccount(request BatchActiveAccountRequest) (resp BatchActiveAccountResponse)
	GetActiveInfoByCode(request GetActiveInfoByCodeRequest) (resp GetActiveInfoByCodeResponse)
	BatchGetActiveInfoByCode(request BatchGetActiveInfoByCodeRequest) (resp BatchGetActiveInfoByCodeResponse)
	ListActivedAccount(request ListActivedAccountRequest) (resp ListActivedAccountResponse)
	GetActiveInfoByUser(request GetActiveInfoByUserRequest) (resp GetActiveInfoByUserResponse)
	BatchTransferLicense(request BatchTransferLicenseRequest) (resp BatchTransferLicenseResponse)
	GetAdminList(request GetAdminListRequest) (resp GetAdminListResponse)
	SetAutoActiveStatus(request SetAutoActiveStatusRequest) (resp internal.BizResponse)
	GetAutoActiveStatus(corpid string) (resp GetAutoActiveStatusResponse)

	GetPermitUserList(corpId string, T int) (resp GetPermitUserListResponse, err error)
	CheckSingleAgree(corpId string, request CheckSingleAgreeRequest) (resp CheckSingleAgreeResponse, err error)
	GetAuditGroupChat(corpId string, roomId string) (resp GetAuditGroupChatResponse, err error)
	// ExecuteCorpApi 用于执行未实现的接口，返回 []byte,error
	ExecuteCorpApi(corpId string, apiUrl string, query url.Values, data H) (body []byte, err error)

	IdConvertExternalTagId(corpId string, tagIdList []string) (resp IdConvertExternalTagIdResponse)
	CorpIdToOpenCorpId(corpId string) (resp CorpIdToOpenCorpIdResponse)
	UserIdToOpenUserId(corpId string, userIdList []string) (resp UserIdToOpenUserIdResponse)
	GetNewExternalUserId(corpId string, userIdList []string) (resp GetNewExternalUserIdResponse)
	GroupChatGetNewExternalUserId(corpId string, request GroupChatGetNewExternalUserIdRequest) (resp GetNewExternalUserIdResponse)

	RemindGroupMsgSend(corpId string, msgid string) (resp internal.BizResponse)
	CancelMomentTask(corpId string, momentId string) (resp internal.BizResponse)

	SetProxy(proxyUrl string)
	SetDebug(debug bool)
}

type weWork struct {
	corpId              string
	providerSecret      string
	suiteId             string
	suiteSecret         string
	suiteTicket         string
	suiteToken          string
	suiteEncodingAesKey string
	cache               *badger.DB
	logger              *zap.Logger
	engine              *gorm.DB
	getAppSecretFunc    func(corpId string) (corpid string, secret string, customizedApp bool)
	getAgentIdFunc      func(corpId string) (appId int)
	httpClient          *resty.Client
}

type WeWorkConfig struct {
	CorpId              string
	ProviderSecret      string
	SuiteId             string
	SuiteSecret         string
	SuiteToken          string
	SuiteEncodingAesKey string
	Dsn                 string
}

const (
	UserAgent   = "haiyiyun-weworksdk"
	ContentType = "application/json;charset=UTF-8"
	qyApiHost   = "https://qyapi.weixin.qq.com"
)

func NewWeWork(c WeWorkConfig) IWeWork {
	var ww = new(weWork)
	ww.corpId = c.CorpId
	ww.providerSecret = c.ProviderSecret
	ww.suiteId = c.SuiteId
	ww.suiteSecret = c.SuiteSecret
	ww.suiteToken = c.SuiteToken
	ww.suiteEncodingAesKey = c.SuiteEncodingAesKey
	ww.cache, _ = badger.Open(badger.DefaultOptions("./" + c.CorpId + "_cache.db").WithIndexCacheSize(10 << 20))
	ww.logger = logger
	ww.httpClient = resty.New().
		SetHeader("User-Agent", UserAgent).
		SetHeader("Content-Type", ContentType).
		SetBaseURL(qyApiHost)
	ww.httpClient.AddRetryCondition(func(r *resty.Response, err error) bool {
		var biz internal.BizResponse
		json.Unmarshal(r.Body(), &biz)
		if biz.ErrCode == 42001 {
			ww.cache.DropPrefix([]byte("corpToken"))
			return true
		}
		return false
	})
	if c.Dsn != "" {
		ww.engine, _ = gorm.Open(mysql.Open(c.Dsn), &gorm.Config{})
	}
	// 默认获取企业token函数
	return ww
}

func (ww *weWork) Logger() *zap.Logger {
	return ww.logger
}
func (ww *weWork) SetProxy(proxyUrl string) {
	if _, ok := url.Parse(proxyUrl); ok == nil {
		ww.httpClient.SetProxy(proxyUrl)
	}
}
func (ww *weWork) SetDebug(debug bool) {
	ww.httpClient.SetDebug(debug)
}

func (ww *weWork) getRequest(corpid string) *resty.Request {
	R := ww.httpClient.R().SetQueryParam("access_token", ww.getCorpToken(corpid))
	if os.Getenv("debug") != "" {
		R.SetQueryParam("debug", "1")
	}
	return R
}

func (ww *weWork) getProviderRequest() *resty.Request {
	R := ww.httpClient.R().SetQueryParam("provider_access_token", ww.getProviderToken())
	if os.Getenv("debug") != "" {
		R.SetQueryParam("debug", "1")
	}
	return R
}

// GetCorpId 返回服务商corpId
func (ww *weWork) GetCorpId() string {
	return ww.corpId
}

// GetSuiteId 返回服务商SuiteId
func (ww *weWork) GetSuiteId() string {
	return ww.suiteId
}

// GetSuiteToken 返回服务商配置的SuiteToken
func (ww *weWork) GetSuiteToken() string {
	return ww.suiteToken
}

// GetSuiteEncodingAesKey 返回服务商配置的EncodingAesKey
func (ww *weWork) GetSuiteEncodingAesKey() string {
	return ww.suiteEncodingAesKey
}

// GetAgentId 获取应用的AgentId;三方或代开发应用会将信息存入数据库中
// 如果修改了表结构，需要配合 SetAgentIdFunc 使用
func (ww *weWork) GetAgentId(corpId string) (appId int) {
	if ww.getAgentIdFunc != nil {
		return ww.getAgentIdFunc(corpId)
	} else {
		return ww.defaultAgentIdFunc(corpId)
	}
}
