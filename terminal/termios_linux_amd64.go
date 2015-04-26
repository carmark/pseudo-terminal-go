// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs termios.go

package terminal

type syscall_Termios struct {
	Iflag	uint32
	Oflag	uint32
	Cflag	uint32
	Lflag	uint32
	Line	uint8
	Cc	[32]uint8
	Pad_cgo_0	[3]byte
	Ispeed	uint32
	Ospeed	uint32
}

const (
	syscall_IGNBRK	= 0x1
	syscall_BRKINT	= 0x2
	syscall_PARMRK	= 0x8
	syscall_ISTRIP	= 0x20
	syscall_INLCR	= 0x40
	syscall_IGNCR	= 0x80
	syscall_ICRNL	= 0x100
	syscall_IXON	= 0x400
	syscall_IXOFF	= 0x1000
	syscall_OPOST	= 0x1
	syscall_ECHO	= 0x8
	syscall_ECHONL	= 0x40
	syscall_ICANON	= 0x2
	syscall_ISIG	= 0x1
	syscall_IEXTEN	= 0x8000
	syscall_CSIZE	= 0x30
	syscall_PARENB	= 0x100
	syscall_CS8	= 0x30
	syscall_VMIN	= 0x6
	syscall_VTIME	= 0x5

	syscall_TCGETS	= 0x5401
	syscall_TCSETS	= 0x5402
)
