/*
 * logicmonitor_sdk
 *
 * LogicMonitor is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. logicmonitor_sdk enables you to manage your LogicMonitor account programmatically.
 *
 * API version: 1.0.0
 * Contact: sdk@logicmonitor.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logicmonitor

type WebPageCollectorAttribute struct {
	Name string `json:"name"`
	Request string `json:"request"`
	Port string `json:"port"`
	FollowRedirect bool `json:"followRedirect"`
	Ip string `json:"ip"`
	ReadTimeout int32 `json:"readTimeout"`
	ConnectTimeout int32 `json:"connectTimeout"`
	UseSSL bool `json:"useSSL"`
}