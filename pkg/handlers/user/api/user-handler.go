package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
	"github.com/fillipehmeireles/user-service/core/domain/user/ports"
	"github.com/fillipehmeireles/user-service/pkg/handlers"
	"github.com/fillipehmeireles/user-service/pkg/handlers/user/dto"
)

const USER_ID_PARAM = "id"

type UserHandler struct {
	userUseCase ports.UserUseCase
}

func NewUserHandler(userUsecase ports.UserUseCase, ws *restful.WebService) *UserHandler {
	userHandler := &UserHandler{
		userUseCase: userUsecase,
	}
	ws.Route(ws.POST("/users").To(userHandler.Create))
	ws.Route(ws.GET("/users").To(userHandler.GetAll))
	ws.Route(ws.GET(fmt.Sprintf("users/{%s}", USER_ID_PARAM)).To(userHandler.GetOne))
	ws.Route(ws.PUT(fmt.Sprintf("users/{%s}", USER_ID_PARAM)).To(userHandler.Update))
	ws.Route(ws.DELETE(fmt.Sprintf("users/{%s}", USER_ID_PARAM)).To(userHandler.Delete))
	return userHandler
}

func (uH *UserHandler) Create(req *restful.Request, resp *restful.Response) {
	var newUser dto.CreateUserRequestDto
	if err := req.ReadEntity(&newUser); err != nil {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	if err := uH.userUseCase.Create(newUser); err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	resp.WriteAsJson(handlers.SuccessResponse{Data: "User Created Successfully."})
}

func (uH *UserHandler) GetAll(req *restful.Request, resp *restful.Response) {
	users, err := uH.userUseCase.GetAll()
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: users})
}

func (uH *UserHandler) GetOne(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(USER_ID_PARAM)
	if userID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoOrderUserIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	users, err := uH.userUseCase.GetOne(id)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: users})
}

func (uH *UserHandler) Delete(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(USER_ID_PARAM)
	if userID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoOrderUserIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	if err := uH.userUseCase.Delete(id); err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: "User deleted successfully."})
}

func (uH *UserHandler) Update(req *restful.Request, resp *restful.Response) {
	userID := req.PathParameter(USER_ID_PARAM)
	if userID == "" {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: handlers.ErrNoOrderUserIDProvided.Error()}))
		return
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	var updatedUser dto.UpdateUserRequestDto
	if err := req.ReadEntity(&updatedUser); err != nil {
		resp.WriteError(http.StatusBadRequest, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}

	if err := uH.userUseCase.Update(id, updatedUser); err != nil {
		resp.WriteError(http.StatusInternalServerError, resp.WriteAsJson(handlers.FailResponse{ErrorReason: err.Error()}))
		return
	}
	resp.WriteAsJson(handlers.SuccessResponse{Data: "User updated successfully."})
}
