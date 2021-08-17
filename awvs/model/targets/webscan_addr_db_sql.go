package targets

const (
	webscan_addr_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for webscan_addr
-- ----------------------------
DROP TABLE IF EXISTS webscan_addr;
CREATE TABLE webscan_addr  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  target_id varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '扫描ID',
  user_id int NOT NULL DEFAULT 0 COMMENT '用户ID',
  is_delete tinyint NOT NULL DEFAULT 0 COMMENT '1删除 0未删除',
  create_time int NOT NULL DEFAULT 0 COMMENT '创建时间',
  admin_user_id int NOT NULL DEFAULT 0 COMMENT 'admin用户ID',
  PRIMARY KEY (id) USING BTREE,
  INDEX user_id(user_id) USING BTREE,
  INDEX admin_user_id(admin_user_id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '漏洞扫描关联用户表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
`
	webscan_report_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for webscan_report
-- ----------------------------
DROP TABLE IF EXISTS webscan_report;
CREATE TABLE webscan_report  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  report_id varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '报告ID',
  user_id int NOT NULL DEFAULT 0 COMMENT '用户ID',
  is_delete tinyint NOT NULL DEFAULT 0 COMMENT '1删除 0未删除',
  create_time int NOT NULL DEFAULT 0 COMMENT '创建时间',
  admin_user_id int NOT NULL DEFAULT 0 COMMENT '管理端用户',
  PRIMARY KEY (id) USING BTREE,
  INDEX user_id(user_id) USING BTREE,
  INDEX admin_user_id(admin_user_id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '漏洞扫描报告关联用户表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
`
)
