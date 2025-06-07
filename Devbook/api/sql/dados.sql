insert into usuarios (nome, nick, email, senha)
values
('Usuario 1', 'usuario_1', 'usuario1@gmail.com', '$2a$10$02g9k1U6nUM1B3Y1YYnscu29jRQEJ/SEOHafCdw25pVvgEvEfw9Lu'),
('Usuario 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$02g9k1U6nUM1B3Y1YYnscu29jRQEJ/SEOHafCdw25pVvgEvEfw9Lu'),
('Usuario 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$02g9k1U6nUM1B3Y1YYnscu29jRQEJ/SEOHafCdw25pVvgEvEfw9Lu');

insert into seguidores (usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes (titulo, conteudo, autor_id)
values
('Publicação do Usuário 1', 'Essa é a publicação do usuário 1! Oba!', 1),
('Publicação do Usuário 2', 'Essa é a publicação do usuário 2! Oba!', 2),
('Publicação do Usuário 3', 'Essa é a publicação do usuário 3! Oba!', 3);