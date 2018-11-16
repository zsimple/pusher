# Pusher

Pusher.js 后端的Golang实现，来自dirmiro1/ipe，继续开发满足公司内部需求。

# Why I wrote this software?

1. 项目中一开始用的云巴，但是它有点贵……

# How to configure?

## The server

```javascript
{
	"Host": ":8080",                    // Required
	"SSL": false,                       // Not Required, default is false
	"Profiling": false,                 // Mount pprof at /debug. Not Required, default is false
	"SSLHost": ":4433",                 // Required if SSL is true
	"SSLKeyFile": "A key.pem file",     // Required if SSL is true
	"SSLCertFile": "A cert.pem file",   // Required if SSL is true
	"Apps": [                           // Required, A Json arrays with multiple apps
		{
			"Disabled": false,                          // Required but can be false
			"Secret": "A really secret random string",  // Required
			"Key": "A random Key string",               // Required
			"OnlySSL": false,                           // Required but can be false
			"Name": "The app name",                     // Required
			"AppID": "The app ID",                      // Required
			"UserEvents": true,                         // Required but can be false
			"WebHooks": true,                           // Required but can be false
			"URLWebHook": "Some URL to send webhooks"   // Required if WebHooks is true
		}
	]
}

```

## Libraries

### Client javascript library

```javascript
var pusher = new Pusher(APP_KEY, {
  wsHost: 'localhost',
  wsPort: 8080,
  wssPort: 4433,    // Required if encrypted is true
  encrypted: false, // Optional. the application must use only SSL connections
  enabledTransports: ["ws", "flash"],
  disabledTransports: ["flash"]
});
```

### Client server libraries

Ruby

```ruby
Pusher.host = 'localhost'
Pusher.port = 8080
```

PHP

```php
$pusher = new Pusher(APP_KEY, APP_SECRET, APP_ID, DEBUG, "http://localhost", "8080");
```

NodeJS

```javascript
var pusher = new Pusher({
  appId: APP_ID,
  key: APP_KEY,
  secret: APP_SECRET
  domain: 'localhost',
  port: 80
});

```

# About Pusher

Pusher 是非常棒的服务，不过在国内不够迅捷稳定……