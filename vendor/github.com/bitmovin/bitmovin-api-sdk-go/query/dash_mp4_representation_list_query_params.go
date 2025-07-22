package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type DashMp4RepresentationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *DashMp4RepresentationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
