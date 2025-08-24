package token

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type UserSession struct {
	Token     string
	UserName  string
	CreatedAt time.Time
}

type TokenManager struct {
	tokens map[string]*UserSession
	mu     sync.RWMutex
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		tokens: make(map[string]*UserSession),
		mu:     sync.RWMutex{},
	}
}

func (tm *TokenManager) GenerateToken(userName string) string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	token := uuid.New().String()

	userSession := &UserSession{
		Token:     token,
		UserName:  userName,
		CreatedAt: time.Now(),
	}

	tm.tokens[token] = userSession
	return token
}

func (tm *TokenManager) GetToken(token string) (*UserSession, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	userSession, exists := tm.tokens[token]
	return userSession, exists
}
