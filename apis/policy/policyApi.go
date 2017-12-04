package policy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "gin-mysql-restful/models"
	"strconv"
	"log"
	. "gin-mysql-restful/boot"
)


func GetPolicyApi(c *gin.Context) {
	if !SignAndUser(c) {
		return
	}

	cid := c.Param("policyId")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Policy {PolicyId: id}
	policy, err := p.GetOne()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"policy": policy,
		},
	})
}
