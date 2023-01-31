while getopts n: flag
do
    case "${flag}" in
        n) name=${OPTARG};;
    esac
done

migrate create -ext sql -dir ../migrations -seq $name