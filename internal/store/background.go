package store

import "time"

func (s *Store) vacuum() {
	defer s.wg.Done()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:	
		}

		s.mu.Lock()
		for key, item := range s.cache {
			if item.expireAt.IsZero() {
				continue
			}
			if time.Now().After(item.expireAt) {
				delete(s.cache, key)
			}
		}
		s.mu.Unlock()
	}
}