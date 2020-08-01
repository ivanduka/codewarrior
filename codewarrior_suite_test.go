package main_test

import (
	. "codewarrior/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCodewarrior(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Should pass all of these")
}

// Change below only:

var _ = Describe("Should pass all of these", func() {
	//It("Should work with simple PIN", func() {
	//	Expect(Crack("827ccb0eea8a706c4c34a16891f84e7b")).To(Equal("12345"))
	//})
	//It("Should work with harder PIN", func() {
	//	Expect(Crack("86aa400b65433b608a9db30070ec60cd")).To(Equal("00078"))
	//})
	It("Should work with a very hard PIN", func() {
		Expect(Crack("283f42764da6dba2522412916b031080")).To(Equal("9999999"))
	})
	//It("Should work with a very hard PIN", func() {
	//	Expect(Crack("ef775988943825d2871e1cfa75473ec0")).To(Equal("99999999"))
	//})
})
