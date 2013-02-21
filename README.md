benchmike
=========

benchmike is a command line benchmarking tool.  It runs your code a specified
number of iterations and gives you back the min, max, 95th and 99th percentile
in milliseconds.

Install
-------

Download the `benchmike` file, make it executable and put it on your `$PATH`.

Usage
-----

    Usage: benchmike [options] COMMAND

    Options:
      --version             show program's version number and exit
      -h, --help            show this help message and exit
      -i ITERATIONS, --iterations=ITERATIONS
                            How many times should the command run?

Example
-------

    $ ./benchmike -i 100 "date"
    Iterations:      100
    95th percentile: 6.99806213379
    99th percentile: 7.36784934998
    min:             5.24401664734
    max:             7.38000869751

License
-------

GPLv3
