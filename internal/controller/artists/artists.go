package artists

import (
	"github/irvantaufik/album/internal/entity"
	"github/irvantaufik/album/internal/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller *artirsController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("artist not found",
			err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	artist, err := controller.artists.FindById(ctx, id)

	if err != nil {
		res := helper.BuildErrorResponse("internal server error",
			err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "success", artist)
	ctx.JSON(http.StatusOK, res)
}

func (controller *artirsController) Create(ctx *gin.Context) {
	var requestBody entity.Artists

	err := ctx.BindJSON(&requestBody)

	if err != nil {
		res := helper.BuildErrorResponse("Bad Request", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)

		return
	}

	artists, err := controller.artists.Create(ctx, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", artists)
	ctx.JSON(http.StatusOK, res)

}
