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
-- Table structure for table `brand`
--

DROP TABLE IF EXISTS `brand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `brand` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(15) NOT NULL DEFAULT '' COMMENT '品牌名称',
  `logo_url` varchar(45) NOT NULL DEFAULT '' COMMENT 'logo地址',
  `maker` varchar(15) NOT NULL DEFAULT '' COMMENT '制造厂商',
  `sort` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `pinyin` varchar(20) NOT NULL DEFAULT '',
  `first` char(1) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `first` (`first`)
) ENGINE=InnoDB AUTO_INCREMENT=211 DEFAULT CHARSET=utf8mb4 COMMENT='汽车品牌';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `brand`
--

LOCK TABLES `brand` WRITE;
/*!40000 ALTER TABLE `brand` DISABLE KEYS */;
INSERT INTO `brand` VALUES (1,'阿尔法·罗密','','',0,'aerfaluomiou','A'),(2,'安驰','','',0,'anchi','A'),(3,'ARCFOX','','',0,'arcfox','A'),(4,'阿斯顿·马丁','','',0,'asidunmading','A'),(5,'奥迪','','',1,'aodi','A'),(6,'别克','','',2,'bieke','B'),(7,'奔驰','','',3,'benchi','B'),(8,'保时捷','','',0,'baoshijie','B'),(9,'宝龙','','',0,'baolong','B'),(10,'比亚迪','','',14,'biyadi','B'),(11,'宝马','','',4,'baoma','B'),(12,'标致','','',11,'biaozhi','B'),(13,'宾利','','',0,'binli','B'),(14,'北汽制造','','',0,'beiqizhizao','B'),(15,'北京汽车','','',0,'beijingqiche','B'),(16,'宝骏','','',0,'baojun','B'),(17,'巴博斯','','',0,'babosi','B'),(18,'宝沃','','',0,'baowo','B'),(19,'比速汽车','','',0,'bisuqiche','B'),(20,'北汽瑞丽','','',0,'beiqiruili','B'),(21,'本田','','',8,'bentian','B'),(22,'北汽新能源','','',0,'beiqixinnengyuan','B'),(23,'成功','','',0,'chenggong','C'),(24,'长安','','',0,'changan','C'),(25,'昌河','','',0,'changhe','C'),(26,'长安商用','','',0,'changanshangyong','C'),(27,'长城','','',0,'changcheng','C'),(28,'传祺','','',0,'chuanzuo','C'),(29,'大众','','',5,'dazhong','D'),(30,'大迪','','',0,'dadi','D'),(31,'大发','','',0,'dafa','D'),(32,'大宇','','',0,'dayu','D'),(33,'道奇','','',0,'daoqi','D'),(34,'东风汽车','','',0,'dongfengqiche','D'),(35,'东风风神','','',0,'dongfengfengshen','D'),(36,'DS','','',0,'ds','D'),(37,'电咖','','',0,'dianka','D'),(38,'大乘汽车','','',0,'dachengqiche','D'),(39,'东南','','',0,'dongnan','D'),(40,'东风风度','','',0,'dongfengfengdu','D'),(41,'大通','','',0,'datong','D'),(42,'东风小康','','',0,'dongfengxiaokang','D'),(43,'东风富康','','',0,'dongfengfukang','D'),(44,'东风风行','','',0,'dongfengfengxing','D'),(45,'东风俊风','','',0,'dongfengjunfeng','D'),(46,'菲亚特','','',0,'feiyate','F'),(47,'丰田','','',6,'fengtian','F'),(48,'福迪','','',0,'fudi','F'),(49,'富奇','','',0,'fuqi','F'),(50,'福田','','',0,'futian','F'),(51,'法拉利','','',0,'falali','F'),(52,'福特','','',7,'fute','F'),(53,'观致','','',0,'guanzhi','G'),(54,'GMC','','',0,'gmc','G'),(55,'光冈','','',0,'guanggang','G'),(56,'国金','','',0,'guojin','G'),(57,'广汽新能源','','',0,'guangqixinnengyuan','G'),(58,'哈飞','','',0,'hafei','H'),(59,'华普','','',0,'huapu','H'),(60,'华泰','','',0,'huatai','H'),(61,'黄海','','',0,'huanghai','H'),(62,'红旗','','',0,'hongqi','H'),(63,'汇众','','',0,'huizhong','H'),(64,'海格','','',0,'haige','H'),(65,'黑豹','','',0,'heibao','H'),(66,'恒天','','',0,'hengtian','H'),(67,'华北','','',0,'huabei','H'),(68,'幻速','','',0,'huansu','H'),(69,'华颂','','',0,'huasong','H'),(70,'汉腾汽车','','',0,'hantengqiche','H'),(71,'华骐','','',0,'huazuo','H'),(72,'红星汽车','','',0,'hongxingqiche','H'),(73,'海马','','',0,'haima','H'),(74,'哈弗','','',0,'hafu','H'),(75,'吉奥','','',0,'jiao','J'),(76,'吉普','','',0,'jipu','J'),(77,'江南','','',0,'jiangnan','J'),(78,'江苏南亚','','',0,'jiangsunanya','J'),(79,'捷豹','','',0,'jiebao','J'),(80,'金程','','',0,'jincheng','J'),(81,'江铃','','',0,'jiangling','J'),(82,'九龙','','',0,'jiulong','J'),(83,'解放','','',0,'jiefang','J'),(84,'金旅','','',0,'jinlv','J'),(85,'君马','','',0,'junma','J'),(86,'吉利','','',0,'jili','J'),(87,'几何汽车','','',0,'jiheqiche','J'),(88,'江淮','','',0,'jianghuai','J'),(89,'金杯','','',0,'jinbei','J'),(90,'江铃新能源','','',0,'jianglingxinnengyuan','J'),(91,'捷途','','',0,'jietu','J'),(92,'捷达','','',0,'jieda','J'),(93,'开瑞','','',0,'kairui','K'),(94,'凯迪拉克','','',0,'kaidilake','K'),(95,'克莱斯勒','','',0,'kelaisile','K'),(96,'卡尔森','','',0,'kaersen','K'),(97,'科尼赛克','','',0,'kenisaike','K'),(98,'凯翼','','',0,'kaiyi','K'),(99,'卡威','','',0,'kawei','K'),(100,'铃木','','',18,'lingmu','L'),(101,'路虎','','',0,'luhu','L'),(102,'猎豹','','',0,'liebao','L'),(103,'兰博基尼','','',0,'lanbojini','L'),(104,'雷诺','','',0,'leinuo','L'),(105,'路特斯','','',0,'lutesi','L'),(106,'莲花','','',0,'lianhua','L'),(107,'陆风','','',0,'lufeng','L'),(108,'罗孚','','',0,'luozuo','L'),(109,'劳伦士','','',0,'laolunshi','L'),(110,'陆地方舟','','',0,'ludifangzhou','L'),(111,'理念','','',0,'linian','L'),(112,'领克','','',0,'lingke','L'),(113,'罗夫哈特','','',0,'luofuhate','L'),(114,'力帆','','',0,'lifan','L'),(115,'领途汽车','','',0,'lingtuqiche','L'),(116,'劳斯莱斯','','',0,'laosilaisi','L'),(117,'雷克萨斯','','',0,'leikesasi','L'),(118,'林肯','','',0,'linken','L'),(119,'零跑汽车','','',0,'lingpaoqiche','L'),(120,'洛克汽车','','',0,'luokeqiche','L'),(121,'理想','','',0,'lixiang','L'),(122,'雷丁','','',0,'leiding','L'),(123,'MINI','','',0,'mini','M'),(124,'MG','','',19,'mg','M'),(125,'玛莎拉蒂','','',0,'mashaladi','M'),(126,'迈巴赫','','',0,'maibahe','M'),(127,'美亚','','',0,'meiya','M'),(128,'马自达','','',12,'mazida','M'),(129,'迈凯伦','','',0,'maikailun','M'),(130,'纳智捷','','',0,'nazhijie','N'),(131,'哪吒汽车','','',0,'nazuoqiche','N'),(132,'欧宝','','',0,'oubao','O'),(133,'欧朗','','',0,'oulang','O'),(134,'讴歌','','',0,'zuoge','O'),(135,'欧拉','','',0,'oula','O'),(136,'欧尚汽车','','',0,'oushangqiche','O'),(137,'Polestar','','',0,'polestar','P'),(138,'帕加尼','','',0,'pajiani','P'),(139,'奇瑞','','',0,'qirui','Q'),(140,'起亚','','',17,'qiya','Q'),(141,'庆铃','','',0,'qingling','Q'),(142,'启腾','','',0,'qiteng','Q'),(143,'乔治·巴顿','','',0,'qiaozhi%C2%B7badun','Q'),(144,'前途汽车','','',0,'qiantuqiche','Q'),(145,'启辰','','',0,'qichen','Q'),(146,'日产','','',9,'richan','R'),(147,'瑞麒','','',0,'ruizuo','R'),(148,'RUF','','',0,'ruf','R'),(149,'荣威','','',20,'rongwei','R'),(150,'Smart','','',0,'smart','S'),(151,'思铭','','',0,'siming','S'),(152,'斯巴鲁','','',0,'sibalu','S'),(153,'萨博','','',0,'sabo','S'),(154,'赛宝','','',0,'saibao','S'),(155,'双环','','',0,'shuanghuan','S'),(156,'双龙','','',0,'shuanglong','S'),(157,'世爵','','',0,'shijue','S'),(158,'陕汽通家','','',0,'shanqitongjia','S'),(159,'斯威汽车','','',0,'siweiqiche','S'),(160,'三菱','','',0,'sanling','S'),(161,'斯柯达','','',10,'sikeda','S'),(162,'绅宝','','',0,'shenbao','S'),(163,'天马','','',0,'tianma','T'),(164,'通田','','',0,'tongtian','T'),(165,'腾势','','',0,'tengshi','T'),(166,'天津一汽','','',0,'tianjinyiqi','T'),(167,'特斯拉','','',0,'tesila','T'),(168,'万丰','','',0,'wanfeng','W'),(169,'威麟','','',0,'weizuo','W'),(170,'沃尔沃','','',0,'woerwo','W'),(171,'五菱','','',0,'wuling','W'),(172,'威旺','','',0,'weiwang','W'),(173,'威兹曼','','',0,'weiziman','W'),(174,'五十铃','','',0,'wushiling','W'),(175,'WEY','','',0,'wey','W'),(176,'威马','','',0,'weima','W'),(177,'蔚来','','',0,'weilai','W'),(178,'新雅途','','',0,'xinyatu','X'),(179,'厦门金龙','','',0,'xiamenjinlong','X'),(180,'雪佛兰','','',16,'xuefolan','X'),(181,'雪铁龙','','',15,'xuetielong','X'),(182,'西雅特','','',0,'xiyate','X'),(183,'新凯','','',0,'xinkai','X'),(184,'新特汽车','','',0,'xinteqiche','X'),(185,'现代','','',13,'xiandai','X'),(186,'小鹏汽车','','',0,'xiaopengqiche','X'),(187,'星途','','',0,'xingtu','X'),(188,'英菲尼迪','','',0,'yingfeinidi','Y'),(189,'永源','','',0,'yongyuan','Y'),(190,'云雀','','',0,'yunque','Y'),(191,'一汽吉林','','',0,'yiqijilin','Y'),(192,'一汽华利','','',0,'yiqihuali','Y'),(193,'依维柯','','',0,'yiweike','Y'),(194,'一汽通用','','',0,'yiqitongyong','Y'),(195,'云度汽车','','',0,'yunduqiche','Y'),(196,'裕路','','',0,'yulu','Y'),(197,'御捷新能源','','',0,'yujiexinnengyuan','Y'),(198,'一汽奔腾','','',0,'yiqibenteng','Y'),(199,'野马','','',0,'yema','Y'),(200,'英致','','',0,'yingzhi','Y'),(201,'易至汽车','','',0,'yizhiqiche','Y'),(202,'银隆新能源','','',0,'yinlongxinnengyuan','Y'),(203,'中华','','',0,'zhonghua','Z'),(204,'中顺','','',0,'zhongshun','Z'),(205,'中兴','','',0,'zhongxing','Z'),(206,'众泰','','',0,'zhongtai','Z'),(207,'中欧','','',0,'zhongou','Z'),(208,'知豆','','',0,'zhidou','Z'),(209,'浙江卡尔森','','',0,'zhejiangkaersen','Z'),(210,'之诺','','',0,'zhinuo','Z');
/*!40000 ALTER TABLE `brand` ENABLE KEYS */;
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
