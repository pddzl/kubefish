# ************************************************************
# Sequel Ace SQL dump
# Version 20046
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 127.0.0.1 (MySQL 8.0.28)
# Database: kop
# Generation Time: 2023-07-03 02:11:04 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table casbin_rule
# ------------------------------------------------------------

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;

INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
VALUES
	(1,'p','1','/api/addApi','POST','','',''),
	(2,'p','1','/api/deleteApi','POST','','',''),
	(3,'p','1','/api/editApi','POST','','',''),
	(4,'p','1','/api/getApis','POST','','',''),
	(5,'p','1','/api/getElTreeApis','POST','','',''),
	(6,'p','1','/base/captcha','POST','','',''),
	(7,'p','1','/base/login','POST','','',''),
	(8,'p','1','/casbin/editCasbin','POST','','',''),
	(9,'p','1','/jwt/joinInBlacklist','POST','','',''),
	(10,'p','1','/menu/addMenu','POST','','',''),
	(11,'p','1','/menu/deleteMenu','POST','','',''),
	(12,'p','1','/menu/editMenu','POST','','',''),
	(13,'p','1','/menu/getElTreeMenus','POST','','',''),
	(14,'p','1','/menu/getMenus','GET','','',''),
	(15,'p','1','/role/addRole','POST','','',''),
	(16,'p','1','/role/deleteRole','POST','','',''),
	(17,'p','1','/role/editRole','POST','','',''),
	(18,'p','1','/role/editRoleMenu','POST','','',''),
	(19,'p','1','/role/getRoles','POST','','',''),
	(20,'p','1','/user/addUser','POST','','',''),
	(21,'p','1','/user/deleteUser','POST','','',''),
	(22,'p','1','/user/editUser','POST','','',''),
	(23,'p','1','/user/getUserInfo','GET','','',''),
	(24,'p','1','/user/getUsers','POST','','',''),
	(25,'p','1','/user/modifyPass','POST','','',''),
	(26,'p','1','/user/switchActive','POST','','','');

/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table jwt_blacklists
# ------------------------------------------------------------

DROP TABLE IF EXISTS `jwt_blacklists`;

CREATE TABLE `jwt_blacklists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



# Dump of table role_menus
# ------------------------------------------------------------

DROP TABLE IF EXISTS `role_menus`;

CREATE TABLE `role_menus` (
  `menu_model_id` bigint unsigned NOT NULL,
  `role_model_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`menu_model_id`,`role_model_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `role_menus` WRITE;
/*!40000 ALTER TABLE `role_menus` DISABLE KEYS */;

INSERT INTO `role_menus` (`menu_model_id`, `role_model_id`)
VALUES
	(1,1),
	(2,1),
	(3,1),
	(4,1),
	(5,1),
	(10,1),
	(11,1),
	(12,1),
	(13,1),
	(14,1),
	(15,1),
	(16,1),
	(17,1),
	(18,1),
	(19,1),
	(20,1),
	(21,1),
	(22,1),
	(23,1),
	(24,1),
	(25,1),
	(26,1),
	(27,1),
	(28,1),
	(29,1),
	(30,1),
	(31,1),
	(32,1),
	(33,1),
	(34,1),
	(35,1),
	(36,1),
	(37,1),
	(38,1),
	(39,1),
	(40,1),
	(41,1),
	(42,1),
	(43,1),
	(44,1),
	(45,1),
	(46,1);

/*!40000 ALTER TABLE `role_menus` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_api
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_api`;

CREATE TABLE `sys_api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_api_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_api` WRITE;
/*!40000 ALTER TABLE `sys_api` DISABLE KEYS */;

