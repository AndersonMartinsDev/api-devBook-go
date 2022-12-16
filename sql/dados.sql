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


insert into publicacao(titulo,conteudo,autor_id)
values
("Publicacao do usuario 1","Essa é a publicacao de usuario 1",1),
("Publicacao do usuario 2","Essa é a publicacao de usuario 2",2),
("Publicacao do usuario 3","Essa é a publicacao de usuario 3",3);