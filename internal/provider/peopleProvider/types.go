package peopleProvider

type Provider struct {
}

type personPage struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	People     []person `json:"data"`
	Support    struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

type person struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Avatar    *string `json:"avatar"`
}

type personResponse struct {
	Person  person `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}
