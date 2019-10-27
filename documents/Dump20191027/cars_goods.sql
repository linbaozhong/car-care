-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: localhost    Database: cars
-- ------------------------------------------------------
-- Server version	5.7.15-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goods`
--

DROP TABLE IF EXISTS `goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '商品分类',
  `name` varchar(45) NOT NULL DEFAULT '' COMMENT '商品名称',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '价格',
  `descr` varchar(45) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `type` (`type`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COMMENT='商品与服务';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goods`
--

LOCK TABLES `goods` WRITE;
/*!40000 ALTER TABLE `goods` DISABLE KEYS */;
INSERT INTO `goods` VALUES (1,105,'内饰真皮养护',48000,'不含深度清洁(40分钟)'),(2,105,'前风挡玻璃镀膜',65800,'清洗\\玻璃过泥\\镀膜(3小时)'),(3,105,'漆面轻度养护',32800,'打蜡\\精洗(2.5小时)'),(4,105,'漆面中度养护',128800,'抛光\\还原\\精洗\\打蜡(5小时)'),(5,105,'漆面高级养护',238800,'抛光\\还原\\精洗\\镀晶'),(6,104,'VF',450000,'起'),(7,104,'车衣裳(CYS)',500000,'起'),(8,104,'3M',750000,'起'),(9,104,'日本电镀STM',900000,'起'),(10,104,'艾丽',600000,'起'),(11,104,'法国骇克斯',800000,'起'),(12,103,'前档-3M',150000,'起'),(13,103,'前档-圣科',250000,'起'),(14,103,'前档-龙膜',300000,'起'),(15,103,'前档-威固',400000,'起'),(16,103,'侧挡-3M',100000,'起'),(17,103,'侧挡-圣科',150000,'起'),(18,103,'侧挡-龙膜',200000,'起'),(19,103,'侧挡-威固',300000,'起'),(20,102,'局部透明车衣',100000,'起/每米'),(21,102,'整车透明车衣-STEK',1500000,'起'),(22,102,'整车透明车衣-国内产品',800000,'起'),(23,101,'灯膜',158000,'起/没对'),(24,101,'局部特效膜',300000,'起'),(25,101,'内饰膜雕刻',138000,'起'),(26,120,'精美洗车',15800,'用时40分钟'),(27,120,'深度精洗',26800,'用时80分钟'),(28,120,'内饰除甲醛',39800,'含精洗,用时90分钟'),(29,120,'内饰深度清洁',88000,'用时4.5小时'),(30,120,'发动机深度清洁',26800,'用时90分钟'),(31,120,'轮毂深度清洁',23800,'用时70分钟');
/*!40000 ALTER TABLE `goods` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-27 21:22:35
