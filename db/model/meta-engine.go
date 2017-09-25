package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MetaSearch struct {
	Engines  []MetaSearchEngine  `gorm:"column:engines" json:"engines,omitempty" yaml:"engines,omitempty"`
	General  *MetaSearchGeneral  `gorm:"column:general" json:"general,omitempty" yaml:"general,omitempty"`
	Locales  *MetaSearchLocales  `gorm:"column:locales" json:"locales,omitempty" yaml:"locales,omitempty"`
	Outgoing *MetaSearchOutgoing `gorm:"column:outgoing" json:"outgoing,omitempty" yaml:"outgoing,omitempty"`
	Search   *MetaSearchSearch   `gorm:"column:search" json:"search,omitempty" yaml:"search,omitempty"`
	Server   *MetaSearchServer   `gorm:"column:server" json:"server,omitempty" yaml:"server,omitempty"`
	UI       *MetaSearchUI       `gorm:"column:ui" json:"ui,omitempty" yaml:"ui,omitempty"`
}

type MetaSearchEngine struct {
	gorm.Model      `json:"-" yaml:"-"`
	BaseURL         string        `gorm:"column:base_url" json:"base_url,omitempty" yaml:"base_url,omitempty"`
	Categories      string        `gorm:"column:categories" json:"categories,omitempty" yaml:"categories,omitempty"`
	ContentQuery    string        `gorm:"column:content_query" json:"content_query,omitempty" yaml:"content_query,omitempty"`
	ContentXpath    string        `gorm:"column:content_xpath" json:"content_xpath,omitempty" yaml:"content_xpath,omitempty"`
	Disabled        bool          `gorm:"column:disabled" json:"disabled,omitempty" yaml:"disabled,omitempty"`
	Engine          string        `gorm:"column:engine" json:"engine,omitempty" example:"archlinux" yaml:"engine,omitempty" example:"archlinux"`
	FirstPageNum    int           `gorm:"column:first_page_num" json:"first_page_num,omitempty" yaml:"first_page_num,omitempty"`
	Name            string        `gorm:"column:name" json:"name,omitempty" example:"arch linux wiki" yaml:"name,omitempty" example:"arch linux wiki"`
	NumberOfResults int           `gorm:"column:number_of_results" json:"number_of_results,omitempty" yaml:"number_of_results,omitempty"`
	PageSize        int           `gorm:"column:page_size" json:"page_size,omitempty" yaml:"page_size,omitempty"`
	Paging          bool          `gorm:"column:paging" json:"paging,omitempty" yaml:"paging,omitempty"`
	ResultsQuery    string        `gorm:"column:results_query" json:"results_query,omitempty" yaml:"results_query,omitempty"`
	ResultsXpath    string        `gorm:"column:results_xpath" json:"results_xpath,omitempty" yaml:"results_xpath,omitempty"`
	SearchType      string        `gorm:"column:search_type" json:"search_type,omitempty" yaml:"search_type,omitempty"`
	SearchURL       string        `gorm:"column:search_url" json:"search_url,omitempty" yaml:"search_url,omitempty"`
	Shortcut        string        `gorm:"column:shortcut" json:"shortcut,omitempty" example:"al" yaml:"shortcut,omitempty" example:"al"`
	SuggestionXpath string        `gorm:"column:suggestion_xpath" json:"suggestion_xpath,omitempty" yaml:"suggestion_xpath,omitempty"`
	Timeout         int           `gorm:"column:timeout" json:"timeout,omitempty" yaml:"timeout,omitempty"`
	TitleQuery      string        `gorm:"column:title_query" json:"title_query,omitempty" yaml:"title_query,omitempty"`
	TitleXpath      string        `gorm:"column:title_xpath" json:"title_xpath,omitempty" yaml:"title_xpath,omitempty"`
	URL             string        `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty"`
	URLQuery        string        `gorm:"column:url_query" json:"url_query,omitempty" yaml:"url_query,omitempty"`
	URLXpath        string        `gorm:"column:url_xpath" json:"url_xpath,omitempty" yaml:"url_xpath,omitempty"`
	Weight          int           `gorm:"column:weight" json:"weight,omitempty" yaml:"weight,omitempty"`
	Timeout         time.Duration `gorm:"column:timeout" json:"timeout" yaml:"timeout" default:"2"`
	ConcurrentCalls int           `gorm:"column:concurrent_calls" json:"concurrent_calls" yaml:"concurrent_calls" default:"10" `
}

