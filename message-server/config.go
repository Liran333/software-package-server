package main

import (
	"github.com/opensourceways/server-common-lib/utils"

	"github.com/opensourceways/software-package-server/common/infrastructure/kafka"
	"github.com/opensourceways/software-package-server/config"
	"github.com/opensourceways/software-package-server/softwarepkg/domain/dp"
	"github.com/opensourceways/software-package-server/softwarepkg/infrastructure/maintainerimpl"
)

func loadConfig(path string) (*Config, error) {
	cfg := new(Config)
	if err := utils.LoadFromYaml(path, cfg); err != nil {
		return nil, err
	}

	cfg.setDefault()

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

type Config struct {
	Kafka          kafka.Config            `json:"kafka"                required:"true"`
	Topics         Topics                  `json:"topics_to_subscribe"  required:"true"`
	GroupName      string                  `json:"group_name"           required:"true"`
	Postgresql     config.PostgresqlConfig `json:"postgresql"           required:"true"`
	Maintainer     maintainerimpl.Config   `json:"maintainer"           required:"true"`
	SoftwarePkg    dp.Config               `json:"software_pkg"         required:"true"`
	TopicsToNotify TopicsToNotify          `json:"topics_to_notify"     required:"true"`
}

type Topics struct {
	SoftwarePkgPRClosed    string `json:"software_pkg_pr_closed"      required:"true"`
	SoftwarePkgPRMerged    string `json:"software_pkg_pr_merged"      required:"true"`
	SoftwarePkgPRCIChecked string `json:"software_pkg_pr_ci_checked"  required:"true"`
	SoftwarePkgRepoCreated string `json:"software_pkg_repo_created"   required:"true"`
}

type TopicsToNotify struct {
	AlreadyClosedSoftwarePkg      string `json:"already_closed_software_pkg"        required:"true"`
	IndirectlyApprovedSoftwarePkg string `json:"indirectly_approved_software_pkg"   required:"true"`
}

type configValidate interface {
	Validate() error
}

type configSetDefault interface {
	SetDefault()
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Postgresql.DB,
		&cfg.Postgresql.Config,
		&cfg.SoftwarePkg,
		&cfg.Maintainer,
		&cfg.Kafka,
	}
}

func (cfg *Config) setDefault() {
	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configSetDefault); ok {
			f.SetDefault()
		}
	}
}

func (cfg *Config) validate() error {
	if _, err := utils.BuildRequestBody(cfg, ""); err != nil {
		return err
	}

	items := cfg.configItems()
	for _, i := range items {
		if f, ok := i.(configValidate); ok {
			if err := f.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
