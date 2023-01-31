echo Please enter the env file
read env

source "$env"
migrate -path ../migrations -database $DATABASE_URL up