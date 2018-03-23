// Copyright 2018 The Kubeflow Authors
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
package fake

import (
	v1alpha1 "github.com/kubeflow/tf-operator/pkg/client/clientset/versioned/typed/kubeflow/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

// FakeKubeflowV1alpha1 represents a fake structure for testing
type FakeKubeflowV1alpha1 struct {
	*testing.Fake
}

// TFJobs returns a fake tf-job instance for testing
func (c *FakeKubeflowV1alpha1) TFJobs(namespace string) v1alpha1.TFJobInterface {
	return &FakeTFJobs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubeflowV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
