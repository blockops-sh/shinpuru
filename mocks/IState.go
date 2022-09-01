// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	discordgo "github.com/bwmarrin/discordgo"
	dgrs "github.com/zekrotja/dgrs"

	mock "github.com/stretchr/testify/mock"
)

// IState is an autogenerated mock type for the IState type
type IState struct {
	mock.Mock
}

// Channel provides a mock function with given fields: id
func (_m *IState) Channel(id string) (*discordgo.Channel, error) {
	ret := _m.Called(id)

	var r0 *discordgo.Channel
	if rf, ok := ret.Get(0).(func(string) *discordgo.Channel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Channels provides a mock function with given fields: guildID, forceFetch
func (_m *IState) Channels(guildID string, forceFetch ...bool) ([]*discordgo.Channel, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Channel
	if rf, ok := ret.Get(0).(func(string, ...bool) []*discordgo.Channel); ok {
		r0 = rf(guildID, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(guildID, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Emoji provides a mock function with given fields: guildID, emojiID
func (_m *IState) Emoji(guildID string, emojiID string) (*discordgo.Emoji, error) {
	ret := _m.Called(guildID, emojiID)

	var r0 *discordgo.Emoji
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Emoji); ok {
		r0 = rf(guildID, emojiID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, emojiID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Emojis provides a mock function with given fields: guildID, forceFetch
func (_m *IState) Emojis(guildID string, forceFetch ...bool) ([]*discordgo.Emoji, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Emoji
	if rf, ok := ret.Get(0).(func(string, ...bool) []*discordgo.Emoji); ok {
		r0 = rf(guildID, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Emoji)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(guildID, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Flush provides a mock function with given fields: subKeys
func (_m *IState) Flush(subKeys ...string) error {
	_va := make([]interface{}, len(subKeys))
	for _i := range subKeys {
		_va[_i] = subKeys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(subKeys...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Guild provides a mock function with given fields: id, hydrate
func (_m *IState) Guild(id string, hydrate ...bool) (*discordgo.Guild, error) {
	_va := make([]interface{}, len(hydrate))
	for _i := range hydrate {
		_va[_i] = hydrate[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *discordgo.Guild
	if rf, ok := ret.Get(0).(func(string, ...bool) *discordgo.Guild); ok {
		r0 = rf(id, hydrate...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Guild)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(id, hydrate...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Guilds provides a mock function with given fields:
func (_m *IState) Guilds() ([]*discordgo.Guild, error) {
	ret := _m.Called()

	var r0 []*discordgo.Guild
	if rf, ok := ret.Get(0).(func() []*discordgo.Guild); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Guild)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Member provides a mock function with given fields: guildID, memberID, forceNoFetch
func (_m *IState) Member(guildID string, memberID string, forceNoFetch ...bool) (*discordgo.Member, error) {
	_va := make([]interface{}, len(forceNoFetch))
	for _i := range forceNoFetch {
		_va[_i] = forceNoFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID, memberID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *discordgo.Member
	if rf, ok := ret.Get(0).(func(string, string, ...bool) *discordgo.Member); ok {
		r0 = rf(guildID, memberID, forceNoFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...bool) error); ok {
		r1 = rf(guildID, memberID, forceNoFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Members provides a mock function with given fields: guildID, forceFetch
func (_m *IState) Members(guildID string, forceFetch ...bool) ([]*discordgo.Member, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Member
	if rf, ok := ret.Get(0).(func(string, ...bool) []*discordgo.Member); ok {
		r0 = rf(guildID, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(guildID, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MembersLimit provides a mock function with given fields: guildID, afterID, limit, forceFetch
func (_m *IState) MembersLimit(guildID string, afterID string, limit int, forceFetch ...bool) ([]*discordgo.Member, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID, afterID, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Member
	if rf, ok := ret.Get(0).(func(string, string, int, ...bool) []*discordgo.Member); ok {
		r0 = rf(guildID, afterID, limit, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Member)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, int, ...bool) error); ok {
		r1 = rf(guildID, afterID, limit, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Message provides a mock function with given fields: channelID, messageID
func (_m *IState) Message(channelID string, messageID string) (*discordgo.Message, error) {
	ret := _m.Called(channelID, messageID)

	var r0 *discordgo.Message
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Message); ok {
		r0 = rf(channelID, messageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channelID, messageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Messages provides a mock function with given fields: channelID, forceFetch
func (_m *IState) Messages(channelID string, forceFetch ...bool) ([]*discordgo.Message, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, channelID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Message
	if rf, ok := ret.Get(0).(func(string, ...bool) []*discordgo.Message); ok {
		r0 = rf(channelID, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(channelID, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Presence provides a mock function with given fields: guildID, userID
func (_m *IState) Presence(guildID string, userID string) (*discordgo.Presence, error) {
	ret := _m.Called(guildID, userID)

	var r0 *discordgo.Presence
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Presence); ok {
		r0 = rf(guildID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Presence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Presences provides a mock function with given fields: guildID
func (_m *IState) Presences(guildID string) ([]*discordgo.Presence, error) {
	ret := _m.Called(guildID)

	var r0 []*discordgo.Presence
	if rf, ok := ret.Get(0).(func(string) []*discordgo.Presence); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Presence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Publish provides a mock function with given fields: channel, payload
func (_m *IState) Publish(channel string, payload interface{}) error {
	ret := _m.Called(channel, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(channel, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReleaseShard provides a mock function with given fields: pool, id
func (_m *IState) ReleaseShard(pool int, id int) error {
	ret := _m.Called(pool, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(pool, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveChannel provides a mock function with given fields: id, dehydrate
func (_m *IState) RemoveChannel(id string, dehydrate ...bool) error {
	_va := make([]interface{}, len(dehydrate))
	for _i := range dehydrate {
		_va[_i] = dehydrate[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...bool) error); ok {
		r0 = rf(id, dehydrate...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveEmoji provides a mock function with given fields: guildID, emojiID
func (_m *IState) RemoveEmoji(guildID string, emojiID string) error {
	ret := _m.Called(guildID, emojiID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, emojiID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveGuild provides a mock function with given fields: id, dehydrate
func (_m *IState) RemoveGuild(id string, dehydrate ...bool) error {
	_va := make([]interface{}, len(dehydrate))
	for _i := range dehydrate {
		_va[_i] = dehydrate[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...bool) error); ok {
		r0 = rf(id, dehydrate...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveMember provides a mock function with given fields: guildID, memberID
func (_m *IState) RemoveMember(guildID string, memberID string) error {
	ret := _m.Called(guildID, memberID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, memberID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveMessage provides a mock function with given fields: channelID, messageID
func (_m *IState) RemoveMessage(channelID string, messageID string) error {
	ret := _m.Called(channelID, messageID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channelID, messageID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePresence provides a mock function with given fields: guildID, userID
func (_m *IState) RemovePresence(guildID string, userID string) error {
	ret := _m.Called(guildID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveRole provides a mock function with given fields: guildID, roleID
func (_m *IState) RemoveRole(guildID string, roleID string) error {
	ret := _m.Called(guildID, roleID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, roleID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUser provides a mock function with given fields: id
func (_m *IState) RemoveUser(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveVoiceState provides a mock function with given fields: guildID, userID
func (_m *IState) RemoveVoiceState(guildID string, userID string) error {
	ret := _m.Called(guildID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReserveShard provides a mock function with given fields: pool, cid
func (_m *IState) ReserveShard(pool int, cid ...int) (int, error) {
	_va := make([]interface{}, len(cid))
	for _i := range cid {
		_va[_i] = cid[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pool)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, ...int) int); ok {
		r0 = rf(pool, cid...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, ...int) error); ok {
		r1 = rf(pool, cid...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Role provides a mock function with given fields: guildID, roleID
func (_m *IState) Role(guildID string, roleID string) (*discordgo.Role, error) {
	ret := _m.Called(guildID, roleID)

	var r0 *discordgo.Role
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.Role); ok {
		r0 = rf(guildID, roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, roleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Roles provides a mock function with given fields: guildID, forceFetch
func (_m *IState) Roles(guildID string, forceFetch ...bool) ([]*discordgo.Role, error) {
	_va := make([]interface{}, len(forceFetch))
	for _i := range forceFetch {
		_va[_i] = forceFetch[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, guildID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*discordgo.Role
	if rf, ok := ret.Get(0).(func(string, ...bool) []*discordgo.Role); ok {
		r0 = rf(guildID, forceFetch...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...bool) error); ok {
		r1 = rf(guildID, forceFetch...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelfUser provides a mock function with given fields:
func (_m *IState) SelfUser() (*discordgo.User, error) {
	ret := _m.Called()

	var r0 *discordgo.User
	if rf, ok := ret.Get(0).(func() *discordgo.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetChannel provides a mock function with given fields: channel
func (_m *IState) SetChannel(channel *discordgo.Channel) error {
	ret := _m.Called(channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.Channel) error); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetEmoji provides a mock function with given fields: guildID, emoji
func (_m *IState) SetEmoji(guildID string, emoji *discordgo.Emoji) error {
	ret := _m.Called(guildID, emoji)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *discordgo.Emoji) error); ok {
		r0 = rf(guildID, emoji)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGuild provides a mock function with given fields: guild
func (_m *IState) SetGuild(guild *discordgo.Guild) error {
	ret := _m.Called(guild)

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.Guild) error); ok {
		r0 = rf(guild)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMember provides a mock function with given fields: guildID, member
func (_m *IState) SetMember(guildID string, member *discordgo.Member) error {
	ret := _m.Called(guildID, member)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *discordgo.Member) error); ok {
		r0 = rf(guildID, member)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMessage provides a mock function with given fields: message
func (_m *IState) SetMessage(message *discordgo.Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPresence provides a mock function with given fields: guildID, presence
func (_m *IState) SetPresence(guildID string, presence *discordgo.Presence) error {
	ret := _m.Called(guildID, presence)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *discordgo.Presence) error); ok {
		r0 = rf(guildID, presence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRole provides a mock function with given fields: guildID, role
func (_m *IState) SetRole(guildID string, role *discordgo.Role) error {
	ret := _m.Called(guildID, role)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *discordgo.Role) error); ok {
		r0 = rf(guildID, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSelfUser provides a mock function with given fields: user
func (_m *IState) SetSelfUser(user *discordgo.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUser provides a mock function with given fields: user
func (_m *IState) SetUser(user *discordgo.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetVoiceState provides a mock function with given fields: guildID, vs
func (_m *IState) SetVoiceState(guildID string, vs *discordgo.VoiceState) error {
	ret := _m.Called(guildID, vs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *discordgo.VoiceState) error); ok {
		r0 = rf(guildID, vs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shards provides a mock function with given fields: pool
func (_m *IState) Shards(pool int) ([]*dgrs.Shard, error) {
	ret := _m.Called(pool)

	var r0 []*dgrs.Shard
	if rf, ok := ret.Get(0).(func(int) []*dgrs.Shard); ok {
		r0 = rf(pool)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dgrs.Shard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(pool)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: channel, handler
func (_m *IState) Subscribe(channel string, handler func(func(interface{}) error)) func() error {
	ret := _m.Called(channel, handler)

	var r0 func() error
	if rf, ok := ret.Get(0).(func(string, func(func(interface{}) error)) func() error); ok {
		r0 = rf(channel, handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// SubscribeDMs provides a mock function with given fields: handler
func (_m *IState) SubscribeDMs(handler func(*dgrs.DirectMessageEvent)) func() error {
	ret := _m.Called(handler)

	var r0 func() error
	if rf, ok := ret.Get(0).(func(func(*dgrs.DirectMessageEvent)) func() error); ok {
		r0 = rf(handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

// User provides a mock function with given fields: id
func (_m *IState) User(id string) (*discordgo.User, error) {
	ret := _m.Called(id)

	var r0 *discordgo.User
	if rf, ok := ret.Get(0).(func(string) *discordgo.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGuilds provides a mock function with given fields: id
func (_m *IState) UserGuilds(id string) ([]string, error) {
	ret := _m.Called(id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Users provides a mock function with given fields:
func (_m *IState) Users() ([]*discordgo.User, error) {
	ret := _m.Called()

	var r0 []*discordgo.User
	if rf, ok := ret.Get(0).(func() []*discordgo.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoiceState provides a mock function with given fields: guildID, userID
func (_m *IState) VoiceState(guildID string, userID string) (*discordgo.VoiceState, error) {
	ret := _m.Called(guildID, userID)

	var r0 *discordgo.VoiceState
	if rf, ok := ret.Get(0).(func(string, string) *discordgo.VoiceState); ok {
		r0 = rf(guildID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.VoiceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoiceStates provides a mock function with given fields: guildID
func (_m *IState) VoiceStates(guildID string) ([]*discordgo.VoiceState, error) {
	ret := _m.Called(guildID)

	var r0 []*discordgo.VoiceState
	if rf, ok := ret.Get(0).(func(string) []*discordgo.VoiceState); ok {
		r0 = rf(guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.VoiceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIState interface {
	mock.TestingT
	Cleanup(func())
}

// NewIState creates a new instance of IState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIState(t mockConstructorTestingTNewIState) *IState {
	mock := &IState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
