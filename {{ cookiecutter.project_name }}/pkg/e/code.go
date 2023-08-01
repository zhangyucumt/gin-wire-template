package e

const (
	SUCCESS               = 200
	ERROR                 = 500
	InvalidParams         = 400
	AuthError             = 401
	ErrorPermissionDenied = 403

	ErrorUserList     = 1001
	ErrorUserDelete   = 1002
	ErrorUserDetail   = 1003
	ErrorUserUpdate   = 1004
	ErrorUserCreate   = 1005
	ErrorUserRole     = 1006
	ErrorPassword     = 1007
	ErrorUserExist    = 1008
	ErrorUserNotExist = 1009
	ErrorUserForbid   = 1010
	ErrDeleteAdmin    = 1011

	ErrorGroupList   = 2001
	ErrorGroupDetail = 2002
	ErrorGroupDelete = 2003
	ErrorGroupCreate = 2004
	ErrorGroupUpdate = 2005

	ErrorActionList = 3001

	ErrorExistBareMetal     = 10001
	ErrorCountBareMetalFail = 10002

	ErrorAuthCheckTokenTimeout = 20001
	ErrorAuthCheckTokenFail    = 20002
	ErrorAuth                  = 20003
	ErrorAuthToken             = 20004
	ErrorParserToken           = 20005

	ErrorImageStatsFail = 30001

	ErrorUpdateRegionInfo = 40001
	ErrorDeleteRegionInfo = 40002
	ErrorGetRegionInfo    = 40003

	ErrorUpdateClusterInfo = 40004
	ErrorDeleteClusterInfo = 40005
	ErrorGetClusterInfo    = 40006
	ErrorGetClusters       = 40007
	ErrorClusterExist      = 40008
	ErrorClusterNotExist   = 40009

	ErrorClusterQuotaSetting = 40010

	ErrorGetReceiverInfo      = 400011
	ErrorUpdateReceiverInfo   = 40012
	ErrorDeleteReceiverInfo   = 40013
	ErrorReceiverNameConflict = 40025

	ErrorShowEmailsInfo   = 40014
	ErrorCreateEmailsInfo = 40015
	ErrorTestEmailsInfo   = 40016

	ErrorUpdateRuleInfo = 40017
	ErrorDeleteRuleInfo = 40018
	ErrorShowRuleInfo   = 40019
	ErrorCreateRuleInfo = 40020
	ErrorGetRuleList    = 400021

	ErrorRestartPrometheus   = 400022
	ErrorRestartAlertManager = 400023
	ErrorReceiverNameInUse   = 400026

	ErrorClusterNoSshConfig      = 400027
	ErrorPrometheusConfigSyncing = 400028

	ErrorGetAlertInfo       = 400050
	ErrorShowSmsInfo        = 400051
	ErrorCreateSmsInfo      = 400052
	ErrorTestSmsInfo        = 400053
	ErrorActionHypervisor   = 50001
	ErrorAggregationList    = 50002
	ErrorAggregationPost    = 50003
	ErrorAvailableZoneList  = 50004
	ErrorAggregationAddHost = 50005
	ErrorAggregationDelHost = 50006
	ErrorAggregationMeta    = 50007
	ErrorAggregationDelete  = 50008
	ErrorAggregationUpdate  = 50009

	ErrorCreateSubnet          = 60001
	ErrorDeleteSubnet          = 60002
	ErrorUpdateSubnet          = 60003
	ErrorGetUserNetwork        = 60004
	ErrorGetNetworkInfo        = 60005
	ErrorGetNetworkSubnet      = 60006
	ErrorGetNetworkSubnetPorts = 60007
	ErrorGetPortRelationInfo   = 60008

	ErrorGetRoutersInfo = 60009
	ErrorGetRouterPorts = 60010

	ErrorGetFloatingIPInfo = 60011

	ErrorCreateDatabase        = 70001
	ErrorDeleteDatabase        = 70002
	ErrorListDatabaseInstances = 70003
	ErrorDatabaseInstanceInfo  = 70004

	ErrorGetImageList      = 80001
	ErrorUploadImage       = 80002
	ErrorDeleteImage       = 80003
	ErrorUpdateImage       = 80004
	ErrorUploadCustomImage = 80005

	ErrorCreateRegion = 90001

	ErrorGetHost       = 100001
	ErrorOperationHost = 100002
	ErrorCreateHost    = 100003
	ErrorGetDeployLog  = 100004

	ErrorCreateShareType = 110001

	ErrorResetHostStatus = 120001

	ErrorServerList           = 130001
	ErrorServerShow           = 130002
	ErrorServerSecurityGroups = 130003
	ErrorServerVNC            = 130004
	ErrorServerActions        = 130005
	ErrorServerMigrate        = 130006
	ErrorServerRescue         = 130007
	ErrorServerUnRescue       = 130008
	ErrorServerMonitor        = 130009
	ErrorServerRebuild        = 130010

	ErrorQueryLog        = 140001
	ErrorQueryOperateLog = 140002
	ErrorQueryLoginLog   = 140003

	ErrorDeployList             = 15001
	ErrorDeployInfo             = 15002
	ErrorDeployCreate           = 15003
	ErrorDeployUpdate           = 15004
	ErrorDeployDelete           = 15005
	ErrorDeployUpdateLabel      = 15006
	ErrorDeployAddHost          = 15007
	ErrorDeployFetchHost        = 15008
	ErrorDeployUpdateGlobal     = 15009
	ErrorDeployFetchGlobal      = 15010
	ErrorDeployFetchRealtimeLog = 15011
	ErrorDeployCluster          = 15012
	ErrorDeployManagerInfo      = 15013
	ErrorDeployFetchStep        = 15014

	ErrorFetchLoadBalancesError = 16001
	ErrorFetchLoadBalanceError  = 16002

	ErrorFetchVolumes = 17001

	ErrorPrometheusServerTimeout       = 18001
	ErrorPrometheusServerError         = 18002
	ErrorPrometheusServerBadData       = 18003
	ErrorPrometheusUrl                 = 18004
	ErrorPrometheusNoCluster           = 18005
	ErrorPrometheusClusterNameConflict = 18006
)
