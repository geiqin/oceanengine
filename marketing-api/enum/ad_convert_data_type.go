package enum

// AdConvertDataType 转化统计方式
type AdConvertDataType string

const (
	// AD_CONVERT_DATA_TYPE_EVERY_ONE 每一次
	AD_CONVERT_DATA_TYPE_EVERY_ONE AdConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE"
	// AD_CONVERT_DATA_TYPE_ONLY_ONE 仅一次
	AD_CONVERT_DATA_TYPE_ONLY_ONE AdConvertDataType = "AD_CONVERT_DATA_TYPE_ONLY_ONE"
	// AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW 每一次(新)
	AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW AdConvertDataType = "AD_CONVERT_DATA_TYPE_EVERY_ONE_NEW"
	// AD_CONVERT_DATA_TYPE_LESS_THAN_TEN 一天内不超过10次
	AD_CONVERT_DATA_TYPE_LESS_THAN_TEN AdConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_TEN"
	// AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED 一天内不超过100次
	AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED AdConvertDataType = "AD_CONVERT_DATA_TYPE_LESS_THAN_HUNDRED"
	// AD_CONVERT_DATA_TYPE_PER_TIMES 每一次
	AD_CONVERT_DATA_TYPE_PER_TIMES AdConvertDataType = "AD_CONVERT_DATA_TYPE_PER_TIMES"
)
