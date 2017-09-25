package config

import (
	"github.com/jinzhu/gorm"
)

type LocalesConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Bg         string `gorm:"column:bg" json:"bg" example:"Български (Bulgarian)" yaml:"bg" toml:"bg"`
	Cs         string `gorm:"column:cs" json:"cs" example:"Čeština (Czech)" yaml:"cs" toml:"cs"`
	De         string `gorm:"column:de" json:"de" example:"Deutsch (German)" yaml:"de" toml:"de"`
	DeDE       string `gorm:"column:de_" json:"de_DE" example:"Deutsch (German_Germany)" yaml:"de_DE" toml:"de_DE"`
	ElGR       string `gorm:"column:el_" json:"el_GR" example:"Ελληνικά (Greek_Greece)" yaml:"el_GR" toml:"el_GR"`
	En         string `gorm:"column:en" json:"en" example:"English" yaml:"en" toml:"en"`
	Eo         string `gorm:"column:eo" json:"eo" example:"Esperanto (Esperanto)" yaml:"eo" toml:"eo"`
	Es         string `gorm:"column:es" json:"es" example:"Español (Spanish)" yaml:"es" toml:"es"`
	Fi         string `gorm:"column:fi" json:"fi" example:"Suomi (Finnish)" yaml:"fi" toml:"fi"`
	Fr         string `gorm:"column:fr" json:"fr" example:"Français (French)" yaml:"fr" toml:"fr"`
	He         string `gorm:"column:he" json:"he" example:"עברית (Hebrew)" yaml:"he" toml:"he"`
	Hu         string `gorm:"column:hu" json:"hu" example:"Magyar (Hungarian)" yaml:"hu" toml:"hu"`
	It         string `gorm:"column:it" json:"it" example:"Italiano (Italian)" yaml:"it" toml:"it"`
	Ja         string `gorm:"column:ja" json:"ja" example:"日本語 (Japanese)" yaml:"ja" toml:"ja"`
	Nl         string `gorm:"column:nl" json:"nl" example:"Nederlands (Dutch)" yaml:"nl" toml:"nl"`
	Pt         string `gorm:"column:pt" json:"pt" example:"Português (Portuguese)" yaml:"pt" toml:"pt"`
	PtBR       string `gorm:"column:pt_" json:"pt_BR" example:"Português (Portuguese_Brazil)" yaml:"pt_BR" toml:"pt_BR"`
	Ro         string `gorm:"column:ro" json:"ro" example:"Română (Romanian)" yaml:"ro" toml:"ro"`
	Ru         string `gorm:"column:ru" json:"ru" example:"Русский (Russian)" yaml:"ru" toml:"ru"`
	Sk         string `gorm:"column:sk" json:"sk" example:"Slovenčina (Slovak)" yaml:"sk" toml:"sk"`
	Sv         string `gorm:"column:sv" json:"sv" example:"Svenska (Swedish)" yaml:"sv" toml:"sv"`
	Tr         string `gorm:"column:tr" json:"tr" example:"Türkçe (Turkish)" yaml:"tr" toml:"tr"`
	Uk         string `gorm:"column:uk" json:"uk" example:"українська мова (Ukrainian)" yaml:"uk" toml:"uk"`
	Zh         string `gorm:"column:zh" json:"zh" example:"中文 (Chinese)" yaml:"zh" toml:"zh"`
}

type LocalesDetailsConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Name       string             `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" example:"Español" toml:"name,omitempty" example:"Español"`
	English    string             `gorm:"column:english" json:"english" yaml:"english" example:"Spanish" toml:"english" example:"Spanish"`
	Slug       string             `gorm:"column:slug" json:"slug" yaml:"slug" slug:"Spanish" toml:"slug" slug:"Spanish"`
	Code       string             `gorm:"column:code" json:"code,omitempty" yaml:"code,omitempty" example:"en_EN" toml:"code,omitempty" example:"en_EN"`
	Lang       string             `gorm:"column:lang" json:"lang" yaml:"lang" example:"en" toml:"lang" example:"en"`
	Currency   string             `gorm:"column:currency" json:"currency" yaml:"currency" example:"USD" toml:"currency" example:"USD"`
	Country    CountryConfig      `gorm:"column:country" json:"country" yaml:"country" example:"United States" toml:"country" example:"United States"`
	Extra      LocalesExtraConfig `gorm:"column:extra" json:"extra" yaml:"extra" toml:"extra"`
}

