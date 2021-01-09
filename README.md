# docker.webp
convert images to webp via docker.   
[DockerHub](https://hub.docker.com/r/takecy/webp/)  

## WebP
[WebP](https://developers.google.com/speed/webp/)

>A new image format for the Web  

ex)
```
-rw-r--r-- 1 takecy staff  31372  1 20 12:56 sample_1.jpg
-rw-r--r-- 1 takecy staff   9256  1 20 18:50 sample_1.webp
-rw-r--r-- 1 takecy staff 218144  1 20 12:56 sample_2.png
-rw-r--r-- 1 takecy staff  62754  1 20 19:03 sample_2.webp
```

## Quick Start

```bash
$ make run
$ docker exec d_webp cwebp -o webp/sample_1.webp sample_1.jpg
Saving file 'webp/sample_1.webp'
File:      sample_1.jpg
Dimension: 420 x 315
Output:    4412 bytes Y-U-V-All-PSNR 43.80 44.01 45.11   44.03 dB
block count:  intra4: 238
              intra16: 302  (-> 55.93%)
              skipped block: 172 (31.85%)
bytes used:  header:            118  (2.7%)
             mode-partition:    754  (17.1%)
 Residuals bytes  |segment 1|segment 2|segment 3|segment 4|  total
    macroblocks:  |       3%|       7%|       4%|      85%|     540
      quantizer:  |      36 |      36 |      33 |      24 |
   filter level:  |      11 |       9 |       6 |       5 |

$ ls test/webp/
sample_1.webp
$ make stop
```

## License
[MIT](./LICENSE)
