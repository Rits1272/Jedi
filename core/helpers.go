package core

import (
  "time"
  "math/rand"
)

func GetCurrentTimestamp() uint32 {
  return uint32(time.Now().Unix()) & 0x00FFFFFF
}

func GetLogCounter(metadata uint32) uint8 {
  return uint8(metadata & 0xFF000000)
}

/*
  Same implementation as Redis

  Higher the value of counter, lesser the probability of it being incremented
*/
// [TODO]: Update this with Morris counter algorithm
func IncrLogCounter(counter uint8) uint8 {
  if (counter == 0xFF) {
    return 0xFF
  }

  randomFactor := rand.Float64()

  factor := 1.0 / float64(counter * uint8(LFU_LOG_FACTOR) + uint8(1))

  if (factor > randomFactor) {
    counter = counter + 1
  }
  
  return counter
}
