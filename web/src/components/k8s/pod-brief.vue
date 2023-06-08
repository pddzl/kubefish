<template>
  <div style="margin-bottom: 20px">
    <el-table :data="pods">
      <el-table-column label="名称" min-width="140">
        <template #default="scope">
          <router-link :to="{ name: 'PodDetail', query: { pod: scope.row.name, namespace: scope.row.namespace } }">
            <el-link type="primary" :underline="false">{{ scope.row.name }}</el-link>
          </router-link>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="状态" min-width="60">
        <template #default="scope">
          <el-tag :type="statusPodFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="node" label="节点" min-width="80" />
      <el-table-column prop="namespace" label="命名空间" min-width="80" />
      <el-table-column label="创建时间">
        <template #default="scope">{{ formatDateTime(scope.row.creationTimestamp) }}</template>
      </el-table-column>
    </el-table>
  </div>
  <div class="pager-wrapper">
    <el-pagination
      background
      :layout="paginationData.layout"
      :page-sizes="paginationData.pageSizes"
      :total="total"
      :page-size="paginationData.pageSize"
      :currentPage="paginationData.currentPage"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    />
  </div>
</template>

<script lang="ts" setup>
import { usePagination } from "@/hooks/usePagination"
import { formatDateTime } from "@/utils/index"
import { statusPodFilter } from "@/utils/k8s/filter"
import { toRefs } from "vue"

defineOptions({
  name: "PodBriefC"
})

// 分页
const { paginationData, changeCurrentPage, changePageSize } = usePagination()

const props = defineProps({
  pods: {
    type: Array,
    default: () => []
  },
  total: {
    type: Number,
    default: 0
  }
})

const { pods, total } = toRefs(props)

const emit = defineEmits(["getRelatedPodsFunc"])

const handleSizeChange = (value: number) => {
  changePageSize(value)
  emit("getRelatedPodsFunc", { currentPage: paginationData.currentPage, pageSize: paginationData.pageSize })
}

const handleCurrentChange = (value: number) => {
  changeCurrentPage(value)
  emit("getRelatedPodsFunc", { currentPage: paginationData.currentPage, pageSize: paginationData.pageSize })
}
</script>
