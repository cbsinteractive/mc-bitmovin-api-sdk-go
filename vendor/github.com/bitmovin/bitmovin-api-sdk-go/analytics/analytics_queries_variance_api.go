package analytics
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
)

type AnalyticsQueriesVarianceApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsQueriesVarianceApi(configs ...func(*common.ApiClient)) (*AnalyticsQueriesVarianceApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsQueriesVarianceApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsQueriesVarianceApi) Create(analyticsVarianceQueryRequest model.AnalyticsVarianceQueryRequest) (*model.AnalyticsResponse, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.AnalyticsResponse
    err := api.apiClient.Post("/analytics/queries/variance", &analyticsVarianceQueryRequest, &responseModel, reqParams)
    return responseModel, err
}

