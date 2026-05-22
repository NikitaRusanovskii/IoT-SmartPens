CONTAINER_NAME="iot_postgres_web"
DB_USER="iot"
DB_NAME="iot_db"

if ! docker ps --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    echo "Ошибка: контейнер ${CONTAINER_NAME} не запущен."
    exit 1
fi

docker exec -i "$CONTAINER_NAME" psql -U "$DB_USER" -d "$DB_NAME" <<EOF
\x on
SELECT * FROM student_groups;
SELECT * FROM subjects;
SELECT * FROM students;
SELECT * FROM teachers;
SELECT * FROM lessons;
SELECT * FROM works;
\x off
EOF