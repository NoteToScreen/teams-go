package teams

// A Fact is a key-value pair that can be used in a Section.
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A Section is a section in a Card.
type Section struct {
	ActivityTitle string `json:"activityTitle"`
	Text          string `json:"text"`
	Facts         []Fact `json:"facts"`
}

// A Card is a message that will be posted to Teams.
type Card struct {
	Type    string `json:"@type"`
	Context string `json:"@context"`

	Summary    string    `json:"summary"`
	ThemeColor string    `json:"themeColor"`
	Title      string    `json:"title"`
	Sections   []Section `json:"sections"`
}
