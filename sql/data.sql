insert into users (name, nick, email, password)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "123456"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "123456"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "123456"),
("Usuário 4", "usuario_4", "usuario4@gmail.com", "123456");

insert into followers(user_id, follower_id)
values
(1, 2),
(1, 3),
(1, 4),
(2, 1),
(2, 3),
(2, 4),
(3, 1),
(3, 2),
(3, 4),
(4, 1),
(4, 2),
(4, 3);
