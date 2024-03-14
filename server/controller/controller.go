package controller

import (
	"log"
	"net/http"
	"server/service"
	"server/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
The Controller handles sending data to the user using g.Context objects.
*/

type Controller struct {
	serv service.Service
}

func New(s service.Service) Controller {
	return Controller{serv: s}
}

func (c *Controller) GetAllCasesForLawyer(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	lIDString := params.Get("lawyer_id")
	if lIDString == "" {
		log.Println("no Lawyer ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no lawyer id provided"})
		return
	}

	lID, _ := strconv.ParseInt(params.Get("lawyer_id"), 10, 64)

	cases, err := c.serv.GetAllCases(int(lID))

	if err == service.ErrInvalidLawyerID {
		log.Println("Invalid Lawyer ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.GetAllCasesForLawyer()")
	ctx.JSON(http.StatusOK, cases)
}

func (c *Controller) CreateNewCase(ctx *gin.Context) {

	var newCase types.Case
	var resp types.NewCaseResp
	var err error

	if err = ctx.BindJSON(&newCase); err != nil {
		log.Println("incorrect case format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect case format"})
		return
	}

	resp.CaseID, resp.LawyerName, err = c.serv.CreateNewCase(newCase)

	if err == service.ErrQueryFailure {
		log.Println("db error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.CreateNewCase()")
	ctx.JSON(http.StatusOK, resp)
}
