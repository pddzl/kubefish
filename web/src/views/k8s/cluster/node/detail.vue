<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain>删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="nodeDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="nodeDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.spec && nodeDetail.status" title="资源信息" name="spec">
          <div class="info-box">
            <div v-if="nodeDetail.spec.podCIDR" class="row">
              <div class="item">
                <p>Pod CIDR</p>
                <div class="content">
                  <span class="shadow">{{ nodeDetail.spec.podCIDR }}</span>
                </div>
              </div>
            </div>
            <div v-if="nodeDetail.status.addresses" class="row">
              <div class="item">
                <p>地址</p>
                <div class="content">
                  <span v-for="item in nodeDetail.status.addresses" :key="item.type" class="shadow"
                    >{{ item.type }}: {{ item.address }}</span
                  >
                </div>
              </div>
            </div>
            <div v-if="nodeDetail.spec.taints" class="row">
              <div class="row">
                <div class="item">
                  <p>污点</p>
                  <div class="content">
                    <span v-for="item in nodeDetail.spec.taints" :key="item.type" class="shadow"
                      >{{ item.key }}={{ item.effect }}</span
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.status?.nodeInfo" title="系统信息" name="nodeInfo">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>操作系统</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.operatingSystem }}</span>
              </div>
              <div class="item">
                <p>操作系统镜像</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.osImage }}</span>
              </div>
              <div class="item">
                <p>架构</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.architecture }}</span>
              </div>
              <div class="item">
                <p>内核版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kernelVersion }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>kubelet运行端口</p>
                <span class="content">{{ nodeDetail.status.Port.kubeletEndpoint.Port }}</span>
              </div>
              <div class="item">
                <p>kubelet版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kubeletVersion }}</span>
              </div>
              <div class="item">
                <p>kube-proxy版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.kubeProxyVersion }}</span>
              </div>
              <div class="item">
                <p>容器runtime版本</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.containerRuntimeVersion }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>机器ID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.machineID }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>系统UUID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.systemUUID }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>启动ID</p>
                <span class="content">{{ nodeDetail.status.nodeInfo.bootID }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.allocatedResources" title="分配" name="allocated">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>CPU分配</p>
                <span class="content">{{ nodeDetail.allocatedResources.allocatedCPU }}m</span>
              </div>
              <div class="item">
                <p>CPU预留(requests)</p>
                <span class="content"
                  >{{ nodeDetail.allocatedResources.cpuRequests }}m ({{
                    nodeDetail.allocatedResources.cpuRequestsPercent * 100
                  }}%)</span
                >
              </div>
              <div class="item">
                <p>CPU限制(limits)</p>
                <span class="content"
                  >{{ nodeDetail.allocatedResources.cpuLimits }}m ({{
                    nodeDetail.allocatedResources.cpuLimitsPercent * 100
                  }}%)</span
                >
              </div>
              <div class="item">
                <p>内存分配</p>
                <span class="content"
                  >{{ Math.round(nodeDetail.allocatedResources.allocatedMemory / 1024 / 1024 / 1024) }}Mi</span
                >
              </div>
              <div class="item">
                <p>内存预留(requests)</p>
                <span class="content"
                  >{{ Math.round(nodeDetail.allocatedResources.memoryRequests / 1024 / 1024 / 1024) }}Mi ({{
                    nodeDetail.allocatedResources.memoryRequestsPercent * 100
                  }}%)</span
                >
              </div>
              <div class="item">
                <p>内存限制(limits)</p>
                <span class="content"
                  >{{ Math.round(nodeDetail.allocatedResources.memoryLimits / 1024 / 1024 / 1024) }}Mi ({{
                    (nodeDetail.allocatedResources.memoryLimitsPercent * 100).toFixed(1)
                  }}%)</span
                >
              </div>
              <div class="item">
                <p>Pods</p>
                <span class="content"
                  >{{ nodeDetail.allocatedResources.podNum }} | Capacity
                  {{ nodeDetail.allocatedResources.allocatedPods }}</span
                >
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="nodeDetail.status?.conditions" title="状态" name="status">
          <div class="info-table">
            <el-table :data="nodeDetail.status.conditions">
              <el-table-column prop="type" label="类别" min-width="120" />
              <el-table-column prop="status" label="状态" />
              <el-table-column label="最近检查时间" width="200">
                <template #default="scope">{{ formatDateTime(scope.row.lastHeartbeatTime) }}</template>
              </el-table-column>
              <el-table-column label="最近迁移时间" width="200">
                <template #default="scope">{{ formatDateTime(scope.row.lastTransitionTime) }}</template>
              </el-table-column>
              <el-table-column prop="reason" label="原因" min-width="180" />
              <el-table-column prop="message" label="信息" min-width="200" />
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <PodBriefC :total="total" :pods="relatedPods" @getRelatedPodsFunc="getRelatedPodsFunc" />
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { getNodeDetailApi, getNodePodsApi } from "@/api/k8s/cluster/node"
import { type PodBrief } from "@/api/k8s/workloads/pod"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import PodBriefC from "@/components/k8s/pod-brief.vue"
import { formatDateTime } from "@/utils/index"
import { viewOrch } from "@/utils/k8s/orch"
import { ref } from "vue"
import { useRoute } from "vue-router"

defineOptions({
  name: "NodeDetail"
})

const activeNames = ref(["metadata", "spec", "nodeInfo", "allocated", "status", "pods"])
const nodeDetail = ref<any>({})

const route = useRoute()
const nodeName = route.query.name as string

// 加载node详情
const getDetail = async () => {
  await getNodeDetailApi({ name: nodeName }).then((res) => {
    if (res.code === 0) {
      nodeDetail.value = res.data
    }
  })
}
getDetail()

// 加载node关联pods
const relatedPods = ref<PodBrief[]>([])
const total = ref(0)

interface pageObj {
  currentPage: number
  pageSize: number
}

const getRelatedPodsFunc = async (obj: pageObj) => {
  const res = await getNodePodsApi({
    page: obj.currentPage,
    pageSize: obj.pageSize,
    node_name: nodeName
  })
  if (res.code === 0) {
    relatedPods.value = res.data.list
    total.value = res.data.total
  }
}
getRelatedPodsFunc({ currentPage: 1, pageSize: 10 })

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(nodeName, "nodes")
  dialogFormVisible.value = true
}
</script>
