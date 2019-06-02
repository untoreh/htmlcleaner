package HtmlCleaner

import (
	"log"
	"net/http"
	"strings"

	"github.com/untoreh/cleanup/tools"
)

type CleanerPostBody struct {
	*Cleaner
}

type CleanerPostTitle struct {
	*Cleaner
}

type LinkifyPost struct {
	*Cleaner
}

type AllPost struct {
	*Cleaner
}

func (cln *CleanerPostBody) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is for bad json, still it shouldn't have any new lines
	//buf := tools.ConvertUtf8(r.Body)
	//decoder := json.NewDecoder(buf)
	// q := r.URL.Query()
	defer r.Body.Close()

	cleaned := cln.CleanBody(r.Body)

	tools.Headers(&w)
	if _, err := w.Write([]byte(cleaned)); err != nil {
		log.Print(err)
		return
	}
}

func (cln *CleanerPostTitle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is for bad json, still it shouldn't have any new lines
	//buf := tools.ConvertUtf8(r.Body)
	//decoder := json.NewDecoder(buf)
	// q := r.URL.Query()
	defer r.Body.Close()

	cleaned := cln.CleanTitle(r.Body)

	tools.Headers(&w)
	if _, err := w.Write([]byte(cleaned)); err != nil {
		log.Print(err)
		return
	}
}

func (cln *LinkifyPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is for bad json, still it shouldn't have any new lines
	//buf := tools.ConvertUtf8(r.Body)
	//decoder := json.NewDecoder(buf)
	// q := r.URL.Query()
	defer r.Body.Close()

	cleaned := cln.Linkify(r.Body)

	tools.Headers(&w)
	if _, err := w.Write([]byte(cleaned)); err != nil {
		log.Print(err)
		return
	}
}

func (cln *AllPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is for bad json, still it shouldn't have any new lines
	//buf := tools.ConvertUtf8(r.Body)
	//decoder := json.NewDecoder(buf)
	// q := r.URL.Query()
	defer r.Body.Close()

	text := cln.CleanBody(r.Body)
	text = cln.CleanTitle(strings.NewReader(text))
	text = cln.Linkify(strings.NewReader(text))

	tools.Headers(&w)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Print(err)
		return
	}
}
