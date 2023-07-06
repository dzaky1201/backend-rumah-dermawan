package main

func main() {
	server := InitializedServer()
	err := server.Run()
	if err != nil {
		return
	}
}
