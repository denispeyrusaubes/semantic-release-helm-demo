package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	appsV1 "k8s.io/api/apps/v1"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/random"
)

func TestRenderDeployment(t *testing.T) {
	// Path to the helm chart we will test
	helmChartPath := "../"

	// Setup the args. For this test, we will set the following input values:
	// - image=nginx:1.15.8
	options := &helm.Options{
		SetValues: map[string]string{},
	}

	releaseName := fmt.Sprintf("example-%s", strings.ToLower(random.UniqueId()))

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/deployment.yaml"})

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var deployment appsV1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	var expectedLabels = map[string]string{
		"helm.sh/chart": "example-0.0.0",
		"app.kubernetes.io/name": "example",
		"app.kubernetes.io/instance": releaseName,
		"app.kubernetes.io/version": "0.0.0",
		"app.kubernetes.io/managed-by": "Helm",
	}
	assert.Equal(t, releaseName, deployment.ObjectMeta.Name)
	assert.Equal(t, expectedLabels, deployment.Labels)
}
