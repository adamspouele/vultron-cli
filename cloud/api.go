package cloud

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	pat = "a672e9dc039db44186f3b9e1bd4a2ac0f4c44844a037b22d2e51770cc5164dc7"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func createDroplet(name string) {
	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: "ams3",
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-14-04-x64",
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		fmt.Printf("error")
	}
}
