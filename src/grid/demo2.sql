-- MySQL dump 10.13  Distrib 8.0.23, for osx10.16 (x86_64)
--
-- Host: 127.0.0.1    Database: test
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article`
--

DROP TABLE IF EXISTS `article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT '1000-01-01 00:00:01',
  `updated_at` datetime NOT NULL DEFAULT '1000-01-01 00:00:01',
  `deleted_at` datetime DEFAULT '1000-01-01 00:00:01',
  `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` longtext,
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `add_time` datetime DEFAULT '1000-01-01 00:00:01',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COMMENT='文章管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article`
--

LOCK TABLES `article` WRITE;
/*!40000 ALTER TABLE `article` DISABLE KEYS */;
INSERT INTO `article` VALUES (1,'2021-03-11 11:13:51','2021-03-19 14:00:02','2021-03-11 14:46:32',1,'test','add for testing','modified by Edit',1,'1000-01-01 00:00:01'),(2,'2021-03-11 11:13:52','2021-03-19 14:00:02',NULL,1,'test','add for testing','modified by Edit',1,'1000-01-01 00:00:01'),(18,'2021-03-11 11:14:51','2021-03-19 14:00:02',NULL,1,'test','add for testing','modified by Edit',1,'1000-01-01 00:00:01'),(19,'2021-03-11 11:17:29','2021-03-19 14:00:03',NULL,1,'test','add for testing','modified by Edit',1,'1000-01-01 00:00:01'),(20,'2021-03-11 15:32:48','2021-03-19 14:00:03',NULL,1,'test','add for testing','modified by Edit',1,'1000-01-01 00:00:01'),(26,'2021-03-19 13:07:01','2021-03-19 14:00:03',NULL,1,'test','add for testing','modified by Edit',1,'2021-03-19 13:07:01');
/*!40000 ALTER TABLE `article` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-19 14:30:41
