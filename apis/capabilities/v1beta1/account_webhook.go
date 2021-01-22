/*
Copyright 2020 Red Hat.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"errors"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var accountlog = logf.Log.WithName("account-resource")

func (r *Account) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:verbs=update,path=/validate-capabilities-3scale-net-v1beta1-account,mutating=false,failurePolicy=fail,groups=capabilities.3scale.net,resources=accounts,versions=v1beta1,name=vaccount.kb.io

var _ webhook.Validator = &Account{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Account) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Account) ValidateUpdate(old runtime.Object) error {
	accountlog.Info("validate update", "name", r.Name)

	oldAccount, ok := old.(*Account)
	if !ok {
		return fmt.Errorf("%T is not a *Account", old)
	}

	if oldAccount.Status.ID != nil {
		if reflect.DeepEqual(r.Status.ID, oldAccount.Status.ID) {
			return errors.New("Account ID is inmutable")
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Account) ValidateDelete() error {
	return nil
}
