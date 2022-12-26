create schema hte;

create table if not exists hte.locations
(
    id          integer generated always as identity
    constraint locations_pk
    primary key,
    device_id   varchar(40) not null,
    battery     integer     not null,
    timestamp   timestamp   not null,
    coordinates point       not null
    );

alter table hte.locations
    owner to postgres;