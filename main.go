package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/searchconsole/v1"
	"os"
)

func main() {
	ctx := context.Background()

	// 1
	//searchconsoleService, err := searchconsole.NewService(ctx, option.WithAPIKey("---"))

	// 2
	//config := &oauth2.Config{
	//	ClientID:     "---",
	//	ClientSecret: "---",
	//	Scopes: []string{
	//		searchconsole.WebmastersScope,
	//		searchconsole.WebmastersReadonlyScope,
	//	},
	//	RedirectURL: "https://vendasta-hackathon.firebaseapp.com/__/auth/handler",
	//	Endpoint: oauth2.Endpoint{
	//		AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
	//		TokenURL: "https://www.googleapis.com/oauth2/v4/token",
	//	},
	//}
	//
	//authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	//token, err := config.Exchange(ctx, authURL)
	//searchconsoleService, err := searchconsole.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	// 3 -
	// server-go@vendasta-hackathon.iam.gserviceaccount.com
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	searchconsoleService, err := searchconsole.NewService(ctx, option.WithCredentialsFile(credentials))

	if err != nil {
		fmt.Printf(err.Error())
	}

	siteUrl := "https://www.geleiatotal.com.br/"

	//getSitemaps(searchconsoleService, siteUrl)
	//getSites(searchconsoleService)
	getSearchanalytics(searchconsoleService, siteUrl)
}


func getSitemaps(searchconsoleService *searchconsole.Service, siteUrl string) {
	response, err := searchconsoleService.Sitemaps.List(siteUrl).Do()
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("sitemaps[%#v] \n", response)

	for _, sitemap := range response.Sitemap {
		fmt.Printf("sitemap[%#v] \n", sitemap)
	}
}

func getSites(searchconsoleService *searchconsole.Service) {
	response, err := searchconsoleService.Sites.List().Do()
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("sites[%#v] \n", response)

	for _, site := range response.SiteEntry {
		fmt.Printf("site[%#v] \n", site)
	}
}

func getSearchanalytics(searchconsoleService *searchconsole.Service, siteUrl string) {

	query := &searchconsole.SearchAnalyticsQueryRequest{
		Dimensions: []string{"QUERY"},
		StartDate: "2020-09-29", // Search 1 day before - OBSERVATION
		EndDate: "2020-10-05",
		RowLimit: 5,
		SearchType: "WEB",
	}

	response, err := searchconsoleService.Searchanalytics.Query(siteUrl, query).Do()
	if err != nil {
		fmt.Printf(err.Error())
	}

	//fmt.Printf("searchanalytics[%#v] \n", response)

	for _, row := range response.Rows {
		fmt.Printf("row[%#v] \n", row)
	}
}