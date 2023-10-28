package interfaces

import "gadget-points/utils/mock"

var (
	userApp   mock.UserAppInterface
	agentApp  mock.AgentAppInterface
	fakeAuth  mock.AuthInterface
	fakeToken mock.TokenInterface

	s  = NewUsers(&userApp, &fakeAuth, &fakeToken)
	f  = NewAgent(&agentApp, &fakeAuth, &fakeToken)
	au = NewAuthenticate(&userApp, &fakeAuth, &fakeToken)
)
