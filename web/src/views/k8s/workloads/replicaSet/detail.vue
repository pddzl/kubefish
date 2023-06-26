<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="replicaSetDetail.metadata" title="元数据" name="metadata">
          <MetaData :metadata="replicaSetDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="replicaSetDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ replicaSetDetail.spec.replicas }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>Selector:</p>
                <div class="content">
                  <span v-for="(label, index) in replicaSetDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="replicaSetDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ replicaSetDetail.status.replicas }}</span>
              </div>
              <div class="item">
                <p>fullyLabeledReplicas</p>
                <span class="content">{{ replicaSetDetail.status.fullyLabeledReplicas }}</span>
              </div>
              <div class="item">
                <p>readyReplicas</p>
                <span class="content">{{ replicaSetDetail.status.readyReplicas }}</span>
              </div>
              <div class="item">
                <p>availableReplicas</p>
                <span class="content">{{ replicaSetDetail.status.availableReplicas }}</span>
              </div>
            </div>
          </div>
          <div v-if="replicaSetDetail.status.conditions" style="margin-top: 20px; margin-right: 25px">
            <el-table :data="replicaSetDetail.status.conditions">
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
            <PodBriefC :total="total" :pods="relatedPods" @getRelatedPodsFunc="getRelatedPodsFunc" />
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 查看编排对话框 -->
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {
  deleteReplicaSetApi,
  getReplicaSetDetailApi,
  getReplicaSetPodsApi,
  type ReplicaSetDetail
} from "@/api/k8s/workloads/replicaSet"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import PodBriefC from "@/components/k8s/pod-brief.vue"
import { formatDateTime } from "@/utils/index"
import { statusRsFilter } from "@/utils/k8s/filter.js"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { ref } from "vue"
import { useRoute } from "vue-router"

defineOptions({
  name: "ReplicaSetDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "spec", "status", "pods", "services"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 加载replicaSet详情
const replicaSetDetail = ref<ReplicaSetDetail>({
  metadata: {},
  spec: {
    replicas: 0,
    selector: {}
  },
  status: {
    replicas: 0,
    fullyLabeledReplicas: 0,
    readyReplicas: 0,
    availableReplicas: 0,
    conditions: []
  }
})

const getDetail = async () => {
  await getReplicaSetDetailApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      replicaSetDetail.value = res.data
    }
  })
}
getDetail()

// 加载replicaSet关联pods

const relatedPods = ref([])
const total = ref(0)

interface pageObj {
  currentPage: number
  pageSize: number
}

const getRelatedPodsFunc = async (obj: pageObj) => {
  const res = await getReplicaSetPodsApi({
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

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "replicasets", namespace)
  dialogViewVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该ReplicaSet, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteReplicaSetApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
