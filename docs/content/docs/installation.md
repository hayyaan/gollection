+++
date = "2016-05-22T19:55:18+02:00"
title = "Installation"
+++

To start you need to fetch the example application or you can bootstrap one by yourself.

    go get https://github.com/MetalMatze/gollection-app

At first you can delete the current `.git` folder. You probably want to start a new git history for your project.

### Compile
Now it is time to compile and run your application for the first time.

    make

This runs every action in the `Makefile` that comes after `all:`.

### Run
_make_ created a binary in the root directory of your application called `app`.

    ./app

If you want to start the http server of your application:

    ./app serve
