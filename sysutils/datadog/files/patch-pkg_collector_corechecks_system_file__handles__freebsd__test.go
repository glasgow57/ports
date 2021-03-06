--- pkg/collector/corechecks/system/file_handles_freebsd_test.go.orig	2020-03-02 22:03:50 UTC
+++ pkg/collector/corechecks/system/file_handles_freebsd_test.go
@@ -0,0 +1,38 @@
+// Unless explicitly stated otherwise all files in this repository are licensed
+// under the Apache License Version 2.0.
+// This product includes software developed at Datadog (https://www.datadoghq.com/).
+// Copyright 2016-2019 Datadog, Inc.
+// +build !freebsd
+
+package system
+
+import (
+	"testing"
+
+	"github.com/DataDog/datadog-agent/pkg/aggregator/mocksender"
+	"github.com/DataDog/datadog-agent/pkg/util/log"
+
+	"bou.ke/monkey"
+)
+
+func TestFhCheckFreeBSD(t *testing.T) {
+
+	fileHandleCheck := new(fhCheck)
+	fileHandleCheck.Configure(nil, nil, "test")
+
+	monkey.PatchInstanceMethod(sysctl.GetInt64, (name string, err error)) {
+		return (65534, nil)
+	})
+
+	mock := mocksender.NewMockSender(fileHandleCheck.ID())
+
+	mock.On("Gauge", "system.fs.file_handles.in_use", 421, "", []string(nil)).Return().Times(1)
+	mock.On("Gauge", "system.fs.file_handles.max", 65534, "", []string(nil)).Return().Times(1)
+	mock.On("Commit").Return().Times(1)
+	fileHandleCheck.Run()
+
+	mock.AssertExpectations(t)
+	mock.AssertNumberOfCalls(t, "Gauge", 2)
+	mock.AssertNumberOfCalls(t, "Commit", 1)
+
+}
