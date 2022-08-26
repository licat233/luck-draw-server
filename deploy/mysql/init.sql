
----------------
-- Mysql创建luck_draw用户
----------------
create user 'luck_draw'@'%' identified by '123123123';

---------------------
-- 将luck_draw用户绑定到luck_draw数据库上
---------------------
grant all privileges on luck_draw.* to 'luck_draw'@'%';

------------------
-- 赋予增删改查的权限
------------------
grant select,delete,update,create,drop on luck_draw.* to 'luck_draw'@'%';

----------
-- 刷新权限
----------
flush privileges;
