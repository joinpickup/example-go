while getopts e: flag
do
    case "${flag}" in
        e) env=${OPTARG};;
    esac
done

source "$env"
echo "$DATABASE_URL"
migrate -path ../migrations -database $DATABASE_URL up