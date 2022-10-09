# mysql

* 登录
  mysql -h host -P port -u username -p password

# 索引

MySQL数据库索引查询

1. 创建数据库和表
   CREATE DATABASE test;
   CREATE table writers (w_id int(10) PRIMARY key not null AUTO_INCREMENT COMMENT ‘编号’,
   w_name VARCHAR(20) NOT NULL COMMENT ‘作者姓名’,
   w_address VARCHAR(50) COMMENT ‘作者地址’,
   w_age int not null COMMENT ‘年龄’,
   w_note TEXT COMMENT ‘说明’,
   UNIQUE INDEX uniqidx (w_id));
2. 使用alter table 语句在w_name字段上建立名为nameidx的普通索引.
   ALTER table writers ADD INDEX nameidx(w_name);
3. 使用create index语句在w_address和w_age字段上建立名为multiidx的组合索引.
   CREATE INDEX multiidx on writers(w_address,w_age);
4. 使用CREATE index语句在 w_note字段上建立名为ftidx的全文索引.
   CREATE FULLTEXT INDEX ftidx on writers(w_note);
   1
5. 删除索引。利用alter table 语句将全文索引ftidx删除.
   alter table writers drop INDEX ftidx;
   利用drop index语句将nameidx索引删除.
   drop INDEX nameidx on writers;
