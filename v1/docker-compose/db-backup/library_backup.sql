-- MariaDB dump 10.19  Distrib 10.7.1-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: library
-- ------------------------------------------------------
-- Server version	10.7.1-MariaDB-1:10.7.1+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fullname` varchar(150) NOT NULL,
  `email` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(300) DEFAULT NULL,
  `roles` set('user','admin') NOT NULL DEFAULT 'user',
  `avatar` varchar(200) DEFAULT NULL,
  `created_at` int(11) unsigned NOT NULL DEFAULT 0,
  `updated_at` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES
(13,'Vallin Plasencia Valdes','vallin.plasencia@gmail.com','vallin.plasencia_6npeu','$2a$04$avTesDT8R8fsmP2dxIWN/uurw3Lojv4xAyPy5lcKf9ky1.kd/PxC2','user,admin','/upload/media/avatars/2021/december/5/PJnUiQFScrViNM7gP.jpeg',1638739703,0),
(14,'Vallin Plasencia Valdes','xvallin.plasencia@gmail.com','xvallin.plasencia_wjal6','$2a$04$KUqGmhMnaKcycakh5RNhI.KNtMJYn/gcTXujdXFuYjNHSlLnV2dXe','user','xxxvvv/media/avatars/2021/december/6/2ORo1w8dGVqVsZ8IX.jpeg',1638810256,0),
(15,'Vallin Plasencia Valdes','avallin.plasencia@gmail.com','avallin.plasencia_yxdvs','$2a$04$keyDWASpYPB2cWAv7MrEiOer.nWGfKj3dmyZEYN0cbZACl/9DRJrW','user','/xxxvvv/media/avatars/2021/december/6/rPZhycAV5Z5cNZTPk.jpeg',1638811309,0);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `original` bit(1) DEFAULT b'0',
  `tags` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`tags`)),
  `published_at` int(11) unsigned NOT NULL DEFAULT 0,
  `created_at` int(11) unsigned NOT NULL DEFAULT 0,
  `updated_at` int(11) unsigned NOT NULL DEFAULT 0,
  `category_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `books_FK` (`category_id`),
  KEY `books_FK_1` (`user_id`),
  CONSTRAINT `books_FK` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE,
  CONSTRAINT `books_FK_1` FOREIGN KEY (`user_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES
(15,'Libro 1','','[\"uno\",\"dos\"]',123,1639489933,0,1,13),
(16,'xLibro 1.1','\0','[\"uno\",\"dos\"]',1231,1639490119,1,2,13);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` varchar(200) DEFAULT NULL,
  `created_at` int(11) unsigned NOT NULL DEFAULT 0,
  `updated_at` int(11) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `categories_UN` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES
(1,'terror','categoria terr bla bla bla',11111,0),
(2,'humor','categoria hum bla bla bla',22222,21111);
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `refresh_token` varchar(200) NOT NULL,
  `device_id` varchar(20) NOT NULL,
  `useragent_str` varchar(200) NOT NULL,
  `useragent` varchar(50) NOT NULL,
  `platform` varchar(50) NOT NULL,
  `ip` varchar(40) NOT NULL,
  `location` varchar(150) DEFAULT NULL,
  `last_access_token_generated_at` int(11) unsigned NOT NULL,
  `created_at` int(11) unsigned NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `sessions_FK` (`user_id`),
  CONSTRAINT `sessions_FK` FOREIGN KEY (`user_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sessions`
--

LOCK TABLES `sessions` WRITE;
/*!40000 ALTER TABLE `sessions` DISABLE KEYS */;
INSERT INTO `sessions` VALUES
(3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxOTA5ODksInVzZXJfaWQiOiIxMyJ9.kMI1zdQRzxUrX7OjM-58NMin0mnLNRvEGswBMxzaFOA','ZTVgyLV1638994189','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1638994189,1638994189,13),
(4,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxOTEwMDIsInVzZXJfaWQiOiIxMyJ9.Ie2gwLD73rYbp7jJrfaDqYgdH4qIpyOAKO8GifrIBeY','IyRMTls1638994202','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1638994202,1638994202,13),
(5,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDIxOTEzMzMsInVzZXJfaWQiOiIxMyJ9.JDqB7lJNFf6zByLCGfG3klauY2PzXWfocHsvZKLFhB8','DkPK0ie1638994533','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1638994533,1638994533,13),
(6,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MDY3NDAsInVzZXJfaWQiOiIxMyJ9.XZS7EZeyQoLpGNpg7ewiPCryP6lILUr4-YT06OLfj5g','kjiHhKJ1639409940','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639409940,1639409940,13),
(7,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MDc4ODcsInVzZXJfaWQiOiIxMyJ9.zE_kWSf-g_RMapQ4IxgJPFWOYyRy2iiUV4QCyQkc5Mw','oWv3d8S1639411087','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639411087,1639411087,13),
(8,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MDc5MjEsInVzZXJfaWQiOiIxMyJ9.j9VYis32vhuS1SYZMMxJLLC8fbcZiLwtHbul7PHIplE','UG65fS71639411121','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639411121,1639411121,13),
(9,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MDgyMDksInVzZXJfaWQiOiIxMyJ9.n04DVEroscbB9CgM4MVbeORbT-0wVjYYMUleH8WY1Pc','SKPnsSj1639411409','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639411409,1639411409,13),
(10,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MTAzNjQsInVzZXJfaWQiOiIxMyJ9._V0mxdjXNYT7akuLOlwLGnKJ4hRgg2WNjvcz1lYOwm0','K98mJWx1639413564','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639413564,1639413564,13),
(11,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MTkzMjQsInVzZXJfaWQiOiIxMyJ9.ZWvV1J-o9Ov6mGne_zXyS87pjcq_BTdWUMe37EUdaE4','jH8Uei81639422524','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639422524,1639422524,13),
(12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MjE0MDcsInVzZXJfaWQiOiIxMyJ9.O-7qx2neF2_cMqleQvEPudqIQ5Je22zMcH_eSRHJQFQ','EzJDhZ81639424607','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639424607,1639424607,13),
(13,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2MjE0MTcsInVzZXJfaWQiOiIxMyJ9.pQu2wfmF20ByKYb3CH6ssfxVLzTSTHZ8tX2y_amAZaY','pCWj2ko1639424617','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639424617,1639424617,13),
(14,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2ODY2NjQsInVzZXJfaWQiOiIxMyJ9.hsW6p0OG6XWWn1mUOdNpnsM2PuxZWXFBTOQ3q9EmStU','WDSPFr51639489864','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639489864,1639489864,13),
(15,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2ODkwMjMsInVzZXJfaWQiOiIxMyJ9.EFXGxvHEmnR_OM4MXAWguMd1wU55oAuPrM5I21dXtK4','nqZY1gN1639492223','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639492223,1639492223,13),
(16,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2OTExNDIsInVzZXJfaWQiOiIxMyJ9.fMEaOaCfC8mnXwh5v_Zj51_zIJtc16EExIcUpmAdspY','LdMrHzs1639494342','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639494342,1639494342,13),
(17,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2OTY4MzAsInVzZXJfaWQiOiIxMyJ9.7HwW6tscKztS5IeSiLSgp8w8McrS-_BbvVp8OO4ePzs','4wqpYJx1639500030','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639500030,1639500030,13),
(18,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI2OTg5NDksInVzZXJfaWQiOiIxMyJ9.EwL-_Z2Bla7dOjO9tB4cZp1QnLt1WkMSUQAD9Pa7wVE','ZktAUda1639502149','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639502149,1639502149,13),
(19,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI3ODM0OTIsInVzZXJfaWQiOiIxMyJ9.Xd9_3aEw5Duoh1TxCqsB3IQ89Hb9mhNGBWVFr2FCjgs','feR3xNC1639586692','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639586692,1639586692,13),
(20,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI3ODc2MTEsInVzZXJfaWQiOiIxMyJ9.XpBPmASU2tvCnq7kLA0JYcMVYqkyUxWMMjvqO7pYw3g','BVrDboD1639590811','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639590811,1639590811,13),
(21,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI4NDk5NzcsInVzZXJfaWQiOiIxMyJ9.LL3oV_06Q9R9LqbnOws3X46fdhoSFzH-HxC-hm4XFQ0','JQjERkJ1639653177','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639653177,1639653177,13),
(22,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI4NTkzOTQsInVzZXJfaWQiOiIxMyJ9.MDXlqjnmfc8_YiDAOwd5EHY9JS9I4DHigeR7A90Znvk','N3UmviL1639662594','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639662594,1639662594,13),
(23,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI4NjExMzIsInVzZXJfaWQiOiIxMyJ9.NQQDNq03NcrMLJ0q0QuerzNxSfz3gr8ZZw3eYdb1v7I','NU5klFp1639664332','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1639664332,1639664332,13),
(24,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDMzODMxMjIsInVzZXJfaWQiOiIxMyJ9.ojO_hrlrVdx_kGLMr22bIxRM6u75qA_PKSaINPP7YhI','BrXZEAN1640186322','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1640186322,1640186322,13),
(25,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQwMTg4MTAsInVzZXJfaWQiOiIxMyJ9.dFBqIECe-PlEjW62l5FhSe2QRLJWb9N3O6zgAhvzsIo','WGyrSDJ1640822010','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1640822010,1640822010,13),
(26,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQwMjA3NjQsInVzZXJfaWQiOiIxMyJ9.QDM6gu9NKmW5Iu4DceXcQ6huB1RfjAIKiNd0zb-ouN0','SnzIn9Z1640823964','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1640823964,1640823964,13),
(27,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQxNjc5NzQsInVzZXJfaWQiOiIxMyJ9.ng3xxKIkKA9A83OwuRhM-qWGUvIuSVcwCPswULkJCzI','uFfhEk41640971174','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1640971174,1640971174,13),
(28,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQxNjk0MzksInVzZXJfaWQiOiIxMyJ9.7zDKdH-PvdTOeT8VSQ7ltmfL8VOXFIRaYT99BeHsDRI','MEEinBe1640972639','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1640972639,1640972639,13),
(29,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQzMjU4NDAsInVzZXJfaWQiOiIxMyJ9.v3MATayODy6XrF5oinIywaHOwOSszsxdMN_guZ4leUM','OauvPfu1641129040','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641129040,1641129040,13),
(30,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQzMjU4NTgsInVzZXJfaWQiOiIxMyJ9.nNH0rU3zakwWSBR8CA15hPlJGYQKPy3_dPIAuIALqzc','rL2I3lB1641129058','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641129058,1641129058,13),
(31,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQzMjg1NzUsInVzZXJfaWQiOiIxMyJ9.S7clGIbdjIbG7baIIJWpjjRzuN3EnifCdBlTMmJ0rSw','H8MfXYq1641131775','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641131775,1641131775,13),
(32,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQzMzAwNTYsInVzZXJfaWQiOiIxMyJ9.6BXfJYaM8BFMaEahAnJvIVD5vqUk3tkiXI0rslXM_Vw','CBoVQUA1641133256','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641133256,1641133256,13),
(33,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQzNDc4ODQsInVzZXJfaWQiOiIxMyJ9.vPNlyiBG8c7sbxKzG2bCHCuKQnjO_hB4y9pnlDvAaTM','UPC4UXZ1641151084','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641151084,1641151084,13),
(34,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ1MzAzOTUsInVzZXJfaWQiOiIxNCJ9.tBuwytJcUVAWMiFoqNmmgFZ_U3ELuivEyqrNiUNNiMA','9xvR9iH1641333595','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641333595,1641333595,14),
(35,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ1NDcxODQsInVzZXJfaWQiOiIxMyJ9.coylHsXPWAC_bbcQrnoOHvTNrLZUZT9qbXmYfxTzPJk','5CBo7Wu1641350384','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36','Chrome','Linux','','country-city-town',1641350384,1641350384,13),
(36,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2ODExMTAsInVzZXJfaWQiOiIxMyJ9.xSgLJh9VR0xtD63yuEsQYEmPW8oYDSOToYpdkaLUuBM','YnuXEX71641484310','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36','Chrome','Linux','','country-city-town',1641484517,1641484310,13),
(37,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2ODI5NjksInVzZXJfaWQiOiIxMyJ9.JHVQGlY7ptRuDS8tgxPpYYVSA4zaDco95bGs523bjy4','tUqmLom1641486169','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36','Chrome','Linux','','country-city-town',1641486169,1641486169,13),
(38,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2ODMwNDMsInVzZXJfaWQiOiIxMyJ9.Pkcw_smL_BKYilz6a_C7lRXa5b-aGNXoz1jW4siAHfc','w2iX4pc1641486243','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36','Chrome','Linux','','country-city-town',1641486243,1641486243,13),
(39,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2ODMxNDAsInVzZXJfaWQiOiIxMyJ9.rDBSvDWWKKjc4qvRBnCs0d6c1jOVwFoN-bdc8JnT35Q','tZOaFGM1641486340','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36','Chrome','Linux','','country-city-town',1641486340,1641486340,13),
(40,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2ODU1MzEsInVzZXJfaWQiOiIxMyJ9.rYILFONtXJ6beWIdMLMOVzIsZnVAdsD1BnMR7BVnh8o','fV4dVo51641488731','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36','Chrome','Linux','','country-city-town',1641488731,1641488731,13),
(41,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDQ2OTIyMjksInVzZXJfaWQiOiIxMyJ9.b52P3rAwEAPqeGCMXue2ovi0S0iOb865Whtb63-t7Qc','QmMkv1c1641495429','Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36','Chrome','Linux','','country-city-town',1641495429,1641495429,13);
/*!40000 ALTER TABLE `sessions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-06 21:22:46
