CREATE TABLE `video`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT,
    `user_id`     bigint       NOT NULL COMMENT '用户id',
    `job_id`      bigint       NOT NULL COMMENT '任务ID',
    `description` varchar(255) NOT NULL COMMENT '作品描述',
    `file_addr`   varchar(255) NOT NULL COMMENT '视频文件url地址',
    `cover_addr`  varchar(255) NOT NULL COMMENT '封面url地址',
    `status`      int          NOT NULL COMMENT '状态',
    `created_at`  datetime(3) DEFAULT NULL,
    `updated_at`  datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

