<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="detail.metadata" title="元数据" name="metadata">
          <MetaData :metadata="detail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="detail.rules && detail.rules.length > 0" title="规则" name="rules">
          <div class="info-table">
            <el-table :data="detail.rules">
              <el-table-column label="资源">
                <template #default="scope">
                  <span v-for="(rs, index) in scope.row.resources" :key="index">
                    {{ rs }}<span v-if="index !== scope.row.resources.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="非资源URL">
                <template #default="scope">
                  <span v-for="(nRsUrl, index) in scope.row.nonResourceURLs" :key="index">
                    {{ nRsUrl }}<span v-if="index !== scope.row.nonResourceURLs.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="资源名">
                <template #default="scope">
                  <span v-for="(rn, index) in scope.row.resourceNames" :key="index">
                    {{ rn }}<span v-if="index !== scope.row.resourceNames.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="动作">
                <template #default="scope">
                  <span v-for="(v, index) in scope.row.verbs" :key="index">
                    {{ v }}<span v-if="index !== scope.row.verbs.length - 1">,</span>
                  </span>
                </template>
              </el-table-column>
              <el-table-column label="API组">
                <template #default="scope">
                  <span v-for="(v, index) in scope.row.apiGroups" :key="index">
                    {{ v }}<span v-if="index !== scope.row.apiGroups.length - 1">,</span>
                  </span>
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
import { ref } from "vue"
import { useRoute } from "vue-router"
import { getClusterRoleDetailApi, deleteClusterRoleApi } from "@/api/k8s/accessControl/clusterRole"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { viewOrch } from "@/utils/k8s/orch"

defineOptions({
  name: "ClusterRoleDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "rules"])

// 路由
const route = useRoute()
const name = route.query.name as string

// 获取clusterRole详情
const detail = ref({ metadata: {}, rules: [] })

const getDetailData = async () => {
  await getClusterRoleDetailApi({ name: name }).then((res) => {
    if (res.code === 0) {
      detail.value = res.data
    }
  })
}
getDetailData()

// 查看
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "clusterroles")
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该ClusterRole, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteClusterRoleApi({ name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>

<style lang="scss" scoped>
.data-container {
  padding: 15px;
  border-radius: 4px;
  background-color: #f9fbfd;
  margin-right: 20px;
}

.key-container {
  font-size: 14px;
  font-weight: 600;
  font-style: normal;
  color: #242e42;
  text-shadow: 0 4px 8px rgba(36, 46, 66, 0.1);
}

.value-container {
  margin-top: 8px;
  padding: 12px;
  border-radius: 4px;
  background-color: #fff;
  border: 1px solid #e3e9ef;
  word-break: break-all;
  white-space: pre-wrap;
}

.data-wrapper:not(:last-child) {
  margin-bottom: 15px;
}
</style>
