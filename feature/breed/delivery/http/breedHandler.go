package delivery

import (
	"breed/constant"
	"breed/domain"
	requestForm "breed/domain/requestForm"
	"breed/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type breedHandler struct {
	breedUsecase domain.BreedUsecase
}

func NewBreedHandler(breedUsecase domain.BreedUsecase) breedHandler {
	return breedHandler{breedUsecase}
}

func (h *breedHandler) ListBreed(c *gin.Context) {
	req := requestForm.ListBreedRequest{}
	if err := c.Bind(&req); err != nil {
		errMessage := err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ErrorMessage(constant.CODE_BAD_REQUEST, &errMessage))
		return
	}

	res, err := h.breedUsecase.ListBreed(req)
	if err != nil {
		errMessage := constant.ERR_INTERNAL_SERVER_ERROR
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ErrorMessage(constant.CODE_INTERNAL_SERVER_ERROR, &errMessage))
		return
	}

	c.JSON(http.StatusOK, res)
}
