package heap 

type Slot struct {
    num     int32
    ref     *heap.Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
    
}

func (self Slots) SetInt(index uint, val int32) {}
func (self Slots) GetInt(index uint) int32 {}
func (self Slots) SetFloat(index uint, val flaot32) {}
func (self Slots) GetFloat(index uint) float32 {}
func (self Slots) SetLong(index uint, val int64) {}
func (self Slots) GetLong(index uint) int64 {}
func (self Slots) SetDouble(index uint, val float64) {}
func (self Slots) GetDouble(index uint) float64 {}
func (self Slots) SetRef(index uint, val *Object) {}
func (self Slots) GetRef(index uint) *Object {}