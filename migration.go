package pangu

import (
	`io/fs`
	`net/http`

	`github.com/go-sql-driver/mysql`
	`github.com/rubenv/sql-migrate`
	`github.com/storezhang/gox`
	`xorm.io/builder`
	`xorm.io/xorm`
)

const noSuchTable = 1146

type migration struct {
	// Id 文件名称
	Id string `xorm:"varchar(64) notnull default('')"`
	// AppliedAt 升级时间
	AppliedAt gox.Timestamp `xorm:"created default('2020-02-04 09:55:52')"`

	migrations  []fs.FS      `xorm:"-"`
	application *Application `xorm:"-"`
}

func newMigration(application *Application) *migration {
	return &migration{
		migrations:  make([]fs.FS, 0, 0),
		application: application,
	}
}

func (m *migration) migrate() (err error) {
	if 0 == len(m.migrations) {
		return
	}

	err = m.application.Get(func(database gox.DatabaseConfig, engine *xorm.Engine) (err error) {
		var migrations migrate.MigrationSource

		// 设置升级记录的表名，默认值是grop_migrations
		migrate.SetTable(database.MigrationTableName)
		migrate.SetIgnoreUnknown(true)

		// 开始升级数据库
		// 如果升级有错误，应退出程序
		// 不应该完成启动，导致数据库错误越来越离谱
		migrations = &migrate.HttpFileSystemMigrationSource{
			FileSystem: http.FS(m.migrations[0]),
		}

		if err = m.cleanDeletedMigrations(migrations, engine); nil != err {
			return
		}
		_, err = migrate.Exec(engine.DB().DB, database.Type, migrations, migrate.Up)

		return
	})

	return
}

func (m *migration) addSource(migration fs.FS) {
	m.migrations = append(m.migrations, migration)
}

func (m *migration) shouldMigration() bool {
	return 0 != len(m.migrations)
}

func (m *migration) cleanDeletedMigrations(ms migrate.MigrationSource, engine *xorm.Engine) (err error) {
	var (
		migrates     []*migrate.Migration
		migrateFiles = make([]string, 0)
	)

	if migrates, err = ms.FindMigrations(); nil != err {
		return
	}

	for _, m := range migrates {
		migrateFiles = append(migrateFiles, m.Id)
	}

	cond := builder.NotIn("id", migrateFiles)
	if _, err = engine.Where(cond).Delete(&migration{}); nil != err {
		// 表不存在不需要清理
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if noSuchTable == mysqlErr.Number {
				err = nil
			}
		}
	}

	return
}
