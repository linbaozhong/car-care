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
-- Table structure for table `dict`
--

DROP TABLE IF EXISTS `dict`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dict` (
  `type` smallint(5) unsigned NOT NULL DEFAULT '0',
  `key` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`type`,`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dict`
--

LOCK TABLES `dict` WRITE;
/*!40000 ALTER TABLE `dict` DISABLE KEYS */;
INSERT INTO `dict` VALUES (0,1,'汽车外观检查项目'),(0,2,'汽车内饰检查项目'),(0,3,'汽车仪表检查项目'),(0,4,'汽车其它检查项目'),(0,5,'汽车上线检查项目'),(0,101,'特效膜'),(0,102,'车衣'),(0,103,'太阳膜'),(0,104,'改色膜'),(0,105,'养护'),(0,120,'清洗'),(1,1,'前保险杠'),(1,2,'前引擎盖'),(1,3,'前挡风玻璃'),(1,4,'左前翼子板'),(1,5,'左前轮毂'),(1,6,'左后视镜'),(1,7,'左前门'),(1,8,'左裙边'),(1,9,'左后门'),(1,10,'后保险杠'),(1,11,'左后轮毂'),(1,12,'左后翼子板'),(1,13,'后备箱盖'),(1,14,'右后翼子板'),(1,15,'右后轮毂'),(1,16,'右后门'),(1,17,'右裙边'),(1,18,'右前门'),(1,19,'右前翼子板'),(1,20,'右前轮毂'),(1,21,'右后视镜'),(2,1,'正驾遮阳板'),(2,2,'仪表盘'),(2,3,'方向盘'),(2,4,'中控'),(2,5,'中央扶手箱'),(2,6,'中控玻璃升降开关'),(2,7,'左前内饰门板'),(2,8,'左前座椅'),(2,9,'左后玻璃升降开关'),(2,10,'左后内饰门板'),(2,11,'左后座椅'),(2,12,'右后玻璃升降开关'),(2,13,'右后内饰门板'),(2,14,'右后座椅'),(2,15,'右前玻璃升降开关'),(2,16,'右前内饰门板'),(2,17,'右前座椅'),(2,18,'副驾遮阳板'),(2,19,'车顶'),(3,1,'驻车指示灯'),(3,2,'ABS指示灯'),(3,3,'车身电子稳定系统'),(3,4,'电瓶指示灯'),(3,5,'机油指示灯'),(3,6,'水温指示灯'),(3,7,'安全气囊指示灯'),(3,8,'发动机故障指示灯'),(3,9,'燃油指示灯'),(3,10,'EPC指示灯'),(3,11,'胎压指示灯'),(3,12,'车辆维修指示灯'),(3,13,'灯泡损坏故障灯'),(3,14,'刹车片磨损指示灯'),(4,1,'驾驶证'),(4,2,'备胎'),(4,3,'随车工具'),(4,4,'贵重物品'),(4,5,'车险到期日期'),(4,6,'投保车险公司'),(5,1,'胎压监测显示'),(5,2,'测试蓄电池状态'),(5,3,'制动盘外观'),(5,4,'制动片磨损'),(5,5,'轮胎检查'),(5,6,'车轮轮毂外观检查'),(5,7,'减震器外观检查'),(5,8,'减震器防尘罩缓冲块');
/*!40000 ALTER TABLE `dict` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-10-27 21:22:36
