create table locations(
    id uuid primary key DEFAULT gen_random_uuid(),
    name varchar(32) not null ,
    location_type varchar(10) not null ,
    organization varchar(50) not null ,
    address varchar(100),
    api_key varchar(100)
);

CREATE UNIQUE INDEX CONCURRENTLY locations_name_org_unique_idx ON locations (name, organization);

create table robots(
    id uuid primary key DEFAULT gen_random_uuid(),
    name varchar(32) not null,
    api_key varchar(100)
);
CREATE UNIQUE INDEX CONCURRENTLY robots_name_unique_idx ON robots (name);
alter table locations add column is_online boolean default false;
alter table locations add column last_seen_online timestamp null;
