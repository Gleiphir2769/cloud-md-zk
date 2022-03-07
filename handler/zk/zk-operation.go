package zk

import (
	"cloud-md-zk/handler"
	"cloud-md-zk/lib/logger"
	"cloud-md-zk/pkg/errno"
	"cloud-md-zk/pkg/zkoperator"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateZKCluster(c *gin.Context) {
	clusterName := c.Param("clusterName")

	var config zkoperator.ZKClusterConfig
	err := c.BindJSON(config)
	if err != nil {
		logger.GetLog().Error(errno.New(errno.ErrBind, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrBind, err), nil)
	}
	status, err := zkoperator.CreateZKCluster(clusterName, &config)
	if err != nil {
		logger.GetLog().Error(errno.New(errno.ErrCreate, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrCreate, err), err)
	}
	handler.SendResponse(c, nil, status)
}

func DeleteZKCluster(c *gin.Context) {
	clusterName := c.Param("clusterName")
	if len(clusterName) == 0 {
		logger.GetLog().Error(errno.New(errno.ErrInvalidParameters, fmt.Errorf("/{clusterName} is needed")).Error())
		handler.SendResponse(c, errno.New(errno.ErrInvalidParameters, fmt.Errorf("/{clusterName} is needed")), nil)
	}
	if err := zkoperator.DeleteZKCluster(clusterName); err != nil {
		logger.GetLog().Error(errno.New(errno.ErrDelete, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrDelete, err), nil)
	}
	handler.SendResponse(c, nil, nil)
}

func UpdateZKCluster(c *gin.Context) {
	clusterName := c.Param("clusterName")

	var config zkoperator.ZKClusterConfig
	err := c.BindJSON(config)
	if err != nil {
		logger.GetLog().Error(errno.New(errno.ErrBind, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrBind, err), nil)
	}
	status, err := zkoperator.UpdateZKCluster(clusterName, &config)
	if err != nil {
		logger.GetLog().Error(errno.New(errno.ErrUpdate, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrUpdate, err), err)
	}
	handler.SendResponse(c, nil, status)
}

func GetZKCluster(c *gin.Context) {
	clusterName := c.Param("clusterName")

	if len(clusterName) == 0 {
		logger.GetLog().Error(errno.New(errno.ErrInvalidParameters, fmt.Errorf("/{gid} is needed")).Error())
		handler.SendResponse(c, errno.New(errno.ErrInvalidParameters, fmt.Errorf("/{gid} is needed")), nil)
	}
	if err := zkoperator.GetZKClusterInfo(clusterName); err != nil {
		logger.GetLog().Error(errno.New(errno.ErrGetInfo, err).Error())
		handler.SendResponse(c, errno.New(errno.ErrGetInfo, err), nil)
	}
	handler.SendResponse(c, nil, nil)
}
