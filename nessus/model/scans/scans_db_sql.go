package scans

const (
	scans_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nessus_scans
-- ----------------------------
DROP TABLE IF EXISTS nessus_scans;
CREATE TABLE nessus_scans  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  scans_id bigint NOT NULL COMMENT '扫描id',
  user_id bigint NOT NULL COMMENT '用户id',
  admin_user_id bigint NOT NULL COMMENT 'admin用户id',
  is_delete tinyint NOT NULL COMMENT '删除1',
  create_time int NOT NULL COMMENT '创建时间',
  description varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '新增备注',
  addr varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '目标地址',
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 34 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
`
	scans_report_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for nessus_scan_report
-- ----------------------------
DROP TABLE IF EXISTS nessus_scan_report;
CREATE TABLE nessus_scan_report  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  scans_id bigint NOT NULL COMMENT '扫描id',
  user_id bigint NOT NULL COMMENT '用户id',
  admin_user_id bigint NOT NULL COMMENT 'admin用户id',
  is_delete tinyint NOT NULL COMMENT '删除1',
  create_time int NOT NULL COMMENT '创建时间',
  addr varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '目标',
  history_id varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'history_id',
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
`
)
