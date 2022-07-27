# tinygo-iskrajs-poc

Скопировать файлы из каталогов tinygo/src и tinygo/targets в одноимённые каталоги с установленным дистрибом tinygo 0.24.0.
На ubuntu 20.04 и установленным пакетом tinygo-0.24.0 скопировать в /usr/local/lib/tinygo
```
https://github.com/tinygo-org/tinygo/releases/download/v0.24.0/tinygo_0.24.0_amd64.deb
```

дальше скомпилить пример из каталога examples, предварительно в него перейдя, и прошить бинарник в контроллер

```
tinygo build -o blinky3.bin -target iskrajs ./blinky3.go
dfu-util -d "0483:df11" -s 0x08000000 -a 0 -D ./blinky3.bin
```
