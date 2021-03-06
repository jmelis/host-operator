package notification

import (
	"context"

	toolchainv1alpha1 "github.com/codeready-toolchain/api/pkg/apis/toolchain/v1alpha1"
	"github.com/codeready-toolchain/toolchain-common/pkg/test"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Assertion struct {
	notification   *toolchainv1alpha1.Notification
	client         client.Client
	namespacedName types.NamespacedName
	t              test.T
}

func (a *Assertion) loadNotificationAssertion() error {
	notification := &toolchainv1alpha1.Notification{}
	err := a.client.Get(context.TODO(), a.namespacedName, notification)
	a.notification = notification
	return err
}

func AssertThatNotification(t test.T, name string, client client.Client) *Assertion {
	return &Assertion{
		client:         client,
		namespacedName: test.NamespacedName(test.HostOperatorNs, name),
		t:              t,
	}
}

func (a *Assertion) HasConditions(expected ...toolchainv1alpha1.Condition) *Assertion {
	err := a.loadNotificationAssertion()
	require.NoError(a.t, err)
	test.AssertConditionsMatch(a.t, a.notification.Status.Conditions, expected...)
	return a
}
