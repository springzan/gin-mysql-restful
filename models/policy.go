package models

import (
	. "gin-mysql-restful/database"
)

type Policy struct {
	PolicyId int `json:"policyId" form:"policyId"`
	CompanyId int `json:"companyId" form:"companyId"`
	BrandId int `json:"brandId" form:"brandId"`
	AgentId int `json:"agentId" form:"agentId"`
}

func (p Policy) GetOne() (policy Policy, err error) {

	row := SqlDB.QueryRow("SELECT policy_id, company_id, brand_id,agent_id FROM policy WHERE policy_id=?", p.PolicyId)
	if row == nil {
		return
	}
	err = row.Scan(&policy.PolicyId, &policy.CompanyId, &policy.BrandId,&policy.AgentId)
	if err != nil {
		return
	}
	return
}