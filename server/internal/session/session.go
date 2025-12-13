package session

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type SessionManager struct {
	userSessionIds map[string]string
	mu             sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		userSessionIds: make(map[string]string),
		mu:             sync.Mutex{},
	}
}

var (
	ErrNotAuthenticated = errors.New("not authenticated")
)

func (sm *SessionManager) AuthUser(userId string, sessionId string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	storedSessionId, exists := sm.userSessionIds[userId]
	if !exists || storedSessionId != sessionId {
		return ErrNotAuthenticated
	}
	return nil
}

func (sm *SessionManager) CreateSession(userId string) (string, error) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, exists := sm.userSessionIds[userId]; exists {
		return "", ErrNotAuthenticated
	}

	sessionId := uuid.NewString()
	sm.userSessionIds[userId] = sessionId
	return sessionId, nil
}
