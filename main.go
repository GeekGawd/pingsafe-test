package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)


type GitHub []struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		Verification struct {
			Verified  bool   `json:"verified"`
			Reason    string `json:"reason"`
			Signature string `json:"signature"`
			Payload   string `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
}

var client = http.Client{
	Timeout: time.Second*5, // Timeout the requests after 5 second
}

type Github_Commit struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		Verification struct {
			Verified  bool   `json:"verified"`
			Reason    string `json:"reason"`
			Signature string `json:"signature"`
			Payload   string `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
	Stats struct {
		Total     int `json:"total"`
		Additions int `json:"additions"`
		Deletions int `json:"deletions"`
	} `json:"stats"`
	Files []struct {
		Sha         string `json:"sha"`
		Filename    string `json:"filename"`
		Status      string `json:"status"`
		Additions   int    `json:"additions"`
		Deletions   int    `json:"deletions"`
		Changes     int    `json:"changes"`
		BlobURL     string `json:"blob_url"`
		RawURL      string `json:"raw_url"`
		ContentsURL string `json:"contents_url"`
		Patch       string `json:"patch"`
	} `json:"files"`
}

func github_requests(uri string, is_patch bool)[]byte {
	fmt.Println(uri)
	req, err := http.NewRequest("GET", uri, nil)

	if err != nil{
		log.Fatal(err)
	}

	if(is_patch){
		req.Header.Set("Accept", "*/*")
		req.Header.Add("Authorization", "Bearer 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")
		req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	}else{
		req.Header.Set("Accept", "application/vnd.github.patch")
		req.Header.Add("Authorization", "Bearer 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")
		req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	}
	res, err := client.Do(req)

	if err != nil{
		log.Fatal(err)
	}

	if res.Body != nil{
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil{
		log.Fatal(readErr)
	}

	return body
}

func main(){
	uri := "https://api.github.com/repos/GeekGawd/pingsafe-test/commits"
	body := github_requests(uri, true)
	github_obj := GitHub{}
	unmarhsal_err := json.Unmarshal(body, &github_obj)

	var github_commits []string

	if unmarhsal_err != nil{
		log.Fatal(unmarhsal_err)
	}else{
		for _, objects := range github_obj{
			github_commits = append(github_commits, objects.URL)
		}
	}

	f, err := os.Create("patch.txt")

	if err != nil{
		log.Fatal(err)
	}

	defer f.Close()

	for i := range github_commits{
		body := string(github_requests(github_commits[i], false))
		_, err := f.WriteString(body)
		if err != nil{
			log.Fatal(err)
		}
	}
	fmt.Println("Patch File Created, It is ready to be scanned")
}