package main

import "fmt"

type observable interface {
	// 3 methods
	registerObserver(ob obeserver)
	unregisterObserver(ob obeserver)
	sampleTest(i int)
	notifyAll()
}

type DataSubject struct {
	dataListners []DataListener
	field        string
}

func (ds *DataSubject) sampleTest(i float32) {
	fmt.Println("WORKING")
}

func (ds *DataSubject) registerObserver(ob DataListener) {
	ds.dataListners = append(ds.dataListners, ob)
}

func (ds *DataSubject) unregisterObserver(ob DataListener) {
	// ds.dataListners = append(ds.dataListners, ob)
	var newDl []DataListener
	for _, dls := range ds.dataListners {
		if dls.Name != ob.Name {
			newDl = append(newDl, dls)
		}
	}
	ds.dataListners = newDl
}

func (ds *DataSubject) notifyAll() {
	for _, dls := range ds.dataListners {
		// if dls != ob {
		dls.onUpdate(ds.field)
		// }
	}
}

func (ds *DataSubject) changeData(data string) {
	ds.field = data
	ds.sampleTest(1)
	ds.notifyAll()
}
