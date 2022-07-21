CREATE DATABASE IF NOT EXISTS dev_book;
USE dev_book;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

/* Timestamp é o tipo Data, e current_timestamp é uma função que ao criar o usuário irá pegar a data e tempo atual da criação */

