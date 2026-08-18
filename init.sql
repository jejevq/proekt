-- 1. Создаём базу
CREATE DATABASE proekt;

-- 2. Создаём пользователя
CREATE USER 'proekt'@'%' IDENTIFIED BY '123456';

-- 3. Даём пользователю права на эту базу
GRANT ALL PRIVILEGES ON proekt.* TO 'proekt'@'%';
FLUSH PRIVILEGES;

-- 4. Переключаемся на базу
USE proekt;

-- 5. Создаём таблицу
CREATE TABLE `proekt` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `out` varchar(40) NOT NULL,
  PRIMARY KEY (`id`)
)
