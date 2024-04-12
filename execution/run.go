package execution

import "github.com/containerd/cgroups"

func isV2() bool {
  return cgroups.Mode() == cgroups.Unified 
}
