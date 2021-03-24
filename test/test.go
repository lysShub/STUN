package main

import (
	"fmt"
	"net"

	"github.com/lysShub/stun"
	"github.com/lysShub/stun/internal/com"
)

func main() {

	var s = new(stun.STUN)
	s.Port = 19986
	s.SeverAddr = "114.116.254.26"
	s.SecondPort = 19987

	fmt.Println("开始")

	fmt.Println(s.Sever())

	// fmt.Println(s.DiscoverClient())

}

func A() {
	laddr, err := net.ResolveUDPAddr("udp", ":19986")
	com.Errorlog(err)
	laddr2, err := net.ResolveUDPAddr("udp", ":19987")
	com.Errorlog(err)

	raddr, err := net.ResolveUDPAddr("udp", "114.116.254.26:19986")
	com.Errorlog(err)

	conn1, err := net.DialUDP("udp", laddr, raddr)
	com.Errorlog(err)
	_, err = conn1.Write([]byte("我是你爸爸"))
	com.Errorlog(err)

	conn2, err := net.DialUDP("udp", laddr2, raddr)
	com.Errorlog(err)
	_, err = conn2.Write([]byte("我是你爸爸"))
	com.Errorlog(err)
}
