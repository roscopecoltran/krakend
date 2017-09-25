// this code is based on https://github.com/markbates/goth/blob/master/examples/main.go

package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/roscopecoltran/goth"
	"github.com/roscopecoltran/goth/gothic"
	"github.com/roscopecoltran/goth/providers/amazon"
	"github.com/roscopecoltran/goth/providers/auth0"
	"github.com/roscopecoltran/goth/providers/battlenet"
	"github.com/roscopecoltran/goth/providers/bitbucket"
	"github.com/roscopecoltran/goth/providers/box"
	"github.com/roscopecoltran/goth/providers/dailymotion"
	"github.com/roscopecoltran/goth/providers/deezer"
	"github.com/roscopecoltran/goth/providers/digitalocean"
	"github.com/roscopecoltran/goth/providers/discord"
	"github.com/roscopecoltran/goth/providers/dropbox"
	"github.com/roscopecoltran/goth/providers/eveonline"
	"github.com/roscopecoltran/goth/providers/facebook"
	"github.com/roscopecoltran/goth/providers/fitbit"
	"github.com/roscopecoltran/goth/providers/github"
	"github.com/roscopecoltran/goth/providers/gitlab"
	"github.com/roscopecoltran/goth/providers/gplus"
	"github.com/roscopecoltran/goth/providers/heroku"
	"github.com/roscopecoltran/goth/providers/instagram"
	"github.com/roscopecoltran/goth/providers/intercom"
	"github.com/roscopecoltran/goth/providers/lastfm"
	"github.com/roscopecoltran/goth/providers/linkedin"
	"github.com/roscopecoltran/goth/providers/meetup"
	"github.com/roscopecoltran/goth/providers/onedrive"
	"github.com/roscopecoltran/goth/providers/openidConnect"
	"github.com/roscopecoltran/goth/providers/paypal"
	"github.com/roscopecoltran/goth/providers/salesforce"
	"github.com/roscopecoltran/goth/providers/slack"
	"github.com/roscopecoltran/goth/providers/soundcloud"
	"github.com/roscopecoltran/goth/providers/spotify"
	"github.com/roscopecoltran/goth/providers/steam"
	"github.com/roscopecoltran/goth/providers/stripe"
	"github.com/roscopecoltran/goth/providers/twitch"
	"github.com/roscopecoltran/goth/providers/twitter"
	"github.com/roscopecoltran/goth/providers/uber"
	"github.com/roscopecoltran/goth/providers/wepay"
	"github.com/roscopecoltran/goth/providers/xero"
	"github.com/roscopecoltran/goth/providers/yahoo"
	"github.com/roscopecoltran/goth/providers/yammer"
)

