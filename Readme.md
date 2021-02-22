	package main

	import (
		"github.com/qq345340585/crchelpercrchelper"

		"github.com/astaxie/beego/logs"
	)

	func main() {
		m_data := []byte{0x01, 0x02, 0x03, 0x04}
		//参数对照https://crccalc.com/
		crc, _ := crchelper.CheckSum(m_data, "1021", "1D0F", "0000", false, false, 16)
		logs.Error("crc:0x%04X", crc)
	}
