package main

type Event struct {
	Type string
	Title string
	Description string
	Addtime int64
	Tiptime int64
	Owner string
}


type ServerConfigs struct {
	addr string
	static string
	dbFile string
}
