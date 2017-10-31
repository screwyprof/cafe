package main

import (
	"log"

	"github.com/screwyprof/cafe/infra/cqrs"
	"github.com/screwyprof/cafe/interface/adaptor"
	"github.com/screwyprof/cafe/interface/cli/handler"
	"github.com/screwyprof/cafe/interface/repo/inmemory"

	"github.com/screwyprof/cafe/usecase/tab"
	"github.com/screwyprof/cafe/usecase/tab/command"
	"github.com/screwyprof/cafe/usecase/tab/query"
)

func main() {

	repo := inmemory.NewTabRepo()

	h := cqrs.NewCompositeHandler()
	h.RegisterHandler("OpenTab", adaptor.FromDomain(command.NewOpenTabHandler(repo)))
	h.RegisterHandler("PlaceOrder", adaptor.FromDomain(command.NewPlaceOrderHandler(repo)))

	showMenuQueryHandler := query.NewShowMenuHandler(inmemory.NewMenuRepo())
	w := handler.NewWaiter(tab.NewUseCase(adaptor.ToDomain(h)), showMenuQueryHandler)

	if err := w.OpenTab("tab777", 7); err != nil {
		log.Fatal(err)
		return
	}

	if err := w.PlaceOrder("tab777", []uint8{1, 2, 3}); err != nil {
		log.Fatal(err)
		return
	}
}
