<template>
  <el-dialog v-model="dialogLogVisible" title="查看日志" width="90%">
    <div>
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="容器">
          <el-select v-model="searchInfo.container" clearable placeholder="请选择">
            <el-option v-for="item in containers" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="行数">
          <el-select v-model="searchInfo.lines" clearable placeholder="请选择">
            <el-option v-for="item in lines" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button type="primary" plain icon="download" @click="donwloadLog">下载日志</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="log">
      <p v-for="(log, key) in podLogList" :key="key" style="color: #b7c4d1">
        {{ log }}
      </p>
    </div>
  </el-dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, watch, toRefs } from "vue"
import { getPodLogApi, getPodDetailApi } from "@/api/k8s/workloads/pod"

defineOptions({
  name: "PodLog"
})

const props = defineProps({
  namespace: {
    type: String,
    default: ""
  },
  pod: {
    type: String,
    default: ""
  }
})

const { namespace, pod } = toRefs(props)

// 查看日志
const dialogLogVisible = ref(false)
const podLogList = ref<string[]>([])
let podLogRaw: string
const lines = ref([20, 50, 100, 200, 500])
const searchInfo = reactive({ container: "", lines: 50 })

const getPodLogData = async () => {
  const podData = await getPodLogApi({ namespace: namespace.value, pod: pod.value, ...searchInfo })
  if (podData.code === 0) {
    podLogRaw = podData.data
    podLogList.value = podData.data.split("\n")
  }
}

const containers = ref<Array<string>>([])
const getPodContainer = async () => {
  const res = await getPodDetailApi({ namespace: namespace.value, pod: pod.value })
  if (res.code === 0) {
    res.data.spec.containers.forEach((element: { name: string }) => {
      containers.value.push(element.name)
    })
  }
}

const viewLog = async () => {
  if (namespace.value && pod.value) {
    getPodContainer()
    searchInfo.container = searchInfo.container
    await getPodLogData()
  }
}
viewLog()

watch([namespace, pod], () => {
  viewLog()
})

const onSubmit = () => {
  getPodLogData()
}

const onReset = () => {
  searchInfo.container = ""
  searchInfo.lines = 50
}

// 下载pod日志
const donwloadLog = () => {
  const url = window.URL.createObjectURL(new Blob([podLogRaw]))
  const a = document.createElement("a")
  a.style.display = "none"
  a.href = url
  a.setAttribute("download", `${props.pod}.txt`)
  document.body.appendChild(a)
  a.click()
  window.URL.revokeObjectURL(a.href)
  document.body.removeChild(a)
}

defineExpose({
  dialogLogVisible
})
</script>

<style lang="scss" scoped>
.log {
  padding-left: 15px;
  background-color: #242e42;
  border-radius: 4px;
  height: calc(100vh - 200px);
  overflow-y: auto;
}
</style>
