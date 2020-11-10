package newsfeed

type Item struct {
	Title string `json:"Title"`
	Post  string `json:"Post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
