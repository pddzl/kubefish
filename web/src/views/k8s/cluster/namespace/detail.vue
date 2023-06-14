<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc" :disabled="testDel()">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="namespaceDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="namespaceDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>状态</p>
                <span class="content">
                  <el-tag :type="NamespaceStatusFilter(namespaceDetail.status)" size="small"
                    >{{ namespaceDetail.status }}
                  </el-tag>
                </span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item
          v-if="Array.isArray(namespaceDetail.resourceQuotas) && namespaceDetail.resourceQuotas.length > 0"
          title="资源配额"
          name="resourceQuota"
        >
          <div style="margin-right: 20px">
            <el-table :data="namespaceDetail.resourceQuotas" border>
              <el-table-column label="名称" align="center" prop="name" />
              <el-table-column label="used" align="center">
                <template #default="scope">
                  <span>{{ scope.row.used }}</span>
                </template>
              </el-table-column>
              <el-table-column label="hard" align="center">
                <template #default="scope">
                  <span>{{ scope.row.hard }}</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item
          v-if="Array.isArray(namespaceDetail.limitRanges) && namespaceDetail.limitRanges.length > 0"
          title="资源限制"
          name="resourceLimit"
        >
          <div
            v-for="(rLimit, keyRl) in namespaceDetail.limitRanges"
            :key="keyRl"
            style="margin-right: 20px; margin-bottom: 20px"
          >
            <div style="margin-bottom: 10px">
              <el-tag size="small">{{ rLimit.name }}</el-tag>
            </div>
            <el-table :data="rLimit.limits">
              <el-table-column label="类型" align="center" prop="type" />
              <el-table-column label="max" align="center">
                <template #default="scope">
                  {{ scope.row.max }}
                </template>
              </el-table-column>
              <el-table-column label="min" align="center">
                <template #default="scope">
                  {{ scope.row.min }}
                </template>
              </el-table-column>
              <el-table-column label="default" align="center">
                <template #default="scope">
                  {{ scope.row.default }}
                </template>
              </el-table-column>
              <el-table-column label="defaultRequest" align="center">
                <template #default="scope">
                  {{ scope.row.defaultRequest }}
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { deleteNamespaceApi, getNamespaceDetailApi } from "@/api/k8s/namespace"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { NamespaceStatusFilter } from "@/hooks/filter"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { ref } from "vue"
import { useRoute } from "vue-router"

defineOptions({
  name: "NamespaceDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "spec", "resourceQuota", "resourceLimit"])

// 路由
const route = useRoute()
const namespace = route.query.name as string

const testDel = (): boolean => {
  if (namespace === "default") {
    return true
  }
  const regex = new RegExp("kube-")
  return regex.test(namespace)
}

// 加载namespace详情
const namespaceDetail = ref<any>({})

const getData = async () => {
  await getNamespaceDetailApi({ name: namespace }).then((res) => {
    if (res.code === 0) {
      namespaceDetail.value = res.data
    }
  })
}
getData()

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(namespace, "namespaces")
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该Namespace, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteNamespaceApi({ name: namespace })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
