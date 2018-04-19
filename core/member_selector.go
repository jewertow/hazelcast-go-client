// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

// MemberSelector is an interface for implementations selecting members
// that are capable of executing a special kind of task.
// The Select method is called for every available
// member in the cluster and it is up to the implementation to decide
// if the member is going to be used or not.
type MemberSelector interface {
	// Select decides if the given member will be part of an operation or not.
	// Select returns true if the member should take part in the operation, false otherwise.
	Select(member IMember) (selected bool)
}

// MemberSelectors is a utility variable to get MemberSelector instances.
var MemberSelectors = &selectors{
	&dataMemberSelector{},
}

type dataMemberSelector struct {
}

func (*dataMemberSelector) Select(member IMember) (selected bool) {
	return !member.IsLiteMember()
}

type selectors struct {
	DataMemberSelector MemberSelector
}
