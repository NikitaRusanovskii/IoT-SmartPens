sudo docker compose up -d

sleep 5
gnome-terminal -- bash -c "cd \$PWD && ./migrate_up.sh && go run ./cmd/app && ./migrate_down.sh; exec bash"

read
sudo docker compose down