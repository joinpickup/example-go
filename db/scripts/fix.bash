echo Please enter the version
read version

echo Please enter the env file
read env

source "$env"
migrate -path ../migrations -database $DATBASE_URL force $version