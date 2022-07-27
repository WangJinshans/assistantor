package global

import (
	"errors"
	"sync"
)

var keyMap map[string]map[string]string
var mutex sync.RWMutex

func init() {
	// uuid - private key : public key
	keyMap = make(map[string]map[string]string)
}

func AddKeyValue(key string, privateKeyStr string, publicKeyStr string) {
	mutex.Lock()
	defer mutex.Unlock()
	pbKey := make(map[string]string)
	pbKey["privateKey"] = privateKeyStr
	pbKey["publicKey"] = publicKeyStr
	keyMap[key] = pbKey
}

// 返回私钥
func GetPrivateKey(key string) (string, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if value, ok := keyMap[key]; ok {
		return value["privateKey"], nil
	}
	err := errors.New("empty error")
	return "", err
}

func DeleteKey(key string) {
	mutex.Lock()
	defer mutex.Unlock()
	
	delete(keyMap, key)
	return
}
