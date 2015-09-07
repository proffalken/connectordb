package webclient

/* Provides a user facing website for CDB. There are three template/site files
associate with this:

- templates -- this is for go templates, none are currently used but we may
  need them in the future (see userweb.go)
- spa -- the location of the single page application, this is a static folder.
- site -- the location for the website. / will be redirected here "/site/"

Copyright 2015 - The ConnectorDB Contributors; see AUTHORS for a list of authors.
All Rights Reserved
*/

import (
	"connectordb/streamdb"
	"net/http"
	"path"

	log "github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/kardianos/osext"
)

var (
	userdb *streamdb.Database
)

// Sets up a static site fetching path
func setupStaticPath(subroutePath string, subroutePrefix *mux.Router) {
	folderPath, _ := osext.ExecutableFolder()
	includepath := path.Join(folderPath, subroutePath)
	fileserver := http.FileServer(http.Dir(includepath))

	httpPath := "/" + subroutePath + "/"

	log.Infof("Setting up subroute %v at %v", httpPath, folderPath)

	subroutePrefix.PathPrefix(httpPath).Handler(http.StripPrefix(httpPath, fileserver))
}

func Setup(subroutePrefix *mux.Router, udb *streamdb.Database) {
	userdb = udb

	setupStaticPath("spa", subroutePrefix)
	setupStaticPath("site", subroutePrefix)

	// Main site comes last
	subroutePrefix.Handle("/", http.RedirectHandler("/site/", 307))
}
