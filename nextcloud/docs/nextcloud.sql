CREATE DATABASE IF NOT EXISTS `nextcloud` DEFAULT CHARACTER SET utf8mb4;

CREATE TABLE IF NOT EXISTS `nextcloud_token` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `user` varchar(63) NOT NULL DEFAULT '' COMMENT '主站用户名',
     `uid` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
     `token` varchar(127) NOT NULL DEFAULT '' COMMENT 'nextcloud token',
     PRIMARY KEY (`id`),
     UNIQUE KEY `user` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据备份token表';

INSERT INTO  `subassemblynode` (
     `name`,`addr`,`type`,`key`,`secret`
) VALUES ('nextcloud','http://182.150.0.80:18002/backup',8,'admin','Dengbao123!@#');