# udock

One config file. One command. Your whole dev stack — running on your phone.

Built entirely on Android with Termux. For developers who don't have a laptop.

## Install

    pkg install golang tmux udocker
    git clone https://github.com/YOURNAME/udock
    cd udock
    go build -o udock .
    ./udock install

## Usage

Create `udock.yaml` in your project root:

    version: "1"
    services:
      web:
        type: native
        command: "nginx -g 'daemon off;'"
        ports:
          - "8080"
      files:
        type: native
        command: "python3 -m http.server 8081"
        ports:
          - "8081"

Then from anywhere in the project:

    udock up       # start all services
    udock status   # see what's running
    udock logs web # tail logs from one service
    udock down     # stop everything
    udock clean    # force-kill when stuck

## Why native mode

Docker doesn't run on Android. udocker does, but port publishing is unreliable through PRoot.

udock handles both:
- Container mode — when an image runs well in udocker
- Native mode — run services directly in Termux

Same config file, same commands, same lifecycle.

## Status

Early. Day 6 of building in public.

Follow on x:@sp1Deer3
