package checkHealthUsecase

import "fp_pinjaman_online/src/checkHealth"

type checkHealthUC struct {
	checkHealthRepo checkHealth.CheckHealthRepository
}

func NewCheckHealthUsecase(checkHealthRepo checkHealth.CheckHealthRepository) checkHealth.CheckHealthUseCase {
	return &checkHealthUC{checkHealthRepo}
}

func (usecase *checkHealthUC) GetVersion() (string, error) {
	version, err := usecase.checkHealthRepo.RetrieveVersion()
	if err != nil {
		return "", err
	}
	return version, nil
}
