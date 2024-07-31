package gorinth

// Represents a Modrinth project. Pointer fields are optional
type Project struct {
	// The slug of the project, used for vanity URLs
	Slug string `json:"slug"`
	// The title or name of the project
	Title string `json:"title"`
	// A short description of the project
	Description string `json:"description"`
	// A list of categories that the project is listed under
	Categories []string `json:"categories"`
	// The client side support of the project
	ClientSide Support `json:"client_side"`
	// The server side support of the project
	ServerSide Support `json:"server_side"`
	// The long form description of the project
	Body string `json:"body"`
	// The status of the project
	Status ProjectStatus `json:"status"`
	// The requested status when submitting for review or scheduling the project for release
	RequestedStatus *RequestedProjectStatus `json:"requested_status"`
	// A list of categories which are searchable but non-primary
	AdditionalCategories []string `json:"additional_categories"`
	// An optional link to where to submit bugs or issues with the project
	IssuesUrl *string `json:"issues_url"`
	// An optional link to the source code of the project
	SourceUrl *string `json:"source_url"`
	// An optional link to the project's wiki page or other relevant information
	WikiUrl *string `json:"wiki_url"`
	// An optional invite link to the project's discord
	DiscordUrl *string `json:"discord_url"`
	// A list of donation links for the project
	DonationUrls []string `json:"donation_urls"`
	// The type of the project
	ProjectType ProjectType `json:"project_type"`
	// The total number of downloads of the project
	Downloads int `json:"downloads"`
	// The URL of the project's icon
	IconUrl *string `json:"icon_url"`
	// The RGB colour of the project, automatically generated from the project icon
	Color *int `json:"color"`
	// The ID of the project, encoded as a base62 string
	Id string `json:"id"`
	// The ID of the team that has ownership of this project
	Team string `json:"team"`
	// The date the project was publised (ISO-8601 format)
	Published string `json:"published"`
	// The date the project was last updated (ISO-8601 format)
	Updated string `json:"updated"`
	// The date the project's status was set to an approved status (ISO-8601 format)
	Approved *string `json:"approved"`
	// The date the project's status was submitted to moderators for review (ISO-8601 format)
	Queued *string `json:"queued"`
	// The total number of users following the project
	Followers int `json:"followers"`
	// The license of the project
	License *License `json:"license"`
	// A list of the version IDs of the project
	Versions []string `json:"versions"`
	// A list of all of the game versions supported by the project
	GameVersions []string `json:"game_versions"`
	// A list of all of the loaders supported by the project
	Loaders []string `json:"loaders"`
	// A list of images that have been uploaded to the project's gallery
	Gallery []GalleryImage `json:"gallery"`
	// Separate from the project in the Modrinth API,
	// stores the image data for the project's icon
	Icon []byte
	auth string
}

// Represents one version of a project. Pointer fields are optional
type Version struct {
	// The name of this version
	Name string `json:"name"`
	// The version number. Ideally will follow semantic versioning
	VersionNumber string `json:"version_number"`
	// The changelog for this version
	Changelog *string `json:"changelog"`
	// A list of specific versions of projects that this version depends on
	Dependencies []Dependency `json:"dependencies"`
	// A list of versions of Minecraft that this version supports
	GameVersions []string `json:"game_versions"`
	// The release type foor this version
	VersionType ReleaseType `json:"version_type"`
	// The mod loaders that this version supports
	Loaders []string `json:"loaders"`
	// Whether this version is featured
	Featured bool `json:"featured"`
	// The status of this version
	Status VersionStatus `json:"status"`
	// The requested status of this version
	RequestedStatus RequestedVersionStatus `json:"requested_status"`
	// The ID of this versions, encoded as a base-62 string
	Id string `json:"id"`
	// The ID of the project this version belongs to
	ProjectId string `json:"project_id"`
	// The ID of the author if this version
	AuthorId string `json:"author_id"`
	// The date that this version was published (ISO-8601 format)
	DatePublished string `json:"date_published"`
	// The number of downloads of this version
	Downloads int `json:"downloads"`
	// The files for this version
	Files     []VersionFile `json:"files"`
	FileParts []string      `json:"file_parts"`
}

// Represents a Modrinth user. Pointer fields are optional
type User struct {
	// The user's username
	Username string `json:"username"`
	// The user's display name
	Name *string `json:"name"`
	// The user's email. Only present if requesting your own account
	Email *string `json:"email"`
	// A description of the user
	Bio string `json:"bio"`
	// Data regarding the user's payout status. Only present if requesting your own account
	PayoutData PayoutData `json:"payout_data"`
	// The user's ID
	Id string `json:"id"`
	// The URL to the user's avatar
	AvatarUrl string `json:"avatar_url"`
	// The time at which the user was created (ISO-8601 format)
	Created string `json:"created"`
	// The user's role
	Role UserRole `json:"role"`
	// A user's badges. Currently unused
	Badges Badges `json:"badges"`
	// A list of authentication providers the user has signed up for.
	// Only present if requesting your own account
	AuthProviders []string `json:"auth_providers"`
	// Whether the user's email is verified.
	// Only present if requesting your own account
	EmailVerified bool `json:"email_verified"`
	// Whether the user has a password associated with their account.
	// Only present if requesting your own account
	HasPassword bool `json:"has_password"`
	// Whether the user TOTP two-factor authentication enabled.
	// Only present if requesting your own account
	HasTOTP bool `json:"has_totp"`
	auth    string
}
