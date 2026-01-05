package rest

import (
	"context"
	"log"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"github.com/Szent7/medovukha-web/api/rest/v1/types"
	"github.com/Szent7/medovukha-web/api/validation"

	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	rpcClient dockerpb.DockerServiceClient
}

func NewAPI(rpcClient dockerpb.DockerServiceClient) *API {
	return &API{
		rpcClient: rpcClient,
	}
}

// Containers
func (a *API) GetContainerList(c *gin.Context) {
	resp, err := a.rpcClient.GetContainerList(c.Request.Context(), &dockerpb.GetContainerListRequest{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) PauseContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.PauseContainerByID(c.Request.Context(), &dockerpb.PauseContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) UnpauseContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.UnpauseContainerByID(c.Request.Context(), &dockerpb.UnpauseContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) KillContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.KillContainerByID(c.Request.Context(), &dockerpb.KillContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) StartContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.StartContainerByID(c.Request.Context(), &dockerpb.StartContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) StopContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.StopContainerByID(c.Request.Context(), &dockerpb.StopContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) RestartContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.RestartContainerByID(c.Request.Context(), &dockerpb.RestartContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) RemoveContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.RemoveContainerByID(c.Request.Context(), &dockerpb.RemoveContainerByIDRequest{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

// Images
func (a *API) GetImageList(c *gin.Context) {
	resp, err := a.rpcClient.GetImageList(c.Request.Context(), &dockerpb.GetImageListRequest{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

// Networks
func (a *API) GetNetworkList(c *gin.Context) {
	resp, err := a.rpcClient.GetNetworkList(c.Request.Context(), &dockerpb.GetNetworkListRequest{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

// Volumes
func (a *API) GetVolumeList(c *gin.Context) {
	resp, err := a.rpcClient.GetVolumeList(c.Request.Context(), &dockerpb.GetVolumeListRequest{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

// Deploy
func (a *API) CreateFromGit(c *gin.Context) {
	var req types.CreateFromGit
	if err := c.ShouldBindJSON(&req); err != nil {
		validation.HandleBindError(c, err)
		log.Printf("parse error: %s\n", err.Error())
		return
	}

	resp, err := a.rpcClient.CreateFromGit(c.Request.Context(), &dockerpb.CreateFromGitRequest{
		Url:           req.URL,
		Dockerfile:    req.Dockerfile,
		DockerCompose: req.DockerCompose,
		DockerRun:     req.DockerRun,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError,
			types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrCore, Message: err.Error()}))
		log.Printf("gRPC error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, types.NewSuccess(resp))
}

func (a *API) StreamBuildLogs(ctx context.Context, buildID string, logCh chan<- *dockerpb.StreamBuildLogsResponse, errCh chan<- error) {
	stream, err := a.rpcClient.StreamBuildLogs(ctx, &dockerpb.StreamBuildLogsRequest{BuildId: buildID})
	if err != nil {
		errCh <- err
		return
	}

	for {
		log, err := stream.Recv()
		if err != nil {
			errCh <- err
			return
		}
		logCh <- log
	}
}

// Events
func (a *API) GetContainerState(ctx context.Context, eventCh chan<- *dockerpb.GetContainerStateResponse, errCh chan<- error) {
	stream, err := a.rpcClient.GetContainerState(ctx, &dockerpb.GetContainerStateRequest{})
	if err != nil {
		errCh <- err
		return
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			errCh <- err
			return
		}
		eventCh <- event
	}
}