type MetaSearchGeneral struct {
	gorm.Model   `json:"-" yaml:"-"`
	Debug        bool      `gorm:"column:debug" json:"debug,omitempty" example:"false" yaml:"debug,omitempty" example:"false"`
	InstanceName string    `gorm:"column:instance_name" json:"instance_name,omitempty" example:"searx" yaml:"instance_name,omitempty" example:"searx"`
	StartedAt    time.Time `gorm:"column:started_at" json:"started_at" yaml:"started_at"`
}

type MetaSearchLocales struct {
	gorm.Model `json:"-" yaml:"-"`
	Bg         string `gorm:"column:bg" json:"bg,omitempty" example:"Български (Bulgarian)" yaml:"bg,omitempty" example:"Български (Bulgarian)"`
	Cs         string `gorm:"column:cs" json:"cs,omitempty" example:"Čeština (Czech)" yaml:"cs,omitempty" example:"Čeština (Czech)"`
	De         string `gorm:"column:de" json:"de,omitempty" example:"Deutsch (German)" yaml:"de,omitempty" example:"Deutsch (German)"`
	DeDE       string `gorm:"column:de_" json:"de_DE,omitempty" example:"Deutsch (German_Germany)" yaml:"de_DE,omitempty" example:"Deutsch (German_Germany)"`
	ElGR       string `gorm:"column:el_" json:"el_GR,omitempty" example:"Ελληνικά (Greek_Greece)" yaml:"el_GR,omitempty" example:"Ελληνικά (Greek_Greece)"`
	En         string `gorm:"column:en" json:"en,omitempty" example:"English" yaml:"en,omitempty" example:"English"`
	Eo         string `gorm:"column:eo" json:"eo,omitempty" example:"Esperanto (Esperanto)" yaml:"eo,omitempty" example:"Esperanto (Esperanto)"`
	Es         string `gorm:"column:es" json:"es,omitempty" example:"Español (Spanish)" yaml:"es,omitempty" example:"Español (Spanish)"`
	Fi         string `gorm:"column:fi" json:"fi,omitempty" example:"Suomi (Finnish)" yaml:"fi,omitempty" example:"Suomi (Finnish)"`
	Fr         string `gorm:"column:fr" json:"fr,omitempty" example:"Français (French)" yaml:"fr,omitempty" example:"Français (French)"`
	He         string `gorm:"column:he" json:"he,omitempty" example:"עברית (Hebrew)" yaml:"he,omitempty" example:"עברית (Hebrew)"`
	Hu         string `gorm:"column:hu" json:"hu,omitempty" example:"Magyar (Hungarian)" yaml:"hu,omitempty" example:"Magyar (Hungarian)"`
	It         string `gorm:"column:it" json:"it,omitempty" example:"Italiano (Italian)" yaml:"it,omitempty" example:"Italiano (Italian)"`
	Ja         string `gorm:"column:ja" json:"ja,omitempty" example:"日本語 (Japanese)" yaml:"ja,omitempty" example:"日本語 (Japanese)"`
	Nl         string `gorm:"column:nl" json:"nl,omitempty" example:"Nederlands (Dutch)" yaml:"nl,omitempty" example:"Nederlands (Dutch)"`
	Pt         string `gorm:"column:pt" json:"pt,omitempty" example:"Português (Portuguese)" yaml:"pt,omitempty" example:"Português (Portuguese)"`
	PtBR       string `gorm:"column:pt_" json:"pt_BR,omitempty" example:"Português (Portuguese_Brazil)" yaml:"pt_BR,omitempty" example:"Português (Portuguese_Brazil)"`
	Ro         string `gorm:"column:ro" json:"ro,omitempty" example:"Română (Romanian)" yaml:"ro,omitempty" example:"Română (Romanian)"`
	Ru         string `gorm:"column:ru" json:"ru,omitempty" example:"Русский (Russian)" yaml:"ru,omitempty" example:"Русский (Russian)"`
	Sk         string `gorm:"column:sk" json:"sk,omitempty" example:"Slovenčina (Slovak)" yaml:"sk,omitempty" example:"Slovenčina (Slovak)"`
	Sv         string `gorm:"column:sv" json:"sv,omitempty" example:"Svenska (Swedish)" yaml:"sv,omitempty" example:"Svenska (Swedish)"`
	Tr         string `gorm:"column:tr" json:"tr,omitempty" example:"Türkçe (Turkish)" yaml:"tr,omitempty" example:"Türkçe (Turkish)"`
	Uk         string `gorm:"column:uk" json:"uk,omitempty" example:"українська мова (Ukrainian)" yaml:"uk,omitempty" example:"українська мова (Ukrainian)"`
	Zh         string `gorm:"column:zh" json:"zh,omitempty" example:"中文 (Chinese)" yaml:"zh,omitempty" example:"中文 (Chinese)"`
}

