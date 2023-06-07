package seed_work

import (
	"errors"
	"hash/fnv"
	"math/rand"
	"reflect"
)

type Entity struct {
	_id           int
	_domainEvents []INotification
}

func NewEntity() *Entity {
	return &Entity{_domainEvents: []INotification{}}
}

func (e Entity) GetId() int {
	return e._id
}

func (e *Entity) SetId(_id int) {
	e._id = _id
}

func (e Entity) GetDomainEvents() []INotification {
	return e._domainEvents
}

func (e *Entity) AddDomainEvent(item INotification) error {
	e._domainEvents = append(e._domainEvents, item)
	return nil
}

func (e *Entity) RemoveDomainEvent(item INotification) error {
	for i, v := range e._domainEvents {
		if reflect.DeepEqual(v, item) {
			e._domainEvents = append(e._domainEvents[:i], e._domainEvents[i+1:])
			return nil
		}
	}
	return errors.New("Domain event not found")
}

func (e *Entity) ClearDomainEvent() {
	e._domainEvents = []INotification{}
}

func (e Entity) IsTransient() bool {
	// 0 is default of int type
	return e._id == 0
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
		hash.Write([]byte(string(e._id)))
		return int(hash.Sum32())
	}

	return rand.Int()
}
