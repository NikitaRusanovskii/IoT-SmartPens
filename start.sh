sudo docker compose up -d

sleep 5
gnome-terminal -- bash -c "cd \$PWD && \
~/go/bin/migrate -path migrations -database postgres://iot:iot@localhost:5544/iot_db?sslmode=disable up\
&& go run ./server_part/cmd/app &&\
~/go/bin/migrate -path migrations -database postgres://iot:iot@localhost:5544/iot_db?sslmode=disable down; exec bash"

read
sudo docker compose down