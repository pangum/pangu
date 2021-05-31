-- +migrate Up
CREATE TABLE IF NOT EXISTS `user`
(
    `id`         BIGINT      NOT NULL DEFAULT 0 COMMENT '编号',
    `nickname`   VARCHAR(32) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar`     CHAR(20)    NOT NULL DEFAULT '' COMMENT '存放在对象存储上的用户头像的文件编号',

    `created_at` DATETIME    NULL     DEFAULT '2020-02-04 09:55:52' COMMENT '创建时间',
    `updated_at` DATETIME    NULL     DEFAULT '2020-02-04 09:55:52' COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户'
  ROW_FORMAT = DYNAMIC;


-- +migrate Down
DROP TABLE user;
