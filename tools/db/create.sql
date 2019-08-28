drop index i_mdatas_mid;

drop table m_datas;

drop index i_mele_dateon;

drop index i_mele_reles;

drop index i_mele_tags;

drop table m_elems;

drop index i_notice;

drop table notices;

drop table objects;

drop index i_sele_miid;

drop index i_sele_reles;

drop index i_sele_tags;

drop table s_elems;

drop index u_uid;

drop table users;

/*==============================================================*/
/* Table: m_datas                                               */
/*==============================================================*/
create table m_datas (
   m_iid                INT8                 not null,
   lang                 INT4                 not null,
   title                VARCHAR(200)         not null,
   text                 TEXT                 not null
);

/*==============================================================*/
/* Index: i_mdatas_mid                                          */
/*==============================================================*/
create  index i_mdatas_mid on m_datas (
m_iid,
lang
);

/*==============================================================*/
/* Table: m_elems                                               */
/*==============================================================*/
create table m_elems (
   iid                  INT8                 not null,
   date_on              DATE                 not null,
   tags                 INT4[]               null,
   reles                INT4[]               null,
   o_json               JSONB                null,
   itype                INT4                 null default 0,
   constraint PK_M_ELEMS primary key (iid)
);

comment on column m_elems.reles is
'相关联IDs';

/*==============================================================*/
/* Index: i_mele_tags                                           */
/*==============================================================*/
create  index i_mele_tags on m_elems using GIN (
tags
);

/*==============================================================*/
/* Index: i_mele_reles                                          */
/*==============================================================*/
create  index i_mele_reles on m_elems using GIN (
reles
);

/*==============================================================*/
/* Index: i_mele_dateon                                         */
/*==============================================================*/
create  index i_mele_dateon on m_elems (
date_on
);

/*==============================================================*/
/* Table: notices                                               */
/*==============================================================*/
create table notices (
   id                   SERIAL               not null,
   user_id              INT4                 not null,
   o_json               JSONB                not null,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_NOTICES primary key (id)
);

/*==============================================================*/
/* Index: i_notice                                              */
/*==============================================================*/
create  index i_notice on notices (
user_id,
created_at
);

/*==============================================================*/
/* Table: objects                                               */
/*==============================================================*/
create table objects (
   iid                  INT8                 not null,
   text                 TEXT                 not null,
   itype                INT4                 not null default 0,
   reles                INT4[]               null,
   o_json               JSONb                not null,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_OBJECTS primary key (iid)
);

comment on table objects is
'itype:
1 people
2 country
';

comment on column objects.reles is
'相关联IDs';

/*==============================================================*/
/* Table: s_elems                                               */
/*==============================================================*/
create table s_elems (
   iid                  INT8                 not null,
   m_iid                INT8                 not null default 0,
   date_on              DATE                 not null,
   user_id              INT4                 not null,
   o_json               JSONB                not null,
   tags                 INT4[]               null,
   reles                INT4[]               null,
   review_json          JSONb                null default '0',
   status               INT2                 not null default 0,
   itype                INT2                 not null default 0,
   updated_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_S_ELEMS primary key (iid)
);

comment on table s_elems is
'编辑状态的内容
';

comment on column s_elems.m_iid is
'主表ID，如果没有，默认为0';

comment on column s_elems.reles is
'相关联IDs';

/*==============================================================*/
/* Index: i_sele_tags                                           */
/*==============================================================*/
create  index i_sele_tags on s_elems using GIN (
tags
);

/*==============================================================*/
/* Index: i_sele_reles                                          */
/*==============================================================*/
create  index i_sele_reles on s_elems using GIN (
rele
);

/*==============================================================*/
/* Index: i_sele_miid                                           */
/*==============================================================*/
create  index i_sele_miid on s_elems (
user_id,
m_iid
);

/*==============================================================*/
/* Table: users                                                 */
/*==============================================================*/
create table users (
   id                   SERIAL not null,
   iid                  INT8                 not null,
   nick                 VARCHAR(200)         not null default '',
   token                VARCHAR(64)          not null,
   avatar               VARCHAR(40)          null default '',
   o_json               JSONB                null,
   status               INT4                 not null default 0,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   updated_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_USERS primary key (id)
);

/*==============================================================*/
/* Index: u_uid                                                 */
/*==============================================================*/
create unique index u_uid on users (
iid
);
