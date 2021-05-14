package command

import (
	`github.com/storezhang/glog`
	`github.com/storezhang/pangu/app`
)

var _ app.Command = (*Migrate)(nil)

// Migrate 描述一个提供服务的命令
type Migrate struct {
	Base

	migration migration
	logger    glog.Logger
}

// NewMigrate 创建服务命令
func NewMigrate(logger *glog.ZapLogger) *Migrate {
	return &Migrate{
		Base: Base{
			Name:    "Migrate",
			Aliases: []string{"m", "M", "Migrate"},
			Usage:   "数据迁移",
		},

		logger: logger,
	}
}

func (m *Migrate) SetMigration(migration migration) {
	m.migration = migration
}

func (m *Migrate) Run(_ *app.Context) error {
	return m.migration.Migrate()
}
