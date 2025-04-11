package auth

type AuthService struct {
	authRepository IAuthRepository
}

type IAuthService interface {
}

func NewAuthService(authRepository IAuthRepository) IAuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}