INSERT INTO `sys_api` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `description`, `api_group`, `method`)
VALUES
	(1,'2023-03-10 06:24:36','2023-03-10 06:24:36',NULL,'/base/captcha','获取验证码（必选）','base','POST'),
	(2,'2023-03-08 06:36:24','2023-03-09 08:50:20',NULL,'/base/login','登录（必选）','base','POST'),
	(3,'2023-03-08 08:56:13','2023-03-10 07:11:53',NULL,'/user/getUserInfo','获取用户信息（必选）','user','GET'),
	(4,'2023-03-08 08:56:54','2023-03-08 08:56:54',NULL,'/user/getUsers','获取所有用户','user','POST'),
	(5,'2023-03-10 06:41:32','2023-03-10 06:41:32',NULL,'/user/deleteUser','删除用户','user','POST'),
	(6,'2023-03-10 06:42:24','2023-03-10 06:42:24',NULL,'/user/addUser','添加用户','user','POST'),
	(7,'2023-03-10 06:47:18','2023-03-10 06:47:18',NULL,'/user/editUser','编辑用户','user','POST'),
	(8,'2023-03-10 06:47:59','2023-03-10 06:47:59',NULL,'/user/modifyPass','修改用户密码','user','POST'),
	(9,'2023-03-10 06:48:43','2023-03-10 06:48:43',NULL,'/user/switchActive','切换用户状态','user','POST'),
	(10,'2023-03-10 06:58:30','2023-03-10 06:58:30',NULL,'/role/getRoles','获取所有角色','role','POST'),
	(11,'2023-03-10 06:59:08','2023-03-10 06:59:08',NULL,'/role/addRole','添加角色','role','POST'),
	(12,'2023-03-10 06:59:54','2023-03-10 06:59:54',NULL,'/role/deleteRole','删除角色','role','POST'),
	(13,'2023-03-10 07:00:14','2023-03-10 07:00:53',NULL,'/role/editRole','编辑角色','role','POST'),
	(14,'2023-03-10 07:01:44','2023-03-10 07:01:44',NULL,'/role/editRoleMenu','编辑角色菜单','role','POST'),
	(15,'2023-03-10 07:14:44','2023-03-10 07:14:44',NULL,'/menu/getMenus','获取所有菜单','menu','GET'),
	(16,'2023-03-10 07:15:25','2023-03-10 07:15:25',NULL,'/menu/addMenu','添加菜单','menu','POST'),
	(17,'2023-03-10 07:15:50','2023-03-10 07:15:50',NULL,'/menu/editMenu','编辑菜单','menu','POST'),
	(18,'2023-03-10 07:16:18','2023-03-10 07:16:18',NULL,'/menu/deleteMenu','删除菜单','menu','POST'),
	(19,'2023-03-10 07:17:13','2023-03-10 07:17:13',NULL,'/menu/getElTreeMenus','获取所有菜单（el-tree结构）','menu','POST'),
	(20,'2023-03-10 07:21:37','2023-03-10 07:21:37',NULL,'/casbin/editCasbin','编辑casbin规则','casbin','POST'),
	(21,'2023-03-10 07:23:21','2023-03-10 07:33:01',NULL,'/api/addApi','添加api','api','POST'),
	(22,'2023-03-10 07:24:00','2023-03-10 07:24:00',NULL,'/api/getApis','获取所有api','api','POST'),
	(23,'2023-03-10 07:24:33','2023-03-10 07:24:33',NULL,'/api/deleteApi','删除api','api','POST'),
	(24,'2023-03-10 07:26:15','2023-03-10 07:26:15',NULL,'/api/editApi','编辑api','api','POST'),
	(25,'2023-03-10 07:34:08','2023-03-10 07:35:04',NULL,'/api/getElTreeApis','获取所有api（el-tree结构）','api','POST'),
	(27,'2023-03-11 13:05:40','2023-03-11 13:05:40',NULL,'/jwt/joinInBlacklist','拉黑token','jwt','POST');

