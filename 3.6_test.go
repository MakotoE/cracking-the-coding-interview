package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type AnimalShelter struct {
	head *Node[Animal]
}

type AnimalType int

const (
	Dog AnimalType = iota
	Cat
)

type Animal struct {
	AnimalType AnimalType
	ID         int
}

func (a *AnimalShelter) enqueue(animal Animal) {
	curr := &a.head
	for *curr != nil {
		curr = &(*curr).next
	}
	*curr = &Node[Animal]{item: animal}
}

func (a *AnimalShelter) dequeueAny() (Animal, bool) {
	if a.head == nil {
		return Animal{}, false
	}

	tmp := a.head.item
	a.head = a.head.next
	return tmp, true
}

func (a *AnimalShelter) dequeueType(animalType AnimalType) (Animal, bool) {
	parent := a.head
	if parent == nil {
		return Animal{}, false
	}

	if parent.item.AnimalType == animalType {
		tmp := parent.item
		parent = parent.next
		return tmp, true
	}

	for parent.next != nil {
		if parent.next.item.AnimalType == Dog {
			tmp := parent.next.item
			parent.next = parent.next.next
			return tmp, true
		}

		parent = parent.next
	}

	return Animal{}, false
}

func TestAnimalShelter(t *testing.T) {
	{
		shelter := &AnimalShelter{}
		_, ok := shelter.dequeueAny()
		assert.False(t, ok)

		a := Animal{
			AnimalType: Dog,
			ID:         0,
		}

		shelter.enqueue(a)

		result, ok := shelter.dequeueAny()
		assert.True(t, ok)
		assert.Equal(t, a, result)

		b := Animal{
			AnimalType: Dog,
			ID:         1,
		}
		c := Animal{
			AnimalType: Dog,
			ID:         2,
		}
		shelter.enqueue(b)
		shelter.enqueue(c)

		result, ok = shelter.dequeueAny()
		assert.True(t, ok)
		assert.Equal(t, b, result)

		result, ok = shelter.dequeueAny()
		assert.True(t, ok)
		assert.Equal(t, c, result)
	}
	{
		shelter := &AnimalShelter{}
		_, ok := shelter.dequeueType(Dog)
		assert.False(t, ok)

		a := Animal{
			AnimalType: Cat,
			ID:         0,
		}
		shelter.enqueue(a)

		_, ok = shelter.dequeueType(Dog)
		assert.False(t, ok)

		b := Animal{
			AnimalType: Dog,
			ID:         1,
		}
		shelter.enqueue(b)

		result, ok := shelter.dequeueType(Dog)
		assert.True(t, ok)
		assert.Equal(t, b, result)

		c := Animal{
			AnimalType: Dog,
			ID:         2,
		}
		shelter.enqueue(c)

		d := Animal{
			AnimalType: Cat,
			ID:         3,
		}
		shelter.enqueue(d)

		result, ok = shelter.dequeueType(Dog)
		assert.True(t, ok)
		assert.Equal(t, c, result)

		e := Animal{
			AnimalType: Cat,
			ID:         4,
		}
		shelter.enqueue(e)

		f := Animal{
			AnimalType: Dog,
			ID:         5,
		}
		shelter.enqueue(f)

		result, ok = shelter.dequeueType(Dog)
		assert.True(t, ok)
		assert.Equal(t, f, result)

		_, ok = shelter.dequeueType(Dog)
		assert.False(t, ok)

		result, ok = shelter.dequeueType(Cat)
		assert.True(t, ok)
		assert.Equal(t, a, result)
	}
}
