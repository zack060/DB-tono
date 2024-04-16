// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package discord_Table

import (
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

type DiscordGatewayEvent struct {
	ID                                  sql.NullInt32
	Type                                sql.NullString
	Content                             sql.NullString
	AuthorID                            sql.NullString
	AuthorUsername                      sql.NullString
	AuthorGlobalName                    sql.NullString
	AuthorAvatar                        sql.NullString
	AuthorBannerColor                   sql.NullString
	Attachments                         pqtype.NullRawMessage
	Embeds                              pqtype.NullRawMessage
	Mentions                            pqtype.NullRawMessage
	MentionRoles                        sql.NullString
	Pinned                              sql.NullBool
	MentionEveryone                     sql.NullBool
	Tts                                 sql.NullBool
	Timestamp                           sql.NullTime
	EditedTimestamp                     sql.NullTime
	Flags                               sql.NullInt32
	Components                          pqtype.NullRawMessage
	Nonce                               sql.NullString
	ReferencedMessage                   pqtype.NullRawMessage
	InviteLink                          sql.NullString
	InviteBy                            sql.NullString
	CampaignName                        sql.NullString
	CampaignID                          sql.NullString
	RoleGain                            sql.NullString
	RoleLoss                            sql.NullString
	EventName                           sql.NullString
	EventCreated                        sql.NullBool
	EventLocation                       sql.NullString
	EventDescription                    sql.NullString
	EventTime                           sql.NullTime
	IsEventStart                        sql.NullBool
	IsEventEnd                          sql.NullBool
	ChannelID                           sql.NullString
	GuildID                             sql.NullString
	Position                            sql.NullInt32
	DiscordType                         sql.NullString
	System                              sql.NullBool
	Stickers                            pqtype.NullRawMessage
	WebhookID                           sql.NullString
	ApplicationID                       sql.NullString
	Activity                            pqtype.NullRawMessage
	Interaction                         pqtype.NullRawMessage
	MentionChannels                     pqtype.NullRawMessage
	GuildJoinedTimestamp                sql.NullInt64
	MemberPremiumSince                  sql.NullInt64
	MemberNickname                      sql.NullString
	MembershipScreeningPending          sql.NullBool
	CommunicationDisabledUntilTimestamp sql.NullInt64
	UserID                              sql.NullString
	MemberAvatarHash                    sql.NullString
	MemberDisplayName                   sql.NullString
	Roles                               sql.NullString
	MemberAvatarUrl                     sql.NullString
	MemberDisplayAvatarUrl              sql.NullString
	RoleIcon                            sql.NullString
	RoleUnicodeEmoji                    sql.NullString
	RoleID                              sql.NullString
	RoleColor                           sql.NullInt32
	RolePosition                        sql.NullInt32
	SeparatelyCategorized               sql.NullBool
	Mentionable                         sql.NullBool
	ManagedByIntegration                sql.NullBool
	RolePermissions                     sql.NullString
	MessageID                           sql.NullString
	UpdateContent                       sql.NullString
	ReactionMe                          sql.NullBool
	ReactionUsers                       sql.NullString
	ReactionCount                       sql.NullInt32
	ReactionEmojiID                     sql.NullString
	Bot                                 sql.NullBool
	Discriminator                       sql.NullString
	Tag                                 sql.NullString
	EmojiID                             sql.NullString
	EmojiName                           sql.NullString
	EmojiAnimated                       sql.NullBool
	EmojiRequiresColons                 sql.NullBool
	EmojiManaged                        sql.NullBool
	EmojiAvailable                      sql.NullBool
	EmojiCreatedTimestamp               sql.NullInt64
	EmojiUrl                            sql.NullString
	EmojiIdentifier                     sql.NullString
	StickerID                           sql.NullString
	StickerDescription                  sql.NullString
	StickerType                         sql.NullString
	StickerFormat                       sql.NullString
	StickerName                         sql.NullString
	StickerPackID                       sql.NullString
	StickerTags                         sql.NullString
	StickerAvailable                    sql.NullBool
	StickerSortValue                    sql.NullInt64
	ThreadID                            sql.NullString
	ThreadType                          sql.NullString
	ThreadMessages                      sql.NullString
	ThreadMembers                       sql.NullString
	ThreadFlags                         sql.NullInt64
	ThreadName                          sql.NullString
	ThreadParentID                      sql.NullString
	ThreadLocked                        sql.NullBool
	ThreadInvitable                     sql.NullBool
	ThreadArchived                      sql.NullBool
	ThreadAutoArchiveDuration           sql.NullInt64
	ThreadArchiveTimestamp              sql.NullInt64
	ThreadLastMessageID                 sql.NullString
	ThreadLastPinTimestamp              sql.NullInt64
	ThreadRateLimitPerUser              sql.NullInt64
	ThreadMessageCount                  sql.NullInt64
	ThreadMemberCount                   sql.NullInt64
	ThreadTotalMessageSent              sql.NullInt64
	ThreadAppliedTags                   sql.NullString
	ThreadCreatedTimestamp              sql.NullInt64
	ThreadIsNewlyCreated                sql.NullBool
	ChannelType                         sql.NullString
	ChannelParentID                     sql.NullString
	ChannelPermissionOverwrites         pqtype.NullRawMessage
	ChannelMessages                     pqtype.NullRawMessage
	ChannelThreads                      pqtype.NullRawMessage
	Nsfw                                sql.NullBool
	ChannelName                         sql.NullString
	ChannelRawPosition                  sql.NullInt32
	ChannelTopic                        sql.NullString
	ChannelLastMessageID                sql.NullString
	ChannelRateLimitPerUser             sql.NullInt32
	EventID                             sql.NullString
	EventStartTime                      sql.NullInt64
	EventEndTime                        sql.NullInt64
	EventPrivacyLevel                   sql.NullString
	EventStatus                         sql.NullString
	EventEntityType                     sql.NullString
	EventUserCount                      sql.NullInt32
	EventImage                          sql.NullString
	ServerDeaf                          sql.NullBool
	ServerMute                          sql.NullBool
	SelfDeaf                            sql.NullBool
	SelfMute                            sql.NullBool
	SelfVideo                           sql.NullBool
	VoiceSessionID                      sql.NullString
	VoiceStreaming                      sql.NullBool
	RequestToSpeakTimestamp             sql.NullInt64
	OldVoiceState                       pqtype.NullRawMessage
	StageInstanceID                     sql.NullString
	StageInstanceTopic                  sql.NullString
	DiscoveryDisabled                   sql.NullString
	InviteCode                          sql.NullString
	InviteTemporary                     sql.NullBool
	InviteMaxAge                        sql.NullInt32
	InviteMaxUses                       sql.NullInt32
	InviteTargetUser                    pqtype.NullRawMessage
	InviteTargetApplication             pqtype.NullRawMessage
	InviteTargetType                    sql.NullString
	InviteStageInstance                 pqtype.NullRawMessage
	InviteGuildScheduledEvent           sql.NullBool
	InviteExpiresTimestamp              sql.NullTime
	Reactions                           pqtype.NullRawMessage
	UserBannerHash                      sql.NullString
	UserBannerUrl                       sql.NullString
	UserAccentColor                     sql.NullInt32
	UserHexAccentColor                  sql.NullString
	UserDefaultAvatarUrl                sql.NullString
	UserVerified                        sql.NullBool
	UserMfaEnabled                      sql.NullBool
	UserCreatedTimestamp                sql.NullInt64
	PresenceStatus                      sql.NullString
	PresenceActivities                  pqtype.NullRawMessage
	PresenceDesktopStatus               sql.NullString
	PresenceMobileStatus                sql.NullString
	PresenceWebStatus                   sql.NullString
	ThreadMembersJoined                 sql.NullString
	ThreadMembersLeft                   sql.NullString
}