var (
	store             sessions.CookieStore
	sessionSecret     string
	GothicSessionName = gothic.SessionName
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {

	// Generate a psuedo-random session secret for testing.
	if sessionSecret == "" {
		sessionSecret = func() string {
			b := make([]byte, 32)
			for i := range b {
				b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
			}
			return string(b)
		}()
	}

	store = sessions.NewCookieStore([]byte(sessionSecret))

}

func main() {

	// _ = godotenv.Load("filenumberone.env", "filenumbertwo.env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// const BaseURL = "http://localhost:3000"
	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),

		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		fitbit.New(os.Getenv("FITBIT_KEY"), os.Getenv("FITBIT_SECRET"), "http://localhost:3000/auth/fitbit/callback"),
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3000/auth/gplus/callback"),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		spotify.New(os.Getenv("SPOTIFY_KEY"), os.Getenv("SPOTIFY_SECRET"), "http://localhost:3000/auth/spotify/callback"),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
		lastfm.New(os.Getenv("LASTFM_KEY"), os.Getenv("LASTFM_SECRET"), "http://localhost:3000/auth/lastfm/callback"),
		twitch.New(os.Getenv("TWITCH_KEY"), os.Getenv("TWITCH_SECRET"), "http://localhost:3000/auth/twitch/callback"),
		dropbox.New(os.Getenv("DROPBOX_KEY"), os.Getenv("DROPBOX_SECRET"), "http://localhost:3000/auth/dropbox/callback"),
		digitalocean.New(os.Getenv("DIGITALOCEAN_KEY"), os.Getenv("DIGITALOCEAN_SECRET"), "http://localhost:3000/auth/digitalocean/callback", "read"),
		bitbucket.New(os.Getenv("BITBUCKET_KEY"), os.Getenv("BITBUCKET_SECRET"), "http://localhost:3000/auth/bitbucket/callback"),
		instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), "http://localhost:3000/auth/instagram/callback"),
		intercom.New(os.Getenv("INTERCOM_KEY"), os.Getenv("INTERCOM_SECRET"), "http://localhost:3000/auth/intercom/callback"),
		box.New(os.Getenv("BOX_KEY"), os.Getenv("BOX_SECRET"), "http://localhost:3000/auth/box/callback"),
		salesforce.New(os.Getenv("SALESFORCE_KEY"), os.Getenv("SALESFORCE_SECRET"), "http://localhost:3000/auth/salesforce/callback"),
		amazon.New(os.Getenv("AMAZON_KEY"), os.Getenv("AMAZON_SECRET"), "http://localhost:3000/auth/amazon/callback"),
		yammer.New(os.Getenv("YAMMER_KEY"), os.Getenv("YAMMER_SECRET"), "http://localhost:3000/auth/yammer/callback"),
		onedrive.New(os.Getenv("ONEDRIVE_KEY"), os.Getenv("ONEDRIVE_SECRET"), "http://localhost:3000/auth/onedrive/callback"),
		battlenet.New(os.Getenv("BATTLENET_KEY"), os.Getenv("BATTLENET_SECRET"), "http://localhost:3000/auth/battlenet/callback"),
		eveonline.New(os.Getenv("EVEONLINE_KEY"), os.Getenv("EVEONLINE_SECRET"), "http://localhost:3000/auth/eveonline/callback"),

		//Pointed localhost.com to http://localhost:3000/auth/yahoo/callback through proxy as yahoo
		// does not allow to put custom ports in redirection uri
		yahoo.New(os.Getenv("YAHOO_KEY"), os.Getenv("YAHOO_SECRET"), "http://localhost.com"),
		slack.New(os.Getenv("SLACK_KEY"), os.Getenv("SLACK_SECRET"), "http://localhost:3000/auth/slack/callback"),
		stripe.New(os.Getenv("STRIPE_KEY"), os.Getenv("STRIPE_SECRET"), "http://localhost:3000/auth/stripe/callback"),
		wepay.New(os.Getenv("WEPAY_KEY"), os.Getenv("WEPAY_SECRET"), "http://localhost:3000/auth/wepay/callback", "view_user"),
		//By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		//in sandbox environment
		paypal.New(os.Getenv("PAYPAL_KEY"), os.Getenv("PAYPAL_SECRET"), "http://localhost:3000/auth/paypal/callback"),
		steam.New(os.Getenv("STEAM_KEY"), "http://localhost:3000/auth/steam/callback"),
		heroku.New(os.Getenv("HEROKU_KEY"), os.Getenv("HEROKU_SECRET"), "http://localhost:3000/auth/heroku/callback"),
		uber.New(os.Getenv("UBER_KEY"), os.Getenv("UBER_SECRET"), "http://localhost:3000/auth/uber/callback"),
		soundcloud.New(os.Getenv("SOUNDCLOUD_KEY"), os.Getenv("SOUNDCLOUD_SECRET"), "http://localhost:3000/auth/soundcloud/callback"),
		gitlab.New(os.Getenv("GITLAB_KEY"), os.Getenv("GITLAB_SECRET"), "http://localhost:3000/auth/gitlab/callback"),
		dailymotion.New(os.Getenv("DAILYMOTION_KEY"), os.Getenv("DAILYMOTION_SECRET"), "http://localhost:3000/auth/dailymotion/callback", "email"),
		deezer.New(os.Getenv("DEEZER_KEY"), os.Getenv("DEEZER_SECRET"), "http://localhost:3000/auth/deezer/callback", "email"),
		discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(os.Getenv("MEETUP_KEY"), os.Getenv("MEETUP_SECRET"), "http://localhost:3000/auth/meetup/callback"),

		//Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(os.Getenv("AUTH0_KEY"), os.Getenv("AUTH0_SECRET"), "http://localhost:3000/auth/auth0/callback", os.Getenv("AUTH0_DOMAIN")),
		xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/xero/callback"),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize it self in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := make(map[string]string)
	m["amazon"] = "Amazon"
	m["bitbucket"] = "Bitbucket"
	m["box"] = "Box"
	m["dailymotion"] = "Dailymotion"
	m["deezer"] = "Deezer"
	m["digitalocean"] = "Digital Ocean"
	m["discord"] = "Discord"
	m["dropbox"] = "Dropbox"
	m["eveonline"] = "Eve Online"
	m["facebook"] = "Facebook"
	m["fitbit"] = "Fitbit"
	m["github"] = "Github"
	m["gitlab"] = "Gitlab"
	m["soundcloud"] = "SoundCloud"
	m["spotify"] = "Spotify"
	m["steam"] = "Steam"
	m["stripe"] = "Stripe"
	m["twitch"] = "Twitch"
	m["uber"] = "Uber"
	m["wepay"] = "Wepay"
	m["yahoo"] = "Yahoo"
	m["yammer"] = "Yammer"
	m["gplus"] = "Google Plus"
	m["heroku"] = "Heroku"
	m["instagram"] = "Instagram"
	m["intercom"] = "Intercom"
	m["lastfm"] = "Last FM"
	m["linkedin"] = "Linkedin"
	m["onedrive"] = "Onedrive"
	m["battlenet"] = "Battlenet"
	m["paypal"] = "Paypal"
	m["twitter"] = "Twitter"
	m["salesforce"] = "Salesforce"
	m["slack"] = "Slack"
	m["meetup"] = "Meetup.com"
	m["auth0"] = "Auth0"
	m["openid-connect"] = "OpenID Connect"
	m["xero"] = "Xero"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	// fmt.Printf("providerIndex:\n %#v", providerIndex)

	tpl := template.New("")
	template.Must(tpl.New("index.html").Parse(indexTemplate))
	template.Must(tpl.New("user.html").Parse(userTemplate))

	r := gin.Default()
	r.SetHTMLTemplate(tpl)

	r.Use(sessions.Sessions("session", store))

	r.GET("/auth/:provider", authBeginHandler)
	r.GET("/auth/:provider/callback", completeUserAuthHandler)
	r.GET("/logout/:provider", LogoutHandler)

	r.GET("/", func(c *gin.Context) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(c.Writer, providerIndex)
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Fatal(r.Run(":3000"))
}

