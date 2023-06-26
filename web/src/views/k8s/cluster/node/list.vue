<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="toolbar-wrapper">
        <div>
          <el-button type="primary" icon="CirclePlus" :disabled="true">新增</el-button>
        </div>
        <div>
          <el-tooltip content="刷新" effect="light">
            <el-button type="primary" icon="RefreshRight" circle plain @click="getTableData" />
          </el-tooltip>
        </div>
      </div>
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column label="名称" prop="name">
            <template #default="scope">
              <router-link :to="{ name: 'NodeDetail', query: { name: scope.row.name } }">
                <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
              </router-link>
            </template>
          </el-table-column>
          <el-table-column prop="internalIP" label="内部IP" />
          <el-table-column prop="roles" label="角色">
            <template #default="scope">
              <span v-for="(role, index) in scope.row.role" :key="index" style="margin-right: 5px">
                <el-tag size="small">{{ role }}</el-tag>
              </span>
            </template>
          </el-table-column>
          <el-table-column label="状态" min-width="120" prop="status">
            <template #default="scope">
              <el-tag :type="nodeStatusTypeFilter(scope.row.status)" size="small">{{
                nodeStatusFilter(scope.row.status)
              }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="CPU (core)" min-width="100" prop="cpu" />
          <el-table-column label="内存 (GB)" min-width="100">
            <template #default="scope">{{
              (parseInt(scope.row.memory.slice(0, -2)) / 1024 / 1024).toFixed(2)
            }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="200">
            <template #default="scope">
              <el-button icon="view" size="small" type="primary" link @click="viewOrchFunc(scope.row.name)"
                >查看</el-button
              >
              <el-button icon="delete" size="small" type="primary" link :disabled="true">删除</el-button>
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
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { nodeStatusTypeFilter, nodeStatusFilter } from "@/hooks/filter"
import { usePagination } from "@/hooks/usePagination"
import { type NodeData, getNodesApi } from "@/api/k8s/cluster/node"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { viewOrch } from "@/utils/k8s/orch"

defineOptions({
  name: "NodeList"
})

const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const loading = ref(false)

const tableData = ref<NodeData[]>([])

const getTableData = async () => {
  loading.value = true
  try {
    const res = await getNodesApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
    })
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
  formatData.value = await viewOrch(name, "nodes")
  dialogFormVisible.value = true
}
</script>
