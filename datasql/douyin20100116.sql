/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 16/01/2020 14:08:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dy_dys
-- ----------------------------
DROP TABLE IF EXISTS `dy_dys`;
CREATE TABLE `dy_dys`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'uid',
  `nickname` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `shortid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '抖音ID',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `signature` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '简介',
  `focusnum` int(10) NOT NULL DEFAULT 0 COMMENT '关注人数',
  `followernum` int(10) NOT NULL DEFAULT 0 COMMENT '粉丝数',
  `likenum` int(10) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `createtime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updatetime` int(10) NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
