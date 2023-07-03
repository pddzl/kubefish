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
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-card v-loading="loading" shadow="never">
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column label="名称" min-width="220">
            <template #default="scope">
              <router-link
                :to="{
                  name: 'ReplicaSetDetail',
                  query: { namespace: scope.row.namespace, name: scope.row.name }
                }"
              >
                <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
              </router-link>
            </template>
          </el-table-column>
          <el-table-column label="命名空间" prop="namespace" min-width="120" />
          <el-table-column label="Pods" prop="pods" min-width="80">
            <template #default="scope">{{ scope.row.availableReplicas }} / {{ scope.row.replicas }}</template>
          </el-table-column>
          <el-table-column label="创建时间" width="200">
            <template #default="scope">{{ formatDateTime(scope.row.creationTimestamp) }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作">
            <template #default="scope">
              <el-popover placement="bottom" trigger="click" :popper-style="pWidth">
                <template #reference>
                  <el-button icon="more" type="primary" link size="small" />
                </template>
                <div style="disply: flex; flex-direction: column; justify-content: center; align-items: center">
                  <div
                    style="
                      width: 75px;
                      padding: 5px;
                      border-bottom: solid;
                      border-width: 1px;
                      border-color: rgba(128, 128, 128, 0.157);
                    "
                  >
                    <el-button
                      icon="view"
                      type="primary"
                      link
                      @click="viewOrchFunc(scope.row.name, scope.row.namespace)"
                      >查 看</el-button
                    >
                  </div>
                  <div style="width: 75px; padding: 5px">
                    <el-button icon="delete" type="primary" link @click="deleteFunc(scope.row)">删 除</el-button>
                  </div>
                </div>
              </el-popover>
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
  </div>
</template>

<script lang="ts" setup>
import { getNamespaceNameApi } from "@/api/k8s/cluster/namespace"
import { deleteReplicaSetApi, getReplicaSetsApi, type replicaSetBrief } from "@/api/k8s/workloads/replicaSet"
import { formatDateTime } from "@/utils/index"
import { reactive, ref } from "vue"
// import { scale } from '@/api/kubernetes/scale'
import VueCodeMirror from "@/components/codeMirror/index.vue"
// import warningBar from '@/components/warningBar/warningBar.vue'
import { usePagination } from "@/hooks/usePagination"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"

defineOptions({
  name: "ReplicaSetList"
})

// 分页
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

const searchInfo = reactive({
  namespace: ""
})
const tableData = ref<replicaSetBrief[]>([])

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
  formatData.value = await viewOrch(name, "replicasets", namespace)
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async (row: replicaSetBrief) => {
  ElMessageBox.confirm("此操作将永久删除该ReplicaSet, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteReplicaSetApi({ namespace: row.namespace, name: row.name })
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

const pWidth = reactive({ "min-width": "100px", width: "100px" })
</script>

<style lang="scss" scoped>
.search-wrapper {
  margin-bottom: 20px;
  :deep(.el-card__body) {
    padding-bottom: 2px;
  }
}
</style>
