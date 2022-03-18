package lib

var cleanups = make(map[string]func(string) (string, bool))
