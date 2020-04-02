package server

import (
	"errors"

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
	c.JSON(statusOK, GraphQL(&payload))
}

func getPayload(c *gin.Context, payload *Payload) error {
	if c.PostForm("operations") != "" {
		return json.Unmarshal([]byte(c.PostForm("operations")), &payload)
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
