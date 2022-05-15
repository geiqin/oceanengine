package enum

// GroupFlowType 落地页组流量分配方式
type GroupFlowType string

const (
	// FLOW_DISTRIBUTION_TYPE_INTELLIGENCE 智能分配
	FLOW_DISTRIBUTION_TYPE_INTELLIGENCE GroupFlowType = "FLOW_DISTRIBUTION_TYPE_INTELLIGENCE"
	// FLOW_DISTRIBUTION_TYPE_AVERAGE 平均分配
	FLOW_DISTRIBUTION_TYPE_AVERAGE GroupFlowType = "FLOW_DISTRIBUTION_TYPE_AVERAGE"
)