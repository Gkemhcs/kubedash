package objects

import (
	"github.com/Gkemhcs/kubedash/internal/k8s"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
	"time"
)

func TestListingDeployments(t *testing.T) {
	testCases := []struct {
		testName        string
		deployments     []runtime.Object
		targetNamespace string
		expectedOutput  [][]string
		expectedSuccess bool
	}{
		{
			testName: "multiple_deployments",
			deployments: []runtime.Object{
				&v1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:              "deployment-1",
						Namespace:         "default",
						CreationTimestamp: metav1.Time{Time: time.Now().Add(-time.Hour)},
					},
					Status: v1.DeploymentStatus{
						AvailableReplicas: 3,
						UpdatedReplicas:   3,
						ReadyReplicas:     3,
						Replicas:          3,
					},
				},
				&v1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:              "deployment-2",
						Namespace:         "default",
						CreationTimestamp: metav1.Time{Time: time.Now().Add(-2 * time.Hour)},
					},
					Status: v1.DeploymentStatus{
						AvailableReplicas: 2,
						UpdatedReplicas:   2,
						ReadyReplicas:     2,
						Replicas:          2,
					},
				},
			},
			targetNamespace: "default",
			expectedOutput: [][]string{
				{"deployment-1", "3", "3", "3/3", "1h"},
				{"deployment-2", "2", "2", "2/2", "2h"},
			},
			expectedSuccess: true,
		},
		{
			testName: "single_deployment",
			deployments: []runtime.Object{
				&v1.Deployment{
					ObjectMeta: metav1.ObjectMeta{
						Name:              "deployment-1",
						Namespace:         "namespace-2",
						CreationTimestamp: metav1.Time{Time: time.Now().Add(-3 * time.Hour)},
					},
					Status: v1.DeploymentStatus{
						AvailableReplicas: 1,
						UpdatedReplicas:   1,
						ReadyReplicas:     1,
						Replicas:          1,
					},
				},
			},
			targetNamespace: "namespace-2",
			expectedOutput: [][]string{
				{"deployment-1", "1", "1", "1/1", "3h"},
			},
			expectedSuccess: true,
		},
		{
			testName:        "no_deployments_found",
			deployments:     []runtime.Object{},
			targetNamespace: "namespace-3",
			expectedOutput:  [][]string{},
			expectedSuccess: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.testName, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.deployments...)
			k8sConfig := k8s.K8sConfig{
				DefaultNamespace: "default",
				Client:           fakeClientSet,
			}

			// Using a context for API calls

			// Call ListDeployments function
			deploymentList, err := ListDeployments(test.targetNamespace, &k8sConfig)

			// Handle errors
			if err != nil && test.expectedSuccess {
				t.Fatalf("unexpected error listing deployments: %v", err)
			}
			if err == nil && !test.expectedSuccess {
				t.Fatalf("expected error but listing deployments")
			}

			// Check if the length of the returned deployment list matches expected
			if err == nil && len(deploymentList) != len(test.expectedOutput) {
				t.Fatalf("expected %d deployments but got %d", len(test.expectedOutput), len(deploymentList))
			}

			// Check if each deployment's details are correct
			for i, deployment := range deploymentList {
				for j, value := range deployment {
					if value != test.expectedOutput[i][j] {
						t.Errorf("expected %v for deployment %s, got %v", test.expectedOutput[i][j], deployment[0], value)
					}
				}
			}
		})
	}
}
