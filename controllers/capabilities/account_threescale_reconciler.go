package controllers

import (
	"encoding/json"
	"reflect"

	capabilitiesv1beta1 "github.com/3scale/3scale-operator/apis/capabilities/v1beta1"
	"github.com/3scale/3scale-operator/pkg/reconcilers"

	threescaleapi "github.com/3scale/3scale-porta-go-client/client"
	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
)

type AccountThreescaleReconciler struct {
	*reconcilers.BaseReconciler
	resource            *capabilitiesv1beta1.Account
	threescaleAPIClient *threescaleapi.ThreeScaleClient
	providerAccountHost string
	logger              logr.Logger
}

func NewAccountThreescaleReconciler(b *reconcilers.BaseReconciler, resource *capabilitiesv1beta1.Account, threescaleAPIClient *threescaleapi.ThreeScaleClient, providerAccountHost string, logger logr.Logger) *AccountThreescaleReconciler {
	return &AccountThreescaleReconciler{
		BaseReconciler:      b,
		resource:            resource,
		threescaleAPIClient: threescaleAPIClient,
		providerAccountHost: providerAccountHost,
		logger:              logger.WithValues("3scale Reconciler", providerAccountHost),
	}
}

func (s *AccountThreescaleReconciler) Reconcile() (*threescaleapi.DeveloperAccount, error) {
	s.logger.V(1).Info("START")
	return nil, nil
}
