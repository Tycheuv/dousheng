
CREATE TABLE `users` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` longtext,
    `follow_count` bigint DEFAULT NULL,
    `follower_count` bigint DEFAULT NULL,
    `is_follow` tinyint(1) DEFAULT NULL,
    `avatar` longtext,
    `background_image` longtext,
    `favorite_count` bigint DEFAULT NULL,
    `signature` longtext,
    `total_favorited` longtext,
    `work_count` bigint DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '用户信息表';

CREATE TABLE `accounts` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint DEFAULT NULL,
    `username` longtext,
    `password` longtext,
    PRIMARY KEY (`id`),
    KEY `fk_accounts_user` (`user_id`),
    CONSTRAINT `fk_accounts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci comment '用户账号表';

CREATE TABLE `videos` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `author_id` bigint DEFAULT NULL,
    `comment_count` bigint DEFAULT NULL,
    `cover_url` longtext,
    `favorite_count` bigint DEFAULT NULL,
    `is_favorite` tinyint(1) DEFAULT NULL,
    `play_url` longtext,
    `title` longtext,
    PRIMARY KEY (`id`),
    KEY `fk_videos_author` (`author_id`),
    CONSTRAINT `fk_videos_author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '视频流表';

CREATE TABLE `favorites` (
    `token` longtext,
    `video_id` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '喜欢表';

CREATE TABLE `comments` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `content` longtext,
    `create_date` longtext,
    `user_id` bigint DEFAULT NULL,
    `video_id` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_comments_user` (`user_id`),
    CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '评论表';

CREATE TABLE `relations` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint DEFAULT NULL,
    `to_user_id` bigint DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '关系表';

CREATE TABLE `messages` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `from_user_id` bigint DEFAULT NULL,
    `to_user_id` bigint DEFAULT NULL,
    `content` longtext,
    `create_time` bigint DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT '聊天消息表';

