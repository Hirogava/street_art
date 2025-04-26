-- Категории
INSERT INTO categories (name, description) VALUES
('Краски', 'Аэрозольные краски для граффити'),
('Маркеры', 'Маркеры для рисования и письма'),
('Скейтборды', 'Аксессуары для скейтбордов и самих скейтов');

-- Бренды
INSERT INTO brands (name, description) VALUES
('Montana', 'Известный бренд аэрозольных красок'),
('Posca', 'Маркерный бренд для художников'),
('Santa Cruz', 'Легендарный бренд скейтбордов');

-- Продукты
INSERT INTO products (name, description, price, category_id, stock, image_url, brand_id) VALUES
-- Категория: Краски
('Краска Montana Black', 'Классическая черная краска Montana для граффити', 450.00, 1, 50, '/static/prod_img/category/1/1.jpg', 1),
('Краска Montana Gold', 'Профессиональная краска Montana Gold для тонкой работы', 480.00, 1, 40, '/static/prod_img/category/1/1.jpg', 1),
-- Категория: Маркеры
('Маркер Posca PC-1M', 'Универсальный маркер для тонких линий', 250.00, 2, 100, '/static/prod_img/category/2/2.jpg', 2),
('Маркер Posca PC-5M', 'Средний универсальный маркер для граффити', 300.00, 2, 80, '/static/prod_img/category/2/2.jpg', 2),
-- Категория: Скейтборды
('Скейтборд Santa Cruz Classic Dot', 'Фирменный скейтборд с принтом Classic Dot', 7500.00, 3, 20, '/static/prod_img/category/3/3.jpg', 3),
('Скейтборд Santa Cruz Screaming Hand', 'Знаменитый скейтборд Screaming Hand', 7900.00, 3, 15, '/static/prod_img/category/3/3.jpg', 3);
