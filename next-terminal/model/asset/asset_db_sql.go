package asset

const asset_db_sql = `SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for next_terminal_assets
-- ----------------------------
DROP TABLE IF EXISTS next_terminal_assets;
CREATE TABLE next_terminal_assets  (
                                         id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
                                         asset_id varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '扫描id',
                                         user_id bigint NOT NULL COMMENT '用户id',
                                         admin_user_id bigint NOT NULL COMMENT 'admin用户id',
                                         is_delete tinyint NOT NULL COMMENT '删除1',
                                         create_time int NOT NULL COMMENT '创建时间',
                                         name varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '新增备注',
                                         proto varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '目标地址',
                                         is_auth tinyint NOT NULL COMMENT '是否是授权资产',
                                         PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;`
