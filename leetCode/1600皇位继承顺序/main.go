package main

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		kingName,
		map[string][]string{},
		map[string]bool{},
	}
}

func (t *ThroneInheritance) Birth(parentName string, childName string) {
	t.edges[parentName] = append(t.edges[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var preorder func(string)
	preorder = func(name string) {
		if !t.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range t.edges[name] {
			preorder(childName)
		}
	}
	preorder(t.king)
	return
}
