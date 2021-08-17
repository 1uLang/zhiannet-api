package cert

const cert_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for next_terminal_cert
-- ----------------------------
DROP TABLE IF EXISTS next_terminal_cert;
CREATE TABLE next_terminal_cert  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  cert_id varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '扫描id',
  user_id bigint NOT NULL COMMENT '用户id',
  admin_user_id bigint NOT NULL COMMENT 'admin用户id',
  is_delete tinyint NOT NULL COMMENT '删除1',
  create_time int NOT NULL COMMENT '创建时间',
  is_auth tinyint NOT NULL COMMENT '是否是授权凭证',
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
`
