package main

import (
	"machine"
	"time"
)

var (
	uart = machine.UART0
	tx   = machine.UART_TX_PIN
	rx   = machine.UART_RX_PIN
)

func main() {

	time.Sleep(time.Millisecond * 10000)

	println("Beginning...")

	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: tx, RX: rx})
	
	println("writing to pmon")
	uart.Write([]byte{'h', '\n'})

	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()
			print(string(data))
		}
		time.Sleep(100 * time.Microsecond)

	}

}