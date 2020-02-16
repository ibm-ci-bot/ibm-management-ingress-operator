package handler

import (
	"fmt"

	rbac "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/IBM/management-ingress-operator/pkg/utils"
)

//NewPolicyRule stubs policy rule
func NewPolicyRule(apiGroups, resources, resourceNames, verbs []string) rbac.PolicyRule {
	return rbac.PolicyRule{
		APIGroups:     apiGroups,
		Resources:     resources,
		ResourceNames: resourceNames,
		Verbs:         verbs,
	}
}

//NewPolicyRules stubs policy rules
func NewPolicyRules(rules ...rbac.PolicyRule) []rbac.PolicyRule {
	return rules
}

//NewRole stubs a role
func NewRole(roleName, namespace string, rules []rbac.PolicyRule) *rbac.Role {
	return &rbac.Role{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: rbac.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleName,
			Namespace: namespace,
		},
		Rules: rules,
	}
}

//NewSubject stubs a new subect
func NewSubject(kind, name string) rbac.Subject {
	return rbac.Subject{
		Kind:     kind,
		Name:     name,
		APIGroup: rbac.GroupName,
	}
}

//NewSubjects stubs subjects
func NewSubjects(subjects ...rbac.Subject) []rbac.Subject {
	return subjects
}

//NewRoleBinding stubs a role binding
func NewRoleBinding(bindingName, namespace, roleName string, subjects []rbac.Subject) *rbac.RoleBinding {
	return &rbac.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: rbac.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      bindingName,
			Namespace: namespace,
		},
		RoleRef: rbac.RoleRef{
			Kind:     "Role",
			Name:     roleName,
			APIGroup: rbac.GroupName,
		},
		Subjects: subjects,
	}
}

//NewClusterRoleBinding stubs a cluster role binding
func NewClusterRoleBinding(bindingName, roleName string, subjects []rbac.Subject) *rbac.ClusterRoleBinding {
	return &rbac.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: rbac.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: bindingName,
		},
		RoleRef: rbac.RoleRef{
			Kind:     "ClusterRole",
			Name:     roleName,
			APIGroup: rbac.GroupName,
		},
		Subjects: subjects,
	}
}

// CreateClusterRole creates a cluser role or returns error
func (ingressRequest *IngressRequest) CreateClusterRole(name string, rules []rbac.PolicyRule) (*rbac.ClusterRole, error) {
	clusterRole := &rbac.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: rbac.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Rules: rules,
	}

	utils.AddOwnerRefToObject(clusterRole, utils.AsOwner(ingressRequest.managementIngress))

	err := ingressRequest.Create(clusterRole)
	if err != nil && !errors.IsAlreadyExists(err) {
		return nil, fmt.Errorf("Failure creating '%s' clusterrole: %v", name, err)
	}
	return clusterRole, nil
}

func (ingressRequest *IngressRequest) CreateClusterRoleBinding(binding *rbac.ClusterRoleBinding) error {
	utils.AddOwnerRefToObject(binding, utils.AsOwner(ingressRequest.managementIngress))

	err := ingressRequest.Create(binding)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("Failure creating ClusterRoleBinding: %v", err)
	}
	return nil
}

//RemoveClusterRoleBinding removes a cluster role binding
func (ingressRequest *IngressRequest) RemoveClusterRoleBinding(name string) error {

	binding := NewClusterRoleBinding(
		name,
		"",
		[]rbac.Subject{},
	)

	err := ingressRequest.Delete(binding)
	if err != nil && !errors.IsNotFound(err) {
		return fmt.Errorf("Failure deleting %q clusterrolebinding: %v", name, err)
	}

	return nil
}
