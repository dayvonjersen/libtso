type SafeMap struct {
	m  map[string]*Request
	mu *sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m:  map[string]*Request{},
		mu: &sync.Mutex{},
	}
}

func (s *SafeMap) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.m)
}

func (s *SafeMap) Get(key string) (*Request, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	req, ok := s.m[key]
	return req, ok
}

func (s *SafeMap) Set(key string, req *Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = req
}

func (s *SafeMap) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) Range() map[string]*Request {
	ret := map[string]*Request{}
	s.mu.Lock()
	defer s.mu.Unlock()
	for k := range s.m {
		ret[k] = s.m[k]
	}
	return ret
}
