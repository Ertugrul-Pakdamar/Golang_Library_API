mkdir -p ~/mongo-data

docker rm -f mongodb 2>/dev/null || true
docker run -d \
  --name mongodb \
  -p 27017:27017 \
  -v ~/mongo-data:/data/db:Z \
  mongo:7 mongod --noauth
