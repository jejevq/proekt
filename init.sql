-- Создаём базу данных (если её ещё нет)
CREATE DATABASE IF NOT EXISTS proekt;

-- Создаём пользователя (пароль будет заменён из install.sh)
CREATE USER IF NOT EXISTS 'proekt'@'%' IDENTIFIED BY '__USER_PASSWORD__';

-- Даём только необходимые права на базу proekt
GRANT SELECT, INSERT, UPDATE, DELETE ON proekt.* TO 'proekt'@'%';
FLUSH PRIVILEGES;

-- Переключаемся на базу proekt
USE proekt;

-- Создаём таблицу proekt (если её нет)
CREATE TABLE IF NOT EXISTS proekt (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `out` VARCHAR(40) NOT NULL,
  PRIMARY KEY (id)
);
