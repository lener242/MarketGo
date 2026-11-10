CREATE TABLE IF NOT EXISTS dishes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category VARCHAR(100) NOT NULL,
    price INT NOT NULL
);

INSERT INTO dishes (name, category, price) VALUES
('Батон кукурузный', 'Хлебобулочные изделия', 5000),
('Тостерный хлеб', 'Хлебобулочные изделия', 6000),
('CocaCola 1л', 'Напитки', 10000),
('Fanta 1л', 'Напитки', 10000),
('Котлеты куриные 6шт', 'Мясные изделия', 25000),
('Котлеты говяжьи 4шт', 'Мясные изделия', 20000);