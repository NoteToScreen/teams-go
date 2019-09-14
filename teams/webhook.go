package teams

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PostToWebhook posts the given Card to the Teams webhook with the given URL.
func PostToWebhook(url string, card Card) error {
	return PostToWebhookWithClient(http.DefaultClient, url, card)
}

// PostToWebhookWithClient uses the given *http.Client to post the given Card to the Teams webhook with the given URL.
func PostToWebhookWithClient(client *http.Client, url string, card Card) error {
	card.Type = "@MessageCard"
	card.Context = "http://schema.org/extensions"

	messageJSON, err := json.Marshal(card)
	if err != nil {
		return err
	}
	log.Println(string(messageJSON))

	resp, err := client.Post(url, "text/json", strings.NewReader(string(messageJSON)))
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println(string(data))

	defer resp.Body.Close()
	return nil
}