type MetaSearchOutgoing struct {
	gorm.Model      `json:"-" yaml:"-"`
	PoolConnections int    `gorm:"column:pool_connections" json:"pool_connections,omitempty" example:"100" yaml:"pool_connections,omitempty" example:"100"`
	PoolMaxsize     int    `gorm:"column:pool_maxsize" json:"pool_maxsize,omitempty" example:"10" yaml:"pool_maxsize,omitempty" example:"10"`
	RequestTimeout  int    `gorm:"column:request_timeout" json:"request_timeout,omitempty" example:"2" yaml:"request_timeout,omitempty" example:"2"`
	UseragentSuffix string `gorm:"column:useragent_suffix" json:"useragent_suffix,omitempty" yaml:"useragent_suffix,omitempty"`
}

type MetaSearchSearch struct {
	gorm.Model   `json:"-" yaml:"-"`
	Autocomplete string `gorm:"column:autocomplete" json:"autocomplete,omitempty" yaml:"autocomplete,omitempty"`
	Language     string `gorm:"column:language" json:"language,omitempty" example:"all" yaml:"language,omitempty" example:"all"`
	SafeSearch   int    `gorm:"column:safe_search" json:"safe_search,omitempty" example:"0" yaml:"safe_search,omitempty" example:"0"`
}

type MetaSearchServer struct {
	gorm.Model          `json:"-" yaml:"-"`
	BaseURL             bool   `gorm:"column:base_url" json:"base_url,omitempty" example:"false" yaml:"base_url,omitempty" example:"false"`
	BindAddress         string `gorm:"column:bind_address" json:"bind_address,omitempty" example:"127.0.0.1" yaml:"bind_address,omitempty" example:"127.0.0.1"`
	HTTPProtocolVersion string `gorm:"column:http_protocol_version" json:"http_protocol_version,omitempty" example:"1.0" yaml:"http_protocol_version,omitempty" example:"1.0"`
	ImageProxy          bool   `gorm:"column:image_proxy" json:"image_proxy,omitempty" example:"false" yaml:"image_proxy,omitempty" example:"false"`
	Port                int    `gorm:"column:port" json:"port,omitempty" example:"8888" yaml:"port,omitempty" example:"8888"`
	SecretKey           string `gorm:"column:secret_key" json:"secret_key,omitempty" example:"ultrasecretkey" yaml:"secret_key,omitempty" example:"ultrasecretkey"`
}

type MetaSearchUI struct {
	gorm.Model    `json:"-" yaml:"-"`
	DefaultLocale string `gorm:"column:default_locale" json:"default_locale,omitempty" yaml:"default_locale,omitempty"`
	DefaultTheme  string `gorm:"column:default_theme" json:"default_theme,omitempty" example:"oscar" yaml:"default_theme,omitempty" example:"oscar"`
	StaticPath    string `gorm:"column:static_path" json:"static_path,omitempty" yaml:"static_path,omitempty"`
	TemplatesPath string `gorm:"column:templates_path" json:"templates_path,omitempty" yaml:"templates_path,omitempty"`
}
