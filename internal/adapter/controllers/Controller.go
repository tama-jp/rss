package controllers

import "github.com/google/wire"

var Set = wire.NewSet(
	NewAccessTokenController,
	NewUserController,
	NewSystemUserController,
)
