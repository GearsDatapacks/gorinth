package gorinth

// A donation link for a project
type ProjectDonationURL struct {
	// The ID of the donation platform
	Id string `json:"id"`
	// The donation platform this link is to
	Platform string `json:"platform"`
	// The URL of the donation platform and user
	Url string `json:"url"`
}

// The license of a project. Pointer fields are optional
type License struct {
	// The SPDX license ID of a project
	Id string `json:"id"`
	// The long name of a license
	Name string `json:"name"`
	// The URL to this license
	Url *string `json:"url"`
}

// An image in the gallery of a project. Pointer fields are optional
type GalleryImage struct {
	// The URL of the gallery image
	Url string `json:"url"`
	// Whether or not the image is featured in the gallery
	Featured bool `json:"featured"`
	// The title of the image
	Title *string `json:"title"`
	// The description of the image
	Description *string `json:"description"`
	// The date and time the gallery image was created (ISO-8601 format)
	Created string `json:"created"`
	// The order of the gallery image.
	// Gallery images are sorted by this field and then alphabetically by title
	Ordering int `json:"ordering"`
}

// A specific version of a project that another version depends on
type Dependency struct {
	// The ID of the version that is depended on
	VersionId *string `json:"version_id"`
	// The ID of the project that is depended on
	ProjectId *string `json:"project_id"`
	// For embedded dependencies. The name of the file which is embedded
	FileName *string `json:"file_name"`
	// The type of the dependency
	DependencyType DependencyType `json:"dependency_type"`
}

type VersionFile struct {
	// A map of hashing algorithms to hashes of the file
	Hashes struct {
		sha512 string
		sha1   string
	} `json:"hashes"`
	// The url of the file
	Url string `json:"url"`
	// The name of the file
	Filename string `json:"filename"`
	// Whether this is the primary file for the version
	Primary bool `json:"primary"`
	// The size of the file in bytes
	Size int `json:"size"`
	// The type of the file
	FileType *FileType `json:"file_type"`
}

// Data regarding a user's payout status
type PayoutData struct {
	// The balance available for withdrawal
	Balance string `json:"balance"`
	// The wallet that the user has selected
	PayoutWallet     PayoutWallet     `json:"payout_wallet"`
	// The type of the user's wallet
	PayoutWallerType PayoutWalletType `json:"payout_wallet_type"`
	// The user's payout address
	PayoutAddress string `json:"payout_address"`
}
