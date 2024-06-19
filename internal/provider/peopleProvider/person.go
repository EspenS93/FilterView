package peopleProvider

import (
	"fmt"

	"github.com/EspenS93/filterview/internal/database/mysql"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func (service *Provider) GetDetail(c *gin.Context) templ.Component {
	detailPerson, err := getPerson(c)

	if err != nil {
		fmt.Println("No response from request")
	}

	if c.GetHeader("HX-Request") == "true" {
		return detail(detailPerson)
	} else {
		return detailPage(detailPerson)
	}
}

func getPerson(c *gin.Context) (person, error) {
	id := c.Param("id")
	/*url := UsersURL + "/" + id

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	*/
	var responsePerson personResponse
	person, err := mysql.New().GetPerson(id)

	responsePerson.Person.Email = person.Email
	responsePerson.Person.FirstName = person.FirstName
	responsePerson.Person.LastName = person.LastName
	responsePerson.Person.ID = person.ID

	/*if err := json.Unmarshal(body, &responsePerson); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}*/

	return responsePerson.Person, err
}
