/*
    usage:
        sessions = newSessionStore()
        go func() {
            for {
                <-time.After(time.Minute)
                sessions.GC()
            }
        }()
        
        func httpHandler(w http.ResponseWriter, r *http.Request) {
            sessId := sessions.Create(w, r)
            sess := sessions.Get(sessId)
            sess.Messages = append(sess.Messages, "Hello")
            sessions.Set(sessId, sess)
        }

        func otherHttpHandler(w http.ResponseWriter, r *http.Request) {
            sessId, ok := sessions.ID(w, r)
            if ok {
                sess := sessions.Get(sessId)
                for _, msg := range sess.Messages {
                    io.WriteString(w, msg)
                }
                sessions.Clear(sessId)
            }
        }
*/
import (
    "fmt"
    "net/http"
    "strings"
    "sync"
    "time"
)

var sessions *sessionStore

type session struct {
	Messages   []string
	LastActive int64
}
type sessionStore struct {
	store map[string]session
	mu    *sync.Mutex
}

func newSessionStore() *sessionStore {
	return &sessionStore{
		store: map[string]session{},
		mu:    &sync.Mutex{},
	}
}
func (s *sessionStore) ID(w http.ResponseWriter, r *http.Request) (string, bool) {
	c, err := r.Cookie("PHPSESS1D")
	if err == nil {
		return strings.TrimPrefix(c.String(), "PHPSESS1D="), true
	}
	return "", false
}
func (s *sessionStore) Create(w http.ResponseWriter, r *http.Request) string {
	sessId, ok := s.ID(w, r)
	if ok {
		return sessId
	}
	b := make([]byte, 32)
	rand.Read(b)
	sessId = fmt.Sprintf("%x", b)
	http.SetCookie(w, &http.Cookie{Name: "PHPSESS1D", Value: sessId, MaxAge: 0, SameSite: http.SameSiteStrictMode})

	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[sessId] = session{
		Messages:   []string{},
		LastActive: time.Now().Unix(),
	}

	return sessId
}
func (s *sessionStore) Get(sessId string) session {
	s.mu.Lock()
	defer s.mu.Unlock()

	if sess, ok := s.store[sessId]; ok {
		return sess
	}
	return session{
		Messages:   []*flashMessage{},
		LastActive: time.Now().Unix(),
	}
}
func (s *sessionStore) Set(sessId string, sess session) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[sessId] = sess
}
func (s *sessionStore) Clear(sessId string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, sessId)
}
func (s *sessionStore) GC() {
	s.mu.Lock()
	defer s.mu.Unlock()
	cmpTime := time.Now().Unix() - 60*60
	for sessId, sess := range s.store {
		if sess.LastActive < cmpTime {
			delete(s.store, sessId)
		}
	}
}
