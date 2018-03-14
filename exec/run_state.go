package exec

import (
	"io"
	"reflect"
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
)

type runState struct {
	artifacts *worker.ArtifactRepository
	results   *sync.Map
	inputs    *sync.Map
}

func NewRunState() RunState {
	return &runState{
		artifacts: worker.NewArtifactRepository(),
		results:   &sync.Map{},
		inputs:    &sync.Map{},
	}
}

func (state *runState) Artifacts() *worker.ArtifactRepository {
	return state.artifacts
}

func (state *runState) Result(id atc.PlanID, to interface{}) bool {
	val, ok := state.results.Load(id)
	if !ok {
		return false
	}

	if reflect.TypeOf(val).AssignableTo(reflect.TypeOf(to).Elem()) {
		reflect.ValueOf(to).Elem().Set(reflect.ValueOf(val))
		return true
	}

	return false
}

func (state *runState) StoreResult(id atc.PlanID, val interface{}) {
	state.results.Store(id, val)
}

func (state *runState) SendUserInput(id atc.PlanID, input io.ReadCloser) {
	i, _ := state.inputs.LoadOrStore(id, make(chan io.ReadCloser))
	inputs := i.(chan io.ReadCloser)

	// send input (blocking) to a reader
	inputs <- input

	// wait for reader to close channel signifying that they're done
	<-inputs

	// clear out the channel
	state.inputs.Delete(id)
}

func (state *runState) ReadUserInput(id atc.PlanID, handler InputHandler) error {
	i, _ := state.inputs.LoadOrStore(id, make(chan io.ReadCloser))
	inputs := i.(chan io.ReadCloser)

	// signal to sender that we're done
	defer close(inputs)

	// wait for stream from request to arrive
	stream := <-inputs

	// synchronously stream in
	return handler(stream)
}
