package model

type ProjectGroupPO struct {
	Id int `gorm:"primaryKey"`
	// 集群ID
	//[Field(Name = "cluster_ids", StorageType = EumStorageType.Array)]
	ClusterIds []int
	// 项目组名称
	Name string
}
