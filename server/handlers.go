package server

import (
	"errors"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	statusOK         = 200
	statusBadRequest = 400
	statusNotFound   = 404
)

// Payload query structure
type Payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

// ImgHandler used by grahpql schemas
// Accepts data from Form Data and Request Payload
func grahpqlHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	// получение variables
	variables, err := getVariables(c)
	if err != nil {
		c.JSON(statusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload := Payload{
		Query:     c.PostForm("query"),
		Variables: variables,
	}

	err = getPayload(c, &payload)
	if err != nil {
		c.JSON(statusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uploadsMap, err := getMap(c)
	if err != nil {
		c.JSON(statusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.Variables["map"] = uploadsMap
	// multipart form
	form, _ := c.MultipartForm()
	if form != nil {
		var files = make(map[string]*multipart.FileHeader)
		for i := range form.File {
			files[i] = form.File[i][0]
		}
		payload.Variables["files"] = files
	}
	// fmt.Println(payload.Variables)
	c.JSON(statusOK, GraphQL(&payload))
}

func getPayload(c *gin.Context, payload *Payload) error {
	if c.PostForm("operations") != "" {
		err := json.Unmarshal([]byte(c.PostForm("operations")), &payload)
		if err != nil {
			return err
		}
		return nil
	}
	// check for existence of data from Form Data
	if payload.Query == "" {
		// if not, then we take from Request Payload
		if c.Request.Body == nil {
			return errors.New("Please send a Payload or Form Data")
		}
		err := json.NewDecoder(c.Request.Body).Decode(&payload)
		if err != nil {
			return err
		}
	}
	return nil
}

func getMap(c *gin.Context) (map[string][]string, error) {
	var uploadsMap = make(map[string][]string)
	if c.PostForm("map") != "" {
		err := json.Unmarshal([]byte(c.PostForm("map")), &uploadsMap)
		if err != nil {
			return nil, err
		}
	}
	return uploadsMap, nil
}

func getVariables(c *gin.Context) (map[string]interface{}, error) {
	var variables = make(map[string]interface{})
	if c.PostForm("variables") != "" {
		err := json.Unmarshal([]byte(c.PostForm("variables")), &variables)
		if err != nil {
			return nil, err
		}
	}
	return variables, nil
}

func optionsHandler(c *gin.Context) {
	c.JSON(statusOK, "")
}
