package main

import (
	"encoding/json"
	"fmt"
	"github.com/durolavoro/snippet-app/dal"
	"github.com/durolavoro/snippet-app/model"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type App struct {
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
	DAL         *dal.DAL
}

func (a *App) Home(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := SendResponse(w, map[string]string{
		"version": "v1",
		"status":  "OK",
	}, http.StatusOK)
	if err != nil {
		a.ServerError(w, err)
		a.ErrorLogger.Println("error sending response for endpoint /", err.Error())
	}
}

func (a *App) ShowSnippet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	idStr := r.URL.Query().Get("id")
	var err error
	if strings.EqualFold(idStr, "") {
		snips, err := a.DAL.Latest(1)
		if err != nil || len(snips) == 0 {
			a.ServerError(w, err)
		}
		snip := *snips[0]
		_ = SendResponse(w, map[string]string{
			"message": fmt.Sprintf("Snippet %d is titled %s, with content %s. It expires in %d seconds",
				snip.ID, snip.Title, snip.Content, snip.Expires.Unix()-snip.Created.Unix()),
		}, http.StatusOK)
		return
	}

	id, convErr := strconv.Atoi(idStr)
	if convErr != nil {
		a.ReqError(w, fmt.Errorf("invalid value for param id"))
		return
	}

	snip, err := a.DAL.Get(id)
	if err != nil {
		a.ServerError(w, err)
	}
	_ = SendResponse(w, map[string]string{
		"message": fmt.Sprintf("Snippet %d is titled %s, with content %s. It expires in %d seconds",
			snip.ID, snip.Title, snip.Content, snip.Expires.Unix()-snip.Created.Unix()),
	}, http.StatusOK)
}

func (a *App) CreateSnippet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	p, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
	}
	snip := &model.CreateSnippetPayload{}
	err = json.Unmarshal(p, &snip)
	if err != nil {
		// TODO
	}
	_, err = a.DAL.Insert(snip.Title, snip.Content, strconv.Itoa(snip.Expires))
	if err != nil {
		// TODO
	}
	SendResponse(w, map[string]string{
		"message": "added new snippet successfully",
	}, http.StatusOK)
}

func (a *App) CreateSnippets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		// TODO
	}
	snips := model.CreateSnippetsPayload{}
	err = json.Unmarshal(p, &snips)
	if err != nil {
	}
	
	errChan := make(chan error, len(snips.Snippets))
	var wg sync.WaitGroup
	for i := 0; i < len(snips.Snippets); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			snip := snips.Snippets[idx]
			expiryTime := time.Now().Add(time.Duration(snip.Expires*24) * time.Hour)
			_, err := a.DAL.InsertError(snip.Title, snip.Content, expiryTime.String())
			if err != nil {
				errChan <- fmt.Errorf("error inserting query snippet %d : %w", idx, err)
			}
		}(i)
	}
	wg.Wait()
	close(errChan)
	a.InfoLogger.Printf("concurrent inserts have concluded with %d errors", len(errChan))
	var errMap = make(map[string]string, 0)
	errCounter := 0
	for errs := range errChan {
		errMap[fmt.Sprintf("err%d", errCounter)] = errs.Error()
		errCounter++
	}
	_ = SendResponse(w, errMap, http.StatusInternalServerError)
}
