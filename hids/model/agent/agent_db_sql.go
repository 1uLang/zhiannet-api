package agent

const agent_db_sql = `
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hids_agents
-- ----------------------------
DROP TABLE IF EXISTS hids_agents;
CREATE TABLE hids_agents  (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  ip varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '扫描id',
  user_id bigint NOT NULL COMMENT '用户id',
  admin_user_id bigint NOT NULL COMMENT 'admin用户id',
  is_delete tinyint NOT NULL COMMENT '删除1',
  create_time int NOT NULL COMMENT '创建时间',
  PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

`
