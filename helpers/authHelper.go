package helpers

import(
	"errors"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c*gin.Context, userId string)(err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err= nil
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}