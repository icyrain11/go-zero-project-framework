CREATE TABLE `upload_task`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`    bigint       NOT NULL,
    `file_addr`  varchar(255) NOT NULL,
    `expired_at` bigint DEFAULT NULL COMMENT '过期时间',
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

