package query

import(
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
)

type Fmp4MuxingListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Fmp4MuxingListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
