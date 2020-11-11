# urlbisect

For URLs with autoincrement ids, find the highest number using binary search.

```shell
$ urlbisect -u "https://freidok.uni-freiburg.de/data/@" -f 0 -t 500000
169476 https://freidok.uni-freiburg.de/data/169476
```

Verbose.

```shell
$ urlbisect -v -u "https://freidok.uni-freiburg.de/data/@" -f 0 -t 500000
2020/11/11 21:53:47 [404] https://freidok.uni-freiburg.de/data/250000
2020/11/11 21:53:47 [200] https://freidok.uni-freiburg.de/data/125000
2020/11/11 21:53:47 [404] https://freidok.uni-freiburg.de/data/187500
2020/11/11 21:53:47 [200] https://freidok.uni-freiburg.de/data/156250
2020/11/11 21:53:47 [404] https://freidok.uni-freiburg.de/data/171875
2020/11/11 21:53:47 [200] https://freidok.uni-freiburg.de/data/164063
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/167969
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169922
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/168946
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/169434
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169678
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169556
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169495
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/169465
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169480
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/169473
2020/11/11 21:53:48 [404] https://freidok.uni-freiburg.de/data/169477
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/169475
2020/11/11 21:53:48 [200] https://freidok.uni-freiburg.de/data/169476
169476 https://freidok.uni-freiburg.de/data/169476
```
