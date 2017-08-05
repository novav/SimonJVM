package heap

type Object struct {
    // todo
    staticVars *Slots
    class   *Class
    fields  Slots
}

/* 6.6.1*/
func newObject(class *Class) *Object {
    return &Object{
        class: class,
        fields: newSlots(class.instanceSlotCount),
    }
}

func (self *Object) IsInstanceOf(class *Class) bool {
    return class.isAssignableFrom(self.class)
}