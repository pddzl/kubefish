<template>
  <div class="container">
    <div v-for="(detail, index) in props.data" :key="index">
      <div class="box">
        <el-tag size="small">{{ detail.name }}</el-tag>
      </div>
      <div class="box">
        <p>镜像</p>
        <span class="content">{{ detail.image }}</span>
      </div>
      <div v-if="detail.status" class="box">
        <p>状态</p>
        <div class="block">
          <div class="column">
            <p>Ready</p>
            <span>{{ detail.status.ready }}</span>
          </div>
          <div class="column">
            <p>Started</p>
            <span>{{ detail.status.started }}</span>
          </div>
          <div class="column">
            <p>重启次数</p>
            <span>{{ detail.status.restartCount }}</span>
          </div>
          <div v-if="detail.status.state.running" class="column">
            <p>创建时间</p>
            <span>
              {{ formatDateTime(detail.status.state.running.startedAt) }}
            </span>
          </div>
          <div v-else-if="detail.status.state.terminated" class="column">
            <p>原因</p>
            <span>{{ detail.status.state.terminated.reason }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.env && detail.env.length > 0" class="box">
        <p>环境变量</p>
        <div class="block">
          <div v-for="(env, indexEnv) in detail.env" :key="indexEnv" class="column">
            <p>{{ env.name }}</p>
            <span>{{ env.value }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.command && detail.command.length > 0" class="box">
        <p>命令</p>
        <div class="div-shadow">
          <div v-for="(command, indexC) in detail.command" :key="indexC">{{ command }}</div>
        </div>
      </div>
      <div v-if="detail.args && detail.args.length > 0" class="box">
        <p>参数</p>
        <div class="div-shadow">
          <div v-for="(arg, indexA) in detail.args" :key="indexA">{{ arg }}</div>
        </div>
      </div>
      <div v-if="detail.volumeMounts && detail.volumeMounts.length > 0" class="box">
        <p>挂载点</p>
        <div class="info-table">
          <el-table :data="detail.volumeMounts">
            <el-table-column prop="name" label="名称" />
            <el-table-column prop="readOnly" label="只读">
              <template #default="scope">
                <span v-if="scope.row.readOnly">{{ scope.row.readOnly }}</span>
                <span v-else>false</span>
              </template>
            </el-table-column>
            <el-table-column prop="mountPath" label="挂载路径" min-width="200" />
            <el-table-column prop="volumeType" label="源类型" />
          </el-table>
        </div>
      </div>
      <div v-if="detail.securityContext" class="box">
        <p>安全性上下文</p>
        <div class="block">
          <div
            v-if="detail.securityContext.capabilities?.add && detail.securityContext.capabilities.add.length > 0"
            class="column"
          >
            <p>Added Capabilities</p>
            <span v-for="(add, indexA) in detail.securityContext.capabilities.add" :key="indexA">
              <span>{{ add }}</span>
              <span v-if="detail.securityContext.capabilities.add.length - 1 != indexA">,</span>
            </span>
          </div>
          <div
            v-if="detail.securityContext.capabilities?.drop && detail.securityContext.capabilities.drop.length > 0"
            class="column"
          >
            <p>Dropped Capabilities</p>
            <span v-for="(drop, indexD) in detail.securityContext.capabilities.drop" :key="indexD">
              <span class="content">{{ drop }}</span>
              <span v-if="detail.securityContext.capabilities.drop.length - 1 != indexD">,</span>
            </span>
          </div>
          <div v-if="detail.securityContext.privileged !== undefined" class="column">
            <p>Privileged</p>
            <span>{{ detail.securityContext.privileged }}</span>
          </div>
          <div v-if="detail.securityContext.seLinuxOptions?.user" class="column">
            <p>SeLinuxOptions User</p>
            <span>{{ detail.securityContext.seLinuxOptions.user }}</span>
          </div>
          <div v-if="detail.securityContext.seLinuxOptions?.role" class="column">
            <p>SeLinuxOptions Role</p>
            <span>{{ detail.securityContext.seLinuxOptions.role }}</span>
          </div>
          <div v-if="detail.securityContext.seLinuxOptions?.type" class="column">
            <p>SeLinuxOptions Type</p>
            <span>{{ detail.securityContext.seLinuxOptions.type }}</span>
          </div>
          <div v-if="detail.securityContext.seLinuxOptions?.level" class="column">
            <p>SeLinuxOptions Level</p>
            <span>{{ detail.securityContext.seLinuxOptions.level }}</span>
          </div>
          <div v-if="detail.securityContext.windowsOptions?.gmsaCredentialSpecName" class="column">
            <p>WindowsOptions GMSACredentialSpecName</p>
            <span>{{ detail.securityContext.windowsOptions.gmsaCredentialSpecName }}</span>
          </div>
          <div v-if="detail.securityContext.windowsOptions?.gmsaCredentialSpec" class="column">
            <p>WindowsOptions GMSACredentialSpec</p>
            <span>{{ detail.securityContext.windowsOptions.gmsaCredentialSpec }}</span>
          </div>
          <div v-if="detail.securityContext.windowsOptions?.runAsUserName" class="column">
            <p>WindowsOptions RunAsUserName</p>
            <span>{{ detail.securityContext.windowsOptions.runAsUserName }}</span>
          </div>
          <div v-if="detail.securityContext.windowsOptions?.hostProcess !== undefined" class="column">
            <p>WindowsOptions HostProcess</p>
            <span>{{ detail.securityContext.windowsOptions.hostProcess }}</span>
          </div>
          <div v-if="detail.securityContext.runAsUser" class="column">
            <p>RunAsUser</p>
            {{ detail.securityContext.runAsUser }}
          </div>
          <div v-if="detail.securityContext.runAsGroup" class="column">
            <p>RunAsGroup</p>
            <span class="content">{{ detail.securityContext.runAsGroup }}</span>
          </div>
          <div v-if="detail.securityContext.runAsNonRoot !== undefined" class="column">
            <p>RunAsNonRoot</p>
            <span class="content">{{ detail.securityContext.runAsNonRoot }}</span>
          </div>
          <div v-if="detail.securityContext.readOnlyRootFilesystem !== undefined" class="column">
            <p>ReadOnlyRootFilesystem</p>
            <span class="conetnt">{{ detail.securityContext.readOnlyRootFilesystem }}</span>
          </div>
          <div v-if="detail.securityContext.allowPrivilegeEscalation !== undefined" class="column">
            <p>AllowPrivilegeEscalation</p>
            <span class="content">{{ detail.securityContext.allowPrivilegeEscalation }}</span>
          </div>
          <div v-if="detail.securityContext.procMount" class="column">
            <p>ProcMount</p>
            <span class="content">{{ detail.securityContext.procMount }}</span>
          </div>
          <div v-if="detail.securityContext.seccompProfile?.type" class="column">
            <p>SeccompProfile Type</p>
            <span class="content">{{ detail.securityContext.seccompProfile.type }}</span>
          </div>
          <div v-if="detail.securityContext.seccompProfile?.localhostProfile" class="column">
            <p>SeccompProfile LocalhostProfile</p>
            <span class="content">{{ detail.securityContext.seccompProfile.localhostProfile }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.livenessProbe" class="box">
        <p>Liveness Probe</p>
        <div class="block">
          <div v-if="detail.livenessProbe.initialDelaySeconds" class="column">
            <p>Initial Delay (Seconds)</p>
            <span>{{ detail.livenessProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.timeoutSeconds" class="column">
            <p>Timeout (Seconds)</p>
            <span>{{ detail.livenessProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.periodSeconds" class="column">
            <p>Probe Period (Seconds)</p>
            <span>{{ detail.livenessProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.livenessProbe.successThreshold" class="column">
            <p>Success Threshold</p>
            <span>{{ detail.livenessProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.livenessProbe.failureThreshold" class="column">
            <p>Failure Threshold</p>
            <span>{{ detail.livenessProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.livenessProbe.httpGet" class="column">
            <p>HTTP Healthcheck URI</p>
            <span class="content-shadow">{{
              detail.livenessProbe.httpGet.scheme +
              "://[host]:" +
              detail.livenessProbe.httpGet.port +
              detail.livenessProbe.httpGet.path
            }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.readinessProbe" class="box">
        <p>ReadinessProbe</p>
        <div class="block">
          <div v-if="detail.readinessProbe.initialDelaySeconds" class="column">
            <p>Initial Delay (Seconds)</p>
            <span>{{ detail.readinessProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.timeoutSeconds" class="column">
            <p>Timeout (Seconds)</p>
            <span class="content">{{ detail.readinessProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.periodSeconds" class="column">
            <p>Probe Period (Seconds)</p>
            <span class="content">{{ detail.readinessProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.readinessProbe.successThreshold" class="column">
            <p>Success Threshold</p>
            <span class="content">{{ detail.readinessProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.readinessProbe.failureThreshold" class="column">
            <p>Failure Threshold</p>
            <span class="content">{{ detail.readinessProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.readinessProbe.httpGet" class="column">
            <p>HTTP Healthcheck URI</p>
            <span class="content-shadow">{{
              detail.readinessProbe.httpGet.scheme +
              "://[host]:" +
              detail.livenessProbe.httpGet.port +
              detail.livenessProbe.httpGet.path
            }}</span>
          </div>
        </div>
      </div>
      <div v-if="detail.startupProbe" class="box">
        <p>StartupProbe</p>
        <div class="block">
          <div v-if="detail.startupProbe.initialDelaySeconds" class="column">
            <p>Initial Delay (Seconds)</p>
            <span>{{ detail.startupProbe.initialDelaySeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.timeoutSeconds" class="column">
            <p>Timeout (Seconds)</p>
            <span>{{ detail.startupProbe.timeoutSeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.periodSeconds" class="column">
            <p>Probe Period (Seconds)</p>
            <span>{{ detail.startupProbe.periodSeconds }}</span>
          </div>
          <div v-if="detail.startupProbe.successThreshold" class="column">
            <p>Success Threshold</p>
            <span>{{ detail.startupProbe.successThreshold }}</span>
          </div>
          <div v-if="detail.startupProbe.failureThreshold" class="column">
            <p>Failure Threshold</p>
            <span>{{ detail.startupProbe.failureThreshold }}</span>
          </div>
          <div v-if="detail.startupProbe.httpGet" class="column">
            <p>HTTP Healthcheck URI</p>
            <span class="content-shadow">{{
              detail.startupProbe.httpGet.scheme +
              "://[host]:" +
              detail.livenessProbe.httpGet.port +
              detail.livenessProbe.httpGet.path
            }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { formatDateTime } from "@/utils/index"

const props = defineProps({
  data: {
    type: Object,
    default: () => {}
  }
})
</script>

<style lang="scss" scoped>
.container {
  > div:not(:last-child) {
    margin-bottom: 40px;
  }
  .box:not(:last-child) {
    margin-bottom: 15px;
  }
  .box {
    > p {
      font-size: 13px;
    }
    .div-shadow {
      background-color: rgba(128, 128, 128, 0.205);
      padding-left: 15px;
      padding-top: 10px;
      padding-bottom: 10px;
      margin-right: 20px;
    }
    .block {
      display: flex;
      flex-wrap: wrap;
      .column {
        p {
          font-size: 11px;
        }
        span {
          font-size: 16px;
        }
        .content-shadow {
          background-color: rgba(128, 128, 128, 0.253);
          padding: 5px 10px 5px 10px;
          border-radius: 5px;
          box-sizing: border-box;
        }
      }
      .column:not(:last-child) {
        margin-right: 40px;
      }
    }
  }
}
</style>
