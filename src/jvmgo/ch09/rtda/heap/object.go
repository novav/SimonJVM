package heap

type Object struct {
    // todo
//     staticVars *Slots
    class   *Class
    // fields  Slots
    data  interface{}
    extra interface{}
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

func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}


func (self *Object) GetRefVar(name, descriptor string) *Object {
    field := self.class.getField(name, descriptor, false)
    slots := self.data.(Slots)
    return slots.GetRef(field.slotId)
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
    field := self.class.getField(name, descriptor, false)
    slots := self.data.(Slots)
    slots.SetRef(field.slotId, ref)
}
