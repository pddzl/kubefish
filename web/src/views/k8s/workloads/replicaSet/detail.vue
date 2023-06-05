<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewReplicaSet">查看</el-button>
      <el-button icon="expand" type="warning" plain @click="openScaleDialog">伸缩 </el-button>
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
                  <el-tag :type="StatusRsFilter(scope.row.status)" size="small">
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
          <div style="margin-right: 25px">
            <el-table :data="replicaSetPods" style="margin-bottom: 20px">
              <el-table-column label="名称" prop="metadata.name" min-width="120">
                <template #default="scope">
                  <router-link
                    :to="{
                      name: 'PodDetail',
                      query: { pod: scope.row.metadata.name, namespace: scope.row.metadata.namespace }
                    }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
                  </router-link>
                </template>
              </el-table-column>
              <el-table-column label="命名空间" prop="metadata.namespace" />
              <el-table-column label="节点" prop="nodeName" />
              <el-table-column label="状态">
                <template #default="scope">
                  <el-tag :type="StatusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
                </template>
              </el-table-column>
              <el-table-column label="创建时间">
                <template #default="scope">{{ formatDateTime(scope.row.metadata.creationTimestamp) }}</template>
              </el-table-column>
            </el-table>
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
        <el-collapse-item title="Services" name="services">
          <div style="margin-right: 25px">
            <el-table :data="replicaSetServices">
              <el-table-column label="名称" prop="metadata.name" min-width="120">
                <!-- <template #default="scope">
                  <router-link
                    :to="{ name: 'services_detail', query: { service: scope.row.metadata.name, namespace: namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
                  </router-link>
                </template> -->
              </el-table-column>
              <el-table-column label="命名空间" prop="metadata.namespace" />
              <el-table-column label="类型" prop="spec.type" />
              <el-table-column label="集群IP" prop="spec.clusterIP" />
              <el-table-column label="创建时间">
                <template #default="scope">{{ formatDateTime(scope.row.metadata.creationTimestamp) }}</template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 查看编排对话框 -->
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="replicaSetFormat" :readOnly="true" />
    </el-dialog>
    <!-- 伸缩对话框 -->
    <el-dialog v-model="dialogScaleVisible" title="伸缩" width="55%" center>
      <p style="font-weight: bold">
        ReplicaSet {{ replicaSet }} will be updated to reflect the desired replicas count.
      </p>
      <div style="margin: 25px 0 25px 0px">
        <span style="margin-right: 10px">Desired Replicas:</span>
        <el-input-number v-model="desiredNum" :min="0" :max="50" style="margin-right: 20px" />
        <span style="margin-right: 10px">Actual Replicas: </span>
        <el-input-number v-model="ActualNum" disabled />
      </div>
      <warning-bar :title="warningTitle" />
      <template #footer>
        <el-button @click="closeScaleDialog">取消</el-button>
        <el-button type="primary" @click="scaleFunc">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from "vue"
import { useRoute } from "vue-router"
import {
  type ReplicaSetDetail,
  getReplicaSetDetailApi,
  getReplicaSetPodsApi,
  getReplicaSetServicesApi,
  deleteReplicaSetApi
} from "@/api/k8s/replicaSet"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { StatusPodFilter, StatusRsFilter } from "@/utils/k8s/filter.js"
// import { scale } from "@/api/kubernetes/scale"
import { formatDateTime } from "@/utils/index"
import MetaData from "@/components/k8s/metadata.vue"
// import warningBar from "@/components/warningBar/warningBar.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { usePagination } from "@/hooks/usePagination"

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

// 加载replicaSet关联services
const replicaSetServices = ref([])

const getReplicaSetServicesFunc = async () => {
  const res = await getReplicaSetServicesApi({ namespace: namespace, replicaSet: replicaSet })
  if (res.code === 0) {
    replicaSetServices.value = res.data
  }
}
getReplicaSetServicesFunc()

// 查看编排
const dialogViewVisible = ref(false)
const replicaSetFormat = ref({})

const viewReplicaSet = async () => {
  const result = await getReplicaSetRaw({ replicaSet: replicaSet, namespace: namespace })
  if (result.code === 0) {
    replicaSetFormat.value = JSON.stringify(result.data)
  }
  dialogViewVisible.value = true
}

// 伸缩
const dialogScaleVisible = ref(false)
const warningTitle = ref("")
const desiredNum = ref(0)
const ActualNum = ref(0)

// -> 打开对话框
const openScaleDialog = () => {
  desiredNum.value = replicaSetDetail.value.status.replicas
  ActualNum.value = replicaSetDetail.value.status.availableReplicas
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} replicaset ${replicaSet} --replicas=${ActualNum.value}`
  dialogScaleVisible.value = true
}

watch(desiredNum, (val) => {
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} replicaset ${replicaSet} --replicas=${val}`
})

// -> 关闭对话框
const closeScaleDialog = () => {
  dialogScaleVisible.value = false
}

// -> 操作
const scaleFunc = async () => {
  const res = await scale({ namespace: namespace, name: replicaSet, kind: "replicaSet", num: desiredNum.value })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "伸缩成功",
      showClose: true
    })
    // 查询刷新数据
    getData()
    getReplicaSetPodsData()
    getReplicaSetServicesData()
  }
  dialogScaleVisible.value = false
}

// 删除
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
