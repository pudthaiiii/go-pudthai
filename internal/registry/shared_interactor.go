package registry

import (
	i "go-ibooking/internal/usecase/interactor/shared"
)

func (r *registry) NewSharedAuthInteractor() i.SharedAuthInteractor {
	return i.NewSharedAuthInteractor(
		r.NewUsersRepository(),
		r.NewOauthAccessTokenRepository(),
		r.NewOauthRefreshTokenRepository(),
		r.cacheManager,
		r.cfg,
	)
}
