rm -rf bin/main.exe
rm -rf bin/public
cp -r src/public bin/public/
env GOOS=linux go build -o bin/ src/main.go
docker-compose build
docker-compose up -d
docker-compose logs -f