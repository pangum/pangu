package command

import (
	`sync`

	`github.com/storezhang/glog`
	`github.com/storezhang/gox/field`
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
			Usage:   "升级服务",
		},

		logger: logger,
	}
}

func (s *Migrate) Adds(Migrates ...Migrate) {
	s.Migrates = append(s.Migrates, Migrates...)
}

func (s *Migrate) SetMigration(migration migration) {
	s.migration = migration
}

func (s *Migrate) Run(ctx *app.Context) (err error) {
	if s.migration.ShouldMigrate() {
		s.logger.Info("服务升级开始", field.Int("count", s.migration.NeedMigrateCount()))

		if err = s.migration.Migrate(); nil != err {
			return
		}

		s.logger.Info("服务升级成功", field.Int("count", s.migration.NeedMigrateCount()))
	}

	MigrateCount := len(s.Migrates)
	if 0 != len(s.Migrates) {
		s.logger.Info("启动服务开始", field.Int("count", MigrateCount))

		if err = s.runMigrates(ctx); nil != err {
			return
		}

		s.logger.Info("启动服务成功", field.Int("count", MigrateCount))
	}

	return
}

func (s *Migrate) runMigrates(_ *app.Context) (err error) {
	wg := sync.WaitGroup{}
	worker := len(s.Migrates)
	wg.Add(worker)

	for _, Migrate := range s.Migrates {
		Migrate := Migrate
		go func() {
			defer wg.Done()

			s.logger.Info("启动服务器成功", field.String("name", Migrate.Name()))
			// 服务器不允许中途有服务器启动错误，如果有，应该立即关掉容器
			// 如果调用者想并行执行，可以使用recover机制来阻止程序退出
			if err = Migrate.Run(); nil != err {
				panic(err)
			}
		}()
	}

	wg.Wait()

	return
}
