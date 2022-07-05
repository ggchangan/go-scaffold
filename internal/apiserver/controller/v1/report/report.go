package report

import (
	"github.com/ggchangan/go-scaffold/internal/apiserver/service"
	"github.com/ggchangan/go-scaffold/internal/apiserver/store"
)

// ReportController create a report handler used to handle request for report resource.
type ReportController struct {
	srv service.Service
}

// NewReportController creates a user handler.
func NewReportController(store store.Factory) *ReportController {
	return &ReportController{
		srv: service.NewService(store),
	}
}
