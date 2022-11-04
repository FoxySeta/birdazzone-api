package twitter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const BearerToken = "AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"

func getRequest(query string) *string {
	max_results := "15"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.twitter.com/2/tweets/search/recent?query="+query+"&max_results="+max_results+"&tweet.fields=public_metrics&expansions=geo.place_id&place.fields=geo,country,country_code&user.fields=username", nil)
	req.Header.Set("Authorization", "Bearer "+BearerToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	//Convert the body to type string
	sb := string(body)
	//fmt.Println(sb)
	return &sb
}

func returnTweetQuery(c *gin.Context, res string) {
	var v interface{}
	json.Unmarshal([]byte(res), &v)
	data := v.(map[string]interface{})
	c.JSON(http.StatusOK, gin.H{
		"tweets": data,
	})
}

// TestTwitter godoc
// @Summary     Test Twitter API
// @Tags        test
// @Produce     json
// @Param       query	path	string	true	"Query to search"
// @Success     200
// @Router      /twitter/{query} [get]
func TestTwitter(ctx *gin.Context) {
	q := ctx.Param("query")
	sb := getRequest(q)
	if sb != nil {
		returnTweetQuery(ctx, *sb)
	}

}

//curl "https://api.twitter.com/2/tweets/search/recent?query=ciao" -H "Authorization: Bearer AAAAAAAAAAAAAAAAAAAAAE4higEAAAAAIAkazaLrT4LHjJx2XFPsdVzEPe8%3DE7HE9wBq5B5b0m4F8uGmcslObTmQFccb9gppULjUwTNBGj1Td3"
