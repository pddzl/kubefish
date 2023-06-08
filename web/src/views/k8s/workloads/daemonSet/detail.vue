<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="daemonSetDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="daemonSetDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="daemonSetDetail.spec" title="资源信息" name="resource">
          <div class="info-box">
            <div class="row">
              <div v-if="daemonSetDetail.spec.selector" class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in daemonSetDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
            <div v-if="daemonSetDetail.spec.updateStrategy" class="row">
              <div class="item">
                <p>Strategy</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.type }}</span>
              </div>
              <div class="item">
                <p>maxUnavailable</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
              </div>
              <div v-if="daemonSetDetail.spec.updateStrategy.type === 'RollingUpdate'" class="item">
                <p>maxSurge</p>
                <span class="content">{{ daemonSetDetail.spec.updateStrategy.rollingUpdate.maxUnavailable }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="daemonSetDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>currentNumberScheduled</p>
                <span class="content">{{ daemonSetDetail.status.currentNumberScheduled }}</span>
              </div>
              <div class="item">
                <p>desiredNumberScheduled</p>
                <span class="content">{{ daemonSetDetail.status.desiredNumberScheduled }}</span>
              </div>
              <div class="item">
                <p>numberReady</p>
                <span class="content">{{ daemonSetDetail.status.numberReady }}</span>
              </div>
              <div class="item">
                <p>numberAvailable</p>
                <span class="content">{{ daemonSetDetail.status.numberAvailable }}</span>
              </div>
            </div>
          </div>
          <div v-if="daemonSetDetail.status.conditions" style="margin-top: 20px; margin-right: 20px">
            <el-table :data="daemonSetDetail.status.conditions">
              <el-table-column label="类别" prop="type" />
              <el-table-column label="状态" prop="status">
                <template #default="scope">
                  <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                    {{ scope.row.status }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="上次迁移时间">
                <template #default="scope">
                  {{ formatDateTime(scope.row.lastTransitionTime) }}
                </template>
              </el-table-column>
              <el-table-column label="原因" prop="reason" />
              <el-table-column label="信息" prop="message" :show-overflow-tooltip="true" />
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <PodBrief :total="total" :pods="relatedPods" @getRelatedPodsFunc="getRelatedPodsFunc" />
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { deleteDaemonSetApi, getDaemonSetDetailApi, getDaemonSetPodsApi } from "@/api/k8s/daemonSet"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import PodBrief from "@/components/k8s/pod-brief.vue"
import { formatDateTime } from "@/utils/index"
import { statusRsFilter } from "@/utils/k8s/filter.js"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { ref } from "vue"
import { useRoute } from "vue-router"

// 折叠面板
const activeNames = ref(["metadata", "resource", "status", "pods"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 获取daemonSet详情
const daemonSetDetail = ref({})

const getDaemonSetDetailData = async () => {
  await getDaemonSetDetailApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      daemonSetDetail.value = res.data
    }
  })
}
getDaemonSetDetailData()

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "daemonsets", namespace)
  dialogViewVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该DaemonSet, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteDaemonSetApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}

// 加载daemonSet关联pods
const relatedPods = ref([])
const total = ref(0)

interface pageObj {
  currentPage: number
  pageSize: number
}

const getRelatedPodsFunc = async (obj: pageObj) => {
  const res = await getDaemonSetPodsApi({
    page: obj.currentPage,
    pageSize: obj.pageSize,
    namespace: namespace,
    name: name
  })
  if (res.code === 0) {
    relatedPods.value = res.data.list
    total.value = res.data.total
  }
}
getRelatedPodsFunc({ currentPage: 1, pageSize: 10 })
</script>
