drop index i_hdata_mid;

drop table h_datas;

drop index i_hele_usr;

drop table h_eles;

drop index i_mdatas_mid;

drop table m_datas;

drop index i_mele_rele;

drop index i_mele_tags;

drop table m_eles;

drop index i_notice;

drop table notices;

drop table objects;

drop index i_sele_uid;

drop index i_sele_rele;

drop index i_sele_tags;

drop table s_eles;

drop index u_uid;

drop table users;

/*==============================================================*/
/* Table: h_datas                                               */
/*==============================================================*/
create table h_datas (
   m_iid                INT8                 not null,
   lang                 INT4                 not null,
   title                VARCHAR(200)         not null,
   text                 TEXT                 not null
);

/*==============================================================*/
/* Index: i_hdata_mid                                           */
/*==============================================================*/
create  index i_hdata_mid on h_datas (
m_iid,
lang
);

/*==============================================================*/
/* Table: h_eles                                                */
/*==============================================================*/
create table h_eles (
   iid                  INT8                 not null,
   user_id              INT4                 not null,
   date_on              DATE                 not null,
   tags                 INT4[]               null,
   reles                 INT4[]               null,
   o_json               JSONB                null,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_H_ELES primary key (iid)
);

comment on table h_eles is
'审核提交后，的历史数据';

comment on column h_eles.reles is
'相关联IDs';

/*==============================================================*/
/* Index: i_hele_usr                                            */
/*==============================================================*/
create  index i_hele_usr on h_eles (

);

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
/* Table: m_eles                                                */
/*==============================================================*/
create table m_eles (
   iid                  INT8                 not null,
   date_on              DATE                 not null,
   tags                 INT4[]               null,
   reles                 INT4[]               null,
   o_json               JSONB                null,
   constraint PK_M_ELES primary key (iid)
);

comment on column m_eles.reles is
'相关联IDs';

/*==============================================================*/
/* Index: i_mele_tags                                           */
/*==============================================================*/
create  index i_mele_tags on m_eles using GIN (
tags
);

/*==============================================================*/
/* Index: i_mele_rele                                           */
/*==============================================================*/
create  index i_mele_rele on m_eles using GIN (
reles
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
   o_json               JSONb                not null,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_OBJECTS primary key (iid)
);

comment on table objects is
'itype:
1 people
2 country
';

/*==============================================================*/
/* Table: s_eles                                                */
/*==============================================================*/
create table s_eles (
   iid                  INT8                 not null,
   m_iid                INT8                 not null default 0,
   date_on              DATE                 not null,
   user_id              INT4                 not null,
   o_json               JSONB                not null,
   tags                 INT4[]               null,
   reles                 INT4[]               null,
   review_json          JSONb                null default '0',
   status               INT4                 not null default 0,
   updated_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_S_ELES primary key (iid)
);

comment on table s_eles is
'编辑状态的内容
';

comment on column s_eles.m_iid is
'主表ID，如果没有，默认为0';

comment on column s_eles.reles is
'相关联IDs';

/*==============================================================*/
/* Index: i_sele_tags                                           */
/*==============================================================*/
create  index i_sele_tags on s_eles using GIN (
tags
);

/*==============================================================*/
/* Index: i_sele_rele                                           */
/*==============================================================*/
create  index i_sele_rele on s_eles using GIN (
reles
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
