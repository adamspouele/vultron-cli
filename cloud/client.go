package cloud

import (
	"context"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

const (
	pat = "a672e9dc039db44186f3b9e1bd4a2ac0f4c44844a037b22d2e51770cc5164dc7"
)

// TokenSource the access token variable
type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// GetDoClient get digitalOcean client and context
func GetDoClient() (*godo.Client, context.Context, error) {
	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	return client, ctx, nil
}
