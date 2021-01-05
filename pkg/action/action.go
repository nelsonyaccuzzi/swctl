/**
 * Copyright © 2014-2020 The SiteWhere Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package action

import (
	"github.com/sitewhere/swctl/pkg/kube"

	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
	//scriptingsitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/scripting.sitewhere.io/v1alpha4"
	//templatessitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/templates.sitewhere.io/v1alpha4"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(apiextv1beta1.AddToScheme(scheme))

	utilruntime.Must(sitewhereiov1alpha4.AddToScheme(scheme))
	//	utilruntime.Must(templatessitewhereiov1alpha4.AddToScheme(scheme))
	//	utilruntime.Must(scriptingsitewhereiov1alpha4.AddToScheme(scheme))
}

// Configuration injects the dependencies that all actions share.
type Configuration struct {
	// RESTClientGetter is an interface that loads Kubernetes clients.
	RESTClientGetter RESTClientGetter

	// KubeClient is a Kubernetes API client.
	KubeClient kube.Interface

	Log func(string, ...interface{})
}

// RESTClientGetter gets the rest client
type RESTClientGetter interface {
	ToRESTConfig() (*rest.Config, error)
	ToDiscoveryClient() (discovery.CachedDiscoveryInterface, error)
	ToRESTMapper() (meta.RESTMapper, error)
}

// DebugLog sets the logger that writes debug strings
type DebugLog func(format string, v ...interface{})

// Init initializes the action configuration
func (c *Configuration) Init(getter genericclioptions.RESTClientGetter, namespace string, log DebugLog) error {
	kc := kube.New(getter)
	kc.Log = log

	c.RESTClientGetter = getter
	c.KubeClient = kc
	c.Log = log

	return nil
}

// KubernetesClientSet creates a new kubernetes ClientSet based on the configuration
func (c *Configuration) KubernetesClientSet() (kubernetes.Interface, error) {
	conf, err := c.RESTClientGetter.ToRESTConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate config for kubernetes client")
	}

	return kubernetes.NewForConfig(conf)
}

// KubernetesAPIExtensionClientSet create a new kubernetes API Extension Clientset
func (c *Configuration) KubernetesAPIExtensionClientSet() (clientset.Interface, error) {
	conf, err := c.RESTClientGetter.ToRESTConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate config for API Extension Clientset")
	}
	return clientset.NewForConfig(conf)
}

// KubernetesDynamicClientSet create a new kubernetes API Extension Clientset
func (c *Configuration) KubernetesDynamicClientSet() (dynamic.Interface, error) {
	conf, err := c.RESTClientGetter.ToRESTConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate config for Dynamic Clientset")
	}
	return dynamic.NewForConfig(conf)
}

// ControllerClient creates a new controller client
func (c *Configuration) ControllerClient() (client.Client, error) {
	conf, err := c.RESTClientGetter.ToRESTConfig()
	if err != nil {
		return nil, errors.Wrap(err, "unable to generate config for kubernetes client")
	}
	return client.New(conf, client.Options{
		Scheme: scheme,
	})
}
