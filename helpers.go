package main

import (
  "crypto/sha256"
  "encoding/hex"
)

// this seems slow, can we use some other lightweight hashing algorithm
func sha256Hash (key string) string {
  hash := sha256.New()
  hash.Write([]byte(key))

  hashBytes := hash.Sum(nil)
  checksum := hex.EncodeToString(hashBytes) 

  return checksum 
}

func max(a, b int) int {
  if a > b {
    return a
  }
  return b
}
