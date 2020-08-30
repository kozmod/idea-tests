package pipe

type Executor func(interface{}) (interface{}, error)

type Pipeline interface {
	Pipe(executor Executor) Pipeline
	Merge() <-chan interface{}
}
