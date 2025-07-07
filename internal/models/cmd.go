package models

type CmdArgs struct {
	InputDirectory  string
	OutputDirectory string
	CPUProfile      string
	MemProfile      string
	Clear           bool
	Serve           bool
	Port            string
	VaultHealth     bool
}
