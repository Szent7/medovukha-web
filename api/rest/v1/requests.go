package rest

import (
	"context"
	"fmt"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"github.com/Szent7/medovukha-web/api/rest/v1/types"
	"google.golang.org/protobuf/encoding/protojson"

	"net/http"

	"github.com/gin-gonic/gin"
)

const mimeTypeJson = "application/json"

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
	resp, err := a.rpcClient.GetContainerList(c.Request.Context(), &dockerpb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetContainerList error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetContainerList error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) PauseContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.PauseContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "PauseContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "PauseContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) UnpauseContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.UnpauseContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "UnpauseContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "UnpauseContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) KillContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.KillContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "KillContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "KillContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) StartContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.StartContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StartContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StartContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) StopContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.StopContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StopContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StopContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) RestartContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.RestartContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RestartContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RestartContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

func (a *API) RemoveContainerByID(c *gin.Context) {
	var req types.BaseID
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.RemoveContainerByID(c.Request.Context(), &dockerpb.BaseID{Id: req.ID})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RemoveContainerByID error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RemoveContainerByID error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

// Images
func (a *API) GetImageList(c *gin.Context) {
	resp, err := a.rpcClient.GetImageList(c.Request.Context(), &dockerpb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetImageList error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetImageList error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

// Networks
func (a *API) GetNetworkList(c *gin.Context) {
	resp, err := a.rpcClient.GetNetworkList(c.Request.Context(), &dockerpb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetNetworkList error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetNetworkList error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

// Volumes
func (a *API) GetVolumeList(c *gin.Context) {
	resp, err := a.rpcClient.GetVolumeList(c.Request.Context(), &dockerpb.Empty{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetVolumeList error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetVolumeList error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

// Deploy
func (a *API) CreateFromGit(c *gin.Context) {
	var req dockerpb.DeployFromGit
	if err := c.BindJSON(&req); err != nil {
		return
	}

	resp, err := a.rpcClient.CreateFromGit(c.Request.Context(), &req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "CreateFromGit error" + err.Error()})
		fmt.Printf("gRPC error: %s\n", err.Error())
		return
	}

	jsonBytes, err := protojson.Marshal(resp)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "CreateFromGit error" + err.Error()})
		fmt.Printf("protobuf marshall error: %s\n", err.Error())
		return
	}

	c.Data(http.StatusOK, mimeTypeJson, jsonBytes)
}

// Events
func (a *API) GetContainerState(ctx context.Context, eventCh chan<- *dockerpb.ContainerState, errCh chan<- error) {
	stream, err := a.rpcClient.GetContainerState(ctx, &dockerpb.Empty{})
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
