package main

import (
	"fmt"
	"go-dl-benchmark/pkg/benchmark_v1/model"
	"go-dl-benchmark/pkg/devices"
	"go-dl-benchmark/pkg/protos"
	"go-dl-benchmark/pkg/terminal"
	"time"
)

func main() {
	ip, port := "127.0.0.1", 10001
	config := &protos.ModelBenchmarkTestArgs{
		ModelPath:        "D:\\code\\Go-DL-Benchmark\\res\\resnet18-12.onnx",
		Framework:        protos.FrameworkType_onnxruntime,
		InputTensorShape: "1,3,512,512",
		InputTensorType:  protos.TensorDataType_float32,
		WarmupRounds:     0,
		RunRounds:        0,
	}

	device := devices.HardwareDevice{
		Description: devices.HardwareDescription{
			DeviceName:     "ROG",
			OSType:         devices.Windows,
			Architecture:   devices.X86_64,
			ComputableChip: devices.Cpu,
			Ip:             ip,
			Port:           port,
		},
	}

	ability := model.BenchmarkAbility{
		IsSupportModelBenchmarkTest:           true,
		IsSupportHardwareBenchmarkTest:        false,
		SupportedFrameworksForRuntimeAnalysis: []protos.FrameworkType{protos.FrameworkType_onnxruntime},
		SupportedFrameworksForStaticAnalysis:  []protos.FrameworkType{protos.FrameworkType_onnxruntime},
	}

	go terminal.LaunchRemoteTerminalService(ip, port)
	time.Sleep(2 * 1e9)
	res, err := ability.ModelBenchmarkTest(config, device)
	if err == nil {
		fmt.Println(res)
	}

	var s string
	_, _ = fmt.Scanln(&s)

}
