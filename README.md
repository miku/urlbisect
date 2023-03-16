# urlbisect

For URLs with (gapless) autoincrement ids, find the upper bound using binary
search. Reports the last successful id and URL.

## Installation

```
$ go install github.com/miku/urlbisect/cmd/urlbisect@latest
```

## Usage

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

Another example.

```
$ urlbisect -v -u https://www.semanticscholar.org/author/@ -f 0 -t 500000000
2020/11/11 22:15:22 [404] https://www.semanticscholar.org/author/250000000
2020/11/11 22:15:22 [404] https://www.semanticscholar.org/author/125000000
2020/11/11 22:15:23 [404] https://www.semanticscholar.org/author/62500000
2020/11/11 22:15:24 [200] https://www.semanticscholar.org/author/31250000
2020/11/11 22:15:25 [200] https://www.semanticscholar.org/author/46875000
2020/11/11 22:15:25 [404] https://www.semanticscholar.org/author/54687500
2020/11/11 22:15:27 [200] https://www.semanticscholar.org/author/50781250
2020/11/11 22:15:27 [404] https://www.semanticscholar.org/author/52734375
2020/11/11 22:15:27 [404] https://www.semanticscholar.org/author/51757813
2020/11/11 22:15:28 [404] https://www.semanticscholar.org/author/51269532
2020/11/11 22:15:29 [200] https://www.semanticscholar.org/author/51025391
2020/11/11 22:15:29 [200] https://www.semanticscholar.org/author/51147462
2020/11/11 22:15:30 [200] https://www.semanticscholar.org/author/51208497
2020/11/11 22:15:31 [200] https://www.semanticscholar.org/author/51239015
2020/11/11 22:15:31 [404] https://www.semanticscholar.org/author/51254274
2020/11/11 22:15:31 [404] https://www.semanticscholar.org/author/51246645
2020/11/11 22:15:32 [200] https://www.semanticscholar.org/author/51242830
2020/11/11 22:15:32 [200] https://www.semanticscholar.org/author/51244738
2020/11/11 22:15:32 [404] https://www.semanticscholar.org/author/51245692
2020/11/11 22:15:33 [200] https://www.semanticscholar.org/author/51245215
2020/11/11 22:15:34 [200] https://www.semanticscholar.org/author/51245454
2020/11/11 22:15:34 [404] https://www.semanticscholar.org/author/51245573
2020/11/11 22:15:35 [200] https://www.semanticscholar.org/author/51245514
2020/11/11 22:15:35 [404] https://www.semanticscholar.org/author/51245544
2020/11/11 22:15:35 [404] https://www.semanticscholar.org/author/51245529
2020/11/11 22:15:36 [200] https://www.semanticscholar.org/author/51245522
2020/11/11 22:15:36 [404] https://www.semanticscholar.org/author/51245526
2020/11/11 22:15:37 [200] https://www.semanticscholar.org/author/51245524
2020/11/11 22:15:37 [200] https://www.semanticscholar.org/author/51245525
51245525 https://www.semanticscholar.org/author/51245525
```

## Example sites

* https://journals.openedition.org/historika/363
* https://culturalanalytics.org/article/11156.pdf
* https://www.semanticscholar.org/author/2116037
* https://cyberleninka.org/article/n/427502.pdf
* https://www.irbnet.de/daten/iconda/CIB20436.pdf
* https://elib.dlr.de/134801/
* https://www.alexandria.unisg.ch/62751/
* https://www.scandinavica.net/article/16602.pdf
* https://www.honda-ri.de/pubs/pdf/1986.pdf
* http://journals.ioffe.ru/articles/viewPDF/44587
* ...

## Observations

* publishing system "scholastica" seems to offer a hosted version and it seems
  to be hosted on a single instance; e.g.
https://www.idrimjournal.com/article/11689,
https://www.scandinavica.net/article/16602 and so on share the same id space;
even though it might be totally unrealated, urls work across domains, e.g.
https://epj.us/article/17577.pdf,
https://www.scandinavica.net/article/17577.pdf and so on will point to the same
file (same with issues: https://www.internetmathematicsjournal.com/issue/2592)
