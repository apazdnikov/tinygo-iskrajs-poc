# tinygo-iskrajs-poc

Скопировать в установленный каталог с дистрибутивом tinygo.
На ubuntu 20.04 и утановленным пакетом tinygo-0.23.deb скопировать в /usr/local/lib/tinygo

дальше использовать

```
tinygo build -o blinky3.bin -target iskrajs ./blinky3.go
dfu-util -d "0483:df11" -s 0x08000000 -a 0 -D ./blinky3.bin
```
