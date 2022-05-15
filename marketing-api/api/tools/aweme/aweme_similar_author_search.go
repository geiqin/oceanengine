package aweme

import (
	"github.com/geiqin/oceanengine/marketing-api/core"
	"github.com/geiqin/oceanengine/marketing-api/model/tools/aweme"
)

// AwemeSimilarAuthorSearch 查询抖音类似帐号
func AwemeSimilarAuthorSearch(clt *core.SDKClient, accessToken string, req *aweme.AwemeSimilarAuthorSearchRequest) ([]aweme.Author, error) {
	var resp aweme.AwemeSimilarAuthorSearchResponse
	err := clt.Get("2/tools/aweme_similar_author_search/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
