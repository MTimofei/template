docker build -t first   -f first.dockerfile     .
docker build -t second  -f second.dockerfile    .
docker build -t third   -f third.dockerfile     .
docker build -t fourth  -f fourth.dockerfile    .

docker-compose up