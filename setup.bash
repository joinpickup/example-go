docker compose down
docker compose up -d
cd db/scripts
bash up.bash -e "../.env.local"