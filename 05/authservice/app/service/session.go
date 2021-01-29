package service

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Session struct {
	SessionID string
	UserID    int64
	Started   time.Time
	Expires   time.Time
}

/*
Sessions Session id to user id map
*/
var sessions map[string]Session
var sessionsLock sync.RWMutex

func CreateSession(userID int64) Session {
	result := Session{
		SessionID: randString(16),
		UserID:    userID,
		Started:   time.Now(),
		Expires:   time.Now().Add(time.Hour * 24),
	}
	sessionsLock.Lock()
	sessions[result.SessionID] = result
	sessionsLock.Unlock()
	return result
}

func GetSession(sessionID string) *Session {
	sessionsLock.RLock()
	session, ok := sessions[sessionID]
	sessionsLock.RUnlock()
	if !ok {
		return nil
	}
	if session.Expires.Before(time.Now()) {
		log.Printf("Session %s expired\n", session.SessionID)
		sessionsLock.Lock()
		delete(sessions, sessionID)
		sessionsLock.Unlock()
		return nil
	}
	return &session
}

func DeleteSession(sessionID string) {
	sessionsLock.Lock()
	delete(sessions, sessionID)
	sessionsLock.Unlock()
}

func init() {
	sessions = make(map[string]Session)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
