package printergomocks

import (
	"errors"
	"fmt"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PrinterI
type PrinterI interface {
	Print() error
}

type Printer struct {
}

func NewPrinter() Printer {
	return Printer{}
}

func (Printer) Print(str string) error {
	if str == "" {
		return errors.New("missing input")
	}

	fmt.Println(str)

	return nil
}

func PrinterPrints(printer PrinterI) error {
	if err := printer.Print(); err != nil {
		return errors.New(fmt.Sprintf("printer errored oh no: %s", err))
	}

	return nil
}
