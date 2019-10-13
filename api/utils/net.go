package utils

import (
	"net"
	"strconv"
	"strings"
)

//ip地址 string 转 int
func INetAtoN(ip net.IP) int64 {
	bits := strings.Split(ip.String(), ".")
	
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	
	var sum int64
	
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	
	return sum
}

//ip地址 int 转 string
func INetNtoA(ip int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)
	
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
