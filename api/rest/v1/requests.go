package v1

import (
	"fmt"
	"medovukha/api/rest/v1/types"

	"net/http"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

type API struct {
	rpcClient   *rpc.Client
	serviceName string
}

func NewAPI(rpcClient *rpc.Client, serviceName string) *API {
	return &API{
		rpcClient:   rpcClient,
		serviceName: serviceName,
	}
}

func (a *API) Call(method string, args any, reply any) error {
	return a.rpcClient.Call(a.serviceName+"."+method, args, reply)
}

// Containers
func (a *API) GetContainerList(c *gin.Context) {
	var conList []types.ContainerBaseInfo
	if err := a.Call("GetContainerList", &types.Empty{}, &conList); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetContainerList error" + err.Error()})
		fmt.Printf("GetContainerList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, conList)
}

func (a *API) PauseContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("PauseContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "PauseContainerByID error" + err.Error()})
		fmt.Printf("PauseContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) UnpauseContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("UnpauseContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "UnpauseContainerByID error" + err.Error()})
		fmt.Printf("UnpauseContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) KillContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("KillContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "KillContainerByID error" + err.Error()})
		fmt.Printf("KillContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) StartContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("StartContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StartContainerByID error" + err.Error()})
		fmt.Printf("StartContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) StopContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("StopContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "StopContainerByID error" + err.Error()})
		fmt.Printf("StopContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) RestartContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("RestartContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RestartContainerByID error" + err.Error()})
		fmt.Printf("RestartContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

func (a *API) RemoveContainerByID(c *gin.Context) {
	var args types.BaseID
	if err := c.BindJSON(&args); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("RemoveContainerByID", &args, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "RemoveContainerByID error" + err.Error()})
		fmt.Printf("RemoveContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}

// Images
func (a *API) GetImageList(c *gin.Context) {
	var imgList []types.ImageBaseInfo
	if err := a.Call("GetImageList", &types.Empty{}, &imgList); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetImageList error" + err.Error()})
		fmt.Printf("GetImageList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}

// Networks
func (a *API) GetNetworkList(c *gin.Context) {
	var ntwList []types.NetworkBaseInfo
	if err := a.Call("GetNetworkList", &types.Empty{}, &ntwList); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetNetworkList error" + err.Error()})
		fmt.Printf("GetNetworkList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ntwList)
}

// Volumes
func (a *API) GetVolumeList(c *gin.Context) {
	var vlmList []types.VolumeBaseInfo
	if err := a.Call("GetVolumeList", &types.Empty{}, &vlmList); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "GetVolumeList error" + err.Error()})
		fmt.Printf("GetVolumeList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, vlmList)
}

// Deploy
func (a *API) CreateFromGit(c *gin.Context) {
	var NewDeploy types.DeployFromGit
	if err := c.BindJSON(&NewDeploy); err != nil {
		return
	}

	var message types.BaseMessage
	if err := a.Call("CreateFromGit", &NewDeploy, &message); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, types.BaseMessage{Message: "CreateFromGit error" + err.Error()})
		fmt.Printf("CreateFromGit error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, message)
}
