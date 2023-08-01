package e

var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	AuthError:                  "认证失败",
	ErrorPermissionDenied:      "权限不足",
	ErrorExistBareMetal:        "已存在该机器",
	ErrorCountBareMetalFail:    "计算总机器数量错误",
	ErrorAuthCheckTokenTimeout: "检查Token超时",
	ErrorAuthCheckTokenFail:    "检查Token错误",
	ErrorAuth:                  "认证失败",
	ErrorAuthToken:             "认证Token错误",
	ErrorParserToken:           "解析Token错误",

	ErrorImageStatsFail: "获取镜像文件Stats失败",

	ErrorUpdateRegionInfo: "更新Region信息错误",
	ErrorDeleteRegionInfo: "删除Region信息错误",
	ErrorGetRegionInfo:    "获取Region信息错误",

	ErrorClusterExist:            "集群名称已存在",
	ErrorClusterNotExist:         "集群不存在",
	ErrorGetClusterInfo:          "获取集群信息错误",
	ErrorUpdateClusterInfo:       "更新集群信息错误",
	ErrorDeleteClusterInfo:       "删除集群信息错误",
	ErrorClusterQuotaSetting:     "集群限制处理错误",
	ErrorGetReceiverInfo:         "获取接收端错误",
	ErrorUpdateReceiverInfo:      "更新接收端错误",
	ErrorDeleteReceiverInfo:      "删除接收端错误",
	ErrorReceiverNameConflict:    "接收端名称冲突",
	ErrorReceiverNameInUse:       "接收端名称正在使用，无法删除",
	ErrorShowRuleInfo:            "获取告警规则信息错误",
	ErrorGetAlertInfo:            "获取告警信息错误",
	ErrorUpdateRuleInfo:          "更新告警规则错误",
	ErrorDeleteRuleInfo:          "删除告警规则错误",
	ErrorCreateRuleInfo:          "创建告警规则错误",
	ErrorClusterNoSshConfig:      "集群没有配置SSH信息",
	ErrorPrometheusConfigSyncing: "prometheus 监控信息正在同步中，请稍后再试",

	ErrorRestartPrometheus:   "重启监控系统失败",
	ErrorRestartAlertManager: "重启告警系统失败",

	ErrorShowEmailsInfo:     "获取邮件服务器信息错误",
	ErrorCreateEmailsInfo:   "配置邮件服务器信息错误",
	ErrorTestEmailsInfo:     "测试邮件服务器信息错误",
	ErrorShowSmsInfo:        "获取短信服务器信息错误",
	ErrorCreateSmsInfo:      "配置短信服务器信息错误",
	ErrorTestSmsInfo:        "测试短信服务器错误",
	ErrorActionHypervisor:   "执行Hypervisor操作错误",
	ErrorAggregationList:    "获取主机聚合列表错误",
	ErrorAggregationPost:    "创建主机聚合错误",
	ErrorAvailableZoneList:  "获取主机列表错误",
	ErrorAggregationAddHost: "聚合添加主机错误",
	ErrorAggregationDelHost: "聚合删除主机错误",
	ErrorAggregationMeta:    "聚合添加metadata错误",
	ErrorAggregationDelete:  "删除主机聚合错误",
	ErrorAggregationUpdate:  "更新主机聚合失败",

	ErrorCreateSubnet:          "创建子网错误",
	ErrorDeleteSubnet:          "删除子网错误",
	ErrorUpdateSubnet:          "更新子网错误",
	ErrorGetUserNetwork:        "获取用户网络错误",
	ErrorGetNetworkInfo:        "获取网络信息错误",
	ErrorGetNetworkSubnet:      "获取网络子网错误",
	ErrorGetNetworkSubnetPorts: "获取网络子网端口错误",
	ErrorGetPortRelationInfo:   "获取端口关联信息错误",
	ErrorGetRoutersInfo:        "获取路由列表信息错误",
	ErrorGetRouterPorts:        "获取路由端口列表信息错误",
	ErrorGetFloatingIPInfo:     "获取浮动IP列表信息错误",

	ErrorCreateDatabase:        "创建数据库版本错误",
	ErrorDeleteDatabase:        "删除数据库版本错误",
	ErrorListDatabaseInstances: "获取数据库列表错误",
	ErrorDatabaseInstanceInfo:  "获取数据库列表错误",

	ErrorGetImageList:      "获取镜像列表错误",
	ErrorUploadImage:       "上传镜像错误",
	ErrorDeleteImage:       "删除镜像错误",
	ErrorUpdateImage:       "更新镜像状态错误",
	ErrorUploadCustomImage: "上传自定义镜像错误",

	ErrorCreateRegion: "添加Region错误",

	ErrorGetHost:       "获取主机错误",
	ErrorOperationHost: "操作主机错误",
	ErrorGetDeployLog:  "获取部署日志错误",

	ErrorCreateShareType: "创建共享类型错误",

	ErrorResetHostStatus: "重置主机状态错误",

	ErrorUserList:     "获取用户列表错误",
	ErrorUserDelete:   "删除用户失败",
	ErrorUserDetail:   "获取用户详情失败",
	ErrorUserUpdate:   "更新用户详情失败",
	ErrorUserCreate:   "创建用户失败",
	ErrorPassword:     "用户名密码不正确",
	ErrorUserNotExist: "用户不存在",
	ErrorUserExist:    "用户已存在",
	ErrorUserForbid:   "用户已禁用",
	ErrDeleteAdmin:    "无法删除admin账户",

	ErrorGroupList:   "获取组列表错误",
	ErrorGroupDetail: "获取组详情错误",
	ErrorGroupDelete: "删除组错误",
	ErrorGroupCreate: "创建组错误",
	ErrorGroupUpdate: "更新组错误",

	ErrorActionList: "获取操作列表错误",

	ErrorUserRole: "获取用户角色错误",

	ErrorServerList:           "获取服务器列表错误",
	ErrorServerShow:           "获取服务器详情错误",
	ErrorServerSecurityGroups: "获取服务器安全组错误",
	ErrorServerVNC:            "获取服务器VNC错误",
	ErrorServerActions:        "获取服务器操作错误",
	ErrorQueryLog:             "获取查询日志错误",
	ErrorQueryOperateLog:      "获取操作日志信息错误",
	ErrorQueryLoginLog:        "获取登陆日志信息错误",

	ErrorServerMigrate:  "执行云主机热迁移错误",
	ErrorServerRescue:   "执行云主机救援错误",
	ErrorServerUnRescue: "执行云主机终止救援错误",
	ErrorServerMonitor:  "获取云主机监控错误",
	ErrorServerRebuild:  "执行云主机重建错误",

	ErrorDeployList:             "获取部署机列表错误",
	ErrorDeployInfo:             "获取部署机详情错误",
	ErrorDeployCreate:           "创建部署机错误",
	ErrorDeployUpdate:           "更新部署机错误",
	ErrorDeployDelete:           "删除部署机错误",
	ErrorDeployUpdateLabel:      "更新部署机标签错误",
	ErrorDeployAddHost:          "添加部署机主机错误",
	ErrorDeployFetchHost:        "获取部署机主机错误",
	ErrorDeployUpdateGlobal:     "更新部署机全局配置错误",
	ErrorDeployFetchGlobal:      "获取部署机全局配置错误",
	ErrorDeployFetchRealtimeLog: "获取部署机实时日志错误",
	ErrorDeployCluster:          "部署机集群错误",
	ErrorDeployManagerInfo:      "获取部署机管理信息错误",
	ErrorDeployFetchStep:        "获取部署步骤信息错误",

	ErrorFetchLoadBalancesError: "获取负载均衡列表错误",
	ErrorFetchLoadBalanceError:  "获取负载均衡详情错误",

	ErrorFetchVolumes:                  "获取卷列表错误",
	ErrorPrometheusServerTimeout:       "Prometheus服务超时",
	ErrorPrometheusServerError:         "无法连接Prometheus服务",
	ErrorPrometheusServerBadData:       "无效的prometheus请求,或者请求数据密度过高",
	ErrorPrometheusUrl:                 "错误的prometheus地址",
	ErrorPrometheusNoCluster:           "请在prometheus的external_labels配置中增加cluster标签来标识集群",
	ErrorPrometheusClusterNameConflict: "集群名称冲突，请检查集群名称是否已存在或者prometheus的external_labels中的cluster名称和已有的集群是否冲突",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

type ApiException struct {
	code       int
	statusCode int
}

func (a *ApiException) Error() string {
	return GetMsg(a.code)
}

func (a *ApiException) GetCode() int {
	return a.code
}

func (a *ApiException) GetStatusCode() int {
	return a.statusCode
}

func RaiseApiException(statusCode int, code int) error {
	return &ApiException{code, statusCode}
}
