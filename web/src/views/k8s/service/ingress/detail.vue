<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="ingressDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="ingressDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="ingressDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div v-if="ingressDetail.spec.ingressClassName" class="row">
              <div class="item">
                <p>Ingress Class Name</p>
                <span class="content">{{ ingressDetail.spec.ingressClassName }}</span>
              </div>
            </div>
            <div v-if="ingressDetail.status.endPoints" class="row">
              <div class="item">
                <p>EndPoints</p>
                <div class="content">
                  <span v-for="endPoint in ingressDetail.status.endPoints" :key="endPoint" class="shadow">
                    {{ endPoint }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="ingressDetail.spec" title="规则" name="rules" class="rules">
          <div v-for="h in ingressDetail.spec.rules" :key="h">
            <span style="margin-right: 10px">主机: {{ h.host }}</span>
            <span>协议: <span v-if="h.http">HTTP</span></span>
            <el-table :data="h.http.paths">
              <el-table-column label="路径" prop="path" />
              <el-table-column label="路径类型" prop="pathType" />
              <el-table-column label="服务" prop="backend.service.name">
                <template #default="scope">
                  <router-link
                    :to="{
                      name: 'ServiceDetail',
                      query: { name: scope.row.backend.service.name, namespace: namespace }
                    }"
                  >
                    <el-link type="primary" :underline="false">{{ scope.row.backend.service.name }}</el-link>
                  </router-link>
                </template>
              </el-table-column>
              <el-table-column label="服务">
                <template #default="scope">
                  <span v-if="scope.row.backend.service.port.number">{{ scope.row.backend.service.port.number }}</span>
                  <span v-else-if="scope.row.backend.service.port.name">{{ scope.row.backend.service.port.name }}</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <!-- 查看编排对话框 -->
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useRoute } from "vue-router"
import { getIngressDetailApi, deleteIngressApi } from "@/api/k8s/service/ingress"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { viewOrch } from "@/utils/k8s/orch"

// 折叠面板
const activeNames = ref(["metadata", "spec", "rules"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 获取ingress详情
const ingressDetail = ref<any>({})

const getIngressDetailFunc = async () => {
  await getIngressDetailApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      ingressDetail.value = res.data
    }
  })
}
getIngressDetailFunc()

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "ingresses", namespace)
  dialogViewVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该Ingress, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteIngressApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>

<style scoped lang="scss">
.rules {
  div {
    margin-right: 20px;
  }
  div:not(:last-child) {
    margin-bottom: 20px;
  }
}
</style>
