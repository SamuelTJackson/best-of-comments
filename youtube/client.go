package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)
type Client struct {
	config *oauth2.Config
	service *youtube.Service
}

func getOauthConfig() (*oauth2.Config, error)  {
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
}

func GetClient() (*Client, error){
	config, err := getOauthConfig()
	if err != nil {
		return nil, err
	}
	cacheFile, err := tokenCacheFile()
	if err != nil {
		return nil, err
	}
	token, err := tokenFromFile(cacheFile)
	if err != nil {
		token, err = getTokenFromWeb(config)
		if err != nil {
			return nil, err
		}
		if err := saveToken(cacheFile, token); err != nil {
			return nil, err
		}
	}
	srv, err := youtube.NewService(context.Background(), option.WithTokenSource(config.TokenSource(context.TODO(),token)))
	if err != nil {
		return nil, err
	}
	return &Client{ service: srv}, nil
}
// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}
// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}
func saveToken(file string, token *oauth2.Token) error{
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}