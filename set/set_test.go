package set

import (
	testifyAssert "github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGeneral(t *testing.T) {
	require := testifyAssert.New(t)
	var set = &Set{}
	set.Init(0)

	require.Equal(0, set.Size)

	couldInsert := set.Insert(1)
	require.True(couldInsert)

	couldInsertSame := set.Insert(1)
	require.False(couldInsertSame, "Attempted to insert an already existing element")

	couldInsertDifferent := set.Insert(3)
	require.True(couldInsertDifferent)

	require.True(set.Exists(1))
	require.False(set.Exists(2))
	require.True(set.Exists(3))

	require.Equal(2, set.Size)

	couldDelete := set.Delete(1)
	require.True(couldDelete)

	couldDeleteSame := set.Delete(1)
	require.False(couldDeleteSame)

	couldDeleteDifferent := set.Delete(2)
	require.False(couldDeleteDifferent)

	require.False(set.Exists(1))
	require.False(set.Exists(2))
	require.True(set.Exists(3))

	require.Equal(1, set.Size)
}

func BenchmarkSet_Insert_True(b *testing.B) {
	set := &Set{}
	set.Init(0)
	for i := 0; i < b.N; i++ {
		set.Insert(i)
	}
}

func BenchmarkMap_Insert_True(b *testing.B) {
	set := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		set[i] = struct{}{}
	}
}

func BenchmarkSet_Insert_False(b *testing.B) {
	set := &Set{}
	set.Init(0)
	set.Insert(1)
	for i := 0; i < b.N; i++ {
		set.Insert(1)
	}
}

func BenchmarkMap_Insert_False(b *testing.B) {
	set := make(map[int]struct{})
	set[1] = struct{}{}
	for i := 0; i < b.N; i++ {
		set[1] = struct{}{}
	}
}

func BenchmarkSet_Insert_Random(b *testing.B) {
	set := &Set{}
	set.Init(0)
	for i := 0; i < b.N; i++ {
		set.Insert(rand.Int())
	}
}

func BenchmarkMap_Insert_Random(b *testing.B) {
	set := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		set[rand.Int()] = struct{}{}
	}
}

func BenchmarkSet_Insert_ManySmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := &Set{}
		set.Insert(rand.Int())
	}
}

func BenchmarkMap_Insert_ManySmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set := make(map[int]struct{})
		set[rand.Int()] = struct{}{}
	}
}
