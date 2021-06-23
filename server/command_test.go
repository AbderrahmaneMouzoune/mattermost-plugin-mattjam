package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mattermost/mattermost-plugin-api/i18n"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCommandHelp(t *testing.T) {
	p := Plugin{
		configuration: &configuration{
			MattJamURL: "http://test",
		},
		botID: "test-bot-id",
	}
	apiMock := plugintest.API{}
	defer apiMock.AssertExpectations(t)
	apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user", Locale: "en"}, nil)
	apiMock.On("GetBundlePath").Return("..", nil)

	p.SetAPI(&apiMock)

	i18nBundle, err := i18n.InitBundle(p.API, filepath.Join("assets", "i18n"))
	require.Nil(t, err)
	p.b = i18nBundle

	helpText := strings.ReplaceAll(`###### Mattermost MattJam Plugin - Slash Command help
* |/mattjam| - Create a new meeting
* |/mattjam start [topic]| - Create a new meeting with specified topic
* |/mattjam help| - Show this help text
* |/mattjam settings see| - View your current user settings for the MattJam plugin
* |/mattjam settings [setting] [value]| - Update your user settings (see below for options)

###### MattJam Settings:
* |/mattjam settings embedded [true/false]|: (Experimental) When true, MattJam meeting is embedded as a floating window inside Mattermost. When false, MattJam meeting opens in a new window.
* |/mattjam settings naming_scheme [words/uuid/mattermost/ask]|: Select how meeting names are generated with one of these options:
    * |words|: Random English words in title case (e.g. PlayfulDragonsObserveCuriously)
    * |uuid|: UUID (universally unique identifier)
    * |mattermost|: Mattermost specific names. Combination of team name, channel name and random text in public and private channels; personal meeting name in direct and group messages channels.
    * |ask|: The plugin asks you to select the name every time you start a meeting`, "|", "`")

	apiMock.On("SendEphemeralPost", "test-user", &model.Post{
		UserId:    "test-bot-id",
		ChannelId: "test-channel",
		Message:   helpText,
	}).Return(nil)
	response, err := p.ExecuteCommand(&plugin.Context{}, &model.CommandArgs{UserId: "test-user", ChannelId: "test-channel", Command: "/mattjam help"})
	require.Equal(t, &model.CommandResponse{}, response)
	require.Nil(t, err)
}

func TestCommandSettings(t *testing.T) {
	p := Plugin{
		configuration: &configuration{
			MattJamURL:          "http://test",
			MattJamEmbedded:     false,
			MattJamNamingScheme: "mattermost",
		},
		botID: "test-bot-id",
	}

	tests := []struct {
		name      string
		command   string
		output    string
		newConfig *UserConfig
	}{
		{
			name:      "set valid setting with valid value",
			command:   "/mattjam settings embedded true",
			output:    "MattJam settings updated",
			newConfig: &UserConfig{Embedded: true, NamingScheme: "mattermost"},
		},
		{
			name:      "set valid setting with invalid value (embedded)",
			command:   "/mattjam settings embedded yes",
			output:    "Invalid `embedded` value, use `true` or `false`.",
			newConfig: nil,
		},
		{
			name:      "set valid setting with invalid value (naming_scheme)",
			command:   "/mattjam settings naming_scheme yes",
			output:    "Invalid `naming_scheme` value, use `ask`, `words`, `uuid` or `mattermost`.",
			newConfig: nil,
		},
		{
			name:      "set invalid setting",
			command:   "/mattjam settings other true",
			output:    "Invalid config field, use `embedded` or `naming_scheme`.",
			newConfig: nil,
		},
		{
			name:      "set invalid number of parameters",
			command:   "/mattjam settings other",
			output:    "Invalid settings parameters",
			newConfig: nil,
		},
		{
			name:      "get current user settings",
			command:   "/mattjam settings see",
			output:    "###### MattJam Settings:\n* Embedded: `false`\n* Naming Scheme: `mattermost`",
			newConfig: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := plugintest.API{}
			defer apiMock.AssertExpectations(t)
			p.SetAPI(&apiMock)

			apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user", Locale: "en"}, nil)
			apiMock.On("GetBundlePath").Return("..", nil)

			i18nBundle, err := i18n.InitBundle(p.API, filepath.Join("assets", "i18n"))
			require.Nil(t, err)
			p.b = i18nBundle

			apiMock.On("KVGet", "config_test-user", mock.Anything).Return(nil, nil)
			apiMock.On("SendEphemeralPost", "test-user", &model.Post{
				UserId:    "test-bot-id",
				ChannelId: "test-channel",
				Message:   tt.output,
			}).Return(nil)
			if tt.newConfig != nil {
				b, _ := json.Marshal(tt.newConfig)
				apiMock.On("KVSet", "config_test-user", b).Return(nil)
				var nilMap map[string]interface{}
				apiMock.On("PublishWebSocketEvent", configChangeEvent, nilMap, &model.WebsocketBroadcast{UserId: "test-user"})
			}
			response, err := p.ExecuteCommand(&plugin.Context{}, &model.CommandArgs{UserId: "test-user", ChannelId: "test-channel", Command: tt.command})
			require.Equal(t, &model.CommandResponse{}, response)
			require.Nil(t, err)
		})
	}
}

