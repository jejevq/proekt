#!/bin/bash
SCRIPT_DIR=$(pwd)
docker compose up -d --scale app=3
echo "Введите пароль от рута в бд"
read -s -p "MySQL root password: " MYSQL_ROOT_PASSWORD
echo
echo "Введите пароль от нового пользователя для проекта в бд"
read -s -p "New user password: " NEW_USER_PASSWORD
echo

sed "s|__USER_PASSWORD__|$NEW_USER_PASSWORD|g" "$SCRIPT_DIR/init.sql" | \ #подставим в запрос что ввел пользователь и отправим в инит скьл для запроса к контейнеру
    docker exec -i sql mysql -u root -p"$MYSQL_ROOT_PASSWORD"

sudo cp "$SCRIPT_DIR/ewok.service" /etc/systemd/system/
sudo cp "$SCRIPT_DIR/bd.service" /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ewok.service bd.service
sudo systemctl start ewok.service bd.service

echo "Done."
