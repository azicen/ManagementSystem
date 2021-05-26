package service

import (
	"management-system-server/core"
	"management-system-server/model"
	"net/http"
	"time"
)

//AddMaintain AddMaintain服务api逻辑处理
func (s *Service) AddMaintain(c *core.Context) {
	info, err := s.dao.BindHTTPAddMaintainInfo(c)
	if err != nil {
		c.Error(http.StatusBadRequest, err)
		return
	}

	m, err := s.dao.AddMaintain(c, info.UserID, info.EquipmentID, time.Unix(int64(info.Date), 0), info.Status, info.Remark)
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPAddMaintainResponse{
		ID: m.ID,
	}
	c.SetData(response)
}

//GetMaintain GetMaintain服务api逻辑处理
func (s *Service) GetMaintain(c *core.Context) {
	m, err := s.dao.GetMaintain(c, c.Param("id"))
	if err != nil {
		c.Error(http.StatusInternalServerError, err)
		return
	}

	response := model.HTTPGetMaintainResponse{
		UserID:      m.UserID,
		EquipmentID: m.EquipmentID,
		Date:        uint64(m.Date.Unix()),
		Status:      m.Status,
		Remark:      m.Remark,
	}
	c.SetData(response)
}