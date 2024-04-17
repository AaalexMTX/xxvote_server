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

 Date: 17/04/2024 18:04:23
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
  `name` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `password` varchar(50) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, NULL, 'admin', '123', '2024-04-16 19:36:04', '2024-04-16 19:36:07');
INSERT INTO `user` VALUES (2, 11, 't', '123', '2024-04-16 20:51:52', '2024-04-16 20:51:52');

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF EXISTS `vote`;
CREATE TABLE `vote`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `time` datetime NULL DEFAULT NULL COMMENT '有效时长',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '是否有效',
  `type` tinyint(1) NULL DEFAULT NULL COMMENT '0单选 1多选',
  `repeat` tinyint(1) NULL DEFAULT NULL COMMENT '0不可重复 1可重复',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote
-- ----------------------------

-- ----------------------------
-- Table structure for vote_option
-- ----------------------------
DROP TABLE IF EXISTS `vote_option`;
CREATE TABLE `vote_option`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vote_id` int(11) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `count` int(11) NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_option
-- ----------------------------

-- ----------------------------
-- Table structure for vote_option_user
-- ----------------------------
DROP TABLE IF EXISTS `vote_option_user`;
CREATE TABLE `vote_option_user`  (
  `id` int(11) NOT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  `vote_id` int(11) NULL DEFAULT NULL,
  `vote_option_id` int(11) NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote_option_user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
