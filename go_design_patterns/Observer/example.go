package main

func main() {
	listener1 := DataListener{"listener1"}
	listener2 := DataListener{"listener2"}
	listener3 := DataListener{"listener3"}

	subj := &DataSubject{} //
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)
	subj.registerObserver(listener3)

	subj.changeData("HELLO")
	subj.unregisterObserver(listener2)
	subj.changeData("WORLD!!")
}
