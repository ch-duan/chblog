使用超级用户登录数据库

创建用户test并使用密码123456

```sql
CREATE USER test WITH PASSWORD '123456';
```

将数据库 testDB 权限授权于 test

```sql
GRANT ALL PRIVILEGES ON DATABASE testDB TO test;
```

将当前数据库下 public schema 的表都授权于 test

```sql
GRANT ALL PRIVILEGES ON all tables in schema public TO test;
```

如果要单独一个权限以及单独一个表，则：

```sql
GRANT SELECT ON TABLE mytable TO test;
```
