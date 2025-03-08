package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/Gharib110/LaraGo/filesystems"
	"github.com/Gharib110/LaraGo/filesystems/miniofilesystem"
	"github.com/Gharib110/LaraGo/filesystems/s3filesystem"
	"github.com/Gharib110/LaraGo/filesystems/sftpfilesystem"
	"github.com/Gharib110/LaraGo/filesystems/webdavfilesystem"
	"io"
	"net/http"
	"net/url"
	"os"
	"test-app/data"

	"github.com/Gharib110/LaraGo"
)

// Handlers is the type for handlers, and gives access to LaraGo and models
type Handlers struct {
	App    *lara.Lara
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) LaraUpload(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "lara-upload", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostLaraUpload(w http.ResponseWriter, r *http.Request) {
	err := h.App.UploadFile(r, "", "formFile", &h.App.S3)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Session.Put(r.Context(), "error", err.Error())
	} else {
		h.App.Session.Put(r.Context(), "flash", "Uploaded!")
	}
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func (h *Handlers) ListFS(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	var list []filesystems.Listing

	fsType := ""
	if r.URL.Query().Get("fs-type") != "" {
		fsType = r.URL.Query().Get("fs-type")
	}

	curPath := "/"
	if r.URL.Query().Get("curPath") != "" {
		curPath = r.URL.Query().Get("curPath")
		curPath, _ = url.QueryUnescape(curPath)
	}

	if fsType != "" {
		switch fsType {
		case "MINIO":
			f := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
			fs = &f
			fsType = "MINIO"

		case "SFTP":
			f := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
			fs = &f
			fsType = "SFTP"

		case "WEBDAV":
			f := h.App.FileSystems["WEBDAV"].(webdavfilesystem.WebDAV)
			fs = &f
			fsType = "WEBDAV"

		case "S3":
			f := h.App.FileSystems["S3"].(s3filesystem.S3)
			fs = &f
			fsType = "S3"
		}

		l, err := fs.List(curPath)
		if err != nil {
			h.App.ErrorLog.Println(err)
			return
		}

		list = l
	}

	vars := make(jet.VarMap)
	vars.Set("list", list)
	vars.Set("fs_type", fsType)
	vars.Set("curPath", curPath)
	err := h.render(w, r, "list-fs", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) UploadToFS(w http.ResponseWriter, r *http.Request) {
	fsType := r.URL.Query().Get("type")

	vars := make(jet.VarMap)
	vars.Set("fs_type", fsType)

	err := h.render(w, r, "upload", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostUploadToFS(w http.ResponseWriter, r *http.Request) {
	fileName, err := getFileToUpload(r, "formFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadType := r.Form.Get("upload-type")

	switch uploadType {
	case "MINIO":
		fs := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
		err = fs.Put(fileName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "SFTP":
		fs := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
		err = fs.Put(fileName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "WEBDAV":
		fs := h.App.FileSystems["WEBDAV"].(webdavfilesystem.WebDAV)
		err = fs.Put(fileName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "S3":
		fs := h.App.FileSystems["S3"].(s3filesystem.S3)
		err = fs.Put(fileName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	h.App.Session.Put(r.Context(), "flash", "File uploaded!")
	http.Redirect(w, r, "/files/upload?type="+uploadType, http.StatusSeeOther)
}

func getFileToUpload(r *http.Request, fieldName string) (string, error) {
	_ = r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./tmp/%s", header.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("./tmp/%s", header.Filename), nil
}

func (h *Handlers) DeleteFromFS(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	fsType := r.URL.Query().Get("fs_type")
	item := r.URL.Query().Get("file")

	switch fsType {
	case "MINIO":
		f := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
		fs = &f
	case "SFTP":
		f := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
		fs = &f
	case "WEBDAV":
		f := h.App.FileSystems["WEBDAV"].(webdavfilesystem.WebDAV)
		fs = &f
	case "S3":
		f := h.App.FileSystems["S3"].(s3filesystem.S3)
		fs = &f
	}

	deleted := fs.Delete([]string{item})
	if deleted {
		h.App.Session.Put(r.Context(), "flash", fmt.Sprintf("%s was deleted", item))
		http.Redirect(w, r, "/list-fs?fs-type="+fsType, http.StatusSeeOther)
	}
}
