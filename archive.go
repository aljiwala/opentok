package opentok

// ArchiveMode denotes different modes for an archive.
// Ref: https://tokbox.com/developer/rest/#start_archive
type ArchiveMode string

const (
	// ArchiveModeManual is the string representation of archive mode status `manual`.
	ArchiveModeManual ArchiveMode = "manual"

	// ArchiveModeAlways is the string representation of archive mode status `always`.
	ArchiveModeAlways ArchiveMode = "always"
)
