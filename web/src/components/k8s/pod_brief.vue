<template>
  <el-table :data="props.pods">
    <el-table-column label="名称" min-width="140">
      <!-- <template #default="scope">
        <router-link
          :to="{ name: 'pod_detail', query: { pod: scope.row.metadata.name, namespace: scope.row.metadata.namespace } }"
        >
          <el-link type="primary" :underline="false">{{ scope.row.metadata.name }}</el-link>
        </router-link>
      </template> -->
    </el-table-column>
    <el-table-column prop="status" label="状态" min-width="60">
      <template #default="scope">
        <el-tag :type="PodStatusFilter(scope.row.status)" size="small">{{ scope.row.status }}</el-tag>
      </template>
    </el-table-column>
    <el-table-column prop="nodeName" label="节点" min-width="80" />
    <el-table-column prop="metadata.namespace" label="命名空间" min-width="80" />
    <el-table-column label="创建时间" width="200">
      <template #default="scope">{{ formatDateTime(scope.row.metadata.creationTimestamp) }}</template>
    </el-table-column>
    <el-table-column fixed="right" label="操作" width="220">
      <template #default="scope">
        <el-button
          icon="tickets"
          type="primary"
          link
          size="small"
          @click="routerPod(scope.row.metadata.name, scope.row.namespace, 'log')"
          >日志</el-button
        >
        <el-button
          icon="ArrowRight"
          type="primary"
          link
          size="small"
          @click="routerPod(scope.row.metadata.name, scope.row.namespace, 'terminal')"
          >终端</el-button
        >
        <el-button icon="delete" type="primary" link size="small">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts" setup>
import { PodStatusFilter } from "@/hooks/filter.js"
import { formatDateTime } from "@/utils/index"
import { useRouter } from "vue-router"

defineOptions({
  name: "PodBrief"
})

const props = defineProps({
  pods: {
    type: Array,
    default: () => []
  }
})

const router = useRouter()

// 跳转日志页面
const routerPod = async (name: string, namespace: string, dest: string) => {
  if (dest === "log") {
    router.push({ name: "pod_log", query: { pod: name, namespace: namespace } })
  } else if (dest === "terminal") {
    router.push({ name: "pod_terminal", query: { pod: name, namespace: namespace } })
  }
}
</script>
