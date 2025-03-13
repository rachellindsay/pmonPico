# Setting up the plant monitor to work with the pico

uart = UART0 (Need to use corresponding uart0 tx and rx gpio pins)

### cable connections
- pmon tx to pico rx GP1 (uart0 rx)
- pmon rx to pico tx GP1 (uart0 tx)
- pmon 3v to pico 3v3
- pmon gnd to pico gnd

### To set up new Go app
`go mod init pmonPico`

### Install dependencies
`go mod tidy`

### Build the app that will run on the Pico
`tinygo build -target=pico -o main.uf2 .`

### Flash/transfer app to the Pico
- plug the pico into the pi with the button down

`cp main.uf2 /media/rachel/RPI-RP2# pmonPico`
