package set

// Set Hash set implemented with an open-listing hash table based on Robin Hood Algorithm
type Set struct {
	Size    int
	Buckets []int
}

func (s *Set) Len() int { return s.Size }

func (s *Set) Cap() int { return cap(s.Buckets) }

func (s *Set) Init(sizePow uint8) {
	s.Buckets = make([]int, 1<<sizePow)
	s.Size = 0
}

func (s *Set) grow() {
	newBuckets := make([]int, 2*s.Cap())

	for _, element := range s.Buckets {
		if element > 0 {
			bucketId := uint(element) & uint(2*s.Cap()-1)
			for newBuckets[bucketId] > 0 {
				if dist(newBuckets[bucketId], bucketId, uint(2*s.Cap()-1)) < dist(element, bucketId, uint(2*s.Cap()-1)) {
					element, newBuckets[bucketId] = newBuckets[bucketId], element
				}
				bucketId = (bucketId + 1) & uint(2*s.Cap()-1)
			}
			newBuckets[bucketId] = element
		}
	}

	s.Buckets = newBuckets

	return
}

// Insert Inserts an element into the set, returns false if its already in the set and true otherwise.
func (s *Set) Insert(element int) bool {

	if len(s.Buckets) == 0 {
		s.Init(0)
	}

	if s.Size*10 >= s.Cap()*9 {
		s.grow()
	}

	bucketId := uint(element) & uint(s.Cap()-1)

	for s.Buckets[bucketId] != 0 {
		if s.Buckets[bucketId] == element {
			return false
		}
		bucketId = (bucketId + 1) & uint(s.Cap()-1)
	}

	s.Size++
	s.Buckets[bucketId] = element

	return true
}

// Delete Deletes an element from the set, returns false if it is not in the set and true otherwise.
func (s *Set) Delete(element int) bool {
	bucketId := uint(element) & uint(s.Cap()-1)
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
			if dist(s.Buckets[next], next, uint(s.Cap()-1)) > 0 {
				s.Buckets[bucketId], s.Buckets[next] = s.Buckets[next], 0
			}
			bucketId = next
		}
	}
	return done
}

// Exists Returns true if the element is in the set and false otherwise.
func (s *Set) Exists(element int) bool {

	if s.Size == 0 {
		return false
	}

	bucketId := uint(element) & uint(s.Cap()-1)
	oBucketId := bucketId
	for s.Buckets[bucketId] != 0 {
		if s.Buckets[bucketId] == element {
			return true
		}
		bucketId = (bucketId + 1) & uint(s.Cap()-1)
		if oBucketId == bucketId {
			break
		}
	}

	return false
}

// Copy returns a deep copy of the set
func (s *Set) Copy() *Set {
	newSet := &Set{Size: s.Size}
	newSet.Buckets = make([]int, s.Cap())
	copy(newSet.Buckets, s.Buckets)
	return newSet
}
