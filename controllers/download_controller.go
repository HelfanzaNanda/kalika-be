package controllers

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/labstack/echo"
)

type DownloadController interface {
	DownloadPdf(ctx echo.Context) error
}
type DownloadControllerImpl struct {
		
}

func NewDownloadController() DownloadController {
	return &DownloadControllerImpl{
		
	}
}

func (dc * DownloadControllerImpl) DownloadPdf(ctx echo.Context) error {
	params,_ := ctx.FormParams()
	filename := strings.TrimSpace(params.Get("path"))
	f, err := os.Open(filename)
    if err != nil {
		return ctx.JSON(500, err.Error())
    }
	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	ctx.Response().Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(ctx.Response(), f); err != nil {
		return ctx.JSON(500, err.Error())
    }
	return err
}