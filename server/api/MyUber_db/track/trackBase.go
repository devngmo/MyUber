/* 
* Generated by
* 
*      _____ _          __  __      _     _
*     / ____| |        / _|/ _|    | |   | |
*    | (___ | | ____ _| |_| |_ ___ | | __| | ___ _ __
*     \___ \| |/ / _` |  _|  _/ _ \| |/ _` |/ _ \ '__|
*     ____) |   < (_| | | | || (_) | | (_| |  __/ |
*    |_____/|_|\_\__,_|_| |_| \___/|_|\__,_|\___|_|
*
* The code generator that works in many programming languages
*
*			https://www.skaffolder.com
*
*
* You can generate the code from the command-line
*       https://npmjs.com/package/skaffolder-cli
*
*       npm install -g skaffodler-cli
*
*   *   *   *   *   *   *   *   *   *   *   *   *   *   *   *
*
* To remove this comment please upgrade your plan here: 
*      https://app.skaffolder.com/#!/upgrade
*
* Or get up to 70% discount sharing your unique link:
*       https://app.skaffolder.com/#!/register?friend=5d525f5a7b158e50f28eb4af
*
* You will get 10% discount for each one of your friends
* 
*/
package track

/**
 * 
 * 
  _____                      _              _ _ _     _   _     _        __ _ _      
 |  __ \                    | |            | (_) |   | | | |   (_)      / _(_) |     
 | |  | | ___    _ __   ___ | |_    ___  __| |_| |_  | |_| |__  _ ___  | |_ _| | ___ 
 | |  | |/ _ \  | '_ \ / _ \| __|  / _ \/ _` | | __| | __| '_ \| / __| |  _| | |/ _ \
 | |__| | (_) | | | | | (_) | |_  |  __/ (_| | | |_  | |_| | | | \__ \ | | | | |  __/
 |_____/ \___/  |_| |_|\___/ \__|  \___|\__,_|_|\__|  \__|_| |_|_|___/ |_| |_|_|\___|
 
 * DO NOT EDIT THIS FILE!! 
 * 
 *  TO CUSTOMIZE TrackBase.go PLEASE EDIT ./Track.go
 * 
 *  -- THIS FILE WILL BE OVERWRITTEN ON THE NEXT SKAFFOLDER'S CODE GENERATION --
 * 
 */

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Data model
type Track struct {
	ID		bson.ObjectId	`bson:"_id,omitempty" json:"_id"`
	BookID	string		`json:"BookID"`
	Status	int		`json:"Status"`


}

// Functions

// CRUD METHODS
	
// CRUD - CREATE
func (config *Config) create(writer http.ResponseWriter, req *http.Request) {

	var body Track
	json.NewDecoder(req.Body).Decode(&body)
	config.Database.C("Track").Insert(body)
	render.JSON(writer, req, &body)
}

// CRUD - DELETE
func (config *Config) delete(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	config.Database.C("Track").RemoveId(bson.ObjectIdHex(id))
	render.JSON(writer, req, id)
}

// CRUD - GET
func (config *Config) get(writer http.ResponseWriter, req *http.Request) {

	var result Track
	id := chi.URLParam(req, "id")
	err := config.Database.C("Track").FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		panic(err)
	}

	render.JSON(writer, req, result)
}
	
// CRUD - GET LIST
func (config *Config) list(writer http.ResponseWriter, req *http.Request) {

	var result []Track
	err := config.Database.C("Track").Find(nil).All(&result)
	if err != nil {
		panic(err)
	}

	render.JSON(writer, req, result)
}

// CRUD - UPDATE
func (config *Config) update(writer http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")
	var body Track
	json.NewDecoder(req.Body).Decode(&body)
	config.Database.C("Track").UpdateId(bson.ObjectIdHex(id), body)
	render.JSON(writer, req, body)
}

/*
 * CUSTOM SERVICES
 * 
 *	These services will be overwritten and implemented in  Ext.go
 */


/*
Name: LatestTrackByCarID
Description: Get latest track by car ID
Params: 
*/
func (config *Config) LatestTrackByCarID(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]string)
	render.JSON(writer, request, response)
}		