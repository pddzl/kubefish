<template>
  <div class="app-container">
    <el-card v-loading="loading" shadow="never">
      <div class="table-wrapper">
        <el-table :data="tableData">
          <el-table-column label="名称" min-width="100">
            <template #default="scope">
              <router-link
                :to="{
                  name: 'ClusterRoleDetail',
                  query: { name: scope.row.name }
                }"
              >
                <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
              </router-link>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" min-width="100">
            <template #default="scope">{{ formatDateTime(scope.row.creationTimestamp) }}</template>
          </el-table-column>
          <el-table-column fixed="right" label="操作" width="200">
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
                    <el-button icon="view" type="primary" link @click="viewOrchFunc(scope.row.name)">查 看</el-button>
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
import { deleteClusterRoleApi, getClusterRoleListApi, type ClusterRoleBrief } from "@/api/k8s/accessControl/clusterRole"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { usePagination } from "@/hooks/usePagination"
import { formatDateTime } from "@/utils/index"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { reactive, ref } from "vue"

defineOptions({
  name: "ClusterRoleList"
})

// 分页
const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const tableData = ref<ClusterRoleBrief[]>([])

const loading = ref(false)
const getTableData = async () => {
  loading.value = true
  try {
    const table = await getClusterRoleListApi({
      page: paginationData.currentPage,
      pageSize: paginationData.pageSize
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

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async (name: string) => {
  formatData.value = await viewOrch(name, "clusterroles")
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = (row: ClusterRoleBrief) => {
  ElMessageBox.confirm("此操作将永久删除该ClusterRole, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteClusterRoleApi({ name: row.name })
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
