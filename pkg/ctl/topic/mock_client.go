// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package topic

import (
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

// MockClient is a mock implementation of the admin client for testing
type MockClient struct {
	MockTopics        *MockTopics
	MockClusters      admin.Clusters
	MockFunctions     admin.Functions
	MockTenants       admin.Tenants
	MockSubscriptions admin.Subscriptions
	MockSources       admin.Sources
	MockSinks         admin.Sinks
	MockNamespaces    admin.Namespaces
	MockSchemas       admin.Schema
	MockNsIsolation   admin.NsIsolationPolicy
	MockBrokers       admin.Brokers
	MockBrokerStats   admin.BrokerStats
	MockResourceQuotas admin.ResourceQuotas
	MockFunctionsWorker admin.FunctionsWorker
	MockPackages      admin.Packages
}

func (m *MockClient) Clusters() admin.Clusters {
	return m.MockClusters
}

func (m *MockClient) Functions() admin.Functions {
	return m.MockFunctions
}

func (m *MockClient) Tenants() admin.Tenants {
	return m.MockTenants
}

func (m *MockClient) Topics() admin.Topics {
	return m.MockTopics
}

func (m *MockClient) Subscriptions() admin.Subscriptions {
	return m.MockSubscriptions
}

func (m *MockClient) Sources() admin.Sources {
	return m.MockSources
}

func (m *MockClient) Sinks() admin.Sinks {
	return m.MockSinks
}

func (m *MockClient) Namespaces() admin.Namespaces {
	return m.MockNamespaces
}

func (m *MockClient) Schemas() admin.Schema {
	return m.MockSchemas
}

func (m *MockClient) NsIsolationPolicy() admin.NsIsolationPolicy {
	return m.MockNsIsolation
}

func (m *MockClient) Brokers() admin.Brokers {
	return m.MockBrokers
}

func (m *MockClient) BrokerStats() admin.BrokerStats {
	return m.MockBrokerStats
}

func (m *MockClient) ResourceQuotas() admin.ResourceQuotas {
	return m.MockResourceQuotas
}

func (m *MockClient) FunctionsWorker() admin.FunctionsWorker {
	return m.MockFunctionsWorker
}

func (m *MockClient) Packages() admin.Packages {
	return m.MockPackages
}

func (m *MockClient) Token() cmdutils.Token {
	return &token{}
}

// token implements the Token interface for testing
type token struct{}

func (t *token) Generate() (string, error) {
	return "mock-token", nil
}

func (t *token) Validate(token string) error {
	return nil
}

// MockTopics implements the admin.Topics interface for testing
type MockTopics struct {
	Topics map[string]*utils.PartitionedTopicMetadata
	Permissions map[string]map[string][]utils.AuthAction
	Stats map[string]*utils.TopicStats
	InternalStats map[string]*MockTopicInternalStats
}

// MockTopicInternalStats represents internal stats for testing
type MockTopicInternalStats struct {
	EntriesAddedCounter int64
	NumberOfEntries     int64
	TotalSize           int64
	CurrentLedgerEntries int64
	CurrentLedgerSize    int64
}

func NewMockTopics() *MockTopics {
	return &MockTopics{
		Topics: make(map[string]*utils.PartitionedTopicMetadata),
		Permissions: make(map[string]map[string][]utils.AuthAction),
		Stats: make(map[string]*utils.TopicStats),
		InternalStats: make(map[string]*MockTopicInternalStats),
	}
}

func (m *MockTopics) Create(topic utils.TopicName, partitions int) error {
	topicFqn := topic.String()
	m.Topics[topicFqn] = &utils.PartitionedTopicMetadata{
		Partitions: partitions,
	}
	return nil
}

func (m *MockTopics) GetMetadata(topic utils.TopicName) (*utils.PartitionedTopicMetadata, error) {
	topicFqn := topic.String()
	if meta, exists := m.Topics[topicFqn]; exists {
		return meta, nil
	}
	return &utils.PartitionedTopicMetadata{
		Partitions: 0,
	}, nil
}

func (m *MockTopics) Delete(topic utils.TopicName) error {
	topicFqn := topic.String()
	delete(m.Topics, topicFqn)
	return nil
}

func (m *MockTopics) GetPermissions(topic utils.TopicName) (map[string][]utils.AuthAction, error) {
	topicFqn := topic.String()
	if perms, exists := m.Permissions[topicFqn]; exists {
		return perms, nil
	}
	return make(map[string][]utils.AuthAction), nil
}

func (m *MockTopics) GrantPermission(topic utils.TopicName, role string, actions []utils.AuthAction) error {
	topicFqn := topic.String()
	if _, exists := m.Permissions[topicFqn]; !exists {
		m.Permissions[topicFqn] = make(map[string][]utils.AuthAction)
	}
	m.Permissions[topicFqn][role] = actions
	return nil
}

func (m *MockTopics) RevokePermission(topic utils.TopicName, role string) error {
	topicFqn := topic.String()
	if perms, exists := m.Permissions[topicFqn]; exists {
		delete(perms, role)
	}
	return nil
}

func (m *MockTopics) GetStats(topic utils.TopicName, getPreciseBacklog bool, subscriptionBacklogSize bool, getEarliestTimeInBacklog bool, excludePublishers bool, excludeConsumers bool) (*utils.TopicStats, error) {
	topicFqn := topic.String()
	if stats, exists := m.Stats[topicFqn]; exists {
		return stats, nil
	}
	return &utils.TopicStats{}, nil
}

func (m *MockTopics) GetInternalStats(topic utils.TopicName) (interface{}, error) {
	topicFqn := topic.String()
	if stats, exists := m.InternalStats[topicFqn]; exists {
		return stats, nil
	}
	return &MockTopicInternalStats{}, nil
}

func (m *MockTopics) Compact(topic utils.TopicName) error {
	return nil
}

func (m *MockTopics) GetPartitionedStats(topic utils.TopicName, perPartition bool, getPreciseBacklog bool, subscriptionBacklogSize bool, getEarliestTimeInBacklog bool, excludePublishers bool, excludeConsumers bool) (*utils.PartitionedTopicStats, error) {
	return &utils.PartitionedTopicStats{}, nil
}

func (m *MockTopics) Terminate(topic utils.TopicName) (utils.MessageID, error) {
	return utils.MessageID{}, nil
}

func (m *MockTopics) Unload(topic utils.TopicName) error {
	return nil
}

func (m *MockTopics) Lookup(topic utils.TopicName) (*utils.LookupData, error) {
	return &utils.LookupData{}, nil
}

// Implement other methods of the admin.Topics interface as needed for tests

// NewMockClient creates a new mock client for testing
func NewMockClient() cmdutils.Client {
	return &MockClient{
		MockTopics: NewMockTopics(),
	}
}