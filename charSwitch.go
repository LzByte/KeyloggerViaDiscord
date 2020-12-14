package main

import "strconv"

func charSwitch(c int) string {
	switch c {
	case 0x11:
		return "[Ctrl]"
	case 0x08:
		return "[Back]"
	case 0x09:
		return "[Tab]"
	case 0x0D:
		return "[Enter]\r\n"
	case 0x10:
		return "" //[Shift]
	case 0x12:
		return "[Alt]"
	case 0x14:
		return "" //[CapsLock]
	case 0x1B:
		return "[Esc]"
	case 0x20:
		return " "
	case 0x21:
		return "[PageUp]"
	case 0x22:
		return "[PageDown]"
	case 0x23:
		return "[End]"
	case 0x24:
		return "[Home]"
	case 0x25:
		return "[Left]"
	case 0x26:
		return "[Up]"
	case 0x27:
		return "[Right]"
	case 0x28:
		return "[Down]"
	case 0x29:
		return "[Select]"
	case 0x2A:
		return "[Print]"
	case 0x2B:
		return "[Execute]"
	case 0x2C:
		return "[PrintScreen]"
	case 0x2D:
		return "[Insert]"
	case 0x2E:
		return "[Delete]"
	case 0x2F:
		return "[Help]"
	case 0x5B:
		return "[LeftWindows]"
	case 0x5C:
		return "[RightWindows]"
	case 0x5D:
		return "[Applications]"
	case 0x5F:
		return "[Sleep]"
	case 0x60:
		return "[Pad 0]"
	case 0x61:
		return "[Pad 1]"
	case 0x62:
		return "[Pad 2]"
	case 0x63:
		return "[Pad 3]"
	case 0x64:
		return "[Pad 4]"
	case 0x65:
		return "[Pad 5]"
	case 0x66:
		return "[Pad 6]"
	case 0x67:
		return "[Pad 7]"
	case 0x68:
		return "[Pad 8]"
	case 0x69:
		return "[Pad 9]"
	case 0x6A:
		return "*"
	case 0xBB:
		return "+"
	case 0x6C:
		return "[Separator]"
	case 0xBD:
		return "-"
	case 0xBC:
		return ","
	case 0x6F:
		return "[Devide]"
	case 0x70:
		return "[F1]"
	case 0x71:
		return "[F2]"
	case 0x72:
		return "[F3]"
	case 0x73:
		return "[F4]"
	case 0x74:
		return "[F5]"
	case 0x75:
		return "[F6]"
	case 0x76:
		return "[F7]"
	case 0x77:
		return "[F8]"
	case 0x78:
		return "[F9]"
	case 0x79:
		return "[F10]"
	case 0x7A:
		return "[F11]"
	case 0x7B:
		return "[F12]"
	case 0x90:
		return "[NumLock]"
	case 0x91:
		return "[ScrollLock]"
	case 0xA0:
		return "" //[LeftShift]
	case 0xA1:
		return "" //[RightShift]
	case 0xA2:
		return "[LeftCtrl]"
	case 0xA3:
		return "[RightCtrl]"
	case 0xA4:
		return "[LeftMenu]"
	case 0xA5:
		return "[RightMenu]"
	case 0xBA:
		return ";"
	case 0xBF:
		return "/"
	case 0xC0:
		return "`"
	case 0xDB:
		return "["
	case 0xDC:
		return "\\"
	case 0xDD:
		return "]"
	case 0xDE:
		return "'"
	case 0xBE:
		return "."
	case 0x30:
		return "0"
	case 0x31:
		return "1"
	case 0x32:
		return "2"
	case 0x33:
		return "3"
	case 0x34:
		return "4"
	case 0x35:
		return "5"
	case 0x36:
		return "6"
	case 0x37:
		return "7"
	case 0x38:
		return "8"
	case 0x39:
		return "9"
	case 0x41:
		return "a"
	case 0x42:
		return "b"
	case 0x43:
		return "c"
	case 0x44:
		return "d"
	case 0x45:
		return "e"
	case 0x46:
		return "f"
	case 0x47:
		return "g"
	case 0x48:
		return "h"
	case 0x49:
		return "i"
	case 0x4A:
		return "j"
	case 0x4B:
		return "k"
	case 0x4C:
		return "l"
	case 0x4D:
		return "m"
	case 0x4E:
		return "n"
	case 0x4F:
		return "o"
	case 0x50:
		return "p"
	case 0x51:
		return "q"
	case 0x52:
		return "r"
	case 0x53:
		return "s"
	case 0x54:
		return "t"
	case 0x55:
		return "u"
	case 0x56:
		return "v"
	case 0x57:
		return "w"
	case 0x58:
		return "x"
	case 0x59:
		return "y"
	case 0x5A:
		return "z"
	}
	if c < 7 {
		return ""
	}
	return "º" + strconv.Itoa(c) + "º"
}
