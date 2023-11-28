package main

import (
	"crypto/sha256"
	"hash"
	"sync"
)

type Logger struct {
	lock         sync.Mutex
	messageCache map[string]int
}

func Constructor() Logger {
	return Logger{
		lock:         sync.Mutex{},
		messageCache: make(map[string]int),
	}
}

func (this *Logger) writeToCache(messageHash hash.Hash, timestamp int) {
	this.lock.Lock()
	hashString := string(messageHash.Sum(nil))
	this.messageCache[hashString] = timestamp
	this.lock.Unlock()
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	msgHash := sha256.New()
	msgHash.Write([]byte(message))
	hashString := string(msgHash.Sum(nil))
	if msgTime, ok := this.messageCache[hashString]; ok {
		if timestamp-msgTime >= 10 {
			this.writeToCache(msgHash, timestamp)
			return true
		}
	} else if !ok {
		this.writeToCache(msgHash, timestamp)
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
