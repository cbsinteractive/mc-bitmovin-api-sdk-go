package model
import (
	"time"
)

type HeAacV1AudioConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Target bitrate for the encoded audio in bps (required)
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate Hz
	Rate *float64 `json:"rate,omitempty"`
	// Channel layout of the audio codec configuration
	ChannelLayout AacChannelLayout `json:"channelLayout,omitempty"`
	// Sets the Spectral Band Replication (SBR) and Parameteric Stereo (PS) signaling style.
	Signaling HeAacV1Signaling `json:"signaling,omitempty"`
}
func (o HeAacV1AudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_HE_AAC_V1
}

