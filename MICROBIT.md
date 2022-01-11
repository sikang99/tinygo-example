## Micro:Bit
> The BBC micro:bit is a tiny programmable computer designed for learning. 
> It is based on the Nordic Semiconductor nRF51822 ARM Cortex MO chip.

### Articles
- 2021/12/27 [Go on hardware: TinyGo in the wild](https://changelog.com/gotime/199)
- 2021/10/26 [Zephyr and the BBC Microbit V2 Tutorial Part 1 : GPIO](https://zephyrproject.org/zephyr-and-the-bbc-microbit-v2-tutorial-part-1-gpio/)
- 2021/08/31 [TinyGo: Good Things Come in Small Packages](https://auth0.com/blog/tinygo-good-things-come-in-small-packages/)
- 2020/11/09 [Introducing BBC micro:bit and MakeCode blocks .. with a Toy Robot Car and PC](https://www.codeproject.com/Articles/5284807/Introducing-BBC-micro-bit-and-MakeCode-blocks-with)
- 2019/05/05 [Pong-Like Retro Clock Using TinyGo and Microbit](https://www.hackster.io/_conejo/pong-like-retro-clock-using-tinygo-and-microbit-682736)


### Information
- [BBC micro:bit MicroPython documentation](https://microbit-micropython.readthedocs.io/en/v2-docs/)
- TinyGo: [BBC Micro:bit](https://tinygo.org/docs/reference/microcontrollers/microbit/)
- [Set up TinyGo development environment for VS Code](https://linuxtut.com/en/fa4daa52fdb9ff47dbe8/)


### Open Source
- [microbit-foundation](https://github.com/microbit-foundation)
- [microbit-foundation/microbit-v2-hardware](https://github.com/microbit-foundation/microbit-v2-hardware) - The schematic and Bill of Material for the BBC micro:bit V2
- [bbcmicrobit/micropython](https://github.com/bbcmicrobit/micropython) - Port of MicroPython for the BBC micro:bit
- [jp-rad/mbed-microbit-template](https://github.com/jp-rad/mbed-microbit-template)
- [nrf-rs/microbit](https://github.com/nrf-rs/microbit) - A Rust crate for BBC micro:bit development
- [lancaster-university/microbit-samples](https://github.com/lancaster-university/microbit-samples)
- [microbit-more/mbit-more-v2](https://github.com/microbit-more/mbit-more-v2) - Full-functional extension of micro:bit for Scratch3
- [tinygo-org/tinygo-zoo](https://github.com/tinygo-org/tinygo-zoo) - Various apps using TinyGo


### Usage
```
$ brew install tinygo
$ go mod init
$ go mod tidy
$ tinygo build -o microbit-blink.hex -target microbit main.go
$ file microbit-blink.hex
microbit-blink.hex: ASCII text
```
