package wechat

import "github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"

func MustNewWeChatClient(appid, secret, token, aesKey, notifyURL string) *officialAccount.OfficialAccount {
	app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:     appid,     // 公众号、小程序的appid
		Secret:    secret,    // 公众号的appsecret
		Token:     token,     // 公众号设置的token
		AESKey:    aesKey,    // 公众号设置的EncodingAESKey
		NotifyURL: notifyURL, // 公众号接收消息的URL

		Log: officialAccount.Log{
			Level:  "debug",
			Stdout: true,
		},

		HttpDebug: true,
		Debug:     false,
	})
	if err != nil {
		panic(err)
	}
	return app
}
