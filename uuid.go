/**
 * Created by goland.
 * User: adam_wang
 * Date: 2023-07-07 00:44:23
 */

package gotool

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// UUID 遵循 RFC4122 标准，UUID为128 bit (16 字节),此版本为V4根据随机数或者伪随机数生成 UUID
type UUID [16]byte

// `rand.Reader`是一个全局、共享的密码用强随机数生成器
var reader = rand.Reader

// Nil 定义一个类型为UUID的空值
var Nil UUID

// UuidNew 生成一个UUID
// @return UUID
// @return error
func UuidNew() (UUID, error) {
	return newRandom()
}

// UuidToString UUID转string
// @param uuid
// @return string
func UuidToString(uuid UUID) string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// 生成一个随机
// @return UUID
// @return error
func newRandom() (UUID, error) {
	return newRandomFromReader(reader)
}

// `io.ReadFull` 从 `rand.Reader` 精确地读取len(uuid)字节数据填充进uuid
// @param r io.Reader
// @return UUID
// @return error
func newRandomFromReader(r io.Reader) (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(r, uuid[:])
	if err != nil {
		return Nil, err
	}
	// 设置uuid版本信息
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

// 按照 8-4-4-4-12 的规则将 uuid 分段编码，使用 - 连接
// @param dst []byte
// @param uuid UUID
func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
