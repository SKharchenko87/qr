package qr

import (
	"slices"
)

//101110001100000
//________ ________
//    0100 01100100
//<Способ кодирования><Количество данных><исходная последовательность>

// Заполнение
// На выходе последовательность бит кратная 8 и
// если количество бит в текущей последовательности байт меньше того,
// которое нужно для выбранной версии, то оно дополнено чередующимися
// байтами 11101100 и 00010001.
func fill(data []byte, level levelCorrection) []byte {
	l := len(data)
	bitSizeData := l * 8
	version, _ := slices.BinarySearch(levelToCountBits[level], bitSizeData)
	bitSizeFieldData := lengthFieldData(version, byteByByte)
	if bitSizeData+bitSizeFieldData+4 > levelToCountBits[level][version] {
		version++
		bitSizeFieldData = lengthFieldData(version, byteByByte)
	}
	res := make([]byte, levelToCountBits[level][version]/8)
	res[0] = byte(byteByByte << 4)
	index := 0
	if bitSizeFieldData == 16 {
		res[index] |= byte(l & 0b0000_0000_0000_0000_1111_0000_0000_0000 >> 12)
		index++
		res[index] |= byte(l & 0b0000_0000_0000_0000_0000_1111_0000_0000 >> 4)
	}
	res[index] |= byte(l & 0b0000_0000_0000_0000_0000_0000_1111_0000 >> 4)
	index++
	res[index] |= byte(l & 0b0000_0000_0000_0000_0000_0000_0000_1111 << 4)
	i := 0
	for ; i < l; i++ {
		res[index] |= byte(data[i] & 0b0000_0000_0000_0000_0000_0000_1111_0000 >> 4)
		index++
		res[index] |= byte(data[i] & 0b0000_0000_0000_0000_0000_0000_0000_1111 << 4)
	}
	index++
	tmp := [2]byte{0b11101100, 0b00010001}
	tmpIndex := 0
	for ; index < levelToCountBits[level][version]/8; i++ {
		res[index] = tmp[tmpIndex]
		tmpIndex = (tmpIndex + 1) % 2
		index++
	}
	return res
}

// Разделение информации на блоки - Определение количество байт в каждом блоке
func getCountByteOfBlock(level levelCorrection, version int) []int {
	bytesData := levelToCountBits[level][version] / 8
	byteOfBlock := bytesData / CountOfBlocks[level][version]
	reminderByteOfBlock := bytesData % CountOfBlocks[level][version]
	res := make([]int, CountOfBlocks[level][version])
	index := CountOfBlocks[level][version] - 1
	for ; reminderByteOfBlock > 0 && index >= 0; index, reminderByteOfBlock = index-1, reminderByteOfBlock-1 {
		res[index] = byteOfBlock + 1
	}
	for ; index >= 0; index-- {
		res[index] = byteOfBlock
	}
	return res
}

// Заполнение блоков
// Блок заполняется байтами из данных полностью.
// Когда текущий блок полностью заполняется, очередь переходит к следующему.
// Байтов данных должно хватить ровно на все блоки, ни больше и ни меньше.
func fillBlocks(data []byte, countByteOfBlock []int) [][]byte {
	l := len(countByteOfBlock)
	res := make([][]byte, l)
	index := 0
	for cnt, i := range countByteOfBlock {
		res[i] = make([]byte, cnt)
		for j := 0; j < cnt; j++ {
			res[i][j] = data[index]
			index++
		}
	}
	return res
}
