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
	curr := a.head
	for curr != nil {
		curr = curr.next
	}
	curr = &Node[Animal]{item: animal}
}

func (a *AnimalShelter) dequeueAny() (Animal, bool) {
	if a.head == nil {
		return Animal{}, false
	}

	tmp := a.head.item
	a.head = a.head.next
	return tmp, true
}

//func (a *AnimalShelter) dequeueDog() Animal {
//
//}
//
//func (a *AnimalShelter) dequeueCat() Animal {
//
//}

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
}
