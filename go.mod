module github.com/gamedb/gamedb

// These two lines are here because ahmdrz/goinsta needs to tag their repo after a fix
replace github.com/ahmdrz/goinsta/v2 => github.com/krylovsk/goinsta/v2 v2.4.0

require github.com/ahmdrz/goinsta/v2 v2.4.2

require (
	cloud.google.com/go v0.41.0
	github.com/Jleagle/go-durationfmt v0.0.0-20190307132420-e57bfad84057
	github.com/Jleagle/influxql v0.0.0-20190502115937-4ac053a1ed8e
	github.com/Jleagle/memcache-go v0.0.0-20190624135135-7b358ad3e388
	github.com/Jleagle/patreon-go v0.0.0-20190513114123-359f6ccef16d
	github.com/Jleagle/recaptcha-go v0.0.0-20190220085232-0e548dc7cc83
	github.com/Jleagle/session-go v0.0.0-20190515070633-3c8712426233
	github.com/Jleagle/sitemap-go v0.0.0-20190405195207-2bdddbb3bd50
	github.com/Jleagle/steam-go v0.0.0-20190620210405-b7520ae60b2f
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/antchfx/htmlquery v1.0.0 // indirect
	github.com/antchfx/xmlquery v1.0.0 // indirect
	github.com/antchfx/xpath v0.0.0-20190129040759-c8489ed3251e // indirect
	github.com/badoux/checkmail v0.0.0-20181210160741-9661bd69e9ad
	github.com/beefsack/go-rate v0.0.0-20180408011153-efa7637bb9b6 // indirect
	github.com/bwmarrin/discordgo v0.19.0
	github.com/cenkalti/backoff v2.2.0+incompatible
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/derekstavis/go-qs v0.0.0-20180720192143-9eef69e6c4e7
	github.com/dghubble/go-twitter v0.0.0-20190621134507-53f972dc4b06
	github.com/dghubble/oauth1 v0.5.0
	github.com/dghubble/sling v1.2.0 // indirect
	github.com/didip/tollbooth v4.0.0+incompatible
	github.com/djherbis/fscache v0.8.1
	github.com/dustin/go-humanize v1.0.0
	github.com/frustra/bbcode v0.0.0-20180807171629-48be21ce690c
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-github v26.0.9+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/sessions v1.1.3
	github.com/gorilla/websocket v1.4.0
	github.com/gosimple/slug v1.5.0
	github.com/influxdata/influxdb1-client v0.0.0-20190621134543-8ff2fc3824fc
	github.com/jinzhu/gorm v1.9.10
	github.com/jinzhu/now v1.0.1
	github.com/jzelinskie/geddit v0.0.0-20190613154237-4b3cacc98c40
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/leekchan/accounting v0.0.0-20190702062932-a09595581342
	github.com/logrusorgru/aurora v0.0.0-20190621143223-cea283e61946
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mxpv/patreon-go v0.0.0-20180807002359-67dbab1ad14c
	github.com/nicklaw5/helix v0.5.2
	github.com/nlopes/slack v0.5.0
	github.com/olekukonko/tablewriter v0.0.1 // indirect
	github.com/oschwald/maxminddb-golang v1.3.1
	github.com/pariz/gountries v0.0.0-20171019111738-adb00f6513a3
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/powerslacker/ratelimit v0.0.0-20190505003410-df2fcffc8e0d
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/robfig/cron v1.2.0
	github.com/rollbar/rollbar-go v1.1.0
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/sendgrid/rest v2.4.1+incompatible
	github.com/sendgrid/sendgrid-go v3.5.0+incompatible
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/tdewolff/minify v2.3.6+incompatible
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
	github.com/tdewolff/test v1.0.0 // indirect
	github.com/temoto/robotstxt v0.0.0-20180810133444-97ee4a9ee6ea // indirect
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	github.com/yohcop/openid-go v1.0.0
	go.mongodb.org/mongo-driver v1.0.3
	go.uber.org/atomic v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20190701105106-4def268fd1a4
	golang.org/x/net v0.0.0-20190628195232-da137c7871d7 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sys v0.0.0-20190626222941-04f50cda93cb // indirect
	google.golang.org/genproto v0.0.0-20190627211713-eb59cef1c072 // indirect
	google.golang.org/grpc v1.21.1
	gopkg.in/djherbis/atime.v1 v1.0.0 // indirect
	gopkg.in/djherbis/stream.v1 v1.2.0 // indirect
	jaytaylor.com/html2text v0.0.0-20190408195923-01ec452cbe43
)
