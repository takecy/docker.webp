# docker.webp
images convert to webp via docker.  
[DockerHub](https://hub.docker.com/r/takecy/webp/)

## Quick Start
```bash
// on your machine
$ make build
$ make run
// on container
$ cd /usr/local/img
$ cwebp -q 90 sample_1.jpg -o webp/sample_1.webp
$ exit
// on your machine
$ ls ./test/webp
```

<br/>
## License
[MIT](./LICENSE)
