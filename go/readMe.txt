go get k8s.io/client-go@v0.28.15
go get k8s.io/api@v0.28.15
go get k8s.io/apimachinery@v0.28.15
go get gopkg.in/yaml.v3

#
go clean -modcache
go mod tidy

go env -w GO111MODULE=on

go list -m -versions k8s.io/apimachinery
go list -m -versions sigs.k8s.io/json

go mod why sigs.k8s.io/json


/home/caiman/go/pkg/mod/k8s.io/apimachinery@v0.32.0/pkg/util/sets/set.go:20:2: package cmp is not in GOROOT (/home/caiman/sdk/go1.20.10/src/cmp)
note: imported by a module that requires go 1.23
/home/caiman/go/pkg/mod/k8s.io/apimachinery@v0.32.0/pkg/labels/selector.go:21:2: package slices is not in GOROOT (/home/caiman/sdk/go1.20.10/src/slices)

/home/caiman/go/pkg/mod/sigs.k8s.io/json@v0.0.0-20241010143419-9aa6b5e7a4b3/internal/golang/encoding/json/encode.go:15:2: package cmp is not in GOROOT (/home/caiman/sdk/go1.20.10/src/cmp)
note: imported by a module that requires go 1.21
/home/caiman/go/pkg/mod/sigs.k8s.io/json@v0.0.0-20241010143419-9aa6b5e7a4b3/internal/golang/encoding/json/encode.go:21:2: package slices is not in GOROOT (/home/caiman/sdk/go1.20.10/src/slices)
note: imported by a module that requires go 1.21


sidecar.istio.io/extraAnnotation.pod.prometheus.io/scrape="false" 这个注解 不会覆盖 Deployment 中已经设置的 prometheus.io/scrape 注解，因为它仅控制 Istio 侧车代理是否会注入注解，而不是修改 PodTemplate 或 Deployment 中已有的注解。
