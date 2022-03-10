package main

func main() {
	r := initRouter()
	r.Debug = true
	r.Logger.Fatal(r.Start(":1323"))

}
