# teams-go
A (very simple) Go library for posting messages to incoming webhooks in Microsoft Teams.

For details about card formatting, see the [MessageCard reference](https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference). (even though it's described as 'legacy', that is actually the only format that the webhooks support) Note that not everything on that page is currently supported.

## Example
```go
package main

import "github.com/NoteToScreen/teams-go/teams"

func main() {
	url := "<insert webhook URL here>"
	card := teams.Card{
		Summary:    "This is a test.",
		ThemeColor: "0000FF",
		Title:      "Test",
		Sections: []teams.Section{
			teams.Section{
				ActivityTitle: "Important section",
				Text:          "This is a description of the section.",
				Facts: []teams.Fact{
					teams.Fact{
						Name:  "Is this a section?",
						Value: "Yes",
					},
					teams.Fact{
						Name:  "Contains text",
						Value: "Also yes",
					},
				},
			},
		},
	}

	err := teams.PostToWebhook(url, card)
	if err != nil {
		panic(err)
	}
}
```