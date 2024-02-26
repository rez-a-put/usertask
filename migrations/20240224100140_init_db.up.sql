create table `users` (
    `id` bigint auto_increment not null,
    `name` varchar(255) null,
    `email` varchar(255) null,
    `password` varchar(255) null,
    `data_status` tinyint default 1 null comment "status of data. 1 = active, 2 = inactive",
    `created_at` timestamp default current_timestamp null,
    `updated_at` timestamp null,
    `deleted_at` timestamp null,
    constraint `users_pk` primary key (`id`),
    constraint `users_un` unique (`email`)
);

create table `tasks` (
    `id` bigint auto_increment not null,
    `user_id` bigint not null,
    `title` varchar(255) null,
    `description` text null,
    `status` tinyint default 1 null comment "status of task. 1 = pending, 2 = finished",
    `data_status` tinyint default 1 null comment "status of data. 1 = active, 2 = inactive",
    `created_at` timestamp default current_timestamp null,
    `updated_at` timestamp null,
    `deleted_at` timestamp null,
    constraint `tasks_pk` primary key (`id`),
    constraint `tasks_fk` foreign key (`id`) references `users`(`id`) on delete cascade
);