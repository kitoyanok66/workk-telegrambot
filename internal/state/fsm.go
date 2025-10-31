package state

import "sync"

type StepData struct {
	Step string
	Data map[string]interface{}
}

var (
	fsmStore = make(map[string]*StepData)
	fsmMutex sync.RWMutex
)

func SetStep(userID string, step string, data map[string]interface{}) {
	fsmMutex.Lock()
	defer fsmMutex.Unlock()
	fsmStore[userID] = &StepData{Step: step, Data: data}
}

func GetStep(userID string) (*StepData, bool) {
	fsmMutex.RLock()
	defer fsmMutex.RUnlock()
	data, ok := fsmStore[userID]
	return data, ok
}

func DeleteStep(userID string) {
	fsmMutex.Lock()
	defer fsmMutex.Unlock()
	delete(fsmStore, userID)
}
