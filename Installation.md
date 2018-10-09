# Installation

The documentation can be found here:

* https://golang.org/doc/install

Binary packages are provided for Windows, macOS, Linux:

* https://golang.org/dl/

The go distribution contains the runtime, tools and standard library.

```shell
$ ls -lh go
total 200K
drwxr-xr-x  2 tir tir 4.0K Oct  1 23:02 api
-rw-r--r--  1 tir tir  54K Oct  1 23:02 AUTHORS
drwxr-xr-x  2 tir tir 4.0K Oct  1 23:15 bin
-rw-r--r--  1 tir tir 1.4K Oct  1 23:02 CONTRIBUTING.md
-rw-r--r--  1 tir tir  70K Oct  1 23:02 CONTRIBUTORS
drwxr-xr-x  8 tir tir 4.0K Oct  1 23:02 doc
-rw-r--r--  1 tir tir 5.6K Oct  1 23:02 favicon.ico
drwxr-xr-x  3 tir tir 4.0K Oct  1 23:02 lib
-rw-r--r--  1 tir tir 1.5K Oct  1 23:02 LICENSE
drwxr-xr-x 15 tir tir 4.0K Oct  1 23:15 misc
-rw-r--r--  1 tir tir 1.3K Oct  1 23:02 PATENTS
drwxr-xr-x  9 tir tir 4.0K Oct  1 23:15 pkg
-rw-r--r--  1 tir tir 1.6K Oct  1 23:02 README.md
-rw-r--r--  1 tir tir   26 Oct  1 23:02 robots.txt
drwxr-xr-x 46 tir tir 4.0K Oct  1 23:02 src
drwxr-xr-x 22 tir tir  12K Oct  1 23:14 test
-rw-r--r--  1 tir tir    8 Oct  1 23:02 VERSION
```

The Go1.11.1 Linux distribution contains 1145 directories, 8729 files.

```
$ find go/ -type f | xargs -I {} basename {} | rev | cut -d '.' -f 1 | rev | sort | uniq -c | sort -nr | head -30
   6277 go
    519 s
    438 a
    247 txt
    106 golden
    106 c
     78 png
     55 src
     55 html
     42 tar
     35 sng
     35 input
     33 out
     26 zip
     25 jpeg
    ...
```

Includes documentation and mascot:

```
$ open go/misc/tour/static/img/gopher.png
```

![](img/gopher.png)

For the target platform, the installation instructions are noted here:

* https://golang.org/doc/install#install

