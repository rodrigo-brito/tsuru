// AUTOMATICALLY GENERATED FILE - DO NOT EDIT!
// Please run 'go generate' to update this file.
//
// Copyright 2015 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package permission

var (
	PermAll                         = PermissionRegistry.get("")
	PermApp                         = PermissionRegistry.get("app")
	PermAppAdmin                    = PermissionRegistry.get("app.admin")
	PermAppAdminQuota               = PermissionRegistry.get("app.admin.quota")
	PermAppAdminRoutes              = PermissionRegistry.get("app.admin.routes")
	PermAppAdminUnlock              = PermissionRegistry.get("app.admin.unlock")
	PermAppCreate                   = PermissionRegistry.get("app.create")
	PermAppDelete                   = PermissionRegistry.get("app.delete")
	PermAppDeploy                   = PermissionRegistry.get("app.deploy")
	PermAppDeployRollback           = PermissionRegistry.get("app.deploy.rollback")
	PermAppRead                     = PermissionRegistry.get("app.read")
	PermAppReadDeploy               = PermissionRegistry.get("app.read.deploy")
	PermAppReadEnv                  = PermissionRegistry.get("app.read.env")
	PermAppReadLog                  = PermissionRegistry.get("app.read.log")
	PermAppReadMetric               = PermissionRegistry.get("app.read.metric")
	PermAppRun                      = PermissionRegistry.get("app.run")
	PermAppUpdate                   = PermissionRegistry.get("app.update")
	PermAppUpdateBind               = PermissionRegistry.get("app.update.bind")
	PermAppUpdateCname              = PermissionRegistry.get("app.update.cname")
	PermAppUpdateCnameAdd           = PermissionRegistry.get("app.update.cname.add")
	PermAppUpdateCnameRemove        = PermissionRegistry.get("app.update.cname.remove")
	PermAppUpdateEnv                = PermissionRegistry.get("app.update.env")
	PermAppUpdateEnvSet             = PermissionRegistry.get("app.update.env.set")
	PermAppUpdateEnvUnset           = PermissionRegistry.get("app.update.env.unset")
	PermAppUpdateGrant              = PermissionRegistry.get("app.update.grant")
	PermAppUpdateLog                = PermissionRegistry.get("app.update.log")
	PermAppUpdatePlan               = PermissionRegistry.get("app.update.plan")
	PermAppUpdatePool               = PermissionRegistry.get("app.update.pool")
	PermAppUpdateRestart            = PermissionRegistry.get("app.update.restart")
	PermAppUpdateRevoke             = PermissionRegistry.get("app.update.revoke")
	PermAppUpdateStart              = PermissionRegistry.get("app.update.start")
	PermAppUpdateStop               = PermissionRegistry.get("app.update.stop")
	PermAppUpdateSwap               = PermissionRegistry.get("app.update.swap")
	PermAppUpdateTeamowner          = PermissionRegistry.get("app.update.teamowner")
	PermAppUpdateUnbind             = PermissionRegistry.get("app.update.unbind")
	PermAppUpdateUnit               = PermissionRegistry.get("app.update.unit")
	PermAppUpdateUnitAdd            = PermissionRegistry.get("app.update.unit.add")
	PermAppUpdateUnitRegister       = PermissionRegistry.get("app.update.unit.register")
	PermAppUpdateUnitRemove         = PermissionRegistry.get("app.update.unit.remove")
	PermAppUpdateUnitStatus         = PermissionRegistry.get("app.update.unit.status")
	PermDebug                       = PermissionRegistry.get("debug")
	PermHealing                     = PermissionRegistry.get("healing")
	PermMachine                     = PermissionRegistry.get("machine")
	PermMachineCreate               = PermissionRegistry.get("machine.create")
	PermMachineDelete               = PermissionRegistry.get("machine.delete")
	PermMachineRead                 = PermissionRegistry.get("machine.read")
	PermMachineTemplate             = PermissionRegistry.get("machine.template")
	PermMachineTemplateCreate       = PermissionRegistry.get("machine.template.create")
	PermMachineTemplateDelete       = PermissionRegistry.get("machine.template.delete")
	PermMachineTemplateRead         = PermissionRegistry.get("machine.template.read")
	PermMachineTemplateUpdate       = PermissionRegistry.get("machine.template.update")
	PermMachineUpdate               = PermissionRegistry.get("machine.update")
	PermNode                        = PermissionRegistry.get("node")
	PermNodeAutoscale               = PermissionRegistry.get("node.autoscale")
	PermNodeBs                      = PermissionRegistry.get("node.bs")
	PermNodeCreate                  = PermissionRegistry.get("node.create")
	PermNodeDelete                  = PermissionRegistry.get("node.delete")
	PermNodeRead                    = PermissionRegistry.get("node.read")
	PermNodeUpdate                  = PermissionRegistry.get("node.update")
	PermPlan                        = PermissionRegistry.get("plan")
	PermPlanCreate                  = PermissionRegistry.get("plan.create")
	PermPlanDelete                  = PermissionRegistry.get("plan.delete")
	PermPlatform                    = PermissionRegistry.get("platform")
	PermPlatformCreate              = PermissionRegistry.get("platform.create")
	PermPlatformDelete              = PermissionRegistry.get("platform.delete")
	PermPlatformUpdate              = PermissionRegistry.get("platform.update")
	PermPool                        = PermissionRegistry.get("pool")
	PermPoolCreate                  = PermissionRegistry.get("pool.create")
	PermPoolDelete                  = PermissionRegistry.get("pool.delete")
	PermPoolUpdate                  = PermissionRegistry.get("pool.update")
	PermRole                        = PermissionRegistry.get("role")
	PermRoleCreate                  = PermissionRegistry.get("role.create")
	PermRoleDelete                  = PermissionRegistry.get("role.delete")
	PermRoleUpdate                  = PermissionRegistry.get("role.update")
	PermRoleUpdateAssign            = PermissionRegistry.get("role.update.assign")
	PermRoleUpdateDissociate        = PermissionRegistry.get("role.update.dissociate")
	PermService                     = PermissionRegistry.get("service")
	PermServiceInstance             = PermissionRegistry.get("service-instance")
	PermServiceInstanceCreate       = PermissionRegistry.get("service-instance.create")
	PermServiceInstanceDelete       = PermissionRegistry.get("service-instance.delete")
	PermServiceInstanceRead         = PermissionRegistry.get("service-instance.read")
	PermServiceInstanceReadStatus   = PermissionRegistry.get("service-instance.read.status")
	PermServiceInstanceUpdate       = PermissionRegistry.get("service-instance.update")
	PermServiceInstanceUpdateBind   = PermissionRegistry.get("service-instance.update.bind")
	PermServiceInstanceUpdateGrant  = PermissionRegistry.get("service-instance.update.grant")
	PermServiceInstanceUpdateProxy  = PermissionRegistry.get("service-instance.update.proxy")
	PermServiceInstanceUpdateRevoke = PermissionRegistry.get("service-instance.update.revoke")
	PermServiceInstanceUpdateUnbind = PermissionRegistry.get("service-instance.update.unbind")
	PermServiceCreate               = PermissionRegistry.get("service.create")
	PermServiceDelete               = PermissionRegistry.get("service.delete")
	PermServiceRead                 = PermissionRegistry.get("service.read")
	PermServiceReadDoc              = PermissionRegistry.get("service.read.doc")
	PermServiceReadPlans            = PermissionRegistry.get("service.read.plans")
	PermServiceUpdate               = PermissionRegistry.get("service.update")
	PermServiceUpdateDoc            = PermissionRegistry.get("service.update.doc")
	PermServiceUpdateGrantAccess    = PermissionRegistry.get("service.update.grant-access")
	PermServiceUpdateProxy          = PermissionRegistry.get("service.update.proxy")
	PermServiceUpdateRevokeAccess   = PermissionRegistry.get("service.update.revoke-access")
	PermTeam                        = PermissionRegistry.get("team")
	PermTeamCreate                  = PermissionRegistry.get("team.create")
	PermTeamDelete                  = PermissionRegistry.get("team.delete")
	PermUser                        = PermissionRegistry.get("user")
	PermUserCreate                  = PermissionRegistry.get("user.create")
	PermUserDelete                  = PermissionRegistry.get("user.delete")
	PermUserUpdate                  = PermissionRegistry.get("user.update")
	PermUserUpdateQuota             = PermissionRegistry.get("user.update.quota")
	PermUserUpdateToken             = PermissionRegistry.get("user.update.token")
)
