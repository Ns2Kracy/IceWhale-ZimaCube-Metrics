package utils

// func monitorService() {
// 	for _, service := range serviceList {
// 		pid := GetPid(service)
// 		if pid == "-1" {
// 			fmt.Println("Service", service, "is not running.")
// 			continue
// 		}

// 		processInfo := GetProcessInfo(pid)
// 		serviceStatus := ServiceStatus{
// 			ServiceName: service,
// 			ProcessInfo: processInfo,
// 			CreatedAt:   time.Now(),
// 		}

// 		DB.Create(&serviceStatus)
// 		// Get max CPU and MEM in the last 5 minutes
// 		// var maxCPU, maxMEM string
// 		// db.Table("service_statuses").
// 		// 	Select("MAX(process_info.cpu) AS max_cpu, MAX(process_info.mem) AS max_mem").
// 		// 	Where("service_name = ? AND created_at >= ?", service, time.Now().Add(-5*time.Minute)).
// 		// 	Scan(&serviceStatus)

// 		// maxCPU = serviceStatus.MaxCPU
// 		// maxMEM = serviceStatus.MaxMEM

// 		// 		rows += strings.Join([]string{
// 		// 			`{
// 		// 	"Service": "` + service + `",
// 		// 	"CPU": ` + processInfo.CPU + "%" + `,
// 		// 	"MEM": ` + processInfo.Mem + "%" + `
// 		// },`,
// 		// 		}, "")

// 		// 		rows += "\n"
// 	}
// 	// rows = strings.TrimRight(rows, "\n")
// 	// rows = strings.TrimRight(rows, ",")

// 	// message := AssembleMessage(rows)
// }
