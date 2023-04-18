package klt_certs

import (
	"context"
	"github.com/Dynatrace/dynatrace-operator/src/eventfilter"
	"github.com/Dynatrace/dynatrace-operator/src/webhook"
	kltcert "github.com/keptn/lifecycle-toolkit/klt-cert-manager/controllers/keptnwebhookcontroller"
	appsv1 "k8s.io/api/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func Add(mgr manager.Manager, ns string, matchLabels map[string]string) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		WithEventFilter(eventfilter.ForObjectNameAndNamespace(webhook.DeploymentName, ns)).
		Complete(newWebhookCertificateController(mgr, ns, nil, matchLabels))
}

func AddBootstrap(mgr manager.Manager, ns string, cancelMgr context.CancelFunc, matchLabels map[string]string) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1.Deployment{}).
		WithEventFilter(eventfilter.ForObjectNameAndNamespace(webhook.DeploymentName, ns)).
		Complete(newWebhookCertificateController(mgr, ns, cancelMgr, matchLabels))
}

func newWebhookCertificateController(mgr manager.Manager, namespace string, cancelMgr context.CancelFunc, matchLabels map[string]string) *kltcert.KeptnWebhookCertificateReconciler {
	return &kltcert.KeptnWebhookCertificateReconciler{
		CancelMgrFunc: cancelMgr,
		Client:        mgr.GetClient(),
		Log:           ctrl.Log.WithName("KeptnWebhookCert Controller"),
		Namespace:     namespace,
		MatchLabels:   matchLabels,
	}
}
