/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queries

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/queries/maintainer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Queries = base.NewQueries(maintainer.Query)