func AuthProvider(c *gin.Context) {
	fn := gothic.GetProviderName
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		provider := c.Params.ByName("provider")
		if provider == "" {
			return fn(req)
		}
		return provider, nil
	}
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(c.Writer, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func getState(req *http.Request) string {
	state := req.URL.Query().Get("state")
	if len(state) > 0 {
		return state
	}
	return "state"
}

func authBeginHandler(c *gin.Context) {

	providerName := c.Params.ByName("provider")
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return
	}

	sess, err := provider.BeginAuth(getState(c.Request))
	if err != nil {
		return
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return
	}

	session := sessions.Default(c)
	session.Set(GothicSessionName, sess.Marshal())
	session.Save()

	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

func completeUserAuthHandler(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("completeUserAuthHandler error: %v", err)
			http.Redirect(c.Writer, c.Request, fmt.Sprintf("/#/?error=%v", err), http.StatusTemporaryRedirect)
			return
		}
	}()

	session := sessions.Default(c)
	providerName := c.Params.ByName("provider")
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		panic(err)
	}

	v := session.Get(GothicSessionName)
	if v == nil {
		panic(errors.New("Unable to get session"))
	}

	sess, err := provider.UnmarshalSession(v.(string))
	if err != nil {
		panic(err)
	}

	_, err = sess.Authorize(provider, c.Request.URL.Query())
	if err != nil {
		panic(err)
	}

	// Fetch the user.
	gothUser, err := provider.FetchUser(sess)
	if err != nil {
		panic(err)
	}

	session.Set("github.access_token", gothUser.AccessToken)
	session.Set("github.user_id", gothUser.UserID)
	session.Save()

	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(c.Writer, gothUser)
	// c.HTML(http.StatusOK, "user.html", nil)
	return
}

func LogoutHandler(c *gin.Context) {

	fn := gothic.GetProviderName
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		provider := c.Params.ByName("provider")
		if provider == "" {
			return fn(req)
		}
		return provider, nil
	}

	session := sessions.Default(c)
	providerName := c.Params.ByName("provider")
	provider, err := goth.GetProvider(providerName)
	if err != nil {
		panic(err)
	}

	v := session.Get(GothicSessionName)
	if v == nil {
		panic(errors.New("Unable to get session"))
	}

	sess, err := provider.UnmarshalSession(v.(string))
	if err != nil {
		panic(err)
	}

	fmt.Printf("GothicSessionName: %#s\n%#v\n", GothicSessionName, v)
	fmt.Printf("session: %#v\n", session)
	fmt.Printf("sess: %#s\n", sess)

	session.Delete("github.user_id")
	session.Delete("github.access_token")
	session.Delete(v.(string))
	session.Save()

	err = gothic.Logout(c.Writer, c.Request)
	if err != nil {
		fmt.Printf("error while trying to logout: %#v", err)
	}

	next := extractNextPath(c.Request.URL.Query().Get("next"))
	http.Redirect(c.Writer, c.Request, next, http.StatusFound)
}

func extractNextPath(next string) string {
	n, err := url.Parse(next)
	if err != nil {
		return "/"
	}
	path := n.Path
	if path == "" {
		path = "/"
	}
	return path
}

/*
func writeJSON(c *gin.Context, data interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	if data != nil {
		enc := json.NewEncoder(c.Writer)
		if err := enc.Encode(data); err != nil {
			return err
		}
	}
	return nil
}
*/

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
