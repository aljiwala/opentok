package opentok

// Archive holds all the information retrieved from the server.
type Archive struct {

	// CreatedAt represents unix timestamp that specified when the archive was
	// created.
	CreatedAt int64 `json:"createdAt"`

	// Duration of the archive in seconds.
	Duration int64 `json:"duration"`

	// ID of the archive. This is a unique id identifier for the archive.
	// It's used to stop, retrieve and delete archives.
	ID string `json:"id"`

	// Name for the archive. The user can choose any name but it doesn't
	// necessarily need to be different between archives.
	Name string `json:"name"`

	// APIKey to which the archive belongs.
	APIKey int `json:"partnerId"`

	// SessionID to which the archive belongs.
	SessionID string `json:"sessionId"`

	// Size of the archives in KB.
	Size int `json:"size"`

	// URL from where the archive can be retrieved. This is only useful if the
	// archive is in status available in the OpenTok S3 Account.
	URL string `json:"url"`

	// Status of the Archive.
	// The possibilities are:
	// - `started`:   if the archive is being recorded
	// - `stopped`:   if the archive has been stopped and it hasn't been uploaded or available
	// - `deleted`:   if the archive has been deleted. Only available archives can be deleted
	// - `uploaded`:  if the archive has been uploaded to the partner storage account
	// - `paused`:    if the archive has not been stopped but it is not recording. It can transition to Started again
	// - `available`: if the archive has been uploaded to the OpenTok S3 account
	// - `expired`:   available archives are removed from the OpenTok S3 account after 3 days. Their status become expired.
	Status string `json:"status"`

	// HasAudio tells whether the archive contains an audio stream.
	HasAudio bool `json:"hasAudio"`

	// HasVideo tells whether the archive contains a video stream.
	HasVideo bool `json:"hasVideo"`
}

// ArchiveList will hold the list of archives retrieved from opentok service.
type ArchiveList struct {
	Count    int       `json:"count"`
	Archives []Archive `json:"items"`
}
