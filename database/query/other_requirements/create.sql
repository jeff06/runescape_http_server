create table other_requirements
(
    id                integer not null
        primary key autoincrement,
    name              text    not null,
    is_member         integer not null,
    is_quest          integer not null,
    is_skill          integer not null,
    id_of_requirement integer not null,
    id_unlock         integer
        constraint other_requirements_unlocks_other_requirements
            references unlocks
            on delete set null
);
