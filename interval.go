package main

import (
	"reflect"
	"time"
)

type intervalTimer struct {
	fn    reflect.Value
	delay time.Duration
	stop  chan struct{}
}

func NewInterval(fn any, delay time.Duration) (*intervalTimer, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, ErrNotFunction
	}

	return &intervalTimer{v, delay, make(chan struct{})}, nil
}

func (t *intervalTimer) Run(args ...any) {
	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	go func(in []reflect.Value) {
		for {
			select {
			case <-time.After(t.delay):
				t.fn.Call(in)
			case <-t.stop:
				return
			}
		}
	}(in)
}

func (t *intervalTimer) Stop() {
	t.stop <- struct{}{}
}
