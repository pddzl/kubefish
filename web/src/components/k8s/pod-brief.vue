<template>
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
</template>

<script lang="ts" setup>
import { toRefs } from "vue"
import { statusPodFilter } from "@/utils/k8s/filter"
import { formatDateTime } from "@/utils/index"

defineOptions({
  name: "PodBrief"
})

const props = defineProps({
  pods: {
    type: Array,
    default: () => []
  }
})

const { pods } = toRefs(props)
</script>
