package auth

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/liuhailove/tc-base-go/protocol/tc"
)

type VideoGrant struct {
	// 对房间的操作
	RoomCreate bool `json:"roomCreate,omitempty"`
	RoomList   bool `json:"roomList,omitempty"`
	RoomRecord bool `json:"roomRecord,omitempty"`

	// 对特定房间的操作
	RoomAdmin bool   `json:"roomAdmin,omitempty"`
	RoomJoin  bool   `json:"roomJoin,omitempty"`
	Room      string `json:"room,omitempty"`

	// 房间内的权限，如果没有明确设置任何权限
	// 它将被授予所有发布和订阅权限
	CanPublish     *bool `json:"canPublish,omitempty"`
	CanSubscribe   *bool `json:"canSubscribe,omitempty"`
	CanPublishData *bool `json:"canPublishData,omitempty"`
	// 参与者可以发布的 TrackSource 类型。
	// 设置后，它将取代 CanPublish。只有此处明确设置的来源才能发布
	CanPublishSources []string `json:"canPublishSources,omitempty"` // 键跟踪每个源
	// 默认情况下，参与者不允许更新自己的元数据
	CanUpdateOwnMetadata *bool `json:"canUpdateOwnMetadata,omitempty"`

	// 对入口的操作
	IngressAdmin bool `json:"ingressAdmin,omitempty"` // 适用于所有入口
	// 参与者对其他参与者不可见
	Hidden bool `json:"hidden,omitempty"`
	// 向房间表明当前参与者是录音者
	Recorder bool `json:"recorder,omitempty"`
}

type ClaimGrants struct {
	Identity string      `json:"-"`
	Name     string      `json:"name,omitempty"`
	Video    *VideoGrant `json:"video,omitempty"`
	// 用于验证消息体的完整性
	Sha256   string `json:"sha256,omitempty"`
	Metadata string `json:"metadata,omitempty"`
}

func (c *ClaimGrants) Clone() *ClaimGrants {
	if c == nil {
		return nil
	}

	clone := *c
	clone.Video = c.Video.Clone()

	return &clone
}

func (v *VideoGrant) SetCanPublish(val bool) {
	v.CanPublish = &val
}

func (v *VideoGrant) SetCanPublishData(val bool) {
	v.CanPublishData = &val
}

func (v *VideoGrant) SetCanSubscribe(val bool) {
	v.CanSubscribe = &val
}

func (v *VideoGrant) SetCanPublishSources(sources []tc.TrackSource) {
	v.CanPublishSources = make([]string, 0, len(sources))
	for _, s := range sources {
		v.CanPublishSources = append(v.CanPublishSources, sourceToString(s))
	}
}

func (v *VideoGrant) SetCanUpdateOnwMetadata(val bool) {
	v.CanUpdateOwnMetadata = &val
}

func (v *VideoGrant) GetCanPublish() bool {
	if v.CanPublish == nil {
		return true
	}
	return *v.CanPublish
}

func (v *VideoGrant) GetCanPublishSource(source tc.TrackSource) bool {
	if !v.GetCanPublish() {
		return false
	}
	// 不要区分 nil 和 unset，因为这种区别在序列化后就无法存在
	if len(v.CanPublishSources) == 0 {
		return true
	}
	sourceStr := sourceToString(source)
	for _, s := range v.CanPublishSources {
		if s == sourceStr {
			return true
		}
	}
	return false
}

func (v *VideoGrant) GetCanPublishSources() []tc.TrackSource {
	if len(v.CanPublishSources) == 0 {
		return nil
	}

	sources := make([]tc.TrackSource, 0, len(v.CanPublishSources))
	for _, s := range v.CanPublishSources {
		sources = append(sources, sourceToProto(s))
	}
	return sources
}

func (v *VideoGrant) GetCanPublishData() bool {
	if v.CanPublishData == nil {
		return v.GetCanPublish()
	}
	return *v.CanPublishData
}

func (v *VideoGrant) GetCanSubscribe() bool {
	if v.CanSubscribe == nil {
		return true
	}
	return *v.CanSubscribe
}

func (v *VideoGrant) GetCanUpdateOwnMetadata() bool {
	if v.CanUpdateOwnMetadata == nil {
		return false
	}
	return *v.CanUpdateOwnMetadata
}

func (v *VideoGrant) MatchesPermission(permission *tc.ParticipantPermission) bool {
	if permission == nil {
		return false
	}

	if v.GetCanPublish() != permission.CanPublish {
		return false
	}
	if v.GetCanPublishData() != permission.CanPublishData {
		return false
	}
	if v.GetCanSubscribe() != permission.CanSubscribe {
		return false
	}
	if v.GetCanUpdateOwnMetadata() != permission.CanUpdateMetadata {
		return false
	}
	if v.Hidden != permission.Hidden {
		return false
	}
	if v.Recorder != permission.Recorder {
		return false
	}
	if !slices.Equal(v.GetCanPublishSources(), permission.CanPublishSources) {
		return false
	}

	return true
}

func (v *VideoGrant) UpdateFromPermission(permission *tc.ParticipantPermission) {
	if permission == nil {
		return
	}

	v.SetCanPublish(permission.CanPublish)
	v.SetCanPublishData(permission.CanPublishData)
	v.SetCanPublishSources(permission.CanPublishSources)
	v.SetCanSubscribe(permission.CanSubscribe)
	v.SetCanUpdateOnwMetadata(permission.CanUpdateMetadata)
	v.Hidden = permission.Hidden
	v.Recorder = permission.Recorder
}

func (v *VideoGrant) ToPermission() *tc.ParticipantPermission {
	pp := &tc.ParticipantPermission{
		CanPublish:        v.GetCanPublish(),
		CanPublishData:    v.GetCanPublishData(),
		CanSubscribe:      v.GetCanSubscribe(),
		CanPublishSources: v.GetCanPublishSources(),
		CanUpdateMetadata: v.GetCanUpdateOwnMetadata(),
		Hidden:            v.Hidden,
		Recorder:          v.Recorder,
	}
	return pp
}

func (v *VideoGrant) Clone() *VideoGrant {
	if v == nil {
		return nil
	}

	clone := *v

	if v.CanPublish != nil {
		canPublish := *v.CanPublish
		clone.CanPublish = &canPublish
	}

	if v.CanSubscribe != nil {
		canSubscribe := *v.CanSubscribe
		clone.CanSubscribe = &canSubscribe
	}

	if v.CanPublishData != nil {
		canPublishData := *v.CanPublishData
		clone.CanPublishData = &canPublishData
	}

	if v.CanPublishSources != nil {
		clone.CanPublishSources = make([]string, len(v.CanPublishSources))
		copy(clone.CanPublishSources, v.CanPublishSources)
	}

	if v.CanUpdateOwnMetadata != nil {
		canUpdateOwnMetadata := *v.CanUpdateOwnMetadata
		clone.CanUpdateOwnMetadata = &canUpdateOwnMetadata
	}

	return &clone
}

func sourceToString(source tc.TrackSource) string {
	return strings.ToLower(source.String())
}

func sourceToProto(sourceStr string) tc.TrackSource {
	switch sourceStr {
	case "camera":
		return tc.TrackSource_CAMERA
	case "microphone":
		return tc.TrackSource_MICROPHONE
	case "screen_share":
		return tc.TrackSource_SCREEN_SHARE
	case "screen_share_audio":
		return tc.TrackSource_SCREEN_SHARE_AUDIO
	default:
		return tc.TrackSource_UNKNOWN
	}
}
