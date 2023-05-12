package controller

import (
	"aBet/model"
	"aBet/usecase/service"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var WHILE_LIST_EXTENDSION = []string{
	"pdf",
	"ppt",
	"pptx",
	"ppsx",
	"pps",
	"pot",
	"potx",
	"docx",
	"doc",
	"dot",
	"dotx",
	"sldx",
	"xls",
	"xlsx",
	"xlt",
	"xltx",
	"jpg",
	"jpeg",
	"png",
	"gif",
	"svg",
	"tiff",
	"pdf",
	"eps",
	"txt",
	"zip",
	"rar",
	"mp3",
	"mp4",
	"mpeg",
	"m4a",
}

type FileController interface {
	AddNewFile(c *Context) error
	ReadFile(c *Context) error
}

type fileController struct {
	fileUploaderService (service.FileUploaderService)
	fileGetterSerivce   service.FileGetterSerivce
}

func NewFileController(
	s service.FileUploaderService,
	dF service.FileGetterSerivce,

) FileController {
	return &fileController{
		fileUploaderService: s,
		fileGetterSerivce:   dF,
	}
}

/*
	nếu có Params fileName -> ưu tiên lấy fileName
	nếu không thì lấy fileName từ file gửi đến, split, chỉ lấy tên, tách typeFile
*/

func (fC *fileController) AddNewFile(c *Context) error {
	fileUpload, er := c.Context.FormFile("file")
	if er != nil {
		return c.Output(http.StatusBadRequest, nil, errors.New("Upload file need file"))
	}
	var fileName string
	uploadFileName := fileUpload.Filename
	fileNameSplit := strings.Split(uploadFileName, ".")
	fileType := fileNameSplit[len(fileNameSplit)-1]
	fileName = "user_avatar_" + c.Context.FormValue("userId") + "." + fileType
	file := model.File{
		Name: fileName,
	}
	newFile, err := fC.fileUploaderService.AddNewFile(&file, fileUpload)
	if err != nil {
		return c.Output(http.StatusBadRequest, nil, err)
	}
	// fileOutGoing := outgoing.ParseFileOutGoing(*newFile)
	return c.Output(http.StatusOK, newFile, err)
}

func (fC *fileController) ReadFile(c *Context) error {
	fullFileName := c.Param("fileName")
	// if c.QueryParam("tenantId") == "" {
	// 	return c.File("uploads/white_background.jpg")
	// }
	// tenantId := obj.TenantId
	file := model.File{
		Name: fullFileName,
	}
	path, e := fC.fileGetterSerivce.GetFileByName(file)
	fmt.Println(path)
	if e == nil {
		return c.File(path)
	}

	// return c.Output(http.StatusNotFound, nil, errors.New("fileName does not exist"))
	return c.File("static/default_avatar.jpg")
}
