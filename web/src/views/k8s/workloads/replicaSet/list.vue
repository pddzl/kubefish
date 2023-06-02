<template>
  <div class="app-container">
    <el-card shadow="never" class="search-wrapper">
      <el-form ref="searchForm" :inline="true">
        <el-form-item label="命名空间">
          <el-select v-model="searchInfo.namespace" clearable placeholder="请选择">
            <el-option v-for="item in namespaceList" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column label="名称" min-width="220">
            <!-- <template #default="scope">
            <router-link
              :to="{ name: 'replicaSet_detail', query: { replicaSet: scope.row.name, namespace: scope.row.namespace } }"
            >
              <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
            </router-link>
          </template> -->
          </el-table-column>
          <el-table-column label="命名空间" prop="namespace" min-width="120" />
          <el-table-column label="Pods" prop="pods" min-width="80">
            <template #default="scope">{{ scope.row.availableReplicas }} / {{ scope.row.replicas }}</template>
          </el-table-column>
          <el-table-column label="创建时间" width="200">
            <template #default="scope">{{ formatDateTime(scope.row.creationTimestamp) }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="240">
            <template #default="scope">
              <el-button icon="view" type="text" size="small" @click="viewOrchFunc(scope.row.name, scope.row.namespace)"
                >查 看</el-button
              >
              <el-button icon="expand" type="text" size="small" @click="openScaleDialog(scope.row)">伸缩</el-button>
              <el-button icon="delete" type="text" size="small" @click="deleteFunc(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
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
    </el-card>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%" :destroy-on-close="true">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
    <!-- 伸缩对话框 -->
    <el-dialog v-model="dialogScaleVisible" title="伸缩" width="55%" center>
      <p style="font-weight: bold">
        ReplicaSet {{ replicasetName }} will be updated to reflect the desired replicas count.
      </p>
      <div style="margin: 25px 0 25px 0px">
        <span style="margin-right: 10px">Desired Replicas:</span>
        <el-input-number v-model="desiredNum" :min="0" :max="50" style="margin-right: 20px" />
        <span style="margin-right: 10px">Actual Replicas: </span>
        <el-input-number v-model="ActualNum" disabled />
      </div>
      <!-- <warning-bar :title="warningTitle" /> -->
      <template #footer>
        <el-button @click="closeScaleDialog">取消</el-button>
        <el-button type="primary" @click="scaleFunc">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, watch } from "vue"
import { formatDateTime } from "@/utils/index"
import { getNamespaceNameApi } from "@/api/k8s/namespace"
import { type replicaSetData, getReplicaSetsApi, deleteReplicaSetApi } from "@/api/k8s/replicaSet"
// import { scale } from '@/api/kubernetes/scale'
import VueCodeMirror from "@/components/codeMirror/index.vue"
// import warningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from "element-plus"
import { usePagination } from "@/hooks/usePagination"
import { viewOrch } from "@/utils/k8s/orch"

defineOptions({
  name: "ReplicaSetList"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

// 获取所有namespace
const namespaceList = ref<string[]>()

const getNamespace = async () => {
  const table = await getNamespaceNameApi()
  if (table.code === 0) {
    namespaceList.value = table.data
  }
}
getNamespace()

// 加载replicaSet数据
// const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const searchInfo = reactive({
  namespace: ""
})
const tableData = ref<replicaSetData[]>([])

const loading = ref(false)
const getTableData = async () => {
  loading.value = true
  try {
    const table = await getReplicaSetsApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize,
      ...searchInfo
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      paginationData.total = table.data.total
    }
  } catch (error) {
    console.log(error)
  }
  loading.value = false
}
getTableData()

// 分页
const handleSizeChange = (value: number) => {
  changePageSize(value)
  getTableData()
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  getTableData()
}

// 查询
const onSubmit = () => {
  paginationData.currentPage = 1
  paginationData.pageSize = 10
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.namespace = ""
}

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async (name: string, namespace: string) => {
  formatData.value = await viewOrch(name, "pods", namespace)
  dialogFormVisible.value = true
}

// 伸缩
const dialogScaleVisible = ref(false)
const warningTitle = ref("")
const replicasetName = ref("")
const desiredNum = ref(0)
const ActualNum = ref(0)
const activeRow = ref({})

// -> 打开对话框
const openScaleDialog = (row) => {
  replicasetName.value = row.name
  desiredNum.value = row.replicas
  ActualNum.value = row.replicas
  activeRow.value = row
  dialogScaleVisible.value = true
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${row.namespace} replicaset ${row.name} --replicas=${row.replicas}`
}

watch(desiredNum, (val) => {
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${activeRow.value.namespace} replicaset ${activeRow.value.name} --replicas=${val}`
})

// -> 关闭对话框
const closeScaleDialog = () => {
  dialogScaleVisible.value = false
}

// -> 操作
// const scaleFunc = async () => {
//   const res = await scale({
//     namespace: activeRow.value.namespace,
//     name: activeRow.value.name,
//     kind: "replicaSet",
//     num: desiredNum.value
//   })
//   if (res.code === 0) {
//     const index = tableData.value.indexOf(activeRow.value)
//     tableData.value[index].availableReplicas = desiredNum.value
//     tableData.value[index].replicas = desiredNum.value
//     ElMessage({
//       type: "success",
//       message: "伸缩成功",
//       showClose: true
//     })
//   }
//   dialogScaleVisible.value = false
// }

// 删除
const deleteFunc = async (row) => {
  ElMessageBox.confirm("此操作将永久删除该ReplicaSet, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteReplicaSetApi({ namespace: row.namespace, replicaSet: row.name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
      const index = tableData.value.indexOf(row)
      tableData.value.splice(index, 1)
    }
  })
}
</script>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 20px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
