<template>
  <div class="app-container">
    <div class="detail-operation">
      <div class="detail-operation">
        <el-button icon="view" type="primary" plain @click="viewOrchFunc">查看</el-button>
        <el-button icon="delete" type="danger" plain @click="deleteFunc">删除 </el-button>
      </div>
    </div>
    <div class="kop-collapse">
      <el-collapse v-model="activeNames">
        <el-collapse-item v-if="detail.metadata" title="元数据" name="metadata">
          <meta-data :metadata="detail.metadata" />
        </el-collapse-item>
        <el-collapse-item v-if="detail.data" title="数据" name="data">
          <div class="data-container">
            <el-icon :size="20" @click="decodeValue">
              <span v-if="decode"><View /></span>
              <span v-else><Hide /></span>
            </el-icon>
            <div v-for="(value, key) in detail.data" :key="key" class="data-wrapper">
              <span class="key-container">{{ key }}</span>
              <div class="value-container">
                <pre v-if="decode">{{ dataDecode(value) }}</pre>
                <span v-else>{{ value }}</span>
              </div>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
    <el-dialog v-model="dialogFormVisible" title="查看资源" width="55%">
      <VueCodeMirror v-model:modelValue="formatData" :readOnly="true" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { useRoute } from "vue-router"
import { getSecretDetailApi, deleteSecretApi } from "@/api/k8s/config/secret"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { viewOrch } from "@/utils/k8s/orch"

defineOptions({
  name: "SecretDetail"
})

// 折叠面板
const activeNames = ref(["metadata", "data"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// secret detail
const detail = ref({ metadata: {}, data: {} })

const getDetailData = async () => {
  await getSecretDetailApi({ namespace: namespace, name: name }).then((res) => {
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
  formatData.value = await viewOrch(name, "secrets", namespace)
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该Secret, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteSecretApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}

const decode = ref(false)
const decodeValue = () => {
  decode.value = !decode.value
}

// 解码
const dataDecode = (str: string) => {
  return atob(str)
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

.value-container > pre {
  word-break: break-all;
  white-space: pre-wrap;
}

.data-wrapper:not(:last-child) {
  margin-bottom: 15px;
}
</style>
