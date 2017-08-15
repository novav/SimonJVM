package heap

type Object struct {
    // todo
//     staticVars *Slots
    class   *Class
    // fields  Slots
    data interface{}
}

/* 6.6.1*/
func newObject(class *Class) *Object {
    return &Object{
        class: class,
        data: newSlots(class.instanceSlotCount),
    }
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}