package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsInputStreamsAudioMixApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsInputStreamsAudioMixApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsInputStreamsAudioMixApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsInputStreamsAudioMixApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsInputStreamsAudioMixApi) Create(encodingId string, audioMixInputStream model.AudioMixInputStream) (*model.AudioMixInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.AudioMixInputStream
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/input-streams/audio-mix", &audioMixInputStream, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsAudioMixApi) Delete(encodingId string, inputStreamId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/input-streams/audio-mix/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsAudioMixApi) Get(encodingId string, inputStreamId string) (*model.AudioMixInputStream, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["input_stream_id"] = inputStreamId
    }

    var responseModel *model.AudioMixInputStream
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/audio-mix/{input_stream_id}", nil, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsInputStreamsAudioMixApi) List(encodingId string, queryParams ...func(*query.AudioMixInputStreamListQueryParams)) (*pagination.AudioMixInputStreamsListPagination, error) {
    queryParameters := &query.AudioMixInputStreamListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.AudioMixInputStreamsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/input-streams/audio-mix", nil, &responseModel, reqParams)
    return responseModel, err
}

