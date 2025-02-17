// +build integration

package integration

import (
	"context"
	"log"
	"os"

	cloudscale "github.com/cloudscale-ch/cloudscale-go-sdk"
	"golang.org/x/oauth2"
)

var (
	client *cloudscale.Client
)

// called when the package initializes
func init() {
	token := os.Getenv("CLOUDSCALE_TOKEN")

	if token == "" {
		log.Fatal("Missing CLOUDSCALE_TOKEN, tests won't run!\n")
	} else {
		tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		))
		client = cloudscale.NewClient(tc)
	}
}
