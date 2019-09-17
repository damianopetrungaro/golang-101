package main

import "time"

type userName struct {
	v string
	t time.Time
}

func new(v string) userName {
	return userName{
		v, time.Now(),
	}
}

func newWithCreationDate(v string, t time.Time) userName {
	return userName{v, t}
}
