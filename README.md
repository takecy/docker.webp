# docker.webp
convert images to webp via docker.   
[DockerHub](https://hub.docker.com/r/takecy/webp/)  

## Image size
```
REPOSITORY    TAG             IMAGE ID       CREATED         SIZE
takecy/webp   latest          fe4411129bb6   6 minutes ago   14.3MB
```

## WebP
[WebP](https://developers.google.com/speed/webp/)

>A new image format for the Web  

### Example of size compression
ex)
```
.rw-r--r--  31k takecy 15 11  2016 sample_1.jpg
.rw-r--r-- 4.6k takecy  9 1 19:56  sample_1.webp
.rw-r--r-- 218k takecy 15 11  2016 sample_2.png
.rw-r--r--  45k takecy  9 1 19:57  sample_2.webp
```

## Quick Start

```bash
$ make run
$ docker exec d_webp cwebp -o webp/sample_1.webp sample_1.jpg
Saving file 'webp/sample_1.webp'
File:      sample_1.jpg
Dimension: 420 x 315
Output:    4562 bytes Y-U-V-All-PSNR 43.83 44.53 45.35   44.16 dB
           (0.28 bpp)
block count:  intra4:        136  (25.19%)
              intra16:       404  (74.81%)
              skipped:        93  (17.22%)
bytes used:  header:             99  (2.2%)
             mode-partition:    752  (16.5%)
 Residuals bytes  |segment 1|segment 2|segment 3|segment 4|  total
    macroblocks:  |       3%|       7%|       4%|      85%|     540
      quantizer:  |      36 |      36 |      33 |      24 |
   filter level:  |      11 |       9 |       6 |      16 |
$ ls test/webp/
sample_1.webp
$ make stop
```

## License
[MIT](./LICENSE)
