create table unlocks
(
    id          integer not null
        primary key autoincrement,
    name        text    not null,
    description text    not null,
    is_member   integer not null,
    level       integer not null,
    id_skill    integer
        constraint unlocks_skills_unlocks
            references skills
            on delete set null
);

