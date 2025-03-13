Setting up the plant monitor to work with the pico

UART0

pmon tx to pico rx GP1 (uart0 rx)
pmon rx to pico tx GP1 (uart0 tx)
pmon 3v to pico 3v3
pmon gnd to pico gnd

go mod init pmonPico
go mod tidy

tinygo build -target=pico -o main.uf2 .

plug the pico into the pi with the button down

cp main.uf2 /media/rachel/RPI-RP2# pmonPico
