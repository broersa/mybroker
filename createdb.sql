create table applications (
  appkey bigserial primary key,
  appname varchar(50) unique not null,
  appeui varchar(16) unique not null
);

create table devices (
  devkey bigserial primary key,
  devapp bigint references applications (appkey),
  devappkey varchar(32) unique not null,
  deveui varchar(16) unique not null
);

create table sessions (
  seskey bigserial primary key,
  sesdev bigint references devices (devkey),
  sesdevnonce varchar(4) not null,
  sesappnonce varchar(6) not null
);
