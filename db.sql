create database enigma_school;
use enigma_school;

create table student (
    id varchar(36) primary key not null,
    name varchar(50) not null
);

create table subject (
    id varchar(36) primary key not null,
    name varchar(50) not null
);

create table teacher (
    id varchar(36) primary key not null,
    name varchar(50) not null,
    subject_id varchar(36) not null
);

create table score (
    subject_id varchar(36) not null,
    student_id varchar(36) not null,
    score int default 0
);

insert into student values 
    ('64f283b3-3a6d-4fbe-80ac-dd58b2ee7cd8', 'Skidipapap'),
    ('a2cdc29f-6f23-4f05-8f3c-db13b549d57b', 'Vivaldy'),
    ('daa90514-7726-48fd-9528-e887dc56379d', 'Ganggang');

insert into subject values 
    ('43ee7b04-2a62-4516-adf6-38e89f3c3881', 'Science'),
    ('886e67ed-f376-420d-b3a3-bdca25e555e5', 'Math');

insert into teacher values
    ('087a7bcb-8aa5-4497-8127-9d1e5aeefda8', 'Gang', '43ee7b04-2a62-4516-adf6-38e89f3c3881'),
    ('5d6c4056-7959-42a3-83e2-46f27d7008e5', 'Lav', '886e67ed-f376-420d-b3a3-bdca25e555e5'),
    ('dd978b2d-c165-4980-9c39-485fa40a046e', 'Didid', '886e67ed-f376-420d-b3a3-bdca25e555e5');

insert into score values
    ('886e67ed-f376-420d-b3a3-bdca25e555e5', 'daa90514-7726-48fd-9528-e887dc56379d', 90),
    ('43ee7b04-2a62-4516-adf6-38e89f3c3881', '64f283b3-3a6d-4fbe-80ac-dd58b2ee7cd8', 80),
    ('886e67ed-f376-420d-b3a3-bdca25e555e5', 'a2cdc29f-6f23-4f05-8f3c-db13b549d57b', 80);
