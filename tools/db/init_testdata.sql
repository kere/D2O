truncate users RESTART IDENTITY;
truncate tags RESTART IDENTITY;
truncate formfields RESTART IDENTITY;
truncate s_elems RESTART IDENTITY;
truncate images RESTART IDENTITY;

insert into users (iid, nick, token, status) values (999, 'd2o', 'abctoken', 1);

insert into tags ("user_id", "name", "cata_id") values (0, 'tag1', 1),(0, 'tag2', 1),(0, 'tag3', 2);
insert into formfields ("user_id", "name", "itype") values (0, 'label1', 0),(0, 'label2', 0),(0, 'label3', 0);
