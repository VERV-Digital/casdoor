// Copyright 2021 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"fmt"
	"github.com/casdoor/casdoor/util"
)

type ShortRole struct {
	Owner     string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name      string `xorm:"varchar(100) notnull pk" json:"name"`
	IsEnabled bool   `json:"isEnabled"`
}

func getShortRole(owner string, name string) *ShortRole {
	if owner == "" || name == "" {
		return nil
	}

	role := ShortRole{Owner: owner, Name: name}
	existed, err := adapter.Engine.Get(&role)
	if err != nil {
		panic(err)
	}

	if existed {
		return &role
	} else {
		return nil
	}
}

func GetShortRole(id string) *ShortRole {
	owner, name := util.GetOwnerAndNameFromId(id)
	return getShortRole(owner, name)
}

func (role *ShortRole) GetId() string {
	return fmt.Sprintf("%s/%s", role.Owner, role.Name)
}

func GetShortRolesByUser(userId string) []*ShortRole {
	roles := []*ShortRole{}
	err := adapter.Engine.Table("role").Where("users like ?", "%"+userId+"%").Find(&roles)
	if err != nil {
		panic(err)
	}

	return roles
}
