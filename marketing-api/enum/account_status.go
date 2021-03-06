package enum

// AccountStatus 广告主状态
type AccountStatus string

const (
	// ACCOUNT_STATUS_DISABLE 已禁用
	ACCOUNT_STATUS_DISABLE AccountStatus = "STATUS_DISABLE"
	// ACCOUNT_STATUS_PENDING_CONFIRM 申请待审核
	ACCOUNT_STATUS_PENDING_CONFIRM AccountStatus = "STATUS_PENDING_CONFIRM"
	// ACCOUNT_STATUS_PENDING_VERIFIED 待验证
	ACCOUNT_STATUS_PENDING_VERIFIED AccountStatus = "STATUS_PENDING_VERIFIED"
	// ACCOUNT_STATUS_CONFIRM_FAIL 审核不通过
	ACCOUNT_STATUS_CONFIRM_FAIL AccountStatus = "STATUS_CONFIRM_FAIL"
	// ACCOUNT_STATUS_ENABLE 已审核
	ACCOUNT_STATUS_ENABLE AccountStatus = "STATUS_ENABLE"
	// ACCOUNT_STATUS_CONFIRM_FAIL_END CRM审核不通过
	ACCOUNT_STATUS_CONFIRM_FAIL_END AccountStatus = "STATUS_CONFIRM_FAIL_END"
	// ACCOUNT_STATUS_PENDING_CONFIRM_MODIFY 修改待审核
	ACCOUNT_STATUS_PENDING_CONFIRM_MODIFY AccountStatus = "STATUS_PENDING_CONFIRM_MODIFY"
	// ACCOUNT_STATUS_CONFIRM_MODIFY_FAIL 修改审核不通过
	ACCOUNT_STATUS_CONFIRM_MODIFY_FAIL AccountStatus = "STATUS_CONFIRM_MODIFY_FAIL"
	// ACCOUNT_STATUS_LIMIT 限制
	ACCOUNT_STATUS_LIMIT AccountStatus = "STATUS_LIMIT"
	// ACCOUNT_STATUS_WAIT_FOR_BPM_AUDIT 等待CRM审核
	ACCOUNT_STATUS_WAIT_FOR_BPM_AUDIT AccountStatus = "STATUS_WAIT_FOR_BPM_AUDIT"
	// ACCOUNT_STATUS_WAIT_FOR_PUBLIC_AUTH 待对公验证
	ACCOUNT_STATUS_WAIT_FOR_PUBLIC_AUTH AccountStatus = "STATUS_WAIT_FOR_PUBLIC_AUTH"
	// ACCOUNT_STATUS_SELF_SERVICE_UNAUDITED 自助开户待验证资质
	ACCOUNT_STATUS_SELF_SERVICE_UNAUDITED AccountStatus = "STATUS_SELF_SERVICE_UNAUDITED"
)
