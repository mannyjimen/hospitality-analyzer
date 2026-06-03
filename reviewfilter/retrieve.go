package reviewfilter

type Review struct {
	business_id string
	text        string
}

var keywords = make(map[string]bool)
