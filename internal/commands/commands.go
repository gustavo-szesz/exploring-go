package commands

import (
	"flag"
	"net"
)
type CmdFlags struct {
	mail string
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	
	flag.StringVar(&cf.mail, "mail", "", "Main you want to know")
	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute() {
	if cf.mail == "" {
		return
	}
	
	mxRecords, err := net.LookupMX("google.com")
	if err != nil {
		panic(err)
	}

	for _, mx := range mxRecords {
		println(mx.Host)
	}
}
