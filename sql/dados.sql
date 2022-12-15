insert into usuarios(nome, nick, email, senha)
values
("Usuario 1","usuario_1","usuario1@gmail.com","$2a$10$TtFkjrqTAwoFzkY.464vLOf2rmMYKiOu6LxB8OuDGRSyoZsAf55BW"), 
("Usuario 2","usuario_2","usuario2@gmail.com","$2a$10$TtFkjrqTAwoFzkY.464vLOf2rmMYKiOu6LxB8OuDGRSyoZsAf55BW"), 
("Usuario 3","usuario_3","usuario3@gmail.com","$2a$10$TtFkjrqTAwoFzkY.464vLOf2rmMYKiOu6LxB8OuDGRSyoZsAf55BW"), 
("Usuario 4","usuario_5","usuario4@gmail.com","$2a$10$TtFkjrqTAwoFzkY.464vLOf2rmMYKiOu6LxB8OuDGRSyoZsAf55BW"); 

insert into seguidores(usuario_id,seguidor_id)
values 
(1,2),
(3,1),
(1,3);