package gorinth

// The client or server side support of a project
type Support string

const (
	SupportRequired    Support = "required"
	SupportOptional    Support = "optiional"
	SupportUnsupported Support = "unsupported"
)

// The status of a project. One of:
// "approved", "archived", "rejected", "draft", "unlisted",
// "processing", "withheld", "scheduled", "private", "unknown"
type ProjectStatus string

const (
	ProjectApproved   ProjectStatus = "approved"
	ProjectArchived   ProjectStatus = "archived"
	ProjectRejected   ProjectStatus = "rejected"
	ProjectDraft      ProjectStatus = "draft"
	ProjectUnlisted   ProjectStatus = "unlisted"
	ProjectProcessing ProjectStatus = "processing"
	ProjectWithheld   ProjectStatus = "withheld"
	ProjectScheduled  ProjectStatus = "scheduled"
	ProjectPrivate    ProjectStatus = "private"
	ProjectUnknown    ProjectStatus = "unknown"
)

// The requested status when submitting for review or scheduling the project for release.
// One of: "approved", "archived", "unlisted", "private", "draft"
type RequestedProjectStatus string

const (
	RequestedProjectApproved RequestedProjectStatus = "approved"
	RequestedProjectArchived RequestedProjectStatus = "archived"
	RequestedProjectUnlisted RequestedProjectStatus = "unlisted"
	RequestedProjectPrivate  RequestedProjectStatus = "private"
	RequestedProjectDraft    RequestedProjectStatus = "draft"
)

// The type of a project. One of: "mod", "modpack", "resourcepack", "shader"
type ProjectType string

const (
	// A mod project. Also includes plugins and datapacks
	ProjectMod ProjectType = "mod"
	// A modpack project
	ProjectModpack ProjectType = "modpack"
	// A resourcepack project
	ProjectResourcepack ProjectType = "resourcepack"
	// A shader project
	ProjectShader ProjectType = "shader"
)

// The type of a dependency. One of: "required"`, "optional", "incompatible", "embedded"
type DependencyType string

const (
	// A required dependency.
	// This project or version is required for the current version to function
	DependencyRequired DependencyType = "required"
	// An optional dependency.
	// This project or version can be added to the current version to enchance or modify it
	DependencyOptional DependencyType = "optional"
	// An incompatible dependency.
	// This project or version is incompatible with the current version,
	// and the two cannot be used together
	DependencyIncompatible DependencyType = "incompatible"
	// An embedded dependency, used for modpacks.
	// The files of this version are contained within the current version,
	// and do not need to be downloaded separately
	DependencyEmbedded DependencyType = "embedded"
)

// The release type for a version. One of: "release", "beta"
type ReleaseType string

const (
	// A full release of a project
	ReleaseRelease ReleaseType = "release"
	// A beta version of a project
	ReleaseBeta ReleaseType = "beta"
	// An alpha version of a project
	ReleaseAlpha ReleaseType = "alpha"
)

// The status of a version. One of:
// "listed", "archived", "draft", "unlisted", "scheduled", "unknown"
type VersionStatus string

const (
	VersionListed    VersionStatus = "listed"
	VersionArchived  VersionStatus = "archived"
	VersionDraft     VersionStatus = "draft"
	VersionUnlisted  VersionStatus = "unlisted"
	VersionScheduled VersionStatus = "scheduled"
	VersionUnkown    VersionStatus = "unknown"
)

// The requested status of a version. One of:
// "listed", "archived", "draft", "unlisted"
type RequestedVersionStatus string

const (
	RequestedVersionListed    VersionStatus = "listed"
	RequestedVersionArchived  VersionStatus = "archived"
	RequestedVersionDraft     VersionStatus = "draft"
	RequestedVersionUnlisted  VersionStatus = "unlisted"
	RequestedVersionScheduled VersionStatus = "scheduled"
	RequestedVersionUnkown    VersionStatus = "unknown"
)

// The type of an additional file, used mainly for adding resource packs to datapacks
// One of: "required-resource-pack", "optional-resource-pack"
type FileType string

const (
	FileRequiredRP FileType = "required-resource-pack"
	FileOptionalRP FileType = "optional-resource-pack"
)

// The wallet that a user has selected. One of: "paypal", "venmo"
type PayoutWallet string

const (
	WalletPaypal PayoutWallet = "paypal"
	WalletVenmo PayoutWallet = "venmo"
)

// The type of a user's wallet. One of: "email", "phone", "user-handle"
type PayoutWalletType string

const (
	WalletTypeEmail PayoutWalletType = "email"
	WalletTypePhone PayoutWalletType = "phone"
	WalletTypeUserHandle PayoutWalletType = "user-handle"
)

// A user's role. One of: "admin", "moderator", "developer"
type UserRole string

const (
	UserAdmin UserRole = "admin"
	UserModerator UserRole = "moderator"
	UserDeveloper UserRole = "developer"
)

// A user's badges, as a bitfield. Currently unused
type Badges uint8

const (
	_ Badges = 1 << iota
	BadgeEarlyModpackAdopter
	BadgeEarlyRespackAdopter
	BadgeEarlyPluginAdopter
	BadgeAlphaTester
	BadgeContributor
	BadgeTranslator
)

// Returns whether the badges bitfield has the BadgeEarlyModpackAdopter bit set
func (b Badges) EarlyModpackAdopter() bool {
	return b & BadgeEarlyModpackAdopter != 0
}

// Returns whether the badges bitfield has the BadgeEarlyRespackAdopter bit set
func (b Badges) EarlyRespackAdopter() bool {
	return b & BadgeEarlyRespackAdopter != 0
}

// Returns whether the badges bitfield has the BadgeEarlyPluginAdopter bit set
func (b Badges) EarlyPluginAdopter() bool {
	return b & BadgeEarlyPluginAdopter != 0
}

// Returns whether the badges bitfield has the BadgeAlphaTester bit set
func (b Badges) AlphaTester() bool {
	return b & BadgeAlphaTester != 0
}

// Returns whether the badges bitfield has the BadgeContributor bit set
func (b Badges) Contributor() bool {
	return b & BadgeContributor != 0
}

// Returns whether the badges bitfield has the BadgeTranslator bit set
func (b Badges) Translator() bool {
	return b & BadgeTranslator != 0
}

