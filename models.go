package gorinth

type Project struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Description string `json:"description"`
	Categories []string `json:"categories"`
	ClientSide string `json:"client_side"`
	ServerSide string `json:"server_side"`
	Body string `json:"body"`
	AdditionalCategories []string `json:"additional_categories"`
	IssuesUrl string `json:"issues_url"`
	SourceUrl string `json:"source_url"`
	WikiUrl string `json:"wiki_url"`
	DiscordUrl string `json:"discord_url"`
	DonationUrls []string `json:"donation_urls"`
	ProjectType string `json:"project_type"`
	Downloads int `json:"downloads"`
	IconUrl string `json:"icon_url"`
	Color int `json:"color"`
	Id string `json:"id"`
	Team string `json:"team"`
	ModeratorMessage map[string]any `json:"moderator_message"`
	Published string `json:"published"`
	Updated string `json:"updated"`
	Approved string `json:"approved"`
	Followers int `json:"followers"`
	Status string `json:"status"`
	License map[string]any `json:"license"`
	Versions []string `json:"versions"`
	GameVersions []string `json:"game_versions"`
	Loaders []string `json:"loaders"`
	Gallery []map[string]any `json:"gallery"`
}
