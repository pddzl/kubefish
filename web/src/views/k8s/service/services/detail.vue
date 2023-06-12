<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="serviceDetail.metadata" title="元数据" name="metadata">
          <MetaData :metadata="serviceDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="serviceDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in serviceDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
          <div class="info-table" style="margin-top: 20px">
            <p>端口</p>
            <el-table :data="serviceDetail.spec.ports">
              <!-- <el-table-column label="名称" prop="name" /> -->
              <el-table-column label="协议" prop="protocol" />
              <el-table-column label="端口" prop="port" />
              <el-table-column label="目标端口" prop="targetPort" />
              <el-table-column label="nodePort" prop="nodePort" />
            </el-table>
          </div>
        </el-collapse-item>
        <!-- <el-collapse-item v-if="serviceDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ serviceDetail.status.replicas }}</span>
              </div>
              <div class="item">
                <p>fullyLabeledReplicas</p>
                <span class="content">{{ serviceDetail.status.fullyLabeledReplicas }}</span>
              </div>
              <div class="item">
                <p>readyReplicas</p>
                <span class="content">{{ serviceDetail.status.readyReplicas }}</span>
              </div>
              <div class="item">
                <p>availableReplicas</p>
                <span class="content">{{ serviceDetail.status.availableReplicas }}</span>
              </div>
            </div>
          </div>
          <div v-if="serviceDetail.status.conditions" style="margin-top: 20px; margin-right: 25px">
            <el-table :data="serviceDetail.status.conditions">
              <el-table-column label="类别" prop="type" />
              <el-table-column label="状态" prop="status">
                <template #default="scope">
                  <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                    {{ scope.row.status }}
                  </el-tag>
                </template>
              </el-table-column>
              <el-table-column label="上次迁移时间">
                <template #default="scope">
                  {{ formatDateTime(scope.row.lastTransitionTime) }}
                </template>
              </el-table-column>
              <el-table-column label="原因" prop="reason" />
              <el-table-column label="信息" prop="message" :show-overflow-tooltip="true" />
            </el-table>
          </div>
        </el-collapse-item> -->
        <!-- <el-collapse-item title="Pods" name="pods">
          <div class="info-table">
            <PodBriefC :total="total" :pods="relatedPods" @getRelatedPodsFunc="getRelatedPodsFunc" />
          </div>
        </el-collapse-item> -->
      </el-collapse>
    </div>
    <!-- 查看编排对话框 -->
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { deleteServiceApi, getServiceDetailApi, getServicePodsApi } from "@/api/k8s/service"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { ref } from "vue"
import { useRoute } from "vue-router"

defineOptions({
  name: "ServicesDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "spec", "ports"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 加载service详情
const serviceDetail = ref<any>({})

const getDetail = async () => {
  await getServiceDetailApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      serviceDetail.value = res.data
    }
  })
}
getDetail()

// 加载service关联pods
const relatedPods = ref([])
const total = ref(0)

interface pageObj {
  currentPage: number
  pageSize: number
}

const getRelatedPodsFunc = async (obj: pageObj) => {
  const res = await getServicePodsApi({
    page: obj.currentPage,
    pageSize: obj.pageSize,
    namespace: namespace,
    name: name
  })
  if (res.code === 0) {
    relatedPods.value = res.data.list
    total.value = res.data.total
  }
}
// getRelatedPodsFunc({ currentPage: 1, pageSize: 10 })

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "services", namespace)
  dialogViewVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该Service, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteServiceApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
