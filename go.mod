module github.com/IBM/ibm-management-ingress-operator

go 1.13

require (
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/jetstack/cert-manager v0.10.0
	github.com/IBM/controller-filtered-cache v0.2.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200403190813-44a64ad78b9b // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0
	sigs.k8s.io/controller-runtime v0.6.2
)

replace k8s.io/client-go => k8s.io/client-go v0.18.8
