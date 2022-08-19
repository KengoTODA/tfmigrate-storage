package gcs

import storage "github.com/minamijoyo/tfmigrate-storage"

// Config is a config for Google Cloud Storage.
// This is expected to have almost the same options as Terraform gcs backend.
// https://www.terraform.io/language/settings/backends/gcs
// However, it has many minor options and it's a pain to test all options from
// first, so we added only options we need for now.
type Config struct {
	// The name of the GCS bucket.
	Bucket string `hcl:"bucket"`
	// Path to the migration history file.
	Key string `hcl:"key"`
	// Local path to Google Cloud Platform account credentials in JSON format.
	// If unset, Google Application Default Credentials are used.
	Credentials string `hcl:"credentials,optional"`
	// GCS prefix inside the bucket.
	// Named states for workspaces are stored in an object called <prefix>/<name>.tfstate.
	Prefix string `hcl:"prefix,optional"`

	Endpoint string `hcl:"endpoint,optional"`
}

// Config implements a storage.Config.
var _ storage.Config = (*Config)(nil)

// NewStorage returns a new instance of storage.Storage.
func (c *Config) NewStorage() (storage.Storage, error) {
	return NewStorage(c, nil)
}
