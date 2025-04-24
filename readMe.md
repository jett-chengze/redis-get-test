# build
docker build -t redis-get-test .

# run
docker run --name=redis-get-test -p 8081:8081 redis-get-test
docker run --name=redis-get-test -p 8081:8081 -e REDIS_DB=1 redis-get-test 

# use
go get -u github.com/redis/go-redis/v9
go get github.com/spf13/viper