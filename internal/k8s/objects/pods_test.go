package objects

import (
	"testing"
	"time"

	"github.com/Gkemhcs/kubedash/internal/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func TestListingPods(t *testing.T) {
	testCases := []struct {
		testName        string
		pods            []runtime.Object
		targetNamespace string
		expectedOutput  [][]string
		expectedSuccess bool
	}{
		{
			testName: "multiple_pods",
			pods: []runtime.Object{
				&corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
						ContainerStatuses: []corev1.ContainerStatus{
							{RestartCount: 1, Ready: true},
							{RestartCount: 0, Ready: true},
						},
						PodIP:     "192.168.1.1",
						StartTime: &metav1.Time{Time: time.Now().Add(-time.Hour)},
					},
					Spec: corev1.PodSpec{
						NodeName: "node-1",
					},
				},
				&corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-2",
						Namespace: "default",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
						ContainerStatuses: []corev1.ContainerStatus{
							{RestartCount: 0, Ready: true},
						},
						PodIP:     "192.168.1.2",
						StartTime: &metav1.Time{Time: time.Now().Add(-2 * time.Hour)},
					},
					Spec: corev1.PodSpec{
						NodeName: "node-2",
					},
				},
			},
			targetNamespace: "default",
			expectedOutput: [][]string{
				{"pod-1", "2/2", "Running", "1", "192.168.1.1", "1h", "node-1"},
				{"pod-2", "1/1", "Running", "0", "192.168.1.2", "2h", "node-2"},
			},
			expectedSuccess: true,
		},
		{
			testName: "single_pod",
			pods: []runtime.Object{
				&corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "pod-1",
						Namespace: "namespace-2",
					},
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
						ContainerStatuses: []corev1.ContainerStatus{
							{RestartCount: 1, Ready: true},
						},
						PodIP:     "192.168.1.3",
						StartTime: &metav1.Time{Time: time.Now().Add(-3 * time.Hour)},
					},
					Spec: corev1.PodSpec{
						NodeName: "node-3",
					},
				},
			},
			targetNamespace: "namespace-2",
			expectedOutput: [][]string{
				{"pod-1", "1/1", "Running", "1", "192.168.1.3", "3h", "node-3"},
			},
			expectedSuccess: true,
		},
		{
			testName:        "no_pods_found",
			pods:            []runtime.Object{},
			targetNamespace: "namespace-3",
			expectedOutput:  [][]string{},
			expectedSuccess: true,
		},
	}

	for _, test := range testCases {

		t.Run(test.testName, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.pods...)
			k8sConfig := k8s.K8sConfig{
				DefaultNamespace: "default",
				Client:           fakeClientSet,
			}

			// Using a context for API calls

			// Call ListPods function
			podList, err := ListPods(test.targetNamespace, &k8sConfig)

			// Handle errors
			if err != nil && test.expectedSuccess {
				t.Fatalf("unexpected error listing pods: %v", err)
			}
			if err == nil && !test.expectedSuccess {
				t.Fatalf("expected error but listing pods")
			}

			// Check if the length of the returned pod list matches expected
			if err == nil && len(podList) != len(test.expectedOutput) {
				t.Fatalf("expected %d pods but got %d", len(test.expectedOutput), len(podList))
			}

			// Check if each pod's details are correct
			for i, pod := range podList {
				for j, value := range pod {
					if value != test.expectedOutput[i][j] {
						t.Errorf("expected %v for pod %s, got %v", test.expectedOutput[i][j], pod[0], value)
					}
				}
			}
		})
	}
}