func TestCommandStartMeeting(t *testing.T) {
	p := Plugin{
		configuration: &configuration{
			MattJamURL: "http://test",
		},
	}

	t.Run("meeting without topic and ask configuration", func(t *testing.T) {
		apiMock := plugintest.API{}
		defer apiMock.AssertExpectations(t)
		p.SetAPI(&apiMock)

		apiMock.On("GetBundlePath").Return("..", nil)
		config := model.Config{}
		config.SetDefaults()
		apiMock.On("GetConfig").Return(&config, nil)

		i18nBundle, err := i18n.InitBundle(p.API, filepath.Join("assets", "i18n"))
		require.Nil(t, err)
		p.b = i18nBundle

		apiMock.On("SendEphemeralPost", "test-user", mock.MatchedBy(func(post *model.Post) bool {
			return post.Props["attachments"].([]*model.SlackAttachment)[0].Text == "Select type of meeting you want to start"
		})).Return(nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		apiMock.On("GetChannel", "test-channel").Return(&model.Channel{Id: "test-channel"}, nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		b, _ := json.Marshal(UserConfig{Embedded: false, NamingScheme: "ask"})
		apiMock.On("KVGet", "config_test-user", mock.Anything).Return(b, nil)

		response, err := p.ExecuteCommand(&plugin.Context{}, &model.CommandArgs{UserId: "test-user", ChannelId: "test-channel", Command: "/mattjam"})
		require.Equal(t, &model.CommandResponse{}, response)
		require.Nil(t, err)
	})

	t.Run("meeting without topic and no ask configuration", func(t *testing.T) {
		apiMock := plugintest.API{}
		defer apiMock.AssertExpectations(t)
		p.SetAPI(&apiMock)

		apiMock.On("GetBundlePath").Return("..", nil)
		config := model.Config{}
		config.SetDefaults()
		apiMock.On("GetConfig").Return(&config, nil)

		i18nBundle, err := i18n.InitBundle(p.API, filepath.Join("assets", "i18n"))
		require.Nil(t, err)
		p.b = i18nBundle

		apiMock.On("CreatePost", mock.MatchedBy(func(post *model.Post) bool {
			return strings.HasPrefix(post.Props["meeting_link"].(string), "http://test/")
		})).Return(&model.Post{}, nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		apiMock.On("GetChannel", "test-channel").Return(&model.Channel{Id: "test-channel"}, nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		apiMock.On("KVGet", "config_test-user", mock.Anything).Return(nil, nil)

		response, err := p.ExecuteCommand(&plugin.Context{}, &model.CommandArgs{UserId: "test-user", ChannelId: "test-channel", Command: "/mattjam"})
		require.Equal(t, &model.CommandResponse{}, response)
		require.Nil(t, err)
	})

	t.Run("meeting with topic", func(t *testing.T) {
		apiMock := plugintest.API{}
		defer apiMock.AssertExpectations(t)
		p.SetAPI(&apiMock)

		apiMock.On("GetBundlePath").Return("..", nil)
		config := model.Config{}
		config.SetDefaults()
		apiMock.On("GetConfig").Return(&config, nil)

		i18nBundle, err := i18n.InitBundle(p.API, filepath.Join("assets", "i18n"))
		require.Nil(t, err)
		p.b = i18nBundle

		apiMock.On("CreatePost", mock.MatchedBy(func(post *model.Post) bool {
			return strings.HasPrefix(post.Props["meeting_link"].(string), "http://test/topic")
		})).Return(&model.Post{}, nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		apiMock.On("GetChannel", "test-channel").Return(&model.Channel{Id: "test-channel"}, nil)
		apiMock.On("GetUser", "test-user").Return(&model.User{Id: "test-user"}, nil)
		apiMock.On("KVGet", "config_test-user", mock.Anything).Return(nil, nil)

		response, err := p.ExecuteCommand(&plugin.Context{}, &model.CommandArgs{UserId: "test-user", ChannelId: "test-channel", Command: "/mattjam start topic"})
		require.Equal(t, &model.CommandResponse{}, response)
		require.Nil(t, err)
	})
}
