package checkHealth

type CheckHealthRepository interface {
	RetrieveVersion() (string, error)
}

type CheckHealthUseCase interface {
	GetVersion() (string, error)
}
