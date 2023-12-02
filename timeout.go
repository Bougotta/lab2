package main

import (
	"reflect"
	"time"
)

type timeoutTimer struct {
	fn    reflect.Value
	delay time.Duration
	stop  chan struct{}
}

func NewTimeout(fn any, delay time.Duration) (*timeoutTimer, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, ErrNotFunction
	}

	return &timeoutTimer{v, delay, make(chan struct{})}, nil
}

func (t *timeoutTimer) Run(args ...any) {
	in := make([]reflect.Value, len(args))
	for i := range args {
		in[i] = reflect.ValueOf(args[i])
	}

	go func(in []reflect.Value) {
		select {
		case <-time.After(t.delay):
			t.fn.Call(in)
		case <-t.stop:
			return
		}
	}(in)
}

func (t *timeoutTimer) Stop() {
	t.stop <- struct{}{}
}
