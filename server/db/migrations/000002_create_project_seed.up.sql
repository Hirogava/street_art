INSERT INTO categories (name, description, image_url) VALUES
('Краски', 'Аэрозольные краски для граффити', '/static/img/category_img/588355d6571bcc74571b4d65.png'),
('Маркеры', 'Маркеры для рисования и письма', '/static/img/category_img/orig.png'),
('Трафареты', 'Трафареты для быстрого граффити', '/static/img/category_img/ee3a60b8e0e437467e2ca0d1c991561c.png'),
('Аксессуары', 'Аксессуары/одежда/защита', '/static/img/category_img/111.png');

INSERT INTO brands (name, description) VALUES
('Molotow', 'Премиальные краски и маркеры для граффити'),
('Ironlak', 'Доступные качественные краски для стрит-арта'),
('KRINK', 'Экспериментальные краски и чернила'),
('MTN', 'Испанский бренд аэрозольных красок'),
('Loop', 'Бренд трафаретов и аксессуаров для стрит-арта');

INSERT INTO products (name, description, price, category_id, stock, image_url, brand_id) VALUES
('Molotow Premium 400ml', 'Премиальная краска с высоким покрытием', 550.00, 1, 35, '/static/prod_img/category/1/3.jpg', 1),
('Ironlak Fat Cap', 'Краска с увеличенным соплом для толстых линий', 420.00, 1, 60, '/static/prod_img/category/1/4.jpg', 2),
('MTN 94 400ml', 'Классическая краска с мягким распылением', 470.00, 1, 45, '/static/prod_img/category/1/5.jpg', 3),
('KRINK K-71', 'Экспериментальная краска с уникальной текстурой', 600.00, 1, 25, '/static/prod_img/category/1/6.jpg', 4);

INSERT INTO products (name, description, price, category_id, stock, image_url, brand_id) VALUES
('Molotow ONE4ALL', 'Универсальный маркер на водной основе', 350.00, 2, 70, '/static/prod_img/category/2/3.jpg', 1),
('KRINK K-60', 'Маркер с масляной основой для перманентных отметок', 400.00, 2, 50, '/static/prod_img/category/2/4.jpg', 3),
('MTN Water Based', 'Экологичный маркер на водной основе', 280.00, 2, 90, '/static/prod_img/category/2/5.jpg', 4),
('Ironlak Marker', 'Прочный маркер с тонким наконечником', 320.00, 2, 65, '/static/prod_img/category/2/6.jpg', 5);

INSERT INTO products (name, description, price, category_id, stock, image_url, brand_id) VALUES
('Loop Starter Kit', 'Набор трафаретов для начинающих', 1200.00, 3, 30, '/static/prod_img/category/3/4.jpg', 5),
('Urban Stencil Pack', 'Набор городских тематических трафаретов', 950.00, 3, 25, '/static/prod_img/category/3/5.jpg', 1),
('Political Stencil Set', 'Политические и социальные трафареты', 850.00, 3, 20, '/static/prod_img/category/3/6.jpg', 2),
('Animal Silhouettes', 'Набор трафаретов с силуэтами животных', 780.00, 3, 35, '/static/prod_img/category/3/7.jpg', 3);

INSERT INTO products (name, description, price, category_id, stock, image_url, brand_id) VALUES
('Респиратор Montana', 'Защитный респиратор для работы с краской', 1200.00, 4, 40, '/static/prod_img/category/4/1.jpg', 1),
('Перчатки Molotov', 'Защитные перчатки для художников', 450.00, 4, 60, '/static/prod_img/category/4/2.jpg', 2),
('Кепка KRINK', 'Фирменная кепка с логотипом KRINK', 1800.00, 4, 30, '/static/prod_img/category/4/3.jpg', 4),
('Рюкзак MTN', 'Стильный рюкзак для переноски баллонов', 3500.00, 4, 25, '/static/prod_img/category/4/4.jpg', 5),
('Футболка Arton', 'Хлопковая футболка с принтом', 2200.00, 4, 50, '/static/prod_img/category/4/5.jpg', 1);

INSERT INTO top_products (product_id) VALUES
(1), (3), (7), (4), (2);

insert into admin_panel(email, password_hash) values('t@t', '$2a$10$JkAj4y.p3YYYMpveNfZmdO.18yDdduAh.nhi.iTyntIlHj4IeImAK');