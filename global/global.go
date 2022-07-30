package global

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"sync"
	"time"
)

var keyMap map[string]map[string]interface{}
var mutex sync.RWMutex
var enforcer *casbin.Enforcer

func init() {
	// uuid - private key : public key
	keyMap = make(map[string]map[string]interface{})
}

func SetEnforcer(e *casbin.Enforcer) {
	enforcer = e
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}

func AddKeyValue(key string, privateKeyStr string, publicKeyStr string) {
	mutex.Lock()
	defer mutex.Unlock()
	pbKey := make(map[string]interface{})
	pbKey["privateKey"] = privateKeyStr
	pbKey["publicKey"] = publicKeyStr
	pbKey["timestamp"] = time.Now().Unix()
	keyMap[key] = pbKey
}

// 返回私钥
func GetPrivateKey(key string) (string, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	if value, ok := keyMap[key]; ok {
		return value["privateKey"].(string), nil
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

func StartCleanKey(ctx context.Context) {
	timer := time.NewTimer(time.Minute * 5)
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			cleanKey() // 清理游离的key
			timer.Reset(time.Minute * 5)
		}
	}
}

func cleanKey() {
	for key, value := range keyMap {
		ts := value["timestamp"].(int64)
		if (time.Now().Unix() - ts) > 60*30 {
			delete(keyMap, key)
		}
	}
}
