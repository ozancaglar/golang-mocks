package printergomocks_test

import (
	"errors"
	"printergomocks/crudgomockfakes"
	"printergomocks/gomockmocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"printergomocks"
)

var _ = Describe("Printer Specs", func() {
	Describe("Printer", func() {
		var (
			printer printergomocks.Printer
			str     string

			outputErr error
		)

		JustBeforeEach(func() {
			outputErr = printer.Print(str)
		})

		When("str is empty", func() {
			BeforeEach(func() {
				str = ""
				printer = printergomocks.NewPrinter()
			})

			It("should error", func() {
				Expect(outputErr).ToNot(BeNil())
			})
		})

		When("str is not empty", func() {
			BeforeEach(func() {
				str = "123"
				printer = printergomocks.NewPrinter()
			})

			It("should not error", func() {
				Expect(outputErr).To(BeNil())
			})
		})
	})

	Describe("PrinterPrints with GoMock", func() {
		var (
			mockPrinter *gomockmocks.MockPrinterI
			ctrl        *gomock.Controller

			outputErr error
		)

		JustBeforeEach(func() {
			outputErr = printergomocks.PrinterPrints(mockPrinter)
		})

		When("str is empty", func() {
			BeforeEach(func() {
				ctrl = gomock.NewController(GinkgoT())
				mockPrinter = gomockmocks.NewMockPrinterI(ctrl)
				mockPrinter.EXPECT().Print().Return(errors.New("some error"))
			})

			It("should error", func() {
				Expect(outputErr).ToNot(BeNil())
			})
		})

		When("str is not empty", func() {
			BeforeEach(func() {
				ctrl = gomock.NewController(GinkgoT())
				mockPrinter = gomockmocks.NewMockPrinterI(ctrl)
				mockPrinter.EXPECT().Print().Return(nil)
			})

			It("should not error", func() {
				Expect(outputErr).To(BeNil())
			})
		})
	})

	Describe("PrinterPrints with GoMock", func() {
		var (
			mockPrinter *crudgomockfakes.FakePrinterI

			outputErr error
		)

		JustBeforeEach(func() {
			outputErr = printergomocks.PrinterPrints(mockPrinter)
		})

		When("string is empty", func() {
			BeforeEach(func() {
				mockPrinter = &crudgomockfakes.FakePrinterI{}
				mockPrinter.PrintReturns(errors.New("error"))
			})

			It("should error", func() {
				Expect(outputErr).To(Not(BeNil()))
			})
		})

		When("string is not empty", func() {
			BeforeEach(func() {
				mockPrinter = &crudgomockfakes.FakePrinterI{}
				mockPrinter.PrintReturns(nil)
			})

			It("should error", func() {
				Expect(outputErr).To(BeNil())
			})
		})
	})
})
