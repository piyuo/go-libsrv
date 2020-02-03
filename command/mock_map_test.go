package command

//TestMap mapping ID and object
type TestMap struct {
}

// IDToObject return object from id
func (r *TestMap) NewObjectByID(id uint16) interface{} {
	switch id {
	case 1001:
		return new(TestAction)
	case 1002:
		return new(TestActionNotRespond)
	case 1003:
		return new(SlowAction)
	case 1004:
		return new(DeadlineAction)
	}
	return nil
}
