package tcppack

import (
	"bytes"
	"encoding/binary"
)

var (
	// 完整的数据包
	packet = make([]byte, 0)
	// 缓存
	cache = make([]byte, 0)
)

const (
	// 每个数据包前4个字节保存一个整形, 该整形记录数据包的字节长度
	HEADER_LEN = 4
)

// 封包
func Pack(buf []byte) []byte {
	// 将消息内容追加到消息体
	return append(intToBytes(len(buf)), buf...)
}

// 解包
func Unpack(buf []byte) []byte {
	buf = append(cache, buf...)
	length := len(buf)

	messageLength := bytesToInt(buf[:HEADER_LEN])
	// 当前数据包的总长度
	total := HEADER_LEN + messageLength
	if length < total {
		cache = buf
		return []byte{}
	} else if length == total {
		cache = []byte{}
		return buf[HEADER_LEN:]
	} else {
		cache = buf[total:]
		return buf[HEADER_LEN:total]
	}
}

//整形转换成字节
func intToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	// 写入消息头 PC 电脑和服务器广泛使用小端序
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func bytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32

	// 写入消息头 PC 电脑 广泛使用小端序
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
