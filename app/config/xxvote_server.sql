/*
 Navicat Premium Data Transfer

 Source Server         : 8.210.175.17
 Source Server Type    : MySQL
 Source Server Version : 50742
 Source Host           : 8.210.175.17:3306
 Source Schema         : xxvote_server

 Target Server Type    : MySQL
 Target Server Version : 50742
 File Encoding         : 65001

 Date: 23/04/2024 00:18:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NULL DEFAULT NULL COMMENT '可展示的uid',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 12, 'admin', '123', '2024-04-16 19:36:04', '2024-04-16 19:36:07');
INSERT INTO `user` VALUES (2, 11, 't', '123', '2024-04-16 20:51:52', '2024-04-16 20:51:52');

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF EXISTS `vote`;
CREATE TABLE `vote`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `time` int(11) NULL DEFAULT NULL COMMENT '有效时长',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '是否有效',
  `type` tinyint(1) NULL DEFAULT NULL COMMENT '0单选 1多选',
  `repeat` tinyint(1) NULL DEFAULT NULL COMMENT '0不可重复 1可重复',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote
-- ----------------------------
INSERT INTO `vote` VALUES (1, 1, 'what to eat tonight', 3600, 1, 1, 1, '2024-04-17 20:52:35', '2024-04-17 20:52:37');
INSERT INTO `vote` VALUES (2, 2, '选择投票', 60, 1, 1, 1, '2024-04-17 22:26:12', '2024-04-17 22:26:15');
INSERT INTO `vote` VALUES (3, 1, '钟意的厂商', 60, 1, 1, 1, '2024-04-18 19:40:14', '2024-04-18 19:40:17');

-- ----------------------------
-- Table structure for vote_option
-- ----------------------------
DROP TABLE IF EXISTS `vote_option`;
CREATE TABLE `vote_option`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vote_id` int(11) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `count` int(11) NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_option
-- ----------------------------
INSERT INTO `vote_option` VALUES (1, 1, 'chicken', 2, '2024-04-22 17:57:54', '2024-04-22 17:57:56');
INSERT INTO `vote_option` VALUES (2, 1, '凉皮', 3, '2024-04-22 17:58:27', '2024-04-22 17:58:30');
INSERT INTO `vote_option` VALUES (3, 3, '华为', 2, '2024-04-22 17:59:03', '2024-04-22 17:59:05');
INSERT INTO `vote_option` VALUES (4, 2, 'A', 4, '2024-04-22 18:13:35', '2024-04-22 18:13:38');
INSERT INTO `vote_option` VALUES (5, 2, 'B', 1, '2024-04-22 18:13:55', '2024-04-22 18:13:58');

-- ----------------------------
-- Table structure for vote_option_user
-- ----------------------------
DROP TABLE IF EXISTS `vote_option_user`;
CREATE TABLE `vote_option_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `vote_id` int(11) NULL DEFAULT NULL,
  `vote_option_id` int(11) NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_option_user
-- ----------------------------
INSERT INTO `vote_option_user` VALUES (1, 1, 1, 1, '2024-04-22 17:59:41', '2024-04-22 17:59:44');

SET FOREIGN_KEY_CHECKS = 1;
