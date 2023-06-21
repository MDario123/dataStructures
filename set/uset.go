package set

// Uset is like Set but for uint
type Uset struct {
	Size    uint
	Buckets []uint
}

func (s *Uset) Len() uint { return s.Size }

func (s *Uset) Cap() uint { return uint(cap(s.Buckets)) }

func (s *Uset) Init(sizePow uint8) {
	s.Buckets = make([]uint, 1<<sizePow)
	s.Size = 0
}

func (s *Uset) grow() {
	newBuckets := make([]uint, 2*s.Cap())

	for _, element := range s.Buckets {
		if element > 0 {
			bucketId := element & (2*s.Cap() - 1)
			for newBuckets[bucketId] > 0 {
				if udist(newBuckets[bucketId], bucketId, 2*s.Cap()-1) < udist(element, bucketId, 2*s.Cap()-1) {
					element, newBuckets[bucketId] = newBuckets[bucketId], element
				}
				bucketId = (bucketId + 1) & (2*s.Cap() - 1)
			}
			newBuckets[bucketId] = element
		}
	}

	s.Buckets = newBuckets

	return
}

// Insert Inserts an element into the set, returns false if its already in the set and true otherwise.
func (s *Uset) Insert(element uint) bool {

	if len(s.Buckets) == 0 {
		s.Init(0)
	}

	if s.Size*10 >= s.Cap()*9 {
		s.grow()
	}

	bucketId := element & (s.Cap() - 1)

	for s.Buckets[bucketId] != 0 {
		if s.Buckets[bucketId] == element {
			return false
		}
		bucketId = (bucketId + 1) & (s.Cap() - 1)
	}

	s.Size++
	s.Buckets[bucketId] = element

	return true
}

// Delete Deletes an element from the set, returns false if it is not in the set and true otherwise.
func (s *Uset) Delete(element uint) bool {
	bucketId := element & (s.Cap() - 1)
	done := false
	for s.Buckets[bucketId] != 0 {
		if s.Buckets[bucketId] == element {
			s.Size--
			s.Buckets[bucketId] = 0
			done = true
			break
		}
		bucketId = (bucketId + 1) & uint(s.Cap()-1)
	}
	if done {
		for s.Buckets[bucketId] == 0 {
			next := (bucketId + 1) & uint(s.Cap()-1)
			if udist(s.Buckets[next], next, uint(s.Cap()-1)) > 0 {
				s.Buckets[bucketId], s.Buckets[next] = s.Buckets[next], 0
			}
			bucketId = next
		}
	}
	return done
}

// Exists Returns true if the element is in the set and false otherwise.
func (s *Uset) Exists(element uint) bool {

	if s.Size == 0 {
		return false
	}

	bucketId := element & (s.Cap() - 1)
	oBucketId := bucketId
	for s.Buckets[bucketId] != 0 {
		if s.Buckets[bucketId] == element {
			return true
		}
		bucketId = (bucketId + 1) & (s.Cap() - 1)
		if oBucketId == bucketId {
			break
		}
	}

	return false
}

// Copy returns a deep copy of the set
func (s *Uset) Copy() *Uset {
	newSet := &Uset{Size: s.Size}
	newSet.Buckets = make([]uint, s.Cap())
	copy(newSet.Buckets, s.Buckets)
	return newSet
}
