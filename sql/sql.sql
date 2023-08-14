CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id int auto_increment primary_key,
    name varchar(50) not null,
    nick varchar(50) unique not null,
    email varchar(50) unique not null,
    password varchar(255) not null,
    createdAt timestamp default current_timestamp()
) = ENGINE=INNODB;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id) REFERENCES users(id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB;