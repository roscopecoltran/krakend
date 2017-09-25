package ps

type Config struct {
	AdminUrl        string
	Email           string
	Password        string
	MaxItemsPerFile int
	Debug           bool
	SkipFirstRecord bool
}
