CONTAINER_NAME="iot_postgres"
DB_USER="iot"
DB_NAME="iot_db"

if ! docker ps --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Ошибка: контейнер ${CONTAINER_NAME} не запущен."
    echo "Запустите его: docker-compose up -d"
    exit 1
fi

docker exec -i "$CONTAINER_NAME" psql -U "$DB_USER" -d "$DB_NAME" <<EOF
\x on
SELECT * FROM lesson;
SELECT * FROM work;
\x off
EOF