type LocalesExtraConfig struct {
	gorm.Model          `json:"-" yaml:"-" toml:"-"`
	CardinalPluralRules string  `gorm:"column:cardinal_plural_rules" json:"cardinal_plural_rules" yaml:"cardinal_plural_rules" toml:"cardinal_plural_rules"`
	PluralRules         string  `gorm:"column:plural_rules" json:"plural_rules" yaml:"plural_rules" toml:"plural_rules"`
	Percent             float64 `gorm:"column:percent" json:"percent" yaml:"percent" toml:"percent"`
	OrdinalPluralRules  string  `gorm:"column:ordinal_plural_rules" json:"ordinal_plural_rules" yaml:"ordinal_plural_rules" toml:"ordinal_plural_rules"`
	RangePluralRules    string  `gorm:"column:range_plural_rules" json:"range_plural_rules" yaml:"range_plural_rules" toml:"range_plural_rules"`
}

type LocalesTimeConfig struct {
	gorm.Model         `json:"-" yaml:"-" toml:"-"`
	Time               string `gorm:"column:time" json:"time" yaml:"time" toml:"time"`
	Date               string `gorm:"column:date" json:"date" yaml:"date" toml:"date"`
	MonthWide          string `gorm:"column:month_wide" json:"months_wide" yaml:"months_wide" toml:"months_wide"`
	MonthAbbreviated   string `gorm:"column:month_abbreviated" json:"months_abbreviated" yaml:"months_abbreviated" toml:"months_abbreviated"`
	MonthNarrow        string `gorm:"column:month_narrow" json:"months_narrow" yaml:"months_narrow" toml:"months_narrow"`
	WeekdayWide        string `gorm:"column:weekly_wide" json:"weekly_wide" yaml:"weekly_wide" toml:"weekly_wide"`
	WeekdayAbbreviated string `gorm:"column:weekly_abbreviated" json:"weekly_abbreviated" yaml:"weekly_abbreviated" toml:"weekly_abbreviated"`
	WeekdayNarrow      string `gorm:"column:weekly_narrow" json:"weekly_narrow" yaml:"weekly_narrow" toml:"weekly_narrow"`
	WeekdayShort       string `gorm:"column:weekly_short" json:"weekly_short" yaml:"weekly_short" toml:"weekly_short"`
}

type CountryConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Name       string `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" example:"España" toml:"name,omitempty" example:"España"`
	English    string `gorm:"column:english" json:"english" yaml:"english" example:"Spain" toml:"english" example:"Spain"`
	Slug       string `gorm:"column:slug" json:"slug" yaml:"slug" example:"espana" toml:"slug" example:"espana"`
	Iso2       string `gorm:"column:iso2" json:"iso2,omitempty" yaml:"iso2,omitempty" example:"US" toml:"iso2,omitempty" example:"US"`
	Iso3       string `gorm:"column:iso3" json:"iso3" yaml:"iso3" example:"USA" toml:"iso3" example:"USA"`
}

type CitiesConfig struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Name       string `gorm:"column:name" json:"name" yaml:"name" example:"Lyon" toml:"name" example:"Lyon"`
	English    string `gorm:"column:english" json:"english" yaml:"english" example:"Lyon" toml:"english" example:"Lyon"`
	Slug       string `gorm:"column:slug" json:"slug" yaml:"slug" example:"lyon" toml:"slug" example:"lyon"`
}
