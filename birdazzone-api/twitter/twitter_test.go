package twitter

import (
	"testing"
	"time"

	"git.hjkl.gq/team13/birdazzone-api/util"
)

func TestToCompleteUrlOnEmptyString(t *testing.T) {
	if toCompleteUrl("") != BaseUrl {
		t.FailNow()
	}
}

func TestToCompleteUrl(t *testing.T) {
	if toCompleteUrl("partialUrl") != BaseUrl+"partialUrl" {
		t.FailNow()
	}
}

func TestGetRequestOnNilPathParams(t *testing.T) {
	util.GetTestingGinEngine()
	_, err := getRequest("/swagger/index.html", nil, util.Pair[string, string]{First: "a", Second: "b"})
	if err == nil {
		t.FailNow()
	}
}

func TestGetRequestOnWrongPath(t *testing.T) {
	util.GetTestingGinEngine()
	response, err := getRequest("/wrongPage.html", nil)
	if response != nil {
		t.Fatalf("Non-null response")
	}
	if err == nil {
		t.Fatalf("Non-null error")
	}
}

func TestGetUserExisting(t *testing.T) {
	_, err := GetUser("quizzettone")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetUserNonExisting(t *testing.T) {
	_, err := GetUser("")
	if err == nil {
		t.Fatal("Expected error in GetUser")
	}
}

func testTweets(t *testing.T, response *ProfileTweets, err error) {
	if err != nil {
		t.Fatal(err.Error())
	}
	if response == nil {
		t.Fatal("Null response")
	}
	tweets := response.Data
	if tweets == nil {
		t.Fatal("Null reponse.Data")
	}
	if len(tweets) == 0 {
		t.Fatal("No tweets found")
	}
}

func TestGetTweets(t *testing.T) {
	response, err := getTweets(
		"https://api.twitter.com/2/tweets/search/recent",
		[]any{},
		util.Pair[string, string]{First: "query", Second: "hello"},
		util.Pair[string, string]{First: "start_time", Second: util.LastInstantAtGivenTime(time.Now(), 18)},
		util.Pair[string, string]{First: "tweet.fields", Second: "author_id,created_at,public_metrics,text"},
		util.Pair[string, string]{First: "expansions", Second: "author_id"},
		util.Pair[string, string]{First: "user.fields", Second: "id,name,profile_image_url,username"},
	)
	testTweets(t, response, err)
}

func TestGetTweetsFromUser(t *testing.T) {
	response, err := GetTweetsFromUser("1499992669480755204", 10, "2020-11-11T22:00:00Z")
	testTweets(t, response, err)
}

func TestGetTweetsFromHashtag(t *testing.T) {
	response, err := GetManyRecentTweetsFromQuery("hello", util.LastInstantAtGivenTime(time.Now(), 18), "")
	testTweets(t, response, err)
}
