docker build -t go-base-prod . --target=production
docker run --rm -it go-base-prod

docker build -t go-base-dev . --target=dev
docker tag go-base macabrabits/go-base:1
docker push macabrabits/go-base:1
docker run --rm -it go-base