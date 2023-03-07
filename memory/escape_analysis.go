package memory

type foo struct {
	id   string
	name string
}

type bar struct {
	id   string
	name string
}

var input foo
var output bar

func init() {
	input = foo{
		id:   "6b45fb3d-6ae1-4cf7-8d2c-1e1718be298c",
		name: "foo",
	}
}

func ptrInValOut(in *foo) bar {
	return bar{
		id:   in.id,
		name: "bar",
	}
}

func ptrInPtrOut(in *foo) *bar {
	return &bar{
		id:   in.id,
		name: "bar",
	}
}
