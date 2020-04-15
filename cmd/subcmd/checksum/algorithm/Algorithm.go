package algorithm

type InterfaceAlgorithm interface {
	GetName() string
	GetHash(data []byte) string
}

type Algorithm struct {
	Name string
}

func (a *Algorithm) GetName() string {
	return a.Name
}