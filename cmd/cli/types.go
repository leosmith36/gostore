package main

type args struct {
	host string
	port int
}

func getDefaultArgs() *args {
	return &args{
		host: "127.0.0.1",
		port: 4535,
	}
}
