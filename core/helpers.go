package core

import (
	"math/rand"
	"time"
)

func GetCurrentTimestamp() uint32 {
	return uint32(time.Now().Unix()) & 0x00FFFFFF
}

func GetTimestamp(metadata uint32) uint32 {
	return metadata & 0x00FFFFFF
}

func GetLogCounter(metadata uint32) uint8 {
	return uint8(metadata&0xFF000000) >> 24
}

func IncrLogCounter(counter uint8) uint8 {
	if counter == 255 {
		return 255
	}
	randomFactor := rand.Float32()
	approxFactor := 1.0 / float32(counter*uint8(LFULogFactor)+1)
	if approxFactor > randomFactor {
		counter++
	}
	return counter
}
