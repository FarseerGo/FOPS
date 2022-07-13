package containerLog

type Repository interface {
	// ToList 读取前500条日志
	ToList(top int) []DomainObject
}
