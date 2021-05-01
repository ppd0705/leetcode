package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func dfs(id int, employeeMap  map[int]*Employee) int {
	res := employeeMap[id].Importance
	for _, subId := range employeeMap[id].Subordinates {
		res += dfs(subId, employeeMap)
	}
	return res
}

func getImportance(employees []*Employee, id int) int {
	employeeMap := map[int]*Employee{}
	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}
	return dfs(id, employeeMap)
}
