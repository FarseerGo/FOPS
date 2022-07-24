package containerLog

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList 读取前500条日志
func (agent) ToList(top int) []PO {
	//ContainerLogContext.Data.ContainerLogPO.Desc(o => o.CreateAt).ToList(top)
	return nil
}
