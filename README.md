# tinygo-iskrajs-poc

Скопировать в установленный каталог с дистрибутивом tinygo.
На ubuntu 20.04 и утановленным пакетом tinygo-0.24.0 скопировать в /usr/local/lib/tinygo
```
https://github.com/tinygo-org/tinygo/releases/download/v0.24.0/tinygo_0.24.0_amd64.deb
```

дальше использовать

```
tinygo build -o blinky3.bin -target iskrajs ./blinky3.go
dfu-util -d "0483:df11" -s 0x08000000 -a 0 -D ./blinky3.bin
```
