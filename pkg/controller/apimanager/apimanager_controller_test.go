package apimanager

import (
	"context"
	"testing"

	appsv1alpha1 "github.com/3scale/3scale-operator/pkg/apis/apps/v1alpha1"
	"github.com/3scale/3scale-operator/pkg/reconcilers"
	appsv1 "github.com/openshift/api/apps/v1"
	imagev1 "github.com/openshift/api/image/v1"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestAPIManagerControllerCreate(t *testing.T) {
	var (
		name           = "example-apimanager"
		namespace      = "operator-unittest"
		wildcardDomain = "test.3scale.net"
	)

	apimanager := &appsv1alpha1.APIManager{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1alpha1.APIManagerSpec{
			APIManagerCommonSpec: appsv1alpha1.APIManagerCommonSpec{
				WildcardDomain: wildcardDomain,
			},
		},
	}

	ctx := context.TODO()

	// Objects to track in the fake client.
	objs := []runtime.Object{apimanager}

	// Register operator types with the runtime scheme.
	s := scheme.Scheme
	s.AddKnownTypes(appsv1alpha1.SchemeGroupVersion, apimanager)
	err := appsv1.AddToScheme(s)
	if err != nil {
		t.Fatalf("Unable to add Apps scheme: (%v)", err)
	}
	err = imagev1.AddToScheme(s)
	if err != nil {
		t.Fatalf("Unable to add Image scheme: (%v)", err)
	}
	err = routev1.AddToScheme(s)
	if err != nil {
		t.Fatalf("Unable to add Route scheme: (%v)", err)
	}

	// Create a fake client to mock API calls.
	cl := fake.NewFakeClient(objs...)
	clientAPIReader := fake.NewFakeClient(objs...)

	// Create a ReconcileMemcached object with the scheme and fake client.
	r := &ReconcileAPIManager{
		BaseReconciler: reconcilers.NewBaseReconciler(cl, s, clientAPIReader, ctx, log),
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: namespace,
		},
	}

	endLoop := false
	for i := 0; i < 100 && !endLoop; i++ {
		res, err := r.Reconcile(req)
		if err != nil {
			t.Fatal(err)
		}

		endLoop = !res.Requeue
	}

	finalAPIManager := &appsv1alpha1.APIManager{}
	err = r.Client().Get(context.TODO(), req.NamespacedName, finalAPIManager)
	if err != nil {
		t.Fatalf("get APIManager: (%v)", err)
	}

	backendListenerExistingReplicas := finalAPIManager.Spec.Backend.ListenerSpec.Replicas
	if backendListenerExistingReplicas == nil {
		t.Errorf("APIManager's backend listener replicas does not have a default value set")

	}

	if *backendListenerExistingReplicas != 1 {
		t.Errorf("APIManager's backend listener replicas size (%d) is not the expected size (%d)", backendListenerExistingReplicas, 1)
	}
}
