/*
 Navicat Premium Data Transfer

 Source Server         : 我的本地连接
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 21/08/2020 18:36:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `article_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `title` varchar(100) NOT NULL COMMENT '文章标题',
  `desc` varchar(255) NOT NULL COMMENT '文章简述',
  `cover_image_url` varchar(255) NOT NULL COMMENT '封面图片地址',
  `content` longtext NOT NULL COMMENT '文章内容',
  `status` tinyint(3) NOT NULL DEFAULT '10' COMMENT '状态 10 启用 5 停用',
  `is_del` tinyint(3) NOT NULL DEFAULT '10' COMMENT '删除状态 10 未删除 5 已删除',
  `deleted_at` datetime NOT NULL COMMENT '删除时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
  `id` int(11) NOT NULL,
  `article_id` int(11) NOT NULL COMMENT '文章ID',
  `tag_id` int(11) NOT NULL COMMENT '标签ID',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '标签名称',
  `status` tinyint(3) NOT NULL DEFAULT '10' COMMENT '状态 10 启用 5 停用',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
INSERT INTO `tag` VALUES (6, 'compress', 10, '2020-08-21 14:26:32', '2020-08-21 14:41:23');
INSERT INTO `tag` VALUES (7, 'program', 10, '2020-08-21 14:26:45', '2020-08-21 14:41:04');
INSERT INTO `tag` VALUES (8, 'Bahringer Bypass', 10, '2020-08-21 14:27:22', '2020-08-21 14:27:22');
INSERT INTO `tag` VALUES (9, 'Aniyah Manor', 10, '2020-08-21 14:27:23', '2020-08-21 14:27:23');
INSERT INTO `tag` VALUES (11, 'Murphy Causeway', 10, '2020-08-21 18:24:03', '2020-08-21 18:24:03');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
