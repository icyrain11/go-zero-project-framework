CREATE TABLE `user`
(
    `id`        bigint                                  NOT NULL AUTO_INCREMENT,
    `username`  varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'The username',
    `password`  varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The user password',
    `mobile`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    `gender`    char(10) COLLATE utf8mb4_general_ci     NOT NULL DEFAULT 'male' COMMENT 'gender,male|female|unknown',
    `nickname`  varchar(255) COLLATE utf8mb4_general_ci          DEFAULT '' COMMENT 'The nickname',
    `type`      tinyint(1) DEFAULT '0' COMMENT 'The user type, 0:normal,1:vip, for test golang keyword',
    `create_at` timestamp NULL DEFAULT NULL,
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `mobile_index` (`mobile`),
    UNIQUE KEY `name_index` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='user table'

