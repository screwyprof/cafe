package adaptor

import (
	"github.com/screwyprof/cafe/infra/cqrs"
	"github.com/screwyprof/cafe/usecase/intf"
)

func ToDomain(h cqrs.Handler) intf.CommandHandler {
	return domainAdaptor{h: h}
}

type domainAdaptor struct {
	h cqrs.Handler
}

func (a domainAdaptor) Handle(c intf.Command) error {
	return a.h.Handle(c.(cqrs.Command))
}

func FromDomain(h intf.CommandHandler) cqrs.Handler {
	return infraAdapter{h: h}
}

type infraAdapter struct {
	h intf.CommandHandler
}

func (a infraAdapter) Handle(c cqrs.Command) error {
	return a.h.Handle(c.(intf.Command))
}
