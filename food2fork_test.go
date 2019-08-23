package food2fork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/DATA-DOG/godog"
)

const defaultContentType = "application/json"

type StepContext struct {
	searchResponse SearchResponse
}

type SearchResponse struct {
	Count   int `json:"count"`
	Recipes []struct {
		Publisher    string  `json:"publisher"`
		F2FURL       string  `json:"f2f_url"`
		Title        string  `json:"title"`
		SourceURL    string  `json:"source_url"`
		RecipeID     string  `json:"recipe_id"`
		ImageURL     string  `json:"image_url"`
		SocialRank   float64 `json:"social_rank"`
		PublisherURL string  `json:"publisher_url"`
	} `json:"recipes"`
}

func (ctx *StepContext) iSearchServiceRunning() error {
	//The service could have a status API
	return nil
}

//Scenario: verify a Recipe is returned
func (ctx *StepContext) iRequestRecipe(arg1 string) error {
	url := "https://www.food2fork.com/api/search"
	query := strings.Replace(arg1, " ", "%20", -1)
	searchString := "?key=e676ea4152b2077f7e7bef634e232fff&q=" + query
	resp, err := http.Get(url + searchString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error response. %s\n", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Status Code. ", resp.StatusCode)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, &ctx.searchResponse)

	return nil
}

func (ctx *StepContext) iRecipeReturned() error {
	if ctx.searchResponse.Count == 0 {
		return fmt.Errorf("No Recipe found")
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	ctx := &StepContext{}
	//Scenario: verify a random advice is returned
	s.Step(`^the food to fork service is running$`, ctx.iSearchServiceRunning)
	s.Step(`^I request a recipe about "([^"]*)"$`, ctx.iRequestRecipe)
	s.Step(`^the recipe is returned$`, ctx.iRecipeReturned)
}
