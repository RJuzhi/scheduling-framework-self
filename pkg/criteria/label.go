package criteria

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
	"log"
)

const Name = "sample"

type sample struct{}

var _ framework.FilterPlugin = &sample{}
var _ framework.PreScorePlugin = &sample{}

func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &sample{}, nil
}

func (pl *sample) Name() string {
	return Name
}

func (pl *sample) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *framework.NodeInfo) *framework.Status {
	log.Printf("filter pod: %v, node: %v", pod.Name, node)
	// not contain gpu=true, unschedulable
	if node.Node().Labels["gpu"] != "true" {
		return framework.NewStatus(framework.Unschedulable, "Node: "+node.Node().Name)
	}
	return framework.NewStatus(framework.Success, "Node: "+node.Node().Name)
}

func (pl *sample) PreScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodes []*v1.Node) *framework.Status {
	log.Println(nodes)
	log.Println(pod)
	return framework.NewStatus(framework.Success, "Pod: "+pod.Name)
}
