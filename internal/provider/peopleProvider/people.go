package peopleProvider

import (
	"net/url"
	"strconv"

	"github.com/EspenS93/filterview/internal/database/mysql"
	"github.com/EspenS93/filterview/internal/provider/providerInterface"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func New() providerInterface.Provider {
	return &Provider{}
}

func (service *Provider) GetPage(c *gin.Context) templ.Component {

	if c.GetHeader("HX-Request") == "true" {
		return getList(c)
	} else {
		return filterViewPage(filter(getCheckedFilterValues(c)), getList(c))
	}
}

func getPeople(c *gin.Context) (personPage, error) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	/*url := UsersURL + "?page=" + pageStr

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	*/
	var responsePersonPage personPage

	people, pageNum, pageSize, totalPages, totalPeople, err := mysql.New().GetPeople(page)
	for _, i := range people {
		// Convert int to string
		p := person{
			ID:        i.ID,
			Email:     i.Email,
			FirstName: i.FirstName,
			LastName:  i.LastName,
			Avatar:    i.Avatar,
		}
		// Append string to slice
		responsePersonPage.People = append(responsePersonPage.People, p)
	}
	responsePersonPage.Page = pageNum
	responsePersonPage.Total = totalPeople
	responsePersonPage.TotalPages = totalPages
	responsePersonPage.PerPage = pageSize
	/*if err := json.Unmarshal(body, &responsePersonPage); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}*/

	return responsePersonPage, err
}

func getList(c *gin.Context) templ.Component {
	people, _ := getPeople(c)

	return list(people)
}

/*func getFilter(c *gin.Context) templ.Component {
	return filter(getCheckedFilterValues(c))
}*/

func getCheckedFilterValues(c *gin.Context) (map[string]bool, map[string]bool) {
	queryValues := c.Request.URL.Query()

	colors := getQueryMap(queryValues, "color")
	shapes := getQueryMap(queryValues, "shape")

	// Default to square if no shapes are selected
	if len(shapes) == 0 {
		shapes = map[string]bool{
			"square": true,
		}
	}
	return colors, shapes
}

func getQueryMap(queryValues url.Values, key string) map[string]bool {

	// Iterate through the query parameters
	queryMap := make(map[string]bool)

	// Retrieve the entire query map

	// Assuming key is the only query parameter we're interested in
	// Each value of key becomes a key in our map
	if queries, exists := queryValues[key]; exists {
		for _, query := range queries {
			queryMap[query] = true
		}
	}
	return queryMap
}
