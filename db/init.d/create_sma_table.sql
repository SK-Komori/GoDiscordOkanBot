CREATE TABLE `players` (
    id VARCHAR(20) NOT NULL COMMENT 'ユーザID',
    rate int NOT NULL COMMENT 'レート',
    created Datetime NOT NULL COMMENT '登録日時',
    updated Datetime NOT NULL COMMENT '更新日時',
    PRIMARY KEY (id)
)