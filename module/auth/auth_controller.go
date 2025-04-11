package auth

type AuthController struct {
	authService IAuthService
}

func NewAuthController(authService IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}