/*!40000 ALTER TABLE `sys_api` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `pid` bigint unsigned DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `redirect` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta` json DEFAULT NULL,
  `sort` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `path` (`path`),
  KEY `idx_sys_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`id`, `created_at`, `updated_at`, `deleted_at`, `pid`, `name`, `path`, `redirect`, `component`, `meta`, `sort`)
VALUES
	(1,NULL,NULL,NULL,0,'Setting','/setting','/setting/user','Layout','{\"title\": \"系统管理\", \"svgIcon\": \"setting\"}',0),
	(2,NULL,NULL,NULL,1,'User','user',NULL,'setting/user/index.vue','{\"title\": \"用户管理\"}',0),
	(3,NULL,NULL,NULL,1,'Role','role',NULL,'setting/role/index.vue','{\"title\": \"角色管理\"}',0),
	(4,NULL,NULL,NULL,1,'Menu','menu',NULL,'setting/menu/index.vue','{\"title\": \"菜单管理\"}',0),
	(5,'2023-03-07 01:50:48','2023-03-11 15:05:15',NULL,1,'Api','api','','setting/api/index.vue','{\"title\": \"接口管理\", \"keepAlive\": true}',0),
	(7,NULL,NULL,NULL,6,'Cenu1','cenu1','/cenu/cenu1/cenu1-1','cenu/cenu1/index.vue','{\"title\": \"cenu1\"}',0),
	(8,NULL,'2023-03-11 14:06:08',NULL,7,'Cenu1-1','cenu1-1','','cenu/cenu1/cenu1-1/index.vue','{\"title\": \"cenu1-1\"}',0),
	(9,'2023-03-13 06:14:27','2023-03-13 06:14:27',NULL,7,'Cenu1-2','cenu1-2','','cenu/cenu1/cenu1-2/index.vue','{\"title\": \"cenu1-2\"}',0),
	(10,'2023-04-18 11:43:00','2023-06-08 09:51:52',NULL,0,'Cluster','/k8s/cluster','/k8s/cluster/nodeList','Layout','{\"title\": \"集群\", \"svgIcon\": \"cluster\"}',0),
	(11,'2023-04-18 11:50:25','2023-06-05 06:15:12',NULL,10,'NodeList','/k8s/cluster/nodeList','','k8s/cluster/node/list.vue','{\"title\": \"节点\"}',0),
	(12,'2023-04-23 03:10:24','2023-04-23 03:11:41',NULL,10,'NodeDetail','/k8s/cluster/nodeDetail','','k8s/cluster/node/detail.vue','{\"title\": \"节点详情\", \"hidden\": true}',0),
	(13,'2023-04-24 08:51:57','2023-04-24 09:13:43',NULL,10,'NamepsaceList','/k8s/cluster/namespaceList','','k8s/cluster/namespace/list.vue','{\"title\": \"命名空间\"}',0),
	(14,'2023-04-26 07:43:47','2023-04-26 07:44:12',NULL,10,'NamespaceDetail','/k8s/cluster/namespaceDetail','','k8s/cluster/namespace/detail.vue','{\"title\": \"命名空间详情\", \"hidden\": true}',0),
	(15,'2023-04-27 07:31:55','2023-06-08 09:43:26',NULL,0,'Workloads','/k8s/workloads','/k8s/workloads/PodList','Layout','{\"title\": \"工作负载\", \"svgIcon\": \"load\"}',0),
	(16,'2023-04-27 07:33:43','2023-06-02 15:40:38',NULL,15,'PodList','pod/list','','k8s/workloads/pod/list.vue','{\"title\": \"Pod\"}',0),
	(17,'2023-05-05 09:14:12','2023-06-02 15:39:12',NULL,15,'PodDetail','pod/detail','','k8s/workloads/pod/detail.vue','{\"title\": \"Pod详情\", \"hidden\": true}',0),
	(18,'2023-06-02 15:03:23','2023-06-02 15:37:50',NULL,15,'ReplicaSetList','replicaSet/list','','k8s/workloads/replicaSet/list.vue','{\"title\": \"ReplicaSet\"}',0),
	(19,'2023-06-05 06:10:50','2023-06-05 06:15:45',NULL,15,'ReplicaSetDetail','replicaSet/detail','','k8s/workloads/replicaSet/detail.vue','{\"title\": \"ReplicaSet详情\", \"hidden\": true}',0),
	(20,'2023-06-06 08:34:23','2023-06-06 08:34:45',NULL,15,'Deployment','deployment/list','','k8s/workloads/deployment/list.vue','{\"title\": \"Deployment\"}',0),
	(21,'2023-06-06 08:45:57','2023-06-06 08:47:25',NULL,15,'DeploymentDetail','deployment/detail','','k8s/workloads/deployment/detail.vue','{\"title\": \"Deployment详情\", \"hidden\": true}',0),
	(22,'2023-06-08 02:20:44','2023-06-08 02:20:44',NULL,15,'DaemonSetList','daemonSet/list','','k8s/workloads/daemonSet/list.vue','{\"title\": \"DaemonSet\"}',0),
	(23,'2023-06-08 02:23:35','2023-06-08 02:23:35',NULL,15,'DaemonSetDetail','daemonSet/detail','','k8s/workloads/daemonSet/detail.vue','{\"title\": \"DaemonSet详情\", \"hidden\": true}',0),
	(24,'2023-06-08 09:58:06','2023-06-08 09:58:06',NULL,0,'Services','/k8s/services','','Layout','{\"title\": \"服务\", \"svgIcon\": \"network\"}',0),
	(25,'2023-06-08 11:52:23','2023-06-12 02:33:34',NULL,24,'ServiceList','service/list','','k8s/service/services/list.vue','{\"title\": \"Service\"}',0),
	(26,'2023-06-12 06:39:50','2023-06-12 06:39:50',NULL,24,'ServiceDetail','service/detail','','k8s/service/services/detail.vue','{\"title\": \"Service详情\", \"hidden\": true}',0),
	(27,'2023-06-15 07:50:10','2023-06-15 07:50:37',NULL,24,'IngressList','ingress/list','','k8s/service/ingress/list.vue','{\"title\": \"Ingress\"}',0),
	(28,'2023-06-16 09:58:53','2023-06-16 09:58:53',NULL,24,'IngressDetail','ingress/detail','','k8s/service/ingress/detail.vue','{\"title\": \"Ingress详情\", \"hidden\": true}',0),
	(29,'2023-06-19 08:27:52','2023-06-19 08:36:22',NULL,0,'Config','/k8s/config','/k8s/config/configMapList','Layout','{\"title\": \"配置\", \"svgIcon\": \"config\"}',0),
	(30,'2023-06-20 01:53:04','2023-06-20 01:53:04',NULL,29,'ConfigMapList','configMap/list','','k8s/config/configMap/list.vue','{\"title\": \"ConfigMap\"}',0),
	(31,'2023-06-20 02:21:19','2023-06-20 02:21:19',NULL,29,'ConfigMapDetail','configMap/detail','','k8s/config/configMap/detail.vue','{\"title\": \"ConfigMap详情\", \"hidden\": true}',0),
	(32,'2023-06-20 09:50:09','2023-06-20 09:50:44',NULL,29,'SecretList','secret/list','','k8s/config/secret/list.vue','{\"title\": \"Secret\"}',0),
	(33,'2023-06-20 02:21:19','2023-06-20 02:21:19',NULL,29,'SecretDetail','secret/detail','','k8s/config/secret/detail.vue','{\"title\": \"Secret详情\", \"hidden\": true}',0),
	(34,'2023-06-25 07:21:30','2023-06-25 07:26:23',NULL,0,'AccessControl','/k8s/accessControl','/k8s/accessControl/serviceAccount/list','Layout','{\"title\": \"访问控制\", \"svgIcon\": \"accessControl\"}',0),
	(35,'2023-06-25 07:33:01','2023-06-25 07:33:01',NULL,34,'ServiceAccountList','serviceAccount/list','','k8s/accessControl/serviceAccount/list.vue','{\"title\": \"ServiceAccount\"}',0),
	(36,'2023-06-25 07:36:38','2023-06-25 08:03:44',NULL,34,'ServiceAccountDetail','serviceAccount/detail','','k8s/accessControl/serviceAccount/detail.vue','{\"title\": \"ServiceAccount详情\", \"hidden\": true}',0),
	(37,'2023-06-25 07:33:01','2023-06-25 07:33:01',NULL,34,'RoleList','role/list','','k8s/accessControl/role/list.vue','{\"title\": \"Role\"}',0),
	(38,'2023-06-25 07:36:38','2023-06-25 08:03:44',NULL,34,'RoleDetail','role/detail','','k8s/accessControl/role/detail.vue','{\"title\": \"Role详情\", \"hidden\": true}',0),
	(39,'2023-06-25 07:33:01','2023-06-25 07:33:01',NULL,34,'RoleBindingList','roleBinding/list','','k8s/accessControl/roleBinding/list.vue','{\"title\": \"RoleBinding\"}',0),
	(40,'2023-06-25 07:36:38','2023-06-25 08:03:44',NULL,34,'RoleBindingDetail','roleBinding/detail','','k8s/accessControl/roleBinding/detail.vue','{\"title\": \"RoleBinding详情\", \"hidden\": true}',0),
	(41,'2023-06-25 07:33:01','2023-06-25 07:33:01',NULL,34,'ClusterRoleList','clusterRole/list','','k8s/accessControl/clusterRole/list.vue','{\"title\": \"ClusterRole\"}',0),
	(42,'2023-06-25 07:36:38','2023-06-25 08:03:44',NULL,34,'ClusterRoleDetail','clusterRole/detail','','k8s/accessControl/clusterRole/detail.vue','{\"title\": \"ClusterRole详情\", \"hidden\": true}',0),
	(43,'2023-06-25 07:33:01','2023-06-25 07:33:01',NULL,34,'ClusterRoleBindingList','clusterRoleBinding/list','','k8s/accessControl/clusterRoleBinding/list.vue','{\"title\": \"ClusterRoleBinding\"}',0),
	(44,'2023-06-25 07:36:38','2023-06-25 08:03:44',NULL,34,'ClusterRoleBindingDetail','clusterRoleBinding/detail','','k8s/accessControl/clusterRoleBinding/detail.vue','{\"title\": \"ClusterRoleBinding详情\", \"hidden\": true}',0),
	(45,'2023-06-27 02:27:22','2023-06-27 02:29:39',NULL,0,'Resource','/k8s/resource','/k8s/resource/create','Layout','{\"title\": \"资源\", \"hidden\": true}',0),
	(46,'2023-06-27 02:28:55','2023-06-27 02:28:55',NULL,45,'CreateResource','create','','k8s/resource/base.vue','{\"title\": \"创建资源\", \"hidden\": true}',0);

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_name` (`role_name`),
  KEY `idx_sys_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`id`, `created_at`, `updated_at`, `deleted_at`, `role_name`)
VALUES
	(1,NULL,'2023-06-27 02:29:12',NULL,'root');

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '密码',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '邮箱',
  `active` tinyint(1) DEFAULT NULL,
  `role_model_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_sys_user_deleted_at` (`deleted_at`),
  KEY `idx_sys_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `phone`, `email`, `active`, `role_model_id`)
VALUES
	(1,'2023-02-20 12:51:58','2023-03-10 09:59:49',NULL,'admin','e10adc3949ba59abbe56e057f20f883e','11111111111','pddzl5@163.com',1,1);

/*!40000 ALTER TABLE `sys_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
