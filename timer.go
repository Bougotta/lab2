package main

import "errors"

type Timer interface {
	Run()
	Stop()
}

var ErrNotFunction = errors.New("fn is not a function")
