<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus">新增</el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column label="名称">
            <template #default="scope">
              <router-link :to="{ name: 'NamespaceDetail', query: { name: scope.row.name } }">
                <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
              </router-link>
            </template>
          </el-table-column>
          <el-table-column label="状态">
            <template #default="scope">
              <el-tag :type="NamespaceStatusFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="创建时间">
            <template #default="scope">{{ formatDateTime(scope.row.creationTimestamp) }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作">
            <template #default="scope">
              <el-button icon="view" type="primary" link size="small" @click="viewOrchFunc(scope.row.name)"
                >查看</el-button
              >
              <el-button icon="delete" type="primary" link size="small" @click="deleteFunc(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogFormVisible" title="编辑资源" :destroy-on-close="true">
      <!-- eslint-disable-next-line vue/attribute-hyphenation -->
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { NamespaceStatusFilter } from "@/hooks/filter"
import { getNamespacesApi } from "@/api/k8s/namespace"
import { usePagination } from "@/hooks/usePagination"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { formatDateTime } from "@/utils/index"
import { viewOrch } from "@/utils/k8s/orch"

defineOptions({
  name: "NamespaceList"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)

// 加载namespace list
const tableData = ref<any[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getNamespacesApi({ page: paginationData.currentPage, pageSize: paginationData.pageSize })
    if (res.code === 0) {
      tableData.value = res.data.list
      paginationData.total = res.data.total
    }
  } catch (error) {
    //
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

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async (name: string) => {
  formatData.value = await viewOrch(name, "namespaces", "")
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async (row) => {
  ElMessageBox.confirm("此操作将永久删除该Namespace, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteNamespace({ name: row.name })
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
