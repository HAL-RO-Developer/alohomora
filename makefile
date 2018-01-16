create:
	cp dbconfig.yml.template dbconfig.yml
	
migrate:
	sql-migrate $(@F) -env="development" -config="dbconfig.yml"

run:
	env lavender=production go run main.go

docker/build:
	make -f .circleci/ci.mk go/build
	make -f .circleci/ci.mk docker/build

docker/run:
	docker run -it --rm 241556795328.dkr.ecr.ap-northeast-1.amazonaws.com/lavender

docker/push:
	env NODE_ENV=production npm run build
	make -f .circleci/ci.mk go/build
	make -f .circleci/ci.mk docker/build
	make -f .circleci/ci.mk login
	make -f .circleci/ci.mk docker/push
