# Tinygo Examples


- Examples
    - tinygo-zoo : some tested for microbit
    - pico-led : tested


### Articles
- 2021/11/18 [CGo improvements in TinyGo](https://aykevl.nl/2021/11/cgo-tinygo)
- 2021/11/18 [0.21.0 Released](https://github.com/tinygo-org/tinygo/releases/tag/v0.21.0)
    - support ESP32-C3, RP2040, M5Stack Core2
- 2021/09/30 [Go on hardware: TinyGo in the wild](https://dev.to/gotime/go-on-hardware-tinygo-in-the-wild)
- 2020/09/24 [Garbage collection in TinyGo](https://aykevl.nl/2020/09/gc-tinygo)
- 2020/09/22 [ESP32 and ESP8266 support in TinyGo](https://aykevl.nl/2020/09/tinygo-esp32)
- 2019/09/12 [Let’s Go tiny with tinyGo](https://hackernoon.com/lets-go-tiny-with-tinygo-uob035t5)
- 2019/02/25 [Goroutines in TinyGo](https://aykevl.nl/2019/02/tinygo-goroutines)
- 2018/12/08 [Interfaces in TinyGo](https://aykevl.nl/2018/12/tinygo-interface)

### Information
- [191 Open Source Microcontroller Software Projects](https://opensourcelibs.com/libs/microcontroller)


### Slides
- 2021 [Run Go applications on Pico using TinyGo](https://www.slideshare.net/CherrieHsieh/run-go-applications-on-pico-using-tinygo)


### Videos
- 2021 [Optimize Go WebAssembly binary size with TinyGo](https://egghead.io/lessons/go-optimize-go-webassembly-binary-size-with-tinygo)
- 2019 [TinyGo: Go / Golang for Small Places (and Devices) by Ron Evans](https://www.youtube.com/watch?v=KY8u9yZ97Tc)


### Open Source
- [tinygo-org/bluetooth](https://github.com/tinygo-org/bluetooth) - Cross-platform Bluetooth API for Go and TinyGo
- [jypelle/xship](https://github.com/jypelle/xship)
- [mramshaw/TinyGo](https://github.com/mramshaw/TinyGo) - Experimenting with TinyGo
- [justinclift/tinygo-wasm-basic-triangle](https://github.com/justinclift/tinygo-wasm-basic-triangle) - Use TinyGo to create the basic WebGL triangle in Wasm. About 9.5kB compressed
- [dImrich/tinygo-wasmserve](https://github.com/dImrich/tinygo-wasmserve) - An HTTP server for testing TinyGo WebAssemblies
- [cashoefman/TinyGo-On-ESP32](https://github.com/cashoefman/TinyGo-On-ESP32)
- [aykevl/ws2812.go](https://gist.github.com/aykevl/47d0a24408cf585f6ba181c4dc663bca)
- [cashoefman/TinyGo-On-ESP32](https://github.com/cashoefman/TinyGo-On-ESP32)
- [**wasm3/embedded-wasm-apps**](https://github.com/wasm3/embedded-wasm-apps) - Run statically-compiled WebAssembly apps on any embedded platform


### Datasheets
- [ESP32-C3 Family](https://cdn.sparkfun.com/assets/1/a/5/a/6/esp32-c3_datasheet_en.pdf)



### Install
```
brew install platfoemio
brew tap tinygo-org/tools
brew install tinygo
```

```
pio run -e esp32 -t upload
pio device monitor
```