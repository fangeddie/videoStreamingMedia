CREATE DATABASE vid_stre_med;

use vid_stre_med;
create table users(
	id int primary key auto_increment,
    login_name varchar(64)  unique key,
    pwd varchar(64)
);

create table video_info(
	id varchar(64) primary key not null,
    author_id int,
    name varchar(150),
    display_ctime text,
    create_time DATETIME
);
create table comments (
	id varchar(64) primary key not null,
    video_id varchar(64),
    author_id int,
    content text,
    time DATETIME
);

create table sessions (
	session_id varchar(64) primary key not null,
    TTL tinytext,
    login_name varchar(64)
)
