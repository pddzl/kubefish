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
        <el-collapse-item title="数据" name="data">
          <div style="padding: 20px; border-radius: 4px; background-color: #f9fbfd; margin-right: 20px">
            <div v-for="(value, key) in detail.data" :key="key">
              <span
                style="
                  font-size: 14px;
                  font-weight: 600;
                  font-style: normal;
                  color: #242e42;
                  text-shadow: 0 4px 8px rgba(36, 46, 66, 0.1);
                "
                >{{ key }}</span
              >
              <pre
                style="
                  margin-top: 8px;
                  padding: 12px;
                  border-radius: 4px;
                  background-color: #fff;
                  border: 1px solid #e3e9ef;
                "
                >{{ value }}</pre
              >
            </div>
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
import { getConfigMapDetailApi, deleteConfigMapApi } from "@/api/k8s/config/configMap"
import VueCodeMirror from "@/components/codeMirror/index.vue"
import MetaData from "@/components/k8s/metadata.vue"
import { ElMessage, ElMessageBox } from "element-plus"
import { viewOrch } from "@/utils/k8s/orch"

// 折叠面板
const activeNames = ref(["metadata", "data"])

// 路由
const route = useRoute()
const namespace = route.query.namespace as string
const name = route.query.name as string

// 获取configMap详情
const detail = ref({ metadata: {}, data: {} })

const getDetailData = async () => {
  await getConfigMapDetailApi({ namespace: namespace, name: name }).then((res) => {
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
  formatData.value = await viewOrch(name, "configmaps", namespace)
  dialogFormVisible.value = true
}

// 删除
const deleteFunc = async () => {
  ElMessageBox.confirm("此操作将永久删除该ConfigMap, 是否继续?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning"
  }).then(async () => {
    const res = await deleteConfigMapApi({ namespace: namespace, name: name })
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功!"
      })
    }
  })
}
</script>
