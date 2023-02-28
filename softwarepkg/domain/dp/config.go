package dp

var config Config

func Init(cfg *Config) {
	config = *cfg
}

type Config struct {
	MaxLengthOfPackageName       int `json:"max_length_of_pkg_name"`
	MaxLengthOfPackageDesc       int `json:"max_length_of_pkg_desc"`
	MaxLengthOfReviewComment     int `json:"max_length_of_review_comment"`
	MaxLengthOfReasonToImportPkg int `json:"max_length_of_reason_to_import_pkg"`
}

func (cfg *Config) SetDefault() {
	if cfg.MaxLengthOfPackageName <= 0 {
		cfg.MaxLengthOfPackageName = 50
	}

	if cfg.MaxLengthOfPackageDesc <= 0 {
		cfg.MaxLengthOfPackageDesc = 1000
	}

	if cfg.MaxLengthOfReviewComment <= 0 {
		cfg.MaxLengthOfReviewComment = 500
	}

	if cfg.MaxLengthOfReasonToImportPkg <= 0 {
		cfg.MaxLengthOfReasonToImportPkg = 1000
	}
}

func (cfg *Config) Validate() error {
	return nil
}