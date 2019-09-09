truncate users RESTART IDENTITY;
truncate tags RESTART IDENTITY;
truncate formfields RESTART IDENTITY;
truncate s_elems RESTART IDENTITY;
truncate images RESTART IDENTITY;

insert into users (iid, nick, token, status) values (999, 'd2o', '4297f44b13955235245b2497399d7a93', 1);

insert into tags ("iid", "user_id", "name", "cata_id") values (1, 0, 'tag1', 1),(2, 0, 'tag2', 1),(3, 0, 'tag3', 2);
insert into formfields ("iid", "user_id", "name", "itype") values (1, 0, '参考连接', 20),(2, 0, '数据统计', 0),(3, 0, 'label3', 0);
