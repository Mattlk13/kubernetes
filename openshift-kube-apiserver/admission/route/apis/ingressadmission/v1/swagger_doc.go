package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_IngressAdmissionConfig = map[string]string{
	"":                     "IngressAdmissionConfig is the configuration for the the ingress controller limiter plugin. It changes the behavior of ingress objects to behave better with openshift routes and routers. *NOTE* This has security implications in the router when handling ingress objects",
	"allowHostnameChanges": "AllowHostnameChanges when false or unset openshift does not allow changing or adding hostnames to ingress objects. If set to true then hostnames can be added or modified which has security implications in the router.",
}

func (IngressAdmissionConfig) SwaggerDoc() map[string]string {
	return map_IngressAdmissionConfig
}