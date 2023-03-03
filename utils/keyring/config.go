package keyring

// ConfigOpt defines a known cosmos keyring option
type ConfigOpt func(c *cosmosKeyringConfig)

type cosmosKeyringConfig struct {
	KeyringDir     string
	KeyringAppName string
	KeyringBackend Backend
	KeyFrom        string
	KeyPassphrase  string
	PrivKeyHex     string
	Mnemonic       string
	UseLedger      bool
}

// Backend defines a known keyring backend name
type Backend string

const (
	// BackendTest is a testing backend, no passphrases required
	BackendTest Backend = "test"
	// BackendFile is a backend where keys are stored as encrypted files
	BackendFile Backend = "file"
	// BackendOS is a backend where keys are stored in the OS key chain. Platform specific.
	BackendOS Backend = "os"
)

// WithKeyringDir option sets keyring path in the filesystem, useful when keyring backend is `file`.
func WithKeyringDir(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.KeyringDir = v
		}
	}
}

// WithKeyringAppName option sets keyring application name (used by Cosmos to separate keyrings)
func WithKeyringAppName(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.KeyringAppName = v
		}
	}
}

// WithKeyringBackend sets the keyring backend. Expected values: test, file, os.
func WithKeyringBackend(v Backend) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.KeyringBackend = v
		}
	}
}

// WithKeyFrom sets the key name to use for signing. Must exist in the provided keyring.
func WithKeyFrom(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.KeyFrom = v
		}
	}
}

// WithKeyPassphrase sets the passphrase for keyring files. Insecure option, use for testing only.
// The package will fallback to os.Stdin if this option was not provided, but pass is required.
func WithKeyPassphrase(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.KeyPassphrase = v
		}
	}
}

// WithPrivKeyHex allows to specify a private key as plaintext hex. Insecure option, use for testing only.
// The package will create a virtual keyring holding that key, to meet all the interfaces.
func WithPrivKeyHex(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.PrivKeyHex = v
		}
	}
}

// WithMnemonic allows to specify a mnemonic pharse as plaintext hex. Insecure option, use for testing only.
// The package will create a virtual keyring to derive the keys and meet all the interfaces.
func WithMnemonic(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		if len(v) > 0 {
			c.Mnemonic = v
		}
	}
}

// WithUseLedger sets the option to use hardware wallet, if available on the system.
func WithUseLedger(b bool) ConfigOpt {
	return func(c *cosmosKeyringConfig) {
		c.UseLedger = b
	}
}
