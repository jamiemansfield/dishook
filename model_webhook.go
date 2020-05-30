package dishook

// See https://discord.com/developers/docs/resources/webhook#execute-webhook
type Webhook struct {
	Content         string          `json:"content,omitempty"`
	Username        string          `json:"username,omitempty"`
	AvatarURL       string          `json:"avatar_url,omitempty"`
	TTS             bool            `json:"tts,omitempty"`
	File            string          `json:"file,omitempty"`
	Embeds          []*Embed        `json:"embeds,omitempty"`
	PayloadJSON     string          `json:"payload_json,omitempty"`
	AllowedMentions *AllowedMention `json:"allowed_mentions,omitempty"`
}

// See https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mention-types
const (
	AllowedMentionTypeRoles    = "roles"
	AllowedMentionTypeUsers    = "users"
	AllowedMentionTypeEveryone = "everyone"
)

// See https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMention struct {
	Parse []string `json:"parse,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Users []string `json:"users,omitempty"`
}
