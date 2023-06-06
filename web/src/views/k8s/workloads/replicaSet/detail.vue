<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc(replicaSet, namespace)">查看</el-button>
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
            <PodBrief :pods="replicaSetPods" style="margin-bottom: 20px" />
            <div class="pager-wrapper">
              <el-pagination
                background
                :layout="paginationData.layout"
                :page-sizes="paginationData.pageSizes"
                :total="paginationData.total"
                :page-size="paginationData.pageSize"
                :currentPage="paginationData.currentPage"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
              />
            </div>
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
import { ref } from "vue"
import { useRoute } from "vue-router"
import {
  type ReplicaSetDetail,
  getReplicaSetDetailApi,
  getReplicaSetPodsApi,
  deleteReplicaSetApi
} from "@/api/k8s/replicaSet"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { statusRsFilter } from "@/utils/k8s/filter.js"
import { formatDateTime } from "@/utils/index"
import MetaData from "@/components/k8s/metadata.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { usePagination } from "@/hooks/usePagination"
import { viewOrch } from "@/utils/k8s/orch"
import PodBrief from "@/components/k8s/pod-brief.vue"

// 分页
const { paginationData, changeCurrentPage, changePageSize } = usePagination()

// 折叠面板
const activeNames = ref(["metadata", "spec", "status", "pods", "services"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const replicaSet = route.query.replicaSet as string

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

const getTableData = async () => {
  await getReplicaSetDetailApi({ namespace: namespace, replicaSet: replicaSet }).then((res) => {
    if (res.code === 0) {
      replicaSetDetail.value = res.data
    }
  })
}
getTableData()

// 加载replicaSet关联pods
const replicaSetPods = ref([])

const getReplicaSetPodFunc = async () => {
  const res = await getReplicaSetPodsApi({
    page: paginationData.currentPage,
    pageSize: paginationData.pageSize,
    namespace: namespace,
    replicaSet: replicaSet
  })
  if (res.code === 0) {
    replicaSetPods.value = res.data.list
    paginationData.total = res.data.total
  }
}
getReplicaSetPodFunc()

// 分页
const handleSizeChange = (value: number) => {
  changePageSize(value)
  getTableData()
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  getTableData()
}

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async (name: string, namespace: string) => {
  formatData.value = await viewOrch(name, "replicasets", namespace)
  dialogViewVisible.value = true
}

const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该ReplicaSet, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteReplicaSetApi({ namespace: namespace, replicaSet: replicaSet })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
