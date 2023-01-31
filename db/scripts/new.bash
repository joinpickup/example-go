echo Please enter the name of the new file
read name
migrate create -ext sql -dir ../migrations -seq $name