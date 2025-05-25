package main

func main(){
	todos := Todos{}

	todos.Add("Learn Go")
	todos.Add("Build a CLI app")
	todos.Add("Write tests")

	todos.Toogle(0)
	todos.Toogle(1)

	todos.Print()
}