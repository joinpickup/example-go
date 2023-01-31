echo Please enter the env file
read env

source "$env"
migrate -path ../migrations -database $DATBASE_URL down -all