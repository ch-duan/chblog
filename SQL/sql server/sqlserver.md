查询数据库

SELECT name, database_id, create_date FROM sys.databases ;（一般sql语句应该是show databases;）

查询数据库中的表

SELECT * from sysobjects where xtype='U';(一般SQL语句应该是show tabels;)
