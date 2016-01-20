build:
	docker build -f Dockerfile -t takecy/webp .

run_f: build
	docker run --rm -it -v `pwd`/test:/usr/local/img takecy/webp /bin/bash

run: build
	docker run -d -it --name d_webp -v `pwd`/test:/usr/local/img takecy/webp /bin/bash

stop:
	docker stop d_webp && docker rm d_webp
