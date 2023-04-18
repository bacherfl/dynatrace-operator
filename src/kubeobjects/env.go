package kubeobjects

import (
	"os"

	corev1 "k8s.io/api/core/v1"
)

const (
	EnvPodNamespace             = "POD_NAMESPACE"
	EnvPodName                  = "POD_NAME"
	EnvCertManagerSelectorKey   = "CERT_MANAGER_SELECTOR_KEY"
	EnvCertManagerSelectorValue = "CERT_MANAGER_SELECTOR_VALUE"
)

func FindEnvVar(envVars []corev1.EnvVar, name string) *corev1.EnvVar {
	for i, envVar := range envVars {
		if envVar.Name == name {
			// returning reference to env var to ease later manipulation of it
			return &envVars[i]
		}
	}
	return nil
}

func EnvVarIsIn(envVars []corev1.EnvVar, envVarToCheck string) bool {
	for _, envVar := range envVars {
		if envVar.Name == envVarToCheck {
			return true
		}
	}
	return false
}

func AddOrUpdate(envVars []corev1.EnvVar, desiredEnvVar corev1.EnvVar) []corev1.EnvVar {
	targetEnvVar := FindEnvVar(envVars, desiredEnvVar.Name)
	if targetEnvVar != nil {
		*targetEnvVar = desiredEnvVar
	} else {
		envVars = append(envVars, desiredEnvVar)
	}
	return envVars
}

func NewEnvVarSourceForField(fieldPath string) *corev1.EnvVarSource {
	return &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: fieldPath}}
}

func DefaultNamespace() string {
	namespace := os.Getenv(EnvPodNamespace)

	if namespace == "" {
		return "dynatrace"
	}
	return namespace
}
