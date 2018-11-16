// Copyright 2014, 2016 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package pusher

// The config file
type configFile struct {
	Host        string // The host, eg: :8080 will start on 0.0.0.0:8080
	User        string
	SSL         bool
	Profiling   bool
	SSLHost     string
	SSLKeyFile  string
	SSLCertFile string

	Apps []configApp
}

type configApp struct {
	Name       string
	AppID      string
	Key        string
	Secret     string
	OnlySSL    bool
	Disabled   bool
	UserEvents bool
	WebHooks   bool
	URLWebHook string
}

func newAppFromConfig(a configApp) *app {
	return newApp(
		a.Name,
		a.AppID,
		a.Key,
		a.Secret,
		a.OnlySSL,
		a.Disabled,
		a.UserEvents,
		a.WebHooks,
		a.URLWebHook,
	)
}
