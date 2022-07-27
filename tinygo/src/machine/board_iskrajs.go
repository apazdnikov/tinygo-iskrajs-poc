//go:build iskrajs
// +build iskrajs

package machine

import (
	"device/stm32"
	"runtime/interrupt"
)

const (
	LED         = LED_GREEN
	LED1        = LED_GREEN
	LED2        = LED_ORANGE
	LED_BUILTIN = LED_GREEN
	LED_GREEN   = PB6
	LED_ORANGE  = PB7
)

const (
	BUTTON = PC4
)

// UART pins
const (
	UART_TX_PIN = PB10
	UART_RX_PIN = PB11
        NUM_UART_INTERFACES = 1
)

var (

	UART3  = &_UART3
	_UART3 = UART{
		Buffer:            NewRingBuffer(),
		Bus:               stm32.USART3,
		TxAltFuncSelector: AF7_USART1_2_3,
		RxAltFuncSelector: AF7_USART1_2_3,
	}

	DefaultUART = UART3
)

// set up RX IRQ handler. Follow similar pattern for other UARTx instances
func init() {
	UART3.Interrupt = interrupt.New(stm32.IRQ_USART3, _UART3.handleInterrupt)
	stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_WWDGEN)
}

// SPI pins
const (
	SPI1_SCK_PIN = PA5
	SPI1_SDI_PIN = PA6
	SPI1_SDO_PIN = PA7
	SPI0_SCK_PIN = SPI1_SCK_PIN
	SPI0_SDI_PIN = SPI1_SDI_PIN
	SPI0_SDO_PIN = SPI1_SDO_PIN
)

// Since the first interface is named SPI1, both SPI0 and SPI1 refer to SPI1.
// TODO: implement SPI2 and SPI3.
var (
	SPI0 = SPI{
		Bus:             stm32.SPI1,
		AltFuncSelector: AF5_SPI1_SPI2,
	}
	SPI1 = &SPI0
)

const (
	I2C0_SCL_PIN = PB6
	I2C0_SDA_PIN = PB9
)

var (
	I2C0 = &I2C{
		Bus:             stm32.I2C1,
		AltFuncSelector: AF4_I2C1_2_3,
	}
)

/*
// USBCDC pins
const (
        USBCDC_HOSTEN_PIN = PA10
        USBCDC_DM_PIN = PA11
        USBCDC_DP_PIN = PA12
)

// USB CDC identifiers
const (
        usb_STRING_PRODUCT      = "IskraJS"
        usb_STRING_MANUFACTURER = "Amperka"
)

var (
        usb_VID uint16 = 0x2341
        usb_PID uint16 = 0x804e
)
*/