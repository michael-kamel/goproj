package uidGenerator 

import (
	"crypto/md5"
	"strconv"
	"time"
	"encoding/hex"
)
type UIDGenerator interface{
	GenerateUID() string
	DeleteUID(uid string) 
}

type DefaultUIDGenerator struct {
	uids map[string]bool
}

func (this *DefaultUIDGenerator) GenerateUID() string {
	hasher := md5.New()
	hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
	uuid := hex.EncodeToString(hasher.Sum(nil))
	if _, ok := this.uids[uuid]; ok {
		return this.GenerateUID()
	} else {
		this.uids[uuid] = true
		return uuid
	}
}
func (this *DefaultUIDGenerator) DeleteUID(uid string) {
	delete(this.uids, uid)
}
func (this *DefaultUIDGenerator) Init() {
	this.uids = make(map[string]bool)
}