USE edges;

CREATE TABLE IF NOT EXISTS `agent_file` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
     `name` varchar(63) NOT NULL DEFAULT '' COMMENT '文件名',
     `describe` varchar(63) NOT NULL DEFAULT '' COMMENT '文件描述信息',
     `size` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '文件大小',
     `format` varchar(127) NOT NULL DEFAULT '' COMMENT '文件格式',
     `state` tinyint(3) unsigned NOT NULL DEFAULT 1 COMMENT '状态 1正常 0删除',
     `path` varchar(127) NOT NULL DEFAULT '' COMMENT '文件存储路径',
     `created_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '文件上传时间',
     `updated_at` int(10) unsigned NOT NULL DEFAULT 0 COMMENT '文件更新时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代理文件详情表';