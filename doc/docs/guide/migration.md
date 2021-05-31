# 数据迁移

`盘古`内置了数据迁移功能，目录只支持关系型数据库的数据迁移，以后可能会增加以下类型的数据库迁移

- Elasticsearch
- Redis
- 其它类型的数据迁移

## 使用数据迁移

要给应用程序增加数据迁移功能非常简单，只需要在pangu.Application上增加迁移源就可以了

<<< @/../example/bootstrap.go#snippet{14,36}

## 定义数据迁移命令格式

`盘古`使用SQL文件来做数据迁移，具体的文件格式如下

<<< @/../example/db/migration/1_init_user.sql

其中

- `-- +migrate Up`中的内容表示升级时执行的脚本
- `-- +migrate Down`中的内容表示降低时执行的脚本
