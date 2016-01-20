build:
	docker build -f Dockerfile -t takecy/webp .

run:
	docker run --rm -it -v `pwd`/test:/usr/local/img takecy/webp /bin/bash
