package enum

// RegionLevel 地域层级
type RegionLevel string

const (
	// REGION_LEVEL_PROVINCE 省级
	REGION_LEVEL_PROVINCE RegionLevel = "REGION_LEVEL_PROVINCE"
	// REGION_LEVEL_CITY 市级
	REGION_LEVEL_CITY RegionLevel = "REGION_LEVEL_CITY"
	// REGION_LEVEL_DISTRICT 区县级
	REGION_LEVEL_DISTRICT RegionLevel = "REGION_LEVEL_DISTRICT"
	// REGION_LEVEL_BUSINESS_DISTRICT 商业区级
	REGION_LEVEL_BUSINESS_DISTRICT RegionLevel = "REGION_LEVEL_BUSINESS_DISTRICT"
)
