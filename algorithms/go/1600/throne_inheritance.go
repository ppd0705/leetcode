package main

type ThroneInheritance struct {
	kingName string
	dead     map[string]bool
	edges    map[string][]string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		kingName: kingName,
		dead: map[string]bool{},
		edges: map[string][]string{},
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	if edge, ok := this.edges[parentName]; ok {
		this.edges[parentName] = append(edge, childName)
	} else {
		this.edges[parentName] = []string{childName}
	}
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	res := make([]string, 0)
	var preOrder func(name string)
	preOrder = func(name string) {
		if !this.dead[name] {
			res = append(res, name)
		}
		for _, child := range this.edges[name] {
			preOrder(child)
		}
	}
	preOrder(this.kingName)
	return res
}
