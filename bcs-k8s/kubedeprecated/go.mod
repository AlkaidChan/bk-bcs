module github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated

go 1.14

replace (
    github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated => ./
)

require (
	k8s.io/apimachinery v0.18.5
	k8s.io/code-generator v0.18.5
	sigs.k8s.io/controller-runtime v0.6.0
)
