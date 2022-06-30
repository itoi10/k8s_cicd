package main

func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start("0.0.0.0:8080"))
}
