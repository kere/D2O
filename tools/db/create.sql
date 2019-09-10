drop table images;

drop index imele_miid;

drop index imele_reles;

drop index imele_tags;

drop index imele_area;

drop table m_elems;

drop index i_notice;

drop table notices;

drop index i_sele_miid;

drop index i_sele_reles;

drop index i_sele_tags;

drop index i_sele_area;

drop table s_elems;

drop index itags_usr;

drop table tags;

drop index u_uid;

drop table users;

/*==============================================================*/
/* Table: images                                                */
/*==============================================================*/
create table images (
   iid                  INT8                 not null,
   name                 VARCHAR(50)          not null,
   status               INT4                 not null default 0,
   dir                  VARCHAR(20)          not null,
   created_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_IMAGES primary key (iid)
);

/*==============================================================*/
/* Table: m_elems                                               */
/*==============================================================*/
create table m_elems (
   iid                  INT8                 not null,
   m_iid                INT8                 not null default 0,
   date_on              DATE                 null,
   user_id              INT4                 not null,
   o_json               JSONB                not null,
   d_json               JSONB                null,
   area                 INT4[]               null,
   tags                 TEXT[]               null,
   reles                INT4[]               null,
   review_json          JSONb                null default '0',
   status               INT2                 not null default 0,
   itype                INT2                 not null default 0,
   updated_at           TIMESTAMP WITH TIME ZONE not null default CURRENT_TIMESTAMP,
   constraint PK_M_ELEMS primary key (iid)
);

comment on table m_elems is
'内容主表
';

comment on column m_elems.m_iid is
'主表ID，如果没有，默认为0';

comment on column m_elems.reles is
'相关联IDs';

/*==============================================================*/
/* Index: imele_area                                            */
/*==============================================================*/
create  index imele_area on m_elems using GIN (
area
);

/*==============================================================*/
/* Index: imele_tags                                            */
/*==============================================================*/
create  index imele_tags on m_elems using GIN (
tags
);

/*==============================================================*/
/* Index: imele_reles                                           */
/*==============================================================*/
create  index imele_reles on m_elems using GIN (
  reles
);

/*==============================================================*/
/* Index: imele_miid                                            */
/*==============================================================*/
create  index imele_miid on m_elems (
user_id,
m_iid
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
/* Table: s_elems                                               */
/*==============================================================*/
create table s_elems (
   iid                  INT8                 not null,
   m_iid                INT8                 not null default 0,
   date_on              DATE                 null,
   user_id              INT4                 not null,
   o_json               JSONB                not null,
   d_json               JSONB                null,
   area                 INT4[]               null,
   tags                 TEXT[]               null,
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
/* Index: i_sele_area                                           */
/*==============================================================*/
create  index i_sele_area on s_elems using GIN (
area
);

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
reles
);

/*==============================================================*/
/* Index: i_sele_miid                                           */
/*==============================================================*/
create  index i_sele_miid on s_elems (
user_id,
m_iid
);

/*==============================================================*/
/* Table: tags                                                  */
/*==============================================================*/
create table tags (
   iid                  INT8                 not null,
   user_id              INT4                 not null,
   name                 VARCHAR(20)          not null,
   o_json               JSONb                null,
   cata_id              INT4                 not null default 0,
   constraint PK_TAGS primary key (iid)
);

/*==============================================================*/
/* Index: itags_usr                                             */
/*==============================================================*/
create  index itags_usr on tags (
user_id
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
