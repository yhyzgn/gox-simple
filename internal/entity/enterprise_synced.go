// Copyright 2019 yhyzgn enp-simple
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-05-15 16:06
// version: 1.0.0
// desc   : 从企业库同步过来的企业数据，待人工清洗

package entity

import (
	"gox-simple/internal/entity/table"
	"time"
)

// 企业基础数据
type EnterpriseBasic struct {
	EnterpriseID       uint64    `gorm:"column:enterprise_id;type:bigint(20);comment:'企业在企业库中的ID'" json:"enterpriseId"`
	RegionCode         string    `gorm:"column:region_code;type:varchar(255);comment:'行政区划编码'" json:"regionCode"`
	CreditCode         string    `gorm:"column:credit_code;type:varchar(255);comment:'统一社会信用代码'" json:"creditCode"`
	BusinessLicense    string    `gorm:"column:business_license;type:varchar(255);comment:'业务经营许可证号'" json:"businessLicense"`
	OrganizationCode   string    `gorm:"column:organization_code;type:varchar(255);comment:'组织机构代码'" json:"organizationCode"`
	Name               string    `gorm:"column:name;type:varchar(255);comment:'门店名称'" json:"name"`
	BusinessStatus     uint      `gorm:"column:business_status;type:int(11);comment:'经营状态'" json:"businessStatus"`
	Address            string    `gorm:"column:address;type:varchar(255);comment:'经营地址'" json:"address"`
	RegisterCode       string    `gorm:"column:register_code;type:varchar(255);comment:'注册号'" json:"registerCode"`
	RegisterName       string    `gorm:"column:register_name;type:varchar(255);comment:'工商注册名称'" json:"registerName"`
	RegisterAddress    string    `gorm:"column:register_address;type:varchar(255);comment:'工商注册地址'" json:"registerAddress"`
	IndustryCategory   uint      `gorm:"column:industry_category;type:int(11);comment:'行业类别'" json:"industryCategory"`
	OnlineName         string    `gorm:"column:online_name;type:varchar(255);comment:'OTA线上名称'" json:"onlineName"`
	LawName            string    `gorm:"column:law_name;type:varchar(255);comment:'法人姓名'" json:"lawName"`
	LawTel             string    `gorm:"column:law_tel;type:varchar(255);comment:'法人联系方式'" json:"lawTel"`
	LawCertificateType string    `gorm:"column:law_certificate_type;type:varchar(255);comment:'法人证件类型'" json:"lawCertificateType"`
	LawCertificateCode string    `gorm:"column:law_certificate_code;type:varchar(255);comment:'法人证件号码'" json:"lawCertificateCode"`
	RegisterCapital    float64   `gorm:"column:register_capital;type:decimal;comment:'注册资本（万元）'" json:"registerCapital"`
	RegisterTime       time.Time `gorm:"column:register_time;type:datetime;comment:'注册时间'" json:"registerTime"`
	TaxpayerCode       string    `gorm:"column:taxpayer_code;type:varchar(255);comment:'纳税人识别号'" json:"taxpayerCode"`
	BusinessScope      string    `gorm:"column:business_scope;type:varchar(255);comment:'经营范围'" json:"businessScope"`
	RegisterAuthority  string    `gorm:"column:register_authority;type:varchar(255);comment:'登记机关'" json:"registerAuthority"`
	BusinessTerm       time.Time `gorm:"column:business_term;type:datetime;comment:'营业期限'" json:"businessTerm"`
	ApprovalDate       time.Time `gorm:"column:approval_date;type:datetime;comment:'核准日期'" json:"approvalDate"`
}

// 刚从企业库同步过来时你的样子，尚未经过手指的洗礼
type EnterpriseSynced struct {
	EnterpriseBasic
	table.Entity
}

func (es EnterpriseSynced) TableName() string {
	return es.NameWithPrefix("enterprise_synced")
}
