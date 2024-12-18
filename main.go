package main

func main() {
	todos := Todos{}
	todos.add("Buy Milk")
	todos.add("Buy Bread")
	todos.toggle(0)
	todos.print()
}
