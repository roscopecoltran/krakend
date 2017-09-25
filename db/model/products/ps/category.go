package ps

type Category struct {
	ID             string `validate:"nonzero"`
	Active         string
	Name           string `validate:"nonzero"`
	ParentCategory string
	Rootcategory   string //(0/1)
	Description    string `validate:"nonzero"`
	Metatitle      string
	Metakeywords   string
	Meta           string
}
