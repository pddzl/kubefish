<template>
  <div class="app-container">
    <div class="detail-operation">
      <el-button icon="view" type="primary" plain @click="viewOrchFunc(pod, namespace)">查看</el-button>
      <el-button icon="expand" type="primary" plain @click="viewLog">日志</el-button>
      <el-button icon="expand" type="primary" plain @click="routerPod()">终端</el-button>
      <el-button icon="delete" type="danger" plain @click="deleteFunc">删除</el-button>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="podDetail.metadata" title="元数据" name="metadata">
          <div class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <span class="content">{{ podDetail.metadata.name }}</span>
              </div>
              <div class="item">
                <p>命名空间</p>
                <span class="content">{{ podDetail.metadata.namespace }}</span>
              </div>
              <div class="item">
                <p>UID</p>
                <span class="content">{{ podDetail.metadata.uid }}</span>
              </div>
              <div class="item">
                <p>创建时间</p>
                <span class="content">{{ formatDateTime(podDetail.metadata.creationTimestamp) }}</span>
              </div>
            </div>
            <div v-if="podDetail.metadata.labels" class="row">
              <div class="item">
                <p>标签</p>
                <div class="content">
                  <span v-for="(label, index) in podDetail.metadata.labels" :key="index" class="shadow"
                    >{{ index }}: {{ label }}</span
                  >
                </div>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.spec && podDetail.status" title="资源信息" name="resource">
          <div class="info-box">
            <div class="row">
              <div v-if="podDetail.spec.nodeName" class="item">
                <p>节点</p>
                <span class="content">{{ podDetail.spec.nodeName }}</span>
              </div>
              <div v-if="podDetail.status.phase" class="item">
                <p>状态</p>
                <el-tag :type="PodStatusFilter(podDetail.status.phase)" size="small">
                  {{ podDetail.status.phase }}
                </el-tag>
              </div>
              <div v-if="podDetail.status.podIP" class="item">
                <p>IP</p>
                <span class="content">{{ podDetail.status.podIP }}</span>
              </div>
              <div v-if="podDetail.status.qosClass" class="item">
                <p>服务质量</p>
                <span class="content">{{ podDetail.status.qosClass }}</span>
              </div>
              <div v-if="podDetail.spec.restartPolicy" class="item">
                <p>重启策略</p>
                <span class="content">{{ podDetail.spec.restartPolicy }}</span>
              </div>
              <div v-if="podDetail.spec.serviceAccountName">
                <p>服务账号</p>
                <span class="content">{{ podDetail.spec.serviceAccountName }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.status?.conditions" title="状态" name="conditions">
          <div class="info-table">
            <el-table :data="podDetail.status.conditions">
              <el-table-column prop="type" label="类别" />
              <el-table-column prop="status" label="状态" />
              <el-table-column label="最后检查时间">
                <template #default="scope">
                  <span v-if="scope.row.lastProbeTime">{{ formatDateTime(scope.row.lastProbeTime) }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
              <el-table-column label="最后迁移时间">
                <template #default="scope">
                  <span v-if="scope.row.lastTransitionTime">{{ formatDateTime(scope.row.lastTransitionTime) }}</span>
                  <span v-else>-</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.metadata?.ownerReferences" title="控制器" name="controller">
          <!-- <div v-for="reference in podDetail.metadata.ownerReferences" :key="reference" class="info-box">
            <div class="row">
              <div class="item">
                <p>名称</p>
                <div class="content">
                  <router-link
                    v-if="reference.kind === 'ReplicaSet'"
                    :to="{ name: 'replicaSet_detail', query: { replicaSet: reference.name, namespace: namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                  <router-link
                    v-else-if="reference.kind === 'DaemonSet'"
                    :to="{ name: 'daemonSet_detail', query: { daemonSet: reference.name, namespace: namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                  <router-link
                    v-else-if="reference.kind === 'Job'"
                    :to="{ name: 'job_detail', query: { job: reference.name, namespace: namespace } }"
                  >
                    <el-link type="primary" :underline="false">{{ reference.name }}</el-link>
                  </router-link>
                </div>
              </div>
              <div class="item">
                <p>类型</p>
                <span class="content">{{ reference.kind }}</span>
              </div>
            </div>
          </div> -->
        </el-collapse-item>
        <el-collapse-item v-if="podDetail.spec?.containers.length > 0" title="容器" name="container">
          <container :data="podDetail.spec.containers" />
        </el-collapse-item>
        <el-collapse-item
          v-if="podDetail.spec?.initContainers && podDetail.spec.initContainers.length > 0"
          title="初始容器"
          name="initContainers"
        >
          <container :data="podDetail.spec.initContainers" />
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <vue-code-mirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
    <PodLog ref="podLogRef" :namespace="namespace" :pod="pod" />
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useRoute } from "vue-router"
import { PodStatusFilter } from "@/hooks/filter.js"
import { getPodDetailApi, deletePodApi } from "@/api/k8s/pod"
import { formatDateTime } from "@/utils/index"
import Container from "./components/container.vue"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { viewOrch } from "@/utils/k8s/orch"
import PodLog from "./components/log.vue"

// 折叠面板
const activeNames = ref(["metadata", "resource", "conditions", "controller", "container", "initContainers"])

// 路由
const route = useRoute()
const pod = route.query.pod as string
const namespace = route.query.namespace as string

// 加载pod详情
const podDetail = ref<any>({})
const containers = ref<string[]>([])

const getData = async () => {
  await getPodDetailApi({ pod: pod, namespace: namespace }).then((res) => {
    if (res.code === 0) {
      podDetail.value = res.data
      res.data.spec.containers.forEach((element: { name: string }) => {
        containers.value.push(element.name)
      })
    }
  })
}
getData()

// 查看编排
const dialogFormVisible = ref(false)
const formatData = ref<string>("")
const viewOrchFunc = async (name: string, namespace: string) => {
  formatData.value = await viewOrch(name, "pods", namespace)
  dialogFormVisible.value = true
}

const podLogRef = ref(null)
const viewLog = async () => {
  podLogRef.value.dialogLogVisible = true
}

// 删除pod
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该Pod, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deletePodApi({ namespace: namespace, pod: pod })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}

// 跳转日志、webssh页面
const routerPod = () => {
  const url = `http://${window.location.host}/#/k8s/pod/webssh?pod=${pod}&namespace=${namespace}`
  const width = 1300
  const height = 800
  // 获取屏幕的宽度和高度
  const screenWidth = window.screen.width
  const screenHeight = window.screen.height

  // 计算新窗口的左上角位置
  const left = (screenWidth - width) / 2
  const top = (screenHeight - height) / 2
  // 构建 features 参数，包括宽度、高度和位置信息
  const features = `width=${width},height=${height},left=${left},top=${top}`
  // window.open(url, "_blank", features)
  window.open(url, "_blank", features)
}
</script>

