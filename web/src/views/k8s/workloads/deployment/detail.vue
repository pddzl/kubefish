<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
      <el-button icon="expand" type="warning" plain @click="openScaleDialog">伸缩</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="deploymentDetail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="deploymentDetail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="deploymentDetail.spec" title="资源信息" name="spec">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ deploymentDetail.spec.replicas }}</span>
              </div>
              <div class="item">
                <p>Selector</p>
                <div class="content">
                  <span v-for="(label, index) in deploymentDetail.spec.selector" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>Strategy</p>
                <span class="content">{{ deploymentDetail.spec.strategy.type }}</span>
              </div>
              <div v-if="deploymentDetail.spec.strategy.type === 'RollingUpdate'" class="item">
                <p>maxUnavailable</p>
                <span class="content">{{ deploymentDetail.spec.strategy.rollingUpdate.maxUnavailable }}</span>
              </div>
              <div v-if="deploymentDetail.spec.strategy.type === 'RollingUpdate'" class="item">
                <p>maxSurge</p>
                <span class="content">{{ deploymentDetail.spec.strategy.rollingUpdate.maxUnavailable }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="deploymentDetail.status" title="状态" name="status">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>replicas</p>
                <span class="content">{{ deploymentDetail.status.replicas }}</span>
              </div>
              <div class="item">
                <p>updatedReplicas</p>
                <span class="content">{{ deploymentDetail.status.updatedReplicas }}</span>
              </div>
              <div class="item">
                <p>readyReplicas</p>
                <span class="content">{{ deploymentDetail.status.readyReplicas }}</span>
              </div>
              <div class="item">
                <p>availableReplicas</p>
                <span class="content">{{ deploymentDetail.status.availableReplicas }}</span>
              </div>
            </div>
            <div style="margin: 20px 20px 0px 0px">
              <el-table :data="deploymentDetail.status.conditions">
                <el-table-column label="类别" prop="type" />
                <el-table-column label="状态" prop="status">
                  <template #default="scope">
                    <el-tag :type="statusRsFilter(scope.row.status)" size="small">
                      {{ scope.row.status }}
                    </el-tag>
                  </template>
                </el-table-column>
                <el-table-column label="上次检测时间">
                  <template #default="scope">
                    {{ formatDateTime(scope.row.lastUpdateTime) }}
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
          </div>
        </el-collapse-item>
        <el-collapse-item title="ReplicaSet" name="replicaSet">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <div class="content">
                  <router-link
                    :to="{
                      name: 'ReplicaSetDetail',
                      query: { replicaSet: newReplicaSet.name, namespace: newReplicaSet.namespace }
                    }"
                  >
                    <el-link type="primary" :underline="false">{{ newReplicaSet.name }} </el-link>
                  </router-link>
                </div>
              </div>
              <div class="item">
                <p>命名空间</p>
                <span class="content">{{ newReplicaSet.namespace }}</span>
              </div>
              <div class="item">
                <p>Replicas</p>
                <span class="content">{{ newReplicaSet.replicas }}</span>
              </div>
              <div class="item">
                <p>创建时间</p>
                <span class="content">{{ formatDateTime(newReplicaSet.creationTimestamp) }}</span>
              </div>
            </div>
            <div class="row">
              <div class="item">
                <p>标签:</p>
                <div class="content">
                  <span v-for="(label, index) in newReplicaSet.labels" :key="index" class="shadow">
                    {{ index }}
                    <span v-if="label">:</span>
                    {{ label }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogViewVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
    <el-dialog v-model="dialogScaleVisible" title="伸缩" width="45%" center>
      <p style="font-weight: bold; font-size: 15px">
        Deployment {{ name }} will be updated to reflect the desired replicas count.
      </p>
      <div style="margin: 25px 0 25px 0px">
        <span style="margin-right: 10px">Desired Replicas:</span>
        <el-input-number v-model="desiredNum" :min="0" :max="50" style="margin-right: 20px" />
        <span style="margin-right: 10px">Actual Replicas: </span>
        <el-input-number v-model="ActualNum" disabled />
      </div>
      <WarningBar :title="warningTitle" />
      <template #footer>
        <el-button @click="closeScaleDialog">取消</el-button>
        <el-button type="primary" @click="scaleFunc">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { deleteDeploymentApi, getDeploymentDetailApi, getDeploymentRsApi, scaleDeployment } from "@/api/k8s/deployment"
import WarningBar from "@/components/WarningBar/warningBar.vue"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { formatDateTime } from "@/utils/index"
import { statusRsFilter } from "@/utils/k8s/filter.js"
import { viewOrch } from "@/utils/k8s/orch"
import { ElMessage, ElMessageBox } from "element-plus"
import { ref, watch } from "vue"
import { useRoute } from "vue-router"

defineOptions({
  name: "DeploymentDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "spec", "status", "replicaSet"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 获取deployment详情
const deploymentDetail = ref({
  metadata: {},
  spec: {
    replicas: 0,
    selector: {},
    strategy: {
      type: "",
      rollingUpdate: {
        maxUnavailable: 0,
        maxSurge: ""
      }
    }
  },
  status: {
    replicas: 0,
    updatedReplicas: 0,
    readyReplicas: 0,
    availableReplicas: 0,
    conditions: []
  }
})

const getDeploymentDetailFunc = async () => {
  await getDeploymentDetailApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      deploymentDetail.value = res.data
    }
  })
}
getDeploymentDetailFunc()

// 获取deployment关联的replicaset
const newReplicaSet = ref({ name: "", namespace: "", labels: {}, creationTimestamp: "", replicas: "" })

const getNewReplicaSetData = async () => {
  await getDeploymentRsApi({ namespace: namespace, name: name }).then((res) => {
    if (res.code === 0) {
      newReplicaSet.value = res.data
    }
  })
}
getNewReplicaSetData()

// 查看编排
const dialogViewVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async () => {
  formatData.value = await viewOrch(name, "deployments", namespace)
  dialogViewVisible.value = true
}

// 伸缩
const dialogScaleVisible = ref(false)
const warningTitle = ref("")
const desiredNum = ref(0)
const ActualNum = ref(0)

const openScaleDialog = async () => {
  desiredNum.value = deploymentDetail.value.status.replicas
  ActualNum.value = deploymentDetail.value.status.availableReplicas
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} deployment ${name} --replicas=${ActualNum.value}`
  dialogScaleVisible.value = true
}

watch(desiredNum, (val) => {
  warningTitle.value = `This action is equivalent to: kubectl scale -n ${namespace} deployment ${name} --replicas=${val}`
})

const closeScaleDialog = () => {
  dialogScaleVisible.value = false
}

const scaleFunc = async () => {
  const res = await scaleDeployment({
    namespace: namespace,
    name: name,
    num: desiredNum.value
  })
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "伸缩成功",
      showClose: true
    })
    // 刷新数据
    getDeploymentDetailFunc()
    getNewReplicaSetData()
  }
  dialogScaleVisible.value = false
}

// 删除
const deleteFunc = () => {
  ElMessageBox.confirm("此操作将永久删除该Deployment, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteDeploymentApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
