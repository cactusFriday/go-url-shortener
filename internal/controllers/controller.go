package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cactusfriday/go-url-shortener/internal/common"
	"github.com/cactusfriday/go-url-shortener/internal/storage"
	"github.com/cactusfriday/go-url-shortener/internal/storage/base"
	"github.com/gorilla/mux"
)

type Url struct {
	Url string `json:"url"`
}

type IUrlController interface {
	GetUrl() (string, error)
	CreateShortUrl() (string, error)
}

type UrlController struct {
	storageDriver base.IStorage
	request       *http.Request
}

// Return URL saved in storage
func (u *UrlController) GetUrl() (string, error) {
	short_url := mux.Vars(u.request)["short_url"]
	url, err := u.storageDriver.GetUrl(short_url)

	if err != nil {
		return "", err
	}
	return url, nil
}

// Encode URL, save in storage and return encoded short URL
func (u *UrlController) CreateShortUrl() (string, error) {
	var reqBody Url
	decoder := json.NewDecoder(u.request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&reqBody)

	if err != nil {
		return "", errors.New(fmt.Sprintf("JSON Decode error: %s", err))
	}

	shortUrl := common.EncodeString(reqBody.Url)

	err = u.storageDriver.PutUrl(shortUrl, reqBody.Url)

	if err != nil {
		return "", errors.New(fmt.Sprintf("Creating Short URL error: %s", err))
	}
	return shortUrl, nil
}

// Controller constructor. Load config, storage and get request
func InstantiateController(r *http.Request) IUrlController {
	st := common.CFG.StorageType
	storage := storage.NewStorage(st)

	return &UrlController{
		storageDriver: storage,
		request:       r,
	}
}
