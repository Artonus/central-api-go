create table locations(
    id uuid primary key,
    name varchar(32) not null ,
    location_type varchar(10) not null ,
    organization varchar(50) not null ,
    address varchar(100)
);

create table robots(
    id uuid primary key ,
    name varchar(32) not null,
    api_key varchar(100)
);
