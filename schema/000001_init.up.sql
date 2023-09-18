CREATE TABLE users(
                         id serial primary key,
                         name VARCHAR(255) not null,
                         surname VARCHAR(255) not null,
                         patronymic VARCHAR(255),
                         age int,
                         gender VARCHAR(10),
                         nationality VARCHAR(50)
);