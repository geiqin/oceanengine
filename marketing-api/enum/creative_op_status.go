package enum

// CreativeOptStatus 创意操作状态
type CreativeOptStatus string

const (
	// CREATIVE_OP_STATUS_ENABLE 启用
	CREATIVE_OP_STATUS_ENABLE CreativeOptStatus = "CREATIVE_STATUS_ENABLE"
	// CREATIVE_OP_STATUS_DISABLE 暂停
	CREATIVE_OP_STATUS_DISABLE CreativeOptStatus = "CREATIVE_STATUS_DISABLE"
	// CREATIVE_OP_STATUS_DELETE 删除
	CREATIVE_OP_STATUS_DELETE CreativeOptStatus = "CREATIVE_STATUS_DELETE"
)
