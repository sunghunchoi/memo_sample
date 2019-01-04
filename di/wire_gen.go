// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/error"
	"memo_sample/adapter/logger"
	"memo_sample/adapter/memory"
	"memo_sample/adapter/view/render"
	"memo_sample/interface/api"
	"memo_sample/usecase"

	"github.com/google/wire"
)

// Injectors from injector.go:

func InjectAPIServer() api.API {
	jsonRender := view.NewJSONRender()
	logger := loggersub.NewLogger()
	errorManager := apperrorsub.NewErrorManager()
	presenter := api.NewPresenter(jsonRender, logger, errorManager)
	transactionRepository := db.NewTransactionRepository()
	memoRepository := db.NewMemoRepository()
	tagRepository := db.NewTagRepository()
	memo := usecase.NewMemo(transactionRepository, memoRepository, tagRepository, errorManager)
	interactor := usecase.NewInteractor(presenter, memo)
	apiAPI := api.NewAPI(interactor, logger)
	return apiAPI
}

// injector.go:

// ProvideAPI inject api using wire
var ProvideAPI = wire.NewSet(
	ProvideUsecaseIterator, api.NewAPI,
)

// ProvidePresenter inject presenter using wire
var ProvidePresenter = wire.NewSet(
	ProvideRender,
	ProvideLog, api.NewPresenter, ProvideErrorManager,
)

// ProvideMemoUsecase inject memo usecase using wire
var ProvideMemoUsecase = wire.NewSet(
	ProvideDBRepository, usecase.NewMemo,
)

// ProvideUsecaseIterator inject usecase itetator using wire
var ProvideUsecaseIterator = wire.NewSet(
	ProvidePresenter,
	ProvideMemoUsecase, usecase.NewInteractor,
)

// ProvideInMemoryRepository inject repository using wire
var ProvideInMemoryRepository = wire.NewSet(memory.NewTransactionRepository, memory.NewMemoRepository, memory.NewTagRepository)

// ProvideDBRepository inject repository using wire
var ProvideDBRepository = wire.NewSet(db.NewTransactionRepository, db.NewMemoRepository, db.NewTagRepository)

// ProvideLog inject log using wire
var ProvideLog = wire.NewSet(loggersub.NewLogger)

// ProvideRender inject render using wire
var ProvideRender = wire.NewSet(view.NewJSONRender)

// ProvideErrorManager inject error manager using wire
var ProvideErrorManager = wire.NewSet(apperrorsub.NewErrorManager)
