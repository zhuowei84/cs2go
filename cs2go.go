package cs2go

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

/**
C#的BitConverterGetBytes
*/
func CsBitConverterGetBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	bytesBufferVal := bytesBuffer.Bytes()
	byteVal := []byte{}
	for i := len(bytesBufferVal) - 1; i >= 0; i-- {
		byteVal = append(byteVal, bytesBufferVal[i])
	}
	return byteVal
}

/**
C#的BitConverterToUInt32
 */
func CsBitConverterToUInt32(bs []byte, index int) (uint32, error) {
	strVal := []string{}
	if len(bs) > index + 3 {
		for i := index + 3; i >= index; i-- {
			byteStr := hex.EncodeToString([]byte{bs[i]})
			strVal = append(strVal, byteStr)
		}
	}
	strs := strings.Join(strVal, "")

	n, err := strconv.ParseUint(strs, 16, 32)
	if err != nil {
		return 0, err
	}

	return uint32(n), nil
}

/**
C#的ArrayCopy
调用示例：arr2 = CsArrayCopy(arr1, start, arr2, end, length)
*/
func CsArrayCopy(arr1 []byte, start int, arr2 []byte, end int, length int) []byte {
	arrLen := len(arr1) + end + start
	if len(arr2) < arrLen {
		for i := 0; i < len(arr2); i++ {
			arr2[i] = arr1[start+i]
		}
		return arr2
	}
	for i := start + end; i < arrLen; i++ {
		arr2[i] = arr1[i - end - start]
	}
	return arr2
}
