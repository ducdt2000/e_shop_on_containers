package seed_work

import (
	"errors"
	"hash/fnv"
	"math/rand"
	"reflect"
)

type Entity struct {
	id           int
	domainEvents []INotification
}

func NewEntity() *Entity {
	return &Entity{domainEvents: []INotification{}}
}

func (e Entity) GetId() int {
	return e.id
}

func (e *Entity) SetId(_id int) {
	e.id = _id
}

func (e Entity) GetDomainEvents() []INotification {
	return e.domainEvents
}

func (e *Entity) AddDomainEvent(item INotification) error {
	e.domainEvents = append(e.domainEvents, item)
	return nil
}

func (e *Entity) RemoveDomainEvent(item INotification) error {
	for i, v := range e.domainEvents {
		if reflect.DeepEqual(v, item) {
			e.domainEvents = append(e.domainEvents[:i], e.domainEvents[i+1:])
			return nil
		}
	}
	return errors.New("Domain event not found")
}

func (e *Entity) ClearDomainEvent() {
	e.domainEvents = []INotification{}
}

func (e Entity) IsTransient() bool {
	// 0 is default of int type
	return e.id == 0
}

func (e Entity) Equals(obj interface{}) bool {
	if obj == nil || reflect.TypeOf(obj) != reflect.TypeOf(e) {
		return false
	}

	item := obj.(*Entity)
	if item.IsTransient() || e.IsTransient() {
		return false
	}

	return item.GetId() == e.GetId()
}

func (e Entity) GetHashCode() int {
	if !e.IsTransient() {
		hash := fnv.New32a()
		hash.Write([]byte(string(e.id)))
		return int(hash.Sum32())
	}

	return rand.Int()
}
