create table skills
(
    id          integer                not null
        primary key autoincrement,
    name        text not null,
    description text not null,
    is_member   bool not null
);