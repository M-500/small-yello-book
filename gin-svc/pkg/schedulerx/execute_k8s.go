package schedulerx

type K8SPodExecutor struct {
}

func NewK8SPodExecutor() *K8SPodExecutor {
	return &K8SPodExecutor{}
}

func (k *K8SPodExecutor) Execute(task TaskInterface) error {
	//TODO implement me
	panic("implement me")
}

func (k *K8SPodExecutor) Kill(task TaskInterface) error {
	//TODO implement me
	panic("implement me")
